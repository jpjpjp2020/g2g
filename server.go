package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// func handleWebhook(w http.ResponseWriter, r *http.Request) {
// 	if os.Getenv("ACTIVATE_TOOL") != "true" {
// 		fmt.Println("Tool is deactivated.")
// 		return
// 	}

// 	// Parse the webhook payload
// 	// Extract necessary data (commit message, author, etc.)

// 	// Authenticate with GitHub
// 	// Create a dummy commit in GitHub

// 	fmt.Fprintf(w, "Commit mirrored to GitHub")
// }

func main() {

	// ANSI color codes
	green := "\033[32m"
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%sError loading .env file%s\n", red, reset)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%sEnter your acc email:%s\n", yellow, reset)
	scanner.Scan()
	gEmail := scanner.Text()

	fmt.Printf("%sEnter your acc password:%s\n", yellow, reset)
	// Note: This will not obscure the input. Consider a more secure way to handle password input.
	scanner.Scan()
	gPassword := scanner.Text()

	gLink, exists := os.LookupEnv("GLINK")
	if exists {
		fmt.Printf("%sActivity tab link found in .env%s\n", green, reset)
	} else {
		fmt.Printf("%sNo activty tab link set in .env%s\n", red, reset)
	}

	fmt.Println(gEmail)
	fmt.Println(gPassword)
	fmt.Println(gLink)

	// http.HandleFunc("/webhook", handleWebhook)
	// http.ListenAndServe(":3000", nil)
}
