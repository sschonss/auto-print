package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func printImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error parsing form: %s\n", err)
		return
	}

	imageName := r.FormValue("image")
	copies := r.FormValue("copies")
	os := r.FormValue("os")
	printer := r.FormValue("printer")

	fmt.Fprintf(w, "Printing image %s %s times on printer %s...\n", imageName, copies, printer)

	var cmd *exec.Cmd

	fmt.Println("OS:", os)

	if os == "linux" {
		fmt.Println("Linux detected")
		cmd = exec.Command("bash", "scripts/autoprinter.sh", imageName, copies, printer)
	} else if os == "win" {
		fmt.Println("Windows detected")
		cmd = exec.Command("powershell", "-File", "scripts/autoprinter.ps1", imageName, copies, printer)
	} else {
		cmd = exec.Command("echo", "OS not detected")
		fmt.Println("OS not detected")
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
