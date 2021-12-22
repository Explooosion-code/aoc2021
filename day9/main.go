package main;

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type basin struct {
	id,size int
}

type bySize []basin;

func (b bySize) Len() int {
	return len(b);
}

func (s bySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s bySize) Less(i, j int) bool {
	return s[i].size < s[j].size;
}

type point struct {
	row, col, hei int
	bas *basin
}

func getAdjacent(grid [][]point, pt point) []point {
	retval := []point{}
	if pt.row != 0 {
		retval = append(retval, grid[pt.row - 1][pt.col])
	}
	if pt.row != len(grid) -1 {
		retval = append(retval, grid[pt.row + 1][pt.col])
	}
	if pt.col != 0 {
		retval = append(retval, grid[pt.row][pt.col - 1]);
	}
	if pt.col != len(grid[pt.row]) - 1 {
		retval = append(retval, grid[pt.row][pt.col + 1]);
	}

	return retval;
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	grid := [][]point{};

	row := -1;
	for scanner.Scan() {
		line := scanner.Text();
		row++;
		grid = append(grid, []point{})
		digits := strings.Split(line, "");
		for col,v := range digits {
			v, _ := strconv.ParseInt(v, 10, 32)
			grid[row] = append(grid[row], point{row,col,int(v), nil});
		}
	}

	lowPoints := make(map[point]bool);
	for _, col := range grid {
		for _, val := range col {
			adj := getAdjacent(grid, val);

			ok := true;
			for _, v := range adj {
				if v.hei <= val.hei {
					ok = false;
				}
			}

			if ok {
				lowPoints[val] = true;
			}
		}
	}

	var currBasinId = -1;
	basins := []basin{};
	for point, _ := range lowPoints {
		point = grid[point.row][point.col];
		if point.bas != nil {
			continue;
		}
		currBasinId++;
		basins = append(basins, basin{currBasinId, 0});
		growBasin(&grid, point, &basins[len(basins) - 1]);
	}

	drawGrid(grid)

	for _, col := range grid {
		for _, val := range col {
			if val.bas == nil {
				continue;
			}

			basins[val.bas.id].size++;
		}
	}

	sort.Sort(bySize(basins))
	maxIndex := len(basins) - 1;
	fmt.Println(basins[maxIndex].size * basins[maxIndex - 1].size * basins[maxIndex - 2].size);
}

var colors = []string{
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
	"\033[35m",
	"\033[36m",
}
func drawGrid(grid [][]point)  {
	for _, col := range grid {
		for _, val := range col {
			color := "\033[0m";
			if val.bas != nil {
				color = colors[val.bas.id % 5];
			}
			fmt.Print(color, val.hei);
		}
		fmt.Println("\033[0m")
	}
}

func growBasin(grid *[][]point, pt point, currBasin *basin)  {
	(*grid)[pt.row][pt.col].bas = currBasin
	adjacent := getAdjacent(*grid, pt);
	for _, v := range adjacent {
		if v.bas == nil && v.hei != 9 {
			growBasin(grid, v, currBasin);
		}
	}
}
