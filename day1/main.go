package main;

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	ans := 0;
	last := -1;

	hm := []int{};

	for scanner.Scan() {
		line := scanner.Text();
		var depth int;
		fmt.Sscanf(line, "%d", &depth);

		hm = append(hm, depth);
		if len(hm) >= 3 {
			sum := 0;
			for _,v := range(hm) {
				sum+=v;
			}

			fmt.Println(sum)
			if last != -1  && sum > last {
				ans++;
			}

			last = sum;

			hm = []int{hm[1], hm[2]};
		}
	}

	fmt.Printf("%d", ans);
}
