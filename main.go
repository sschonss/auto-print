package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
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

	fmt.Fprintf(w, "Printing image %s %s times...\n", imageName, copies)

	var cmd *exec.Cmd

	fmt.Println("OS:", os)

	if os == "linux" {
		fmt.Println("Linux detected")
		cmd = exec.Command("bash", "scripts/autoprinter.sh", imageName, copies)
	} else if os == "win" {
		fmt.Println("Windows detected")
		cmd = exec.Command("powershell", "-File", "scripts/autoprinter.ps1", imageName, copies)

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

		To print an image, send a POST request to /print with the following parameters:
		- image: Image name
		- copies: Number of copies
		- os: Operating System ("linux" or "win")

		Example:
		curl -X POST -d "image=123456789&copies=2&os=linux" http://localhost:8080/print
	`
	fmt.Fprintln(w, explanation)
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", explainAPI)
	http.HandleFunc("/print", printImage)
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
