package main

import (
	"fmt"
	"github.com/carprks/website/backend"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func _main(args []string) int {
	if len(args) >= 1 {
		if args[0] == "localDev" {
			err := godotenv.Load()
			if err != nil {
				fmt.Println(fmt.Sprintf(".env err: %v", err))
			}
		}
	}

	port := "80"
	if len(os.Getenv("PORT")) >= 2 {
		port = os.Getenv("PORT")
	}

	fmt.Println(fmt.Sprintf("Starting Server on Port: %s", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), backend.Routes()); err != nil {
		fmt.Println(fmt.Sprintf("HTTP Err: %v", err))
		return 1
	}

	return 0
}

func main() {
	os.Exit(_main(os.Args[1:]))
}