package main

import (
    "bufio"
    "fmt"
    "os"
	"regexp"
	"strconv"
)

func main() {

	var result int64

	// Open and read file
	readFile, err := os.Open("input_day03.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	// Read file
	fileScanner := bufio.NewScanner(readFile)

	// Split to strings
	fileScanner.Split(bufio.ScanLines)

	// Special chars and numbers per line
	special_chars_per_line := make(map[int][]int)
	numbers_per_line := make(map[int][][]int)

	// Lines map
	lines := make(map[int]string)

	// Line number tracking
	n := 0

	for fileScanner.Scan() {

		current_line := fileScanner.Text()

		lines[n] = current_line
		// Find all special chars indexes in the line
		special_chars_re := regexp.MustCompile(`[*$#@%+=&/-]+`)
		special_chars_found := special_chars_re.FindAllStringSubmatchIndex(current_line, -1)
		special_chars_indexes := []int{}
		special_chars_per_line[n] = []int{}

		for _, subset := range special_chars_found {
			subset_slice := subset[:1]
			special_chars_indexes = append(special_chars_indexes, subset_slice...)
			special_chars_per_line[n] = special_chars_indexes	
		}

		// Find all number indexes in the line
		numbers_re := regexp.MustCompile(`[0-9]+`)
		numbers_found := numbers_re.FindAllStringSubmatchIndex(current_line, -1)
		numbers_per_line[n] = numbers_found

		n += 1
	}

	// Iterate over rows
	for i := 0; i < n; i ++ {
		fmt.Println("\n", i, lines[i])

		// Iterate over numbers
		number:
		for _, number_position := range numbers_per_line[i] {

			// Analyze special chars in the same line
			for _, special_chars_position := range special_chars_per_line[i] {
				if special_chars_position >= (number_position[0] - 1) && special_chars_position <= (number_position[1] + 1){
					part_number, _ := strconv.ParseInt(string(lines[i][number_position[0]:number_position[1]]), 10, 0)
					result += part_number

					fmt.Println("Same line")
					fmt.Println(strconv.FormatInt(part_number, 10), string(lines[i][special_chars_position]))

					continue number
				}
			}
		
			// Analyze special chars in the previous line
			if i > 0 {
				for _, special_chars_position := range special_chars_per_line[i-1] {
					if special_chars_position >= (number_position[0] - 1) && special_chars_position <= (number_position[1]){
						part_number, _ := strconv.ParseInt(string(lines[i][number_position[0]:number_position[1]]), 10, 0)
						result += part_number
						
						fmt.Println("Prevoius line")
						fmt.Println(strconv.FormatInt(part_number, 10), string(lines[i-1][special_chars_position]))

						continue number
					}
				}
			}
			// Analyze special chars in the next line
			if i < n {
				for _, special_chars_position := range special_chars_per_line[i+1] {
					if special_chars_position >= (number_position[0] - 1) && special_chars_position <= (number_position[1]){
						part_number, _ := strconv.ParseInt(string(lines[i][number_position[0]:number_position[1]]), 10, 0)
						result += part_number

						fmt.Println("Next line")
						fmt.Println(strconv.FormatInt(part_number, 10), string(lines[i+1][special_chars_position]))
						
						continue number
					}
				}
			}
		}
	}

	fmt.Println(result)
}