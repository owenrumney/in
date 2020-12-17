package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fatal("you need to pass a time and command")
	}
	args := os.Args[1:]
	re, err := regexp.Compile("^(\\d+)\\s?(.+?)\\s(.+)")
	if err != nil || re == nil {
		fatal("the time provided was unrecognisable")
	}
	components := re.FindStringSubmatch(strings.Join(args, " "))
	if len(components) < 4 {
		fatal("not enough information provided")
	}

	t := getTime(components)
	command := strings.Split(components[3], " ")

	time.Sleep(t)
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	_ = cmd.Start()
}

func getTime(components []string) time.Duration {
	num, err := strconv.Atoi(components[1])
	if err != nil {
		fatal("time format not recognised")
	}
	duration := strings.ToLower(components[2])

	multiplier := time.Duration(num)
	switch duration {
	case "s", "secs", "sec", "second", "seconds":
		return multiplier * time.Second
	case "m", "mins", "minute", "minutes", "min":
		return multiplier * time.Minute
	}
	fatal("time format not recognised")
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
