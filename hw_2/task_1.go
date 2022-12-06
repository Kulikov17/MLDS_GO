package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./hw_2/task_1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 256)

	var text = ""

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		text += string(data[:n])
	}

	candidats := strings.Split(text, "\n")

	var dict = make(map[string]int)

	for _, candidat := range candidats {
		if candidat == "" {
			continue
		}

		info := strings.Split(candidat, " ")

		intChoice, _ := strconv.Atoi(info[1])
		dict[info[0]] += intChoice
	}

	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, dict[k])
	}
}
