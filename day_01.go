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

	readFile, err := os.Open("input_day01.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

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

		// Continue if none of digits or words found
		if len(keys_of_map) == 0 {
			continue
		}

		// Get first and last indexes
		first := found_digits_indexes[Min(keys_of_map)]
		last := found_digits_indexes[Max(keys_of_map)]

		// Create final strig
		temp_combination := first + last

		// Replace string value to digit in final string
		fmt.Println(fileScanner.Text())
		fmt.Println(found_digits_indexes)
		fmt.Println(temp_combination)

		temp_combination_int, err := strconv.ParseInt(temp_combination, 10, 0)

		if err != nil {
			fmt.Println(err)
		}

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
