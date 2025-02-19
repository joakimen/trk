package main

// generated by ChatGPT, leaving as-is

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type InvocationData map[string]int

func getAbsolutePath(script string) (string, error) {
	if filepath.IsAbs(script) {
		return script, nil
	}
	absPath, err := filepath.Abs(script)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func readJSONFile(filePath string) (InvocationData, error) {
	data := InvocationData{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, return an empty map
			return data, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func writeJSONFile(filePath string, data InvocationData) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, fileData, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycli <script_path>")
		os.Exit(1)
	}

	scriptWithArgs := os.Args[1]
	scriptPath := strings.Fields(scriptWithArgs)[0] // Strip arguments

	var key string
	if filepath.IsAbs(scriptPath) {
		key = scriptPath
	} else if strings.HasPrefix(scriptPath, ".") || strings.HasPrefix(scriptPath, "/") {
		absScriptPath, err := getAbsolutePath(scriptPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting absolute script path: %v\n", err)
			os.Exit(1)
		}
		key = absScriptPath
	} else {
		key = scriptPath
	}

	stateDir := filepath.Join(os.Getenv("HOME"), ".local", "share", "trk")
	stateFilePath := filepath.Join(stateDir, "data.json")

	err := os.MkdirAll(stateDir, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating state directory: %v\n", err)
		os.Exit(1)
	}

	data, err := readJSONFile(stateFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading JSON file: %v\n", err)
		os.Exit(1)
	}

	data[key]++
	if err := writeJSONFile(stateFilePath, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to JSON file: %v\n", err)
		os.Exit(1)
	}
}
