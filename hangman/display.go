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
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		____
		|	|
		|	o
		|  /|\
		|	|
		|  / \
	   _|_
	  |   |`
	case 1:
		draw = `
		____
		|	|
		|	o
		|  /|\
		|	|
		|   
	   _|_
	  |   |`
	case 2:
		draw = `
		____
		|	|
		|	o
		|  	|
		|	|
		|  
	   _|_
	  |   |`
	case 3:
		draw = `
	  ____
	  |   |
	  |	  o
	  |  
	  |	
	  |  
     _|_
	|   |`
	case 4:
		draw = `
	  ____
	  |	  |
	  |	
	  | 
	  |
	  |  
	 _|_
	|   |`
	case 5:
		draw = `
		____
		|	
		|	
		|  
		|	
		|  
	   _|_
	  |   |`
	case 6:
		draw = `
		|	
		|	
		|  
		|	
		|  
	   _|_
	  |   |
		`
	case 7:
		draw = `

	   _|_
	  |   |`
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

	switch g.State {
	case "goodGuess":
		fmt.Println("Bien joué!")
	case "alreadyGuessed":
		fmt.Printf("La lettre '%s' a déjà été utilisée\n", guess)
	case "badGuess":
		fmt.Printf("Hé non !, '%s' n'est pas dans le mot\n", guess)
	case "lost":
		fmt.Println("Tu as perdu :( ! Le mot était:")
		drawLetters(g.Letters)
	case "won":
		fmt.Println("Tu as gagné ! Félicitation ! Le mot était:")
		drawLetters(g.Letters)

	}

}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}
