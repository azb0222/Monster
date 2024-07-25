package scanner

import (
	"monster/pkg/concurrencyUtils"
	"runtime"
	"strings"
	"time"
)

// based off pattern in https://go.dev/doc/codewalk/sharemem/

const (
	logInterval = 60 * time.Second
)

var (
	CPUNum = runtime.NumCPU()
)

// MOVE TYPES SOMEWHERE ELSE?
type FastScanner struct {
	Code string
	//does it make sense to have it as a private field?
	codeBlocks map[string]string
	Tokens     []Token
}

func NewFastScanner(code string, tokens []Token) *FastScanner {
	return &FastScanner{
		Code:   code,
		Tokens: tokens,
	}
}

type codeBlock struct {
	code   string
	tokens []Token
	indexer
}

func newCodeBlock(code string) *codeBlock {
	return &codeBlock{
		code:   code,
		tokens: make([]Token, 0),
		indexer: indexer{
			start: 0,
			curr:  0,
			line:  1,
		},
	}
}

type codeBlockState struct {
	codeBlock *codeBlock
	isDone    bool
}

func newCodeBlockState(codeBlock *codeBlock) *codeBlockState {
	return &codeBlockState{
		codeBlock: codeBlock,
		isDone:    false,
	}
}


func (fs *FastScanner) Scan() {
		pending, complete := make(chan *codeBlock), make(chan *codeBlock)
		stateMonitor := concurrencyUtils.StateMonitor(logInterval)

		for i := 0; i < CPUNum; i++ {
			go cbScanner(pending, complete, status)
		}

		go func() {
			for _, cb := range codeBlocks {
				pending <- &CodeBlock{
					code: cb,
					indexer: indexer{
						start: 0,
						curr:  0,
					},
				}
			}
		}()
	}
}

// TODO: rewrite the comments so they are better
// splitCode() splits the code into logical codeBlocks that can be processed by the scanner concurrently based on CPUNum
func (fs *FastScanner) splitCode() {
	lexemes := strings.Fields(fs.Code)

	var codeBlocks []*codeBlock

	for i := 0; i < len(lexemes); i += CPUNum {
		end := i + CPUNum

		if end > len(lexemes) {
			end = len(lexemes)
		}

		code := strings.Join(lexemes[i:end], " ")
		codeBlocks = append(codeBlocks, newCodeBlock(code))
	}
}

// StateMonitor should be generic?
// CodeBlockStateMonitor has a map w/ state of CodeBlock
//func CodeBlockStateMonitor(updateInterval time.Duration) chan<- CodeBlockState {
//	updates := make(chan CodeBlockState)
//	//this map key type is wrong
//	codeBlockStatus := make(map[string]bool)
//
//	ticker := time.NewTicker(updateInterval)
//	go func() {
//		for {
//			select {
//			case <-ticker.C:
//				logState(codeBlockStatus)
//			case s := <-updates: //if receive an update on updates channel
//				codeBlockStatus[s.codeBlock.code] = s.isComplete
//			}
//		}
//	}()
//	return updates
//}

//TODO: instead of the actual codeBlock.code, could do first Token.lexeme + "..."
// TODO: shouldn't print the actual codeBlock.code, should have some sort of index
// TODO: implement logging everywhere else too instead of just prints
//func logState(codeBlockStatus map[string]bool) {
//	log.Println("Scanner Current State:")
//	for k, v := range codeBlockStatus {
//		log.Printf("\t%s: %t\n", k, v)
//	}
//}

//func (cb *CodeBlock) scanTokens() {
//	for !cb.isAtEnd() {
//		cb.start = cb.curr
//		cb.scanToken()
//	}
//	//cb.Tokens = append(cb.Tokens, *newToken(EOF, "", nil, s.line))
//}
//
//func (cb *CodeBlock) scanToken() {
//
//}
//
//func (cb *CodeBlock) isAtEnd() bool {
//	return cb.curr >= len(cb.code)
//}
//
//func cbScanner(in <-chan *CodeBlock, out chan<- *CodeBlock, status chan<- CodeBlockState {
//	//for each token in the readOnly in channel
//	for cb := range in {
//		cb.scanTokens()
//		status <- CodeBlockState{
//			cb,
//			true,
//		}
//		out <- cb
//	}
//}
