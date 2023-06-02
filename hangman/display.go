package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(` __    __   ______   __    __   ______   __       __   ______ 
	/  |  /  | /      \ /  \  /  | /      \ /  \     /  | /      \ /  \  /  |
	$$ |  $$ |/$$$$$$  |$$  \ $$ |/$$$$$$  |$$  \   /$$ |/$$$$$$  |$$  \ $$ |
	$$ |__$$ |$$ |__$$ |$$$  \$$ |$$ | _$$/ $$$  \ /$$$ |$$ |__$$ |$$$  \$$ |
	$$    $$ |$$    $$ |$$$$  $$ |$$ |/    |$$$$  /$$$$ |$$    $$ |$$$$  $$ |
	$$$$$$$$ |$$$$$$$$ |$$ $$ $$ |$$ |$$$$ |$$ $$ $$/$$ |$$$$$$$$ |$$ $$ $$ |
	$$ |  $$ |$$ |  $$ |$$ |$$$$ |$$ \__$$ |$$ |$$$/ $$ |$$ |  $$ |$$ |$$$$ |
	$$ |  $$ |$$ |  $$ |$$ | $$$ |$$    $$/ $$ | $/  $$ |$$ |  $$ |$$ | $$$ |
	$$/   $$/ $$/   $$/ $$/   $$/  $$$$$$/  $$/      $$/ $$/   $$/ $$/   $$/ `)

}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	showTurnLeft(g.TurnsLeft)
	showTotalTurn(g.TotalTurn)
	showIndice(g.Indice)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Deviné: ")
	drawLetters((g.FoundLetters))

	fmt.Print("utilisé: ")
	drawLetters((g.UsedLetters))

	fmt.Print("nombre d'essai restant : ")

	switch g.State {
	case "goodGuess":
		fmt.Println("Bien joué!")
	case "alreadyGuessed":
		fmt.Printf("La lettre '%s' a déjà été utilisée\n", guess)
	case "badGuess":
		fmt.Printf("Hé non !, '%s' n'est pas dans le mot\n", guess)
	case "lost":
		fmt.Print("Tu as perdu :( ! Le mot était: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Println("Tu as gagné ! Félicitation ! Le mot était: ")
		drawLetters(g.Letters)

	}

}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}

func showTurnLeft(tl int) {
	fmt.Printf("Il vous reste %v essais - tape INDICE pour revelé une lettre (3 max)\n", tl)
}

func showTotalTurn(tt int) {
	fmt.Printf("Tu as joué %v fois\n", tt)
}

func showIndice(i int) {
	if i > 3 {
		fmt.Println("Tu as déjà utilisé tous tes indices !")
	} else {
		r := 3 - i
		fmt.Printf("Il te reste %d indices\n", r)
	}
}
