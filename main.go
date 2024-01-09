package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func printImage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	err := r.ParseMultipartForm(10 << 20)
	fmt.Fprintf(w, "Request: %s\n", r)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error parsing form: %s\n", err)
		return
	}

	file, handler, err := r.FormFile("image")
	fmt.Fprintf(w, "File: %s\n", file)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error retrieving the file: %s\n", err)
		return
	}
	defer file.Close()

	imageName := handler.Filename
	fmt.Fprintf(w, "Image name: %s\n", imageName)
	imagePath := "./files/" + imageName 
	fmt.Fprintf(w, "Image path: %s\n", imagePath)
	outputFile, err := os.Create(imagePath)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error creating the file: %s\n", err)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error copying file: %s\n", err)
		return
	}

	os := r.FormValue("os")
	printer := r.FormValue("printer")

	var cmd *exec.Cmd

	if os == "linux" {
		cmd = exec.Command("bash", "scripts/autoprinter.sh", imageName, "1", printer)
	} else if os == "win" {
		cmd = exec.Command("powershell", "-File", "scripts/autoprinter.ps1", imageName, "1", printer)
	} else {
		cmd = exec.Command("echo", "OS not detected")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error printing image: %s\n", err)
		fmt.Fprintf(w, "Output: %s\n", output)
		return
	}

	fmt.Fprintf(w, "Print request processed successfully.\n%s\n", output)
}

func explainAPI(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	explanation := `
		Welcome to the AutoPrint API!

		To get a list of available printers, send a GET request to /printers with the following parameters:
		- os: Operating System ("linux" or "win")
		

		To print an image, send a POST request to /print with the following parameters:
		- image: Image name
		- copies: Number of copies
		- os: Operating System ("linux" or "win")
		- printer: Printer name

		Example:
		curl -X POST -d "image=123456789&copies=2&os=linux&printer=MyPrinter" http://localhost:8080/print
	`
	fmt.Fprintln(w, explanation)
}

type Printers struct {
	PrintersAvailable []string `json:"printers"`
}

func getPrinters(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	os := r.URL.Query().Get("os")
	var printers []string

	if os == "linux" {
		cmd := exec.Command("bash", "-c", "lpstat -a | awk '{print $1}'")
		output, err := cmd.Output()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error getting printers: %s\n", err)
			return
		}
		printers = strings.Fields(string(output))
	} else if os == "win" {
		cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Query \"SELECT Name FROM Win32_Printer\" | ForEach-Object { $_.Name }")
		output, err := cmd.Output()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error getting printers: %s\n", err)
			return
		}
		printers = strings.Fields(string(output))
	} else {
		fmt.Fprintf(w, "Invalid OS specified")
		return
	}

	response := Printers{PrintersAvailable: printers}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error encoding JSON: %s\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", explainAPI)
	http.HandleFunc("/print", printImage)
	http.HandleFunc("/printers", getPrinters)
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
