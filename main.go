package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/abtsousa/hook/internal/gemini"
	"github.com/abtsousa/hook/safe_error"
)

// Path to the API key.
const API_PATH = ".api-key"

var API_KEY string

func main() {
	// Check parameters
	params := strings.Join(os.Args[1:], " ")
	if len(params) == 0 {
		log.Fatal("You must provide a query.")
	}

	// Get API key
	if os.Getenv("GEMINI_API_KEY") != "" {
		API_KEY = os.Getenv("GEMINI_API_KEY")
	} else {
		api_key, err := os.ReadFile(API_PATH)
		if err != nil {
			log.Fatalf("Please save your API key in the file %s", API_PATH)
		}
		API_KEY = strings.TrimSpace(string(api_key))
	}
	safe_error.API_KEY = API_KEY
	gemini.API_KEY = API_KEY

	// Run
	run(params)
}

func run(params string) {
	client, err := gemini.NewClient(time.Second * 45)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := client.Query(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
