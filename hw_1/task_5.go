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
	numbers := strings.Split(text, " ")

	var dict = make(map[string]int)

	for _, val := range numbers {
		if val == "" {
			continue
		}

		if dict[val] == 0 {
			dict[val] += 1
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}

}
