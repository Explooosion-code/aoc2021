package main;

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var openingBrackets = map[string]bool{
	"(": true,
	"[": true,
	"{": true,
	"<": true,
}

var closingBrackets = map[string]string{
	")":"(",
	"]":"[",
	"}":"{",
	">":"<",
}

var score = map[string]int{
	"(":1,
	"[":2,
	"{":3,
	"<":4,
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);

	answers := []int{};
	for scanner.Scan() {
		line := scanner.Text();

		stack := []string{};
		corrupted := false;
		for _, r := range line {
			s := string(r);
			_, ok := openingBrackets[s];
			if ok {
				stack = append(stack, s);
			} else {
				matching := closingBrackets[s];
				if matching == stack[len(stack) - 1] {
					stack = stack[:len(stack) - 1];
				} else {
					corrupted = true;
					break;
				}
			}
		}

		if corrupted {
			continue;
		}
		ans := 0;
		for i := len(stack) - 1; i >= 0; i-- {
			v := stack[i];
			ans *= 5;
			ans += score[v];
		}

		answers = append(answers, ans);
	}

	sort.Ints(answers);
	fmt.Println(answers[len(answers)/2])
}
