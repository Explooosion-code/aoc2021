package main;

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type boardEntry struct {
	num int
	selected bool
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	scanner.Scan();
	numbersData := scanner.Text();
	numbersArr := strings.Split(numbersData, ",");
	var numbers []int;

	for _,v := range numbersArr {
		x, _ := strconv.ParseInt(v, 10, 32);
		numbers = append(numbers, int(x));
	}

	curBoard := -1;
	var boards [][]boardEntry;
	for scanner.Scan() {
		line := scanner.Text();

		if line == "" {
			curBoard++;
			boards = append(boards, []boardEntry{});
		} else {
			line := strings.ReplaceAll(line, "  ", " ");
			line = strings.TrimSpace(line);
			nums := strings.Split(line, " ");

			for _, v := range nums {
				var x int;
				fmt.Sscanf(v, "%d", &x)
				entry := boardEntry{num: x}
				boards[curBoard] = append(boards[curBoard], entry);
			}
		}
	}

	alreadyInBoard := make(map[int]bool)
	winningBoards := []int{};
	for _,v := range numbers {
		for k, board := range boards {
			for k2, entry := range board {
				if entry.num == v {
					boards[k][k2].selected = true;
				}
				if winCheck(board) && !alreadyInBoard[k] {
					fmt.Println(len(winningBoards))
					alreadyInBoard[k] = true;
					winningBoards = append(winningBoards, score(board, v));
				}
			}
		}
	}

	fmt.Println(winningBoards[len(winningBoards) - 1])
}

func nicePrint(board []boardEntry) {
	for k,v := range board {
		if k %5 == 0 {
			fmt.Println();
		}
		color := "\033[31m";
		if v.selected {
			color = "\033[0m";
		}
		fmt.Print(" ", color, v.num);
	}
	fmt.Println("\n");
}

func score(board []boardEntry, num int) int {
	sum := 0;
	for _, v := range board {
		if !v.selected {
			sum += v.num;
		}
	}
	return sum * num;
}

func winCheck(board []boardEntry) bool {
	for i := 0; i < 25; i += 5 {
		if board[i].selected &&
		board[i + 1].selected &&
		board[i + 2].selected &&
		board[i + 3].selected &&
		board[i + 4].selected {
			return true;
		}
	}

	for i := 0; i < 5; i++ {
		if board[i].selected &&
			board[i + 5].selected &&
			board[i + 10].selected &&
			board[i + 15].selected &&
			board[i + 20].selected {
			return true;
		}
	}

	return false
}