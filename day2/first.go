package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH = "./day2/input.txt"

type Range struct {
	Start int
	End   int
}

func canIncludeInvalidIDs(r Range) bool {
	// old length numbers can't have twice repeating numbers
	if len(strconv.Itoa(r.Start))%2 == 1 && len(strconv.Itoa(r.Start)) == len(strconv.Itoa(r.End)) {
		return false
	}
	return true
}

func readRanges() []Range {
	data, err := os.ReadFile(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	raw_ranges := strings.Split(strings.TrimSpace(string(data)), ",")

	ranges := []Range{}
	for _, raw_range := range raw_ranges {
		parts := strings.Split(raw_range, "-")
		start, serr := strconv.Atoi(parts[0])
		end, eerr := strconv.Atoi(parts[1])

		if serr != nil || eerr != nil {
			log.Fatal(serr, eerr)
		}

		r := Range{Start: start, End: end}
		if canIncludeInvalidIDs(r) {
			ranges = append(ranges, r)
		}
	}
	return ranges
}

func sumInvalidIDsFromRange(r Range) int {
	sum := 0
	invalidIDs := []string{}
	for i := r.Start; i <= r.End; i++ {
		id_str := strconv.Itoa(i)
		str_len := len(id_str)

		if str_len%2 == 1 { // avoid odd length numbers
			continue
		}

		half_way := str_len / 2
		if id_str[:half_way] == id_str[half_way:] {
			invalidIDs = append(invalidIDs, id_str)
		}
	}

	for _, ids := range invalidIDs {
		num, _ := strconv.Atoi(ids)
		sum += num
	}

	return sum
}

func First() {
	data := readRanges()
	sum := 0
	for _, r := range data {
		sum += sumInvalidIDsFromRange(r)
	}
	fmt.Println(sum)
}
