package main;

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mostCommon(read string) string {
	bit := 0;
	for _, x := range read {
		if x == '1' {
			bit ++;
		}
	}

	if float64(bit) >= float64(len(read))/2.0 {
		return "1";
	} else {
		return "0";
	}
}

func oxygen(flat []string, iter int) []string {
	if len(flat) == 1 {
		return flat;
	}

	x := ""
	for _,v := range flat {
		x+= string(v[iter % len(v)]);
	}

	common := mostCommon(x);
	filtered := []string{};

	for _,v := range flat {
		if string(v[iter % len(v)]) == common {
			filtered = append(filtered, v);
		}
	}
	return oxygen(filtered, iter + 1);
}


func co2(flat []string, iter int) []string {
	if len(flat) == 1 {
		return flat;
	}

	x := ""
	for _,v := range flat {
		x+= string(v[iter % len(v)]);
	}

	common := mostCommon(x);
	if common == "1" {common = "0"} else {common = "1"};
	filtered := []string{};

	for _,v := range flat {
		if string(v[iter % len(v)]) == common {
			filtered = append(filtered, v);
		}
	}
	return co2(filtered, iter + 1);
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	flat := []string{}
	buf := []string{}
	for scanner.Scan() {
		line := scanner.Text();
		flat = append(flat, line);
		for pos,char := range line {
			if len(buf) - 1 < pos {
				buf = append(buf, "");
			}
			buf[pos] += string(char);
		}
	}

	dec1, err := strconv.ParseInt(oxygen(flat, 0)[0], 2, 32);
	dec2, err := strconv.ParseInt(co2(flat, 0)[0], 2, 32);

	fmt.Println(dec1*dec2);
}