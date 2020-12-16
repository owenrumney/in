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

	time.Sleep(parseTime(t))
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	_ = cmd.Start()
}

func parseTime(t string) time.Duration {
	t = strings.ToLower(t)
	if strings.HasSuffix(t, "s") {
		parsedTime, _ := strconv.Atoi(strings.TrimSuffix(t, "s"))
		return time.Duration(parsedTime) * time.Second
	}
	if strings.HasSuffix(t, "m") {
		parsedTime, _ := strconv.Atoi(strings.TrimSuffix(t, "m"))
		return time.Duration(parsedTime) * time.Minute
	}
	if strings.HasSuffix(t, "h") {
		parsedTime, _ := strconv.Atoi(strings.TrimSuffix(t, "h"))
		return time.Duration(parsedTime) * time.Hour
	}
	return time.Second
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
