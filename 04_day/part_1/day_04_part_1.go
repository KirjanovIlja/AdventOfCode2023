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

	for fileScanner.Scan() {

		current_line := fileScanner.Text()

		line_result := 0

 		// Retrieve cards
		cards := strings.Split(current_line, ": ")
		cards_splitted := strings.Split(cards[1], " | ")
		winning_cards := strings.Split(cards_splitted[0], " ")
		my_cards := strings.Split(cards_splitted[1], " ")
		
		winning_cards_int, _ := sliceAtoi(winning_cards)
		my_cards_int, _ := sliceAtoi(my_cards)

		intersection := intersection(winning_cards_int, my_cards_int)
		
		if len(intersection) > 0 {
			line_result = 1
			for i := 0; i < len(intersection) - 1; i++ {
				line_result = 2 * line_result
			}
		}

		result += line_result
		fmt.Println(intersection)
		fmt.Println(line_result)
	}
	fmt.Println(result)
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

func intersection(s1, s2 []int) (inter []int) {
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
