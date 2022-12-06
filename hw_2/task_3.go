package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isUserCreate(db map[string]int, key string) bool {
	_, ok := db[key]

	if ok {
		return true
	}

	return false
}

func deposit(db map[string]int, user string, value int) {
	if !isUserCreate(db, user) {
		db[user] = 0
	}

	db[user] += value
}

func withdraw(db map[string]int, user string, value int) {
	if !isUserCreate(db, user) {
		db[user] = 0
	}

	db[user] -= value
}

func transfer(db map[string]int, userFrom string, userTo string, value int) {
	withdraw(db, userFrom, value)
	deposit(db, userTo, value)
}

func balance(db map[string]int, user string) {
	if isUserCreate(db, user) {
		fmt.Println(db[user])
	} else {
		fmt.Println("ERROR")
	}
}

func Percent(percent int, all int) float64 {
	return (float64(all) * float64(percent)) / float64(100)
}

func income(db map[string]int, percent int) {
	for user, val := range db {
		if val > 0 {
			db[user] += int(Percent(percent, db[user]))
		}
	}
}

func main() {
	file, err := os.Open("./hw_2/task_3.txt")
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

	transactions := strings.Split(text, "\n")

	var db = make(map[string]int)

	for _, transaction := range transactions {
		if transaction == "" {
			continue
		}

		tInfo := strings.Split(transaction, " ")

		tType := tInfo[0]

		switch tType {
		case "DEPOSIT":
			user := tInfo[1]
			value, _ := strconv.Atoi(tInfo[2])
			deposit(db, user, value)
		case "WITHDRAW":
			user := tInfo[1]
			value, _ := strconv.Atoi(tInfo[2])
			withdraw(db, user, value)
		case "TRANSFER":
			userFrom := tInfo[1]
			userTo := tInfo[2]
			value, _ := strconv.Atoi(tInfo[3])
			transfer(db, userFrom, userTo, value)
		case "INCOME":
			percent, _ := strconv.Atoi(tInfo[1])
			income(db, percent)
		case "BALANCE":
			user := tInfo[1]
			balance(db, user)
		}
	}
}
