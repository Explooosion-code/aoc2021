package main;

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumFuel(heights []int, height int) int64 {
	sum := int64(0);
	for _,v := range heights {
		dist := int64(math.Abs(float64(v - height)));
		fuel := fuelForDist(dist);
		sum += fuel;
	}

	return sum
}

func fuelForDist(dist int64) int64 {
	if dist == 0 {
		return 0;
	}

	if dist == 1 {
		return 1;
	}
	return dist + fuelForDist(dist - 1);
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	heights := []int{};
	scanner := bufio.NewScanner(f);
	scanner.Scan();
	line := scanner.Text();
	nums := strings.Split(line, ",");

	for _,v := range nums {
		val, err := strconv.ParseInt(v, 10, 32);
		if err != nil {
			panic(err);
		}

		heights = append(heights, int(val));
	}

	sort.Ints(heights);
	smallest := heights[0];
	highest := heights[len(heights) - 1];

	smallestFuel := int64(1000000000000000);
	for i := smallest; i <= highest; i ++ {
		fuel := sumFuel(heights, i);

		if fuel < smallestFuel {
			smallestFuel = fuel;
		}
	}
	fmt.Println(smallestFuel)
}
