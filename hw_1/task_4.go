package main

import "fmt"

func canPutQueenOnBoard(board [8][8]int, posX int, posY int) bool {
	// y = x + b1
	b1 := posY - posX

	// y = -x + b2
	b2 := posY + posX

	// i - ось х
	// j - ось y
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i == posX) || (j == posY) || (j == i+b1) || (j == -i+b2) {
				if board[i][j] == 1 {
					return false
				}
			}
		}
	}

	return true
}

func main() {

	board := [8][8]int{}

	var posX, posY int

	fmt.Scanf("%d %d", &posX, &posY)

	board[posX-1][posY-1] = 1

	for i := 0; i < 7; i++ {
		fmt.Scanf("%d %d", &posX, &posY)

		if canPutQueenOnBoard(board, posX-1, posY-1) {
			board[posX-1][posY-1] = 1
		} else {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}
