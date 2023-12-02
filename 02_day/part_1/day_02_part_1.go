package main

import (
    "bufio"
    "fmt"
    "os"
	"regexp"
	"strings"
	"strconv"
)

func main() {

	var result int64

	// Possible number of cubes
	colors_numbers_map := map[string]int64 {"red" : 12, "green": 13, "blue": 14}

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
		
		// Retrive Game Id
		game_id_re := regexp.MustCompile(`[0-9]+`)
		game_id := game_id_re.FindStringSubmatch(current_line)[0]
		game_id_int, err := strconv.ParseInt(game_id, 10, 0)
		result += game_id_int
		if err != nil {}

		// Retrieve cubes sets
		cubes_sets := strings.Split(current_line, ": ")
		cubes_sets_separated := strings.Split(cubes_sets[1], "; ")

		out:
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
						result -= game_id_int
						break out
					}
				}
			}
	}
	fmt.Println(result)
}