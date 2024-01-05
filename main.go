package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func printImage(w http.ResponseWriter, r *http.Request) {
    // Parse parameters from the POST form
    err := r.ParseForm()
    if err != nil {
        log.Println(err)
        fmt.Fprintf(w, "Error parsing form: %s\n", err)
        return
    }

    imageName := r.FormValue("image")
    copies := r.FormValue("copies")
    height := r.FormValue("height")
    width := r.FormValue("width")

    // Execute the Bash script with the provided parameters
    cmd := exec.Command("bash", "autoprinter.sh", imageName, copies, height, width)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Println(err)
        fmt.Fprintf(w, "Error printing image: %s\n", err)
        return
    }
    
    fmt.Fprintf(w, "Print request processed successfully.\n%s\n", output)
}

func main() {
	fmt.Println("Starting server...")
    http.HandleFunc("/print", printImage)
	fmt.Println("Listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
