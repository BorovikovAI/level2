package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const ShellToUse = "bash"

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func ExecShellCommand(line string) {
	out, err, errout := Shellout(line)
	if err != "" {
		log.Printf("ERROR: %v\n", err)
	}
	fmt.Println("STDOUT:")
	fmt.Println(out)
	fmt.Println("STDERR:")
	fmt.Println(errout)
}

func Cmd(cmd string, shell bool) []byte {

	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	}
}

func main() {
	commands := []string{"ls", "ls -ltr", "echo 'cock'", "cd"}
	for _, line := range commands {
		ExecShellCommand(line)
	}

	result := Cmd("cd ..", true)
	fmt.Println(string(result))
}
