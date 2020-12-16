package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fatal("you need to pass a time and command")
	}
	args := os.Args[1:]
	t := args[0]
	command := args[1:]
	timeInSeconds := getTimeInSeconds(t)
	time.Sleep(timeInSeconds)
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stderr
	cmd.Start()
}

func getTimeInSeconds(t string) time.Duration {
	t = strings.ToLower(t)
	if strings.HasSuffix(t, "s") {
		parsedTime, _ := strconv.Atoi(strings.TrimSuffix(t, "s"))
		return time.Duration(parsedTime) * time.Second
	}
	return 1 * time.Second
}

func fatal(message string) {
	fmt.Print(`Usage:

in <time> <command>

example - 
in 5s echo hello

`)
	fmt.Println(message)
	os.Exit(1)
}
