package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("Qu'elle est ta lettre ? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 && guess != "INDICE" {
			fmt.Printf("lettre non valide. lettre =%v, len=%v\n", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}
