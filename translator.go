package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bregydoc/gtranslate"
)

func translator() string {

	for {
		red := "\033[31m"
		yellow := "\033[33m"
		reset := "\033[0m"
		green := "\033[32m"
		brightest_blue := "\033[36m"
	start:
		fmt.Print(red+"Special Notice: Press Enter If You Want To Exit Off The Program\n\n\n", reset)
		reader := bufio.NewReader(os.Stdin)

		fmt.Println(green+"Input The Word You Want To Translate:", reset)
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)
		if word == "" {
			break
		}

		var ask string

		fmt.Print(yellow+"Which Language Do You Want It To Be Translated To?\n", reset)
		fmt.Scanln(&ask)

		result, err := gtranslate.TranslateWithParams(
			word,
			gtranslate.TranslationParams{
				From: "auto",
				To:   ask,
			},
		)
		if err != nil {
			return fmt.Sprintf(red+"Translation Failed: %v\n", reset, err)

		}
		fmt.Println(brightest_blue+result, reset)
		goto start
	}
	return "You Have Been Exited From The Program"
}

func main() {
	fmt.Println("\033[33m"+"result:", "\033[0m", translator())
}
