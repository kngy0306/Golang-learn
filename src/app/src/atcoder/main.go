package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
	stringSplited := strings.Split(stringTargeted, delim)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}

	return
}

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func SplitStrStdin(delim string) (stringReturned []string) {
	// 文字列で1行取得
	stringInput := StrStdin()

	// 空白で分割
	// stringSplited := strings.Split(stringInput, delim)
	stringReturned = SplitWithoutEmpty(stringInput, delim)

	return
}

func IntStdin() (int, error) {
	stringInput := StrStdin()
	return strconv.Atoi(strings.TrimSpace(stringInput))
}

func main() {
	var words []string

	len, _ := IntStdin()
	for idx := 0; idx < len; idx += 1 {
		wordlist := SplitStrStdin(" ")
		words = append(words, strings.Join(wordlist, ""))
	}
	for _, word := range words {
		fmt.Printf("%s\n", word)
	}
}
