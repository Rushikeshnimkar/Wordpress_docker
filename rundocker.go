package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Define the command (docker-compose up -d)
	cmd := exec.Command("docker-compose", "-f", "./wp-docker/docker-compose.yml", "up", "-d")

	// Set the working directory if needed
	// cmd.Dir = "/path/to/your/docker/project"

	// Set environment variables if needed
	// cmd.Env = append(os.Environ(), "ENV_VAR=value")

	// Redirect standard output and standard error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker-compose: %v", err)
	}

	log.Println("Docker Compose up completed successfully.")
}
