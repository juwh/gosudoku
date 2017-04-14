package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

const max_dim = 9
const box_dim = 3
const filePath = "sudoku-test1.txt"
var board[9][9]int

func readFile (filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowNum := 0
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

func checkRow (row int, num int) bool {
	for i:=0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	return true
}

func checkCol (col int, num int) bool {
	for i:=0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	return true
}

func checkSquare (row int, col int, num int) bool {
	rowStart := row - (row % 3)
	colStart := col - (col % 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if num == board[rowStart+i][colStart+j] {
				return false
			}
		}
	}
	return true
}

func check(row int, col int, num int) bool {
	return checkRow(row,num) && checkCol(col,num) && checkSquare(row,col,num)
}

func main() {
	readFile(filePath)
}