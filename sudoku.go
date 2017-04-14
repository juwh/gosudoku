//Name: William Ju and Nicole Kim
//Email: william.h.ju@vanderbilt.edu and nicole.e.kim@vanderbilt.edu
//VUnet ID: juwh and kimne
//Class: CS3270
//Date: 4/13/17
//Sudoku solver in Golang

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

const max_dim = 9 //Constant for max dimensions of puzzle
const box_dim = 3 //Constant for max dimensions of box
const filePath = "sudoku-test1.txt"
var board[max_dim][max_dim]int //2D array that holds board

//Loads values from specified file into board
//Assumes the input file has the specified format (empty spaces are represented by 0's)
func readFile (file string) {
	rowNum := 0
	boardFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer boardFile.Close()

	scanner := bufio.NewScanner(boardFile)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		// index and value
		for i, v := range row {
			board[rowNum][i], err = strconv.Atoi(v)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
		}
		rowNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//Checks if the current assignment of a spot conflicts with any other existing item in its row
func checkRow (row int, num int) bool {
	for i:=0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	return true
}

//Checks if the current assignment of a spot conflicts with any other existing item in its col
func checkCol (col int, num int) bool {
	for i:=0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	return true
}

//Checks if the current assignment of a spot conflicts with any other existing item in its 3x3 square
func checkSquare (row int, col int, num int) bool {
	rowStart := row - (row % box_dim)
	colStart := col - (col % box_dim)

	for i := 0; i < box_dim; i++ {
		for j := 0; j < box_dim; j++ {
			if num == board[rowStart+i][colStart+j] {
				return false
			}
		}
	}
	return true
}

//Combines all 3 check (row, column, square) to make sure there are no conflicts
func check(row int, col int, num int) bool {
	return checkRow(row,num) && checkCol(col,num) && checkSquare(row,col,num)
}

//The entry point for the solver. Returns true if a solution was found, otherwise returns false.
//Parameters are the row and column that a value is currently being placed in
func solve(row int, col int) bool {
	if row == max_dim {
		return true
	}
	if col == max_dim {
		return solve(row+1,0)
	}
	if board[row][col] != 0 {
		return solve(row,col+1)
	}
	for num := 1; num <= max_dim; num++ {
		if check(row, col, num) {
			board[row][col] = num
			if solve(row,col+1) {
				return true
			}
		}
	}
	board[row][col] = 0
	return false
}

//Prints the board to the screen in a nicely formatted manner
func printBoard (board[9][9]int) {
	var result= ""
	for i := 0; i < max_dim; i++ {
		if i%box_dim == 0 && i != 0 {
			result = result + "------+-------+------\n"
		}
		for j := 0; j < max_dim; j++ {
			if j%box_dim == 0 && j != 0 {
				result = result + "| "
			}
			stringNum := strconv.Itoa(board[i][j])
			result = result + stringNum + " "
		}
		result = result + "\n"
	}
	fmt.Println(result)
}

func main() {
	readFile(filePath)
	fmt.Println("Here is the initial board:")
	printBoard(board)
	if solve(0,0) {
		fmt.Println("Here is the solved board:")
		printBoard(board)
	} else {
		fmt.Println("No solution.")
	}
}