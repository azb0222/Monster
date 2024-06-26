package monster

import (
	"fmt"
	"os"
)

var (
	hadError = false
)

func main() {
	if fileName := os.Args[1]; fileName != "" {
		fileExtension := fileName[len(fileName)-4:]
		if fileExtension != ".mon" {
			fmt.Println("Please provide a .mon file")
			os.Exit(1)
		}
		runFile(fileName)
	} else {
		runPrompt()
	}

}

// https://www.scaler.com/topics/golang/golang-read-file/
func runFile(fileName string) {
	sourceBytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error Opening file")
	}
	run(string(sourceBytes))
	//TODO: error handling
}

// REPL
func runPrompt() {
	fmt.Print("WELCOME TO MONSTER! ʘ‿ʘ")
	fmt.Println()

	var line string
	for {
		fmt.Print("> ")
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		run(line)
		//TODO: erorr handline
	}
}

func run(source string) {
	fmt.Println(source)
	//IF HAD ERROR, SYSTEM EXIT
	//TO IMPLEMENT
	/*
		Scanner scanner = new Scanner(source);
		List<Token> tokens = scanner.scanTokens();
		// For now, just print the tokens.
		for (Token token : tokens) {
			System.out.println(token);
		}
	*/
}

// TODO: create ErrorReporter interface?
func error(lineNum int, message string) {
	report(lineNum, "", message)
}

func report(lineNum int, where string, message string) {
	fmt.Printf("[LINE %d] Error %s: %s\n", lineNum, where, message)
	hadError = true
	fmt.Println(hadError)
}
