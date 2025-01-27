package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define the repository URL and the directory name
	repoURL := "https://github.com/trickest/wordlists.git"
	dirName := "wordlists"

	// Check if the directory already exists
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		fmt.Println("Directory exists, removing old directory...")
		removeOldDir(dirName)
	}

	fmt.Println("Cloning repository...")
	cloneRepo(repoURL)
}

func removeOldDir(dirName string) {
	err := os.RemoveAll(dirName)
	if err != nil {
		fmt.Printf("Error removing old directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Old directory removed successfully.")
}

func cloneRepo(repoURL string) {
	cmd := exec.Command("git", "clone", repoURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Repository cloned successfully.")
}
