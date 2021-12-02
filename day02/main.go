package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction struct {
	X int
	Z int
}

var (
	Forward = Direction{X: 1, Z: 0}
	Up      = Direction{X: 0, Z: -1}
	Down    = Direction{X: 0, Z: 1}
)

type Command struct {
	Direction Direction
	Units     int
}

func ReadCommands(input io.ReadCloser) []Command {
	commands := []Command{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		l := strings.Split(txt, " ")
		direction := l[0]
		units, err := strconv.Atoi(l[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			commands = append(commands, Command{Forward, units})
		case "up":
			commands = append(commands, Command{Up, units})
		case "down":
			commands = append(commands, Command{Down, units})
		}
	}
	return commands
}

func SubmarineCanDive(commands []Command) int {
	X, Z := 0, 0
	for _, command := range commands {
		X += command.Direction.X * command.Units
		Z += command.Direction.Z * command.Units
	}
	return X * Z
}

func SubmarineCanDiveWithAim(commands []Command) int {
	X, Z, Aim := 0, 0, 0
	for _, command := range commands {
		if command.Direction == Forward {
			Z += Aim * command.Units
		}
		X += command.Direction.X * command.Units
		Aim += command.Direction.Z * command.Units
	}
	return X * Z
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := ReadCommands(f)

	fmt.Printf("result: %d\n", SubmarineCanDive(c))
	fmt.Printf("result: %d\n", SubmarineCanDiveWithAim(c))
}
