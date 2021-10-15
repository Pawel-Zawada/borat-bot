package commands

import "fmt"

var board [3][3]int

func main() {
	printBoard()
	setPlayerOne(0, 0)
	setPlayerOne(1, 1)
	setPlayerOne(2, 2)

	printBoard()
}

func printBoard() {
	for i := range board {
		x := board[i]
		fmt.Println("-------------")
		fmt.Printf("| %v | %v | %v | \n", coordToString(x[0]), coordToString(x[1]), coordToString(x[2]))
		fmt.Println("-------------")
	}
}

func coordToString(i int) string {
	if i == 1 {
		return "X"
	} else if i == -1 {
		return "O"
	}

	return " "
}

func setPlayerOne(x int, y int) {
	if board[y][x] != -1 {
		board[y][x] = 1
		checkWin(x, y, 1)
	} else {
		fmt.Println("Tried to set spot that was already set")
	}
}

func setPlayerTwo(x int, y int) {
	if board[y][x] != 1 {
		board[y][x] = -1
		checkWin(x, y, 1)
	} else {
		fmt.Println("Tried to set spot that was already set")
	}
}

func checkWin(x int, y int, player int) (int, bool) {
	n := 3
	won := false
	// Check column
	for i := 0; i < n; i++ {
		if board[i][x] != player {
			break
		}
		if i == n-1 {
			won = true
		}
	}

	//check row
	for i := 0; i < n; i++ {
		if board[y][i] != player {
			break
		}
		if i == n-1 {
			won = true
		}
	}

	//check diag
	if x == y {
		//we're on a diagonal
		for i := 0; i < n; i++ {
			if board[i][i] != player {
				break
			}
			if i == n-1 {
				won = true
			}
		}
	}

	//check anti diag
	if x+y == n-1 {
		for i := 0; i < n; i++ {
			if board[(n-1)-i][i] != player {
				break
			}
			if i == n-1 {
				won = true
			}
		}
	}
	return player, won
}
