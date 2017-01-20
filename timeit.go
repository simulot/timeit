package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const TSFmt = "2006-01-02 15:04:05"

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "%s command line\nMissing command line", os.Args[0])
	}

	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	cmd := exec.Command(os.Args[1], args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()
	err := cmd.Run()
	end := time.Now()
	if err != nil {
		fmt.Println("Execution failed:", err)
	}
	duration := time.Duration(int(100*end.Sub(start).Seconds())) * time.Second / 100
	fmt.Println("Execution start at:", start.Format(TSFmt))
	fmt.Println("Execution stop at: ", end.Format(TSFmt))
	fmt.Println("Duration:          ", duration)
}
