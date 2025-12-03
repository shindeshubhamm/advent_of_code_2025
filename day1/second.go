package day1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countPassThroughZeroRight(position, distance int) int {
	return (position + distance) / 100
}

func countPassThroughZeroLeft(position, distance int) int {
	if position == 0 {
		return distance / 100
	}
	if distance >= position {
		return (distance - position + 100) / 100
	}
	return 0
}

func Second() {
	data, err := os.ReadFile(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	file_data := string(data)
	lines_data := strings.Split(file_data, "\n")

	dial_position := 50
	passes_through_zero := 0

	for _, line := range lines_data {
		line = strings.TrimSpace(line)

		direction := string(line[0])
		distance, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "R":
			passes_through_zero += countPassThroughZeroRight(dial_position, distance)
			dial_position = (dial_position + distance) % 100
		case "L":
			passes_through_zero += countPassThroughZeroLeft(dial_position, distance)
			dial_position = (dial_position - distance) % 100
			if dial_position < 0 {
				dial_position += 100
			}
		}
	}

	fmt.Println(passes_through_zero)
}
