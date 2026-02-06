package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		http.Error(w, "Please provide a 'file' parameter", http.StatusBadRequest)
		return
	}

	// Unsafe: The 'file' parameter is used directly in a command.
	// This is a command injection vulnerability.
	// A malicious user could provide input like "; ls -la"
	cmd := exec.Command("sh", "-c", "ls "+file)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	fmt.Fprintf(w, "Output:\n%s\n", out)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	fmt.Println("Visit http://localhost:8080/?file=.")
	fmt.Println("To exploit, visit http://localhost:8080/?file=.;ls")
	http.ListenAndServe(":8080", nil)
}
