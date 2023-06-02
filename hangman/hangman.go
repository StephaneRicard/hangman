package hangman

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	State        string
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnsLeft    int
	TotalTurn    int
	Indice       int
}

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("word '%s must be at least 3 characters. got=%v ", word, len(word))
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_ "
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
		TotalTurn:    0,
		Indice:       0,
	}

	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	g.TotalTurn++

	switch g.State {
	case "won", "lost":
		return
	}

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else if guess == "INDICE" {
		g.giveIndice(guess)
	} else {
		g.State = "badGuess"
		g.LooseTurn((guess))

		if g.TurnsLeft == 0 {
			g.State = "lost"
		}
	}

}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func (g *Game) LooseTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) giveIndice(guess string) {
	g.Indice++
	if g.Indice < 4 {
		var l string
		found := true
		for found {
			l = PickLetter(g.Letters)
			found = false
			for _, foundLetter := range g.FoundLetters {
				if strings.Contains(foundLetter, l) {
					found = true
					break
				}
			}
		}

		g.RevealLetter(l)
	}
}

func PickLetter(letters []string) string {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(letters))
	return letters[i]
}
