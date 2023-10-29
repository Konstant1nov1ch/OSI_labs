package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

var fileMutex sync.Mutex

func runStressCommand(netlinkTask int, fileName string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	cmd := exec.Command("sudo", "stress-ng", "--netlink-task", fmt.Sprintf("%d", netlinkTask), "--metrics", "-t", "30s")
	cmd.Stdout = file
	cmd.Stderr = file

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error obtaining stdin pipe:", err)
		return
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter sudo password: ")
	password, _ := reader.ReadString('\n')
	_, err = io.WriteString(stdin, password)
	if err != nil {
		fmt.Println("Error writing to stdin:", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}

func main() {
	for i := 1; i <= 10; i += 2 {
		fileName := fmt.Sprintf("stress_test_netlink_task_%d.txt", i)
		runStressCommand(i, fileName)
	}
}
