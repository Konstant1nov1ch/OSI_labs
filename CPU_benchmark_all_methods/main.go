package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func runStressCommand(wg *sync.WaitGroup, cpu int, method string, fileName string) {
	defer wg.Done()

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
	var wg sync.WaitGroup

	wg.Add(2)

	// Файл для метода gray
	go func() {
		runStressCommand(&wg, 14, "gray", "stress_test_cpu_gray.txt")
	}()

	// Файл для метода ipv4checksum
	go func() {
		runStressCommand(&wg, 14, "ipv4checksum", "stress_test_cpu_ipv4checksum.txt")
	}()

	wg.Wait()
}
