package scanner

import (
	"runtime"
	"time"
	"monster/pkg/concurrencyUtils"
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
	Code   string
	//does it make sense to have it as a private field?
	codeBlocks map[string]string
	Tokens []Token
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

func (fs *FastScanner) Scan() {
		pending, complete := make(chan *codeBlock), make(chan *codeBlock)
		stateMonitor := concurrencyUtils.StateMonitor(logInterval)



		//TODO: this runtime.NumCPU() should be in one variable
		for i := 0; i < runtime.NumCPU(); i++ {
			go cbScanner(pending, complete, status)
		}

		//TODO: replace with actual generation function
		codeBlocks := []string{
			"print 'hello'",
			"var x = 10",
			"test 23 = 12",
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

func (fs *FastScanner) splitCode() {

	testCode := "asdf asdflkj lejl lkjwe wlkj xnm ljkl wljel lkjsd wlkejl dkljf"

	for

}