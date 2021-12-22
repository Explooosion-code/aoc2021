package main;

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type screen struct {
	input, output []string
}

func containsCharacters(a, b string) bool {
	for _, v := range b {
		if !strings.ContainsAny(a, string(v)) {
			return false;
		}
	}

	return true;
}

func sortString(input string) string {
	chars := strings.Split(input, "");
	sort.Strings(chars);
	input = strings.Join(chars, "");
	return input;
}

func join(a, b string) string {
	joined := a + b;
	retval := "";

	dups := make(map[rune]bool)
	for _,v := range joined {
		if _, ok := dups[v]; !ok {
			retval += string(v);
			dups[v] = true;
		}
	}

	retval = sortString(retval)

	return retval;
}

func (s screen) getoutput() int {
	values := map[string]int{};
	valuesByInt := map[int]string{}
	all := append(s.input, s.output...);

	for k, v := range all {
		all[k] = sortString(v);
	}

	for _, v := range all {
		if len(v) == 2 {
			values[v] = 1;
			valuesByInt[1] = v;
		} else if len(v) == 4 {
			values[v] = 4;
			valuesByInt[4] = v;
		} else if len(v) == 3 {
			values[v] = 7;
			valuesByInt[7] = v;
		} else if len(v) == 7 {
			values[v] = 8;
			valuesByInt[8] = v;
		}
	}

	for _, v := range all {
		if len(v) == 6 && !containsCharacters(v, valuesByInt[1]) {
			values[v] = 6;
			valuesByInt[6] = v;
		}
	}

	for _, v := range all {
		if len(v) == 5 && containsCharacters(valuesByInt[6], v) {
			values[v] = 5;
			valuesByInt[5] = v;
		}
	}

	nine := join(valuesByInt[5], valuesByInt[1]);
	values[nine] = 9;
	valuesByInt[9] = nine;

	for _, v := range all {
		if len(v) == 7 && containsCharacters(v, valuesByInt[9]) {
			values[v] = 8;
			valuesByInt[8] = v;
		}
	}

	for _, v := range all {
		if len(v) == 6 && v != valuesByInt[6] && v != valuesByInt[9] {
			values[v] = 0;
			valuesByInt[0] = v;
		}
	}

	for _, v := range all {
		if len(v) == 5 && v != valuesByInt[5] {
			if containsCharacters(valuesByInt[9], v) {
				values[v] = 3;
				valuesByInt[3] = v;
			} else {
				values[v] = 2;
				valuesByInt[2] = v;
			}
		}
	}

	retval := 0;
	for _,v := range s.output {
		retval *= 10;
		retval += values[sortString(v)];
	}

	return retval;
}

func main() {
	f, err := os.Open("input");
	if err != nil {
		panic(err);
	}

	scanner := bufio.NewScanner(f);
	screens := []screen{};
	for scanner.Scan() {
		line := scanner.Text();
		data := strings.Split(line, " | ");
		in, out := data[0], data[1];

		scr := screen{strings.Split(in, " "), strings.Split(out, " ")};
		screens = append(screens, scr);
	}

	sum := 0;
	for _, v := range screens {
		output := v.getoutput();
		fmt.Println(output)
		sum += output;
	}

	fmt.Println(sum);
}
