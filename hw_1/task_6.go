package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./hw_1/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	var text = ""

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		text += string(data[:n])
	}

	text = strings.ReplaceAll(text, "\n", " ")

	words := strings.Split(text, " ")

	var dict = make(map[string]int)

	for _, val := range words {
		if val == "" {
			continue
		}

		dict[val] += 1
	}

	minVal := 0
	minKey := ""

	for key, val := range dict {
		if val > minVal {
			minVal = val
			minKey = key

			continue
		}

		if val == minVal {
			if strings.Compare(minKey, key) == 1 {
				minKey = key
			}
		}
	}

	fmt.Println(minKey)
}
