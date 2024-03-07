package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("ACTIVE")
	http.HandleFunc("/webhook", handleWebhook)
	http.ListenAndServe(":3000", nil)

}

func handleWebhook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Webhook received")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// read the payload
	payload, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error reading request body")
		return
	}

	// handle the signature from the header
	signature := r.Header.Get("X-Hub-Signature-256")

	// check for correct
	if !isValidSignature(payload, signature, os.Getenv("WHSEC")) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid signature")
		return
	}

	updateAndPushReadme()

	fmt.Fprint(w, "Webhook received and processed")

}

func isValidSignature(payload []byte, providedSignature, secret string) bool {

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)

	expectedSignature := "sha256=" + hex.EncodeToString(expectedMAC)

	return hmac.Equal([]byte(providedSignature), []byte(expectedSignature))

}

func updateAndPushReadme() {

	readmePath := "./README.md"

	file, err := os.OpenFile(readmePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening README:", err)
		return
	}
	defer file.Close()

	// append new log entry at the end
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	newEntry := fmt.Sprintf("Push: %s\n", timestamp)
	if _, err := file.WriteString(newEntry); err != nil {
		fmt.Println("Error writing to README:", err)
		return
	}

	// git commands to push webhook triggers
	if err := exec.Command("git", "add", readmePath).Run(); err != nil {
		log.Printf("Error adding README to git: %v", err)
		return
	}

	if err := exec.Command("git", "commit", "-m", "Update README").Run(); err != nil {
		log.Printf("Error committing README: %v", err)
		return
	}

	if err := exec.Command("git", "push", "origin", "main").Run(); err != nil {
		log.Printf("Error pushing README to remote: %v", err)
		return
	}

}
