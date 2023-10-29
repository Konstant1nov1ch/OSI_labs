package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var fileMutex sync.Mutex

func runStressCommand(method, fileName string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	cmd := exec.Command("stress-ng", "--misaligned", "2", "--misaligned-method", method, "--metrics", "-t", "30s")
	cmd.Stdout = file
	cmd.Stderr = file

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}

func main() {
	methods := []string{"int32rd", "int32wr", "int32inc", "int32atomic", "int32wrnt", "all"}

	for _, method := range methods {
		fileName := fmt.Sprintf("stress_test_misaligned_%s.txt", method)
		runStressCommand(method, fileName)
	}
}
