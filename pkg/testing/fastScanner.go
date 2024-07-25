package testing

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var codes = []string{
	"print(hello)",
	"var x = 232",
	"lalala",
}

type Token struct {
	literal string
}

type CodeBlock struct {
	code string
}

var tokens []Token
var tempTokens [][]Token

func main() {
	var wg sync.WaitGroup

	//input and output channels
	pending, complete := make(chan *CodeBlock), make(chan *CodeBlock)

	//launch number of goroutines calling Scanner based on number of CPUs
	CPUNum := runtime.NumCPU()
	for i := 0; i < CPUNum; i++ {
		wg.Add(1)
		go Scanner(pending, complete, &wg)
	}
	/*
	   some of the Scanner()'s might finish before the others
	   before combining all the returned token slices together, I must wait for all of them to finish execution
	*/
	go func() {
		for _, code := range codes {
			pending <- &CodeBlock{code: code}
		}
	}()

	go func() {
		wg.Wait()
	}()

	for _, ts := range tempTokens {
		tokens = append(tokens, ts...)
	}

	fmt.Println(tokens)
}

// in: read only channel, out: write only chan
func Scanner(in <-chan *CodeBlock, out chan<- *CodeBlock, wg *sync.WaitGroup) {
	defer wg.Done()
	for codeBlock := range in { //for each codeBlock in input channel
		codeBlock.Scan()
		out <- codeBlock
	}
}

func (c CodeBlock) Scan() {
	lexemes := strings.Fields(c.code)
	var t []Token
	for _, lexeme := range lexemes {
		token := Token{lexeme}
		t = append(t, token)
	}
	tempTokens = append(tempTokens, t)
}
