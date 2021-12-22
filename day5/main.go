package main;

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x,y int
}

type pipe struct {
	start, end point
}

func (p pipe) isDiagonal() bool {
	return math.Abs(float64(p.start.x - p.end.x)) == math.Abs(float64(p.start.y - p.end.y));
}

func (p pipe) isHorizontal() bool {
	return p.start.x == p.end.x || p.start.y == p.end.y;
}

func (p pipe) isValid() bool {
	return p.isDiagonal() || p.isHorizontal();
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	var pipes []pipe;

	for scanner.Scan() {
		line := scanner.Text();
		var x1,y1,x2,y2 int;
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1,&y1,&x2,&y2);
		p := pipe{};
		p.start = point{x1,y1};
		p.end = point{x2,y2};

		if p.isValid() {
			pipes = append(pipes, p);
		}
	}

	covered := make(map[point]int);
	for _, pipe := range pipes {
		if pipe.isHorizontal() {
			if pipe.start.y == pipe.end.y {
				sign := 1;
				if pipe.start.x > pipe.end.x {
					sign = -1;
				}
				for x := pipe.start.x; x != pipe.end.x; x += sign {
					covered[point{x, pipe.start.y}]++;
				}
				covered[point{pipe.end.x, pipe.end.y}]++;
			} else {
				sign := 1;
				if pipe.start.y > pipe.end.y {
					sign = -1;
				}
				for y := pipe.start.y; y != pipe.end.y; y += sign {
					covered[point{pipe.start.x, y}]++;
				}
				covered[point{pipe.end.x, pipe.end.y}]++;
			}
			continue;
		}

		signX := 1;
		signY := 1;

		if pipe.start.x > pipe.end.x {
			signX = -1;
		}

		if pipe.start.y > pipe.end.y {
			signY = -1;
		}
		amount := int(math.Abs(float64(pipe.start.x - pipe.end.x)));
		for a := 0; a < amount; a++ {
			fmt.Println(pipe, pipe.start.x + signX * a, pipe.start.y + signY * a);
			covered[point{pipe.start.x + signX * a, pipe.start.y + signY * a}]++;
		}
		covered[point{pipe.end.x, pipe.end.y}]++;
	}
	count := 0;

	for _,v := range covered {
		if v >= 2 {
			count++;
		}
	}

	fmt.Println(count)
}
