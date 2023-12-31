package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"index/suffixarray"
)

func main() {

	var result int64
	digits_map := map[string]int{"one": 1,"two": 2,"three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	// Open and read file
	readFile, err := os.Open("input_day01.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	// Read file
	fileScanner := bufio.NewScanner(readFile)

	// Split to strings
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		
		// Create suffix array object based on read string
		temp_string_suffix_array := suffixarray.New([]byte(fileScanner.Text()))

		// Declare Map object for found digits and words
		var found_digits_indexes = make(map[int]string)

		// Search for digits and words - save to map
		for word, digit := range digits_map{
			offset_digit := temp_string_suffix_array.Lookup([]byte(strconv.Itoa(digit)), -1)
			offset_word := temp_string_suffix_array.Lookup([]byte(word), -1)

			for _, index := range offset_digit{
				found_digits_indexes[index] = strconv.Itoa(digit)
			}
			for _, index := range offset_word{
				found_digits_indexes[index] = strconv.Itoa(digit)
			}
		}

		// Get found indexes
		keys_of_map := make([]int, 0, len(found_digits_indexes))
		for k, _ := range found_digits_indexes {
			keys_of_map = append(keys_of_map, k)
		}

		// Get first and last indexes
		first := found_digits_indexes[Min(keys_of_map)]
		last := found_digits_indexes[Max(keys_of_map)]

		// Create final number
		temp_combination_int, err := strconv.ParseInt(first + last, 10, 0)
		if err != nil {
			fmt.Println(err)
		}

		// Add to final result sum
		result += temp_combination_int
	}

	fmt.Println(result)
	readFile.Close()
} 

func Min(array []int) (int) {
    var min int = array[0]
    for _, value := range array {
        if min > value {
            min = value
        }
    }
    return min
}

func Max(array []int) (int) {
    var max int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
    }
    return max
}
