package main;

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	cmd string
	arg int
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	horizontalPos := 0;
	depth := 0;
	aim := 0;

	for scanner.Scan() {
		line := scanner.Text();

		cmd := command{};
		fmt.Sscanf(line, "%s %d", &cmd.cmd, &cmd.arg);

		if cmd.cmd == "forward" {
			horizontalPos += cmd.arg;
			depth += cmd.arg * aim;
		} else if cmd.cmd == "down" {
			aim += cmd.arg;
		} else if cmd.cmd == "up" {
			aim -= cmd.arg;
		}
	}
	fmt.Println(depth * horizontalPos)
}