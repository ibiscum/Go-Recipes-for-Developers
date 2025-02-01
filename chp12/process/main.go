package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Run "go build" to build the subprocess in the "sub" directory
func buildProgram() {
	cmd := exec.Command("go", "build", "-o", "subprocess", ".")
	cmd.Dir = "sub"
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	// The build command will not print anything if successful. So if
	// there is any output, it is a failure.
	if len(output) > 0 {
		panic(string(output))
	}
}

// Run the subprocess built before for 10ms, reading from the output
// and error pipes concurrently
func runSubProcessStreamingOutputs() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sub/subprocess")
	fmt.Println("Streaming output running for 10 ms")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	// Read from stderr from a separate goroutine
	go func() {
		io.Copy(os.Stderr, stderr)
	}()
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, stdout)
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

// Run the build subprocess for 10 ms with combined output
func runSubProcessCombinedOutput() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sub/subprocess")
	fmt.Println("Streaming combined output running for 10 ms")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

func filterSubprocessOutput() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sub/subprocess")
	fmt.Println("Filtering stdout running for 10 ms")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// Read from the pipe in a separate goroutine
	go func() {
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Index(line, "0") != -1 {
				fmt.Printf("Filtered line: %s\n", line)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner error: %v", err)
		}
	}()
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	buildProgram()
	runSubProcessStreamingOutputs()
	runSubProcessCombinedOutput()
	filterSubprocessOutput()
}
