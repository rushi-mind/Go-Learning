/*
	Problem: 			Tic-Tac-Toe in GoLang
	Date: 				28-06-2022
	Last Modified:		28-06-2022
	Author:				Rushiraj Parekh
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Globals:
var current []string = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
var available [9]bool
var (
	movesPlayer1 []int
	movesPlayer2 []int
	finished     bool = false
)
var winConditions [][]int = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{1, 4, 7},
	{2, 5, 8},
	{3, 6, 9},
	{1, 5, 9},
	{3, 5, 7},
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func displayGrid() {

	clearConsole()

	fmt.Printf("\n\n\n\n****************************************** TIC TAC TOE ******************************************\n\n\n\n")

	fmt.Printf("\t\t\t\t\t     |     |     \n")
	fmt.Printf("\t\t\t\t\t  %s  |  %s  |  %s  \n", current[0], current[1], current[2])
	fmt.Printf("\t\t\t\t\t_____|_____|_____\n")
	fmt.Printf("\t\t\t\t\t     |     |     \n")
	fmt.Printf("\t\t\t\t\t  %s  |  %s  |  %s  \n", current[3], current[4], current[5])
	fmt.Printf("\t\t\t\t\t_____|_____|_____\n")
	fmt.Printf("\t\t\t\t\t     |     |     \n")
	fmt.Printf("\t\t\t\t\t  %s  |  %s  |  %s  \n", current[6], current[7], current[8])
	fmt.Printf("\t\t\t\t\t     |     |     \n")

	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func checkFinished() int {

	// Check if player 1 wins..!!
	for i := 0; i < 8; i++ {
		flag1 := false
		flag2 := false
		flag3 := false
		for j := 0; j < 3; j++ {
			for k := 0; k < len(movesPlayer1); k++ {
				if winConditions[i][j] == movesPlayer1[k] {
					if j == 0 {
						flag1 = true
					} else if j == 1 {
						flag2 = true
					} else {
						flag3 = true
					}
					break
				}
			}
		}
		if flag1 && flag2 && flag3 {
			return 1
		}
	}

	//  Check if player 2 wins..!!
	for i := 0; i < 8; i++ {
		flag1 := false
		flag2 := false
		flag3 := false
		for j := 0; j < 3; j++ {
			for k := 0; k < len(movesPlayer2); k++ {
				if winConditions[i][j] == movesPlayer2[k] {
					if j == 0 {
						flag1 = true
					} else if j == 1 {
						flag2 = true
					} else {
						flag3 = true
					}
					break
				}
			}
		}
		if flag1 && flag2 && flag3 {
			return 2
		}
	}

	//  Check if match draw..!!
	if len(movesPlayer1)+len(movesPlayer2) == 9 {
		return 9
	}

	return 0
}

func solve() {
	clearConsole()
	fmt.Printf("\n\n\n\n************************************ LET THE GAME BEGIN ************************************\n\n\n\n")
	var (
		player1     string
		player2     string
		currentTurn bool
		result      int
		player      string
	)
	fmt.Print("\t\t\tPlayer-1, Enter your name: ")
	fmt.Scan(&player1)
	fmt.Print("\t\t\tPlayer-2, Enter your name: ")
	fmt.Scan(&player2)

	for {
		displayGrid()
		result = checkFinished()
		if result != 0 {
			break
		}

		if currentTurn {
			player = player2
		} else {
			player = player1
		}

		var inputChoice int
		fmt.Printf("\t%s, Enter any number =>\t", player)
		_, err := fmt.Scan(&inputChoice)
		if err != nil {
			continue
		}
		if inputChoice < 1 || inputChoice > 9 || available[inputChoice-1] {
			continue
		}

		available[inputChoice-1] = true
		if currentTurn {
			current[inputChoice-1] = "O"
			movesPlayer2 = append(movesPlayer2, inputChoice)
			currentTurn = false
		} else {
			current[inputChoice-1] = "X"
			movesPlayer1 = append(movesPlayer1, inputChoice)
			currentTurn = true
		}

	}

	if result == 1 {
		fmt.Printf("\n\n\t###########################  %v won the match ###########################\n\n", player)
		fmt.Printf("\n\n\t#######################  %v - Better luck next time #######################\n\n\n\n\n\n\n", player2)
	} else if result == 2 {
		fmt.Printf("\n\n\t###########################  %v won the match ###########################\n\n", player)
		fmt.Printf("\n\n\t#######################  %v - Better luck next time #######################\n\n\n\n\n\n\n", player1)
	} else {
		fmt.Printf("\n\n\t###########################  Ufff....!!! Match Draw...!!  ###########################\n\n\n\n\n\n\n")
	}
}

func main() {
	solve()
}
