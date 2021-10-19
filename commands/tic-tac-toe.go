package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var board [3][3]int

var TicTacToe = Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "tic-tac-toe",
		Description: "Play a simple game of tic tac toe",
	},
	Handler: func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

		session.ChannelMessageSend(interaction.ChannelID, printBoard())

	},
}

func printBoard() string {
	boardUI := ""
	for i := range board {
		x := board[i]
		boardUI = boardUI + "-------------\n"
		boardUI = boardUI + fmt.Sprintf("| %v | %v | %v |\n", coordToString(x[0]), coordToString(x[1]), coordToString(x[2]))
		boardUI = boardUI + "-------------\n"
	}

	return boardUI
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
