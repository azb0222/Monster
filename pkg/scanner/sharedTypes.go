package scanner

type indexer struct {
	start int //first character of lexeme being currently scanned
	curr  int //unconsumed char being currently considered
	line  int //current line
}
