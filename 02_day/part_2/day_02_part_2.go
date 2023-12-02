package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
)

func main() {

	var result int64

	// Open and read file
	readFile, err := os.Open("input_day02.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	// Read file
	fileScanner := bufio.NewScanner(readFile)

	// Split to strings
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		
		current_line := fileScanner.Text()

		// Retrieve cubes sets
		cubes_sets := strings.Split(current_line, ": ")
		cubes_sets_separated := strings.Split(cubes_sets[1], "; ")

		// Max number of cubes
		colors_numbers_map := map[string]int64 {"red" : 0, "green": 0, "blue": 0}

		// Iterate over sets
		for _, set := range cubes_sets_separated {
			colors_and_numbers := strings.Split(set, ", ")
			
			// Iterate over colors
			for _, color_and_number := range colors_and_numbers {
				color_and_number_splitted := strings.Split(color_and_number, " ")
				number := color_and_number_splitted[0]
				number_int, err := strconv.ParseInt(number, 10, 0)
				if err != nil {}
				color := color_and_number_splitted[1]

				if number_int > colors_numbers_map[color] {
					colors_numbers_map[color] = number_int
				}
			}
		}
		result += colors_numbers_map["red"] * colors_numbers_map["green"] * colors_numbers_map["blue"]
	}
	fmt.Println(result)
}