package main;

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func step(state map[int64]int64) map[int64]int64 {
	retval := make(map[int64]int64);
	adit6 := int64(0);
	adit8 := int64(0);
	for k,v := range state {
		if k == 0 {
			adit8 = v;
			adit6 += v;
			continue
		}
		retval[k-1] = v;
	}

	retval[6] += adit6;
	retval[8] = adit8;

	return retval;
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	scanner.Scan();
	line := scanner.Text();
	nums := strings.Split(line, ",");

	state := make(map[int64]int64);
	for _,v := range nums {
		val, err := strconv.ParseInt(v, 10, 32);
		if err != nil {
			panic(err);
		}

		state[val]++;
	}

	for i := 0; i < 256; i++ {
		state = step(state);
	}

	count := 0;
	for _, v := range state {
		count += int(v);
	}

	fmt.Println(count);
}
