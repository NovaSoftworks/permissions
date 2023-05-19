package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/NovaSoftworks/permissions/internal/config"
)

func main() {
	log.Println("permissions-service starting...")

	definitionsPath := getDefinitionsPath()
	definitionsStr := getDefinitions(definitionsPath)
	permissionsConfig, err := config.NewPermissionsConfig(definitionsStr)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(len(permissionsConfig.Permissions))
}

func getDefinitionsPath() string {
	var definitionsFlag string
	flag.StringVar(&definitionsFlag, "d", "definitions.yaml", "The path of the permissions definitions file.")
	flag.StringVar(&definitionsFlag, "definitions", "./definitions.yaml", "The path of the permissions definitions file.")
	flag.Parse()

	return definitionsFlag
}

func getDefinitions(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return string(data)
}
