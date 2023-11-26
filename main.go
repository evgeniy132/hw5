package main

import (
	"fmt"
	"strconv"
)

var board = make(map[string]string)
var currentPlayer = "X"
var gameState = "ongoing"
var totalMoves = 0
var boardSize = 4

func main() {
	initializeBoard()
	for gameState == "ongoing" {
		printBoard()
		makeMove()
		checkWinCondition()
		switchPlayers()
		totalMoves++
	}
	printBoard()
	if gameState == "win" {
		fmt.Println("Player", currentPlayer, "is the winner!")
	} else {
		fmt.Println("It's a draw!")
	}
}

func initializeBoard() {
	for i := 1; i <= boardSize; i++ {
		for j := 1; j <= boardSize; j++ {
			board[fmt.Sprintf("%d%d", i, j)] = " "
		}
	}
}

func printBoard() {
	fmt.Println("-------------")
	for i := 1; i <= boardSize; i++ {
		for j := 1; j <= boardSize; j++ {
			fmt.Printf("| %s ", board[fmt.Sprintf("%d%d", i, j)])
		}
		fmt.Println("|\n-------------")
	}
}

func makeMove() {
	var position string
	for {
		fmt.Print("Player ", currentPlayer, ", enter position (for example, 21): ")
		fmt.Scanln(&position)

		row, _ := strconv.Atoi(string(position[0]))
		col, _ := strconv.Atoi(string(position[1]))

		if row > 0 && row <= boardSize && col > 0 && col <= boardSize && board[position] == " " {
			board[position] = currentPlayer
			break
		} else {
			fmt.Println("This cell does not exist or is occupied. Try again.")
		}
	}
}

func checkWinCondition() {
	if totalMoves < (2*boardSize)-1 {
		return
	}

	// Rows
	for i := 1; i <= boardSize; i++ {
		win := true
		for j := 1; j <= boardSize; j++ {
			if board[fmt.Sprintf("%d%d", i, j)] != currentPlayer {
				win = false
				break
			}
		}
		if win {
			gameState = "win"
			return
		}
	}

	// Columns
	for j := 1; j <= boardSize; j++ {
		win := true
		for i := 1; i <= boardSize; i++ {
			if board[fmt.Sprintf("%d%d", i, j)] != currentPlayer {
				win = false
				break
			}
		}
		if win {
			gameState = "win"
			return
		}
	}

	// Diagonals
	win := true
	for i := 1; i <= boardSize; i++ {
		if board[fmt.Sprintf("%d%d", i, i)] != currentPlayer {
			win = false
			break
		}
	}
	if win {
		gameState = "win"
		return
	}

	win = true
	for i := 1; i <= boardSize; i++ {
		if board[fmt.Sprintf("%d%d", i, boardSize-i+1)] != currentPlayer {
			win = false
			break
		}
	}
	if win {
		gameState = "win"
		return
	}

	if totalMoves == (boardSize * boardSize) {
		gameState = "draw"
	}
}

func switchPlayers() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}
