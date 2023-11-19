package main

import "fmt"

var board = make(map[string]string)
var currentPlayer = "X"
var gameState = "ongoing"
var totalMoves = 0

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
		fmt.Println("Player", currentPlayer, "winner!")
	} else {
		fmt.Println("draw!")
	}
}

func initializeBoard() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			board[fmt.Sprintf("%d%d", i, j)] = " "
		}
	}
}

func printBoard() {
	fmt.Println("-------------")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
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

		if board[position] == " " {
			board[position] = currentPlayer
			break
		} else {
			fmt.Println("This cell does not exist or is occupied. Try again.")
		}
	}
}

func checkWinCondition() {
	if totalMoves < 5 {
		return
	}

	var winningMoves = [8][3]string{
		{"11", "12", "13"},
		{"21", "22", "23"},
		{"31", "32", "33"},
		{"11", "21", "31"},
		{"12", "22", "32"},
		{"13", "23", "33"},
		{"11", "22", "33"},
		{"13", "22", "31"},
	}

	for i := 0; i < len(winningMoves); i++ {
		combination := winningMoves[i]
		if board[combination[0]] == currentPlayer && board[combination[1]] == currentPlayer && board[combination[2]] == currentPlayer {
			gameState = "win"
			return
		}
	}

	if totalMoves == 9 {
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
