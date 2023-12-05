package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
)

func main() {

	var result int

	// Open and read file
	readFile, err := os.Open("input_day04.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	// Read file
	fileScanner := bufio.NewScanner(readFile)

	// Split to strings
	fileScanner.Split(bufio.ScanLines)

	// Map of won numbers
	won_numbers := make(map[int]int)

	// Map of instance number
	instance_numbers := make(map[int]int)

	// Number of line
	n := 0

	for fileScanner.Scan() {

		current_line := fileScanner.Text()
		intersection_size := 0

 		// Retrieve cards
		cards := strings.Split(current_line, ": ")
		cards_splitted := strings.Split(cards[1], " | ")
		winning_cards := strings.Split(cards_splitted[0], " ")
		my_cards := strings.Split(cards_splitted[1], " ")
		
		winning_cards_int, _ := sliceAtoi(winning_cards)
		my_cards_int, _ := sliceAtoi(my_cards)

		intersection_size = len(intersection_of_arrays(winning_cards_int, my_cards_int))

		won_numbers[n] = intersection_size
		instance_numbers[n] = 1
		n += 1
	}

	for j := 0; j < len(instance_numbers); j++ {
		for a := 0; a < instance_numbers[j]; a++ {
			for i := 1; i < won_numbers[j] + 1; i++ {
				if _, ok := instance_numbers[j+i]; ok {
					instance_numbers[j+i] += 1
				}
			} 
		}

		result += instance_numbers[j]
	}
	fmt.Println(result)
}

func process_line_by_line(n int, k int) (int){
	if n == 0 {
		return 1
	}

	return 1
}

func sliceAtoi(sa []string) ([]int, error) {
    si := []int{}
    for _, a := range sa {
		if a == "" {
			continue
		}
        i, err := strconv.Atoi(strings.TrimSpace(a))
        if err != nil {
            fmt.Println(err)
        }
        si = append(si, i)
    }
    return si, nil
}

func intersection_of_arrays(s1, s2 []int) (inter []int) {
    hash := make(map[int]bool)
    for _, e := range s1 {
        hash[e] = true
    }
    for _, e := range s2 {
        // If elements present in the hashmap then append intersection list.
        if hash[e] {
            inter = append(inter, e)
        }
    }
    //Remove dups from slice.
    //inter = removeDups(inter)
    return
}
