package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH = "./input.txt"

func main() {
	data, err := os.ReadFile(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	file_data := string(data)
	lines_data := strings.Split(file_data, "\n")

	dial_position := 50
	position_on_zero := 0

	for _, line := range lines_data {
		line = strings.TrimSpace(line)

		direction := string(line[0])
		distance, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "R":
			dial_position = (dial_position + distance) % 100
		case "L":
			dial_position = (dial_position - distance) % 100
			if dial_position < 0 {
				dial_position += 100
			}
		}

		if dial_position == 0 {
			position_on_zero += 1
		}
	}

	fmt.Println(position_on_zero)
}
