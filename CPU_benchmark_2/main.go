package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var fileMutex sync.Mutex

func runStressCommand(cpu int, method string, fileName string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	cmd := exec.Command("stress-ng", "--cpu", fmt.Sprintf("%d", cpu), "--cpu-method", method, "--metrics", "-t", "30s")
	cmd.Stdout = file
	cmd.Stderr = file

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}

func main() {
	for i := 1; i <= 15; i += 2 {
		fileName := fmt.Sprintf("stress_test_cpu_ipv4checksum_%d.txt", i)
		runStressCommand(i, "ipv4checksum", fileName)
	}
}
