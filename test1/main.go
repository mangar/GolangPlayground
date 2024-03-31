package main

import "fmt"

func main() {

	fmt.Printf("\n\n----------------\n")

	entry := "abcdefaghbjkclmn"
	replaceDuplicatedWith("_", entry)

	repeatCount(entry)

}

func repeatCount(entry string) map[string]int {

	// fmt.Printf(">> repeatCount ---->>> \n")
	// fmt.Println("ENTRY: " + entry)

	index := make(map[string]int)

	for _, ch := range entry {
		_ch := string(ch)

		index[_ch]++

		// value, exist := index[_ch]
		// if exist {
		// 	index[_ch] = value + 1
		// } else {
		// 	index[_ch] = 1
		// }

	}

	// for k, v := range index {
	// 	if v > 1 {
	// 		fmt.Printf("%v -> %v\n", k, v)
	// 	}
	// }

	return index
}

func repeatCountIFExist(entry string) map[string]int {

	index := make(map[string]int)

	for _, ch := range entry {
		_ch := string(ch)

		value, exist := index[_ch]
		if exist {
			index[_ch] = value + 1
		} else {
			index[_ch] = 1
		}

	}

	// for k, v := range index {
	// 	if v > 1 {
	// 		fmt.Printf("%v -> %v\n", k, v)
	// 	}
	// }

	return index
}

func replaceDuplicatedWith(replacech string, entry string) string {
	// fmt.Println("ENTRY: " + entry)

	index := []string{}
	new_entry := ""

	for _, ch := range entry {
		_ch := string(ch)
		if !stringInSlice(_ch, index) {
			index = append(index, _ch)
			new_entry += _ch
		} else {
			new_entry += replacech
		}

	}
	// fmt.Println("  NEW: " + new_entry)
	return new_entry
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}