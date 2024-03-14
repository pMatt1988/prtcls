package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: prtcls <port>")
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid port number:", err)
		os.Exit(1)
	}

	pid, err := findPIDByPort(port)
	if err != nil {
		fmt.Println("Error finding PID:", err)
		os.Exit(1)
	}

	if pid != -1 {
		fmt.Println("Found process with PID:", pid)

		if err := killProcess(pid); err != nil {
			fmt.Println("Error killing process:", err)
			os.Exit(1)
		} else {
			fmt.Println("Process killed successfully")
			os.Exit(0)
		}
	} else {
		fmt.Println("No process found using that port.")
	}
}

func findPIDByPort(port int) (int, error) {
	cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%d", port))
	output, err := cmd.Output()
	if err != nil {
		return -1, err
	}

	lines := strings.Split(string(output), "\n")

	if len(lines) < 2 {
		return -1, nil
	}

	fields := strings.Fields(lines[1])
	pid, err := strconv.Atoi(fields[1])
	if err != nil {
		return -1, err
	}

	return pid, nil

}

func killProcess(pid int) error {
	process, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		return process.Kill()
	}

	return nil
}
