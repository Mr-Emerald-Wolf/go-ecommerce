package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var directories = []string{
	"cmd/myapp",
	"internal/api/middleware",
	"internal/api/handlers",
	"internal/config",
	"internal/database/migrations",
	"internal/models",
	"internal/repository",
	"internal/service",
	"internal/util",
	"migrations",
	"scripts",
	"pkg",
	"static",
	"templates",
}

var files = []string{
	".env.sample",
	".gitignore",
	"go.mod",
	"go.sum",
	"README.md",
	"LICENSE",
	"cmd/myapp/main.go",
	"internal/api/api.go",
	"internal/api/handlers/user_handler.go",
	"internal/config/config.go",
	"internal/database/db.go",
	"internal/models/user.go",
	"internal/repository/user_repository.go",
	"internal/service/user_service.go",
	"internal/util/utils.go",
}

func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Create directories
	for _, dir := range directories {
		path := filepath.Join(projectRoot, dir)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
		fmt.Printf("Created directory: %s\n", dir)
	}

	// Create files
	for _, file := range files {
		path := filepath.Join(projectRoot, file)
		_, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", file, err)
			return
		}
		fmt.Printf("Created file: %s\n", file)
	}

	fmt.Println("Project structure created successfully!")
}
