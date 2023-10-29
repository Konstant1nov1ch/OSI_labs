package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var fileMutex sync.Mutex

func runStressCommand(iodepth int, fileName string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	cmd := exec.Command("stress-ng", "--iomix", fmt.Sprintf("%d", iodepth), "--metrics", "-t", "30s")
	cmd.Stdout = file
	cmd.Stderr = file

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		fileName := fmt.Sprintf("stress_test_iomix_%d.txt", i)
		runStressCommand(i, fileName)
	}
}
