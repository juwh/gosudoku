package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	const max_dim = 9
	const box_dim = 3
	var board[9][9]int

	file, err := os.Open("sudoku-test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		fmt.Println(row[0] + row[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}