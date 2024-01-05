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

    // Execute the Bash script with the provided parameters
    cmd := exec.Command("bash", "autoprinter.sh", imageName, copies)
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

		Example:
		curl -X POST -d "image=123456789&copies=2" http://localhost:8080/print
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
