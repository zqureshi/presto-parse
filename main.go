package main

import (
	"fmt"
	"github.com/zqureshi/presto-parse/presto"
	"os"
)

func main() {
	scanner := presto.NewScanner(os.Stdin)

loop:
	for {
		tok, lit := scanner.Scan()
		switch tok {
		case presto.EOF:
			break loop
		case presto.ILLEGAL:
			fmt.Println("ILLEGAL", lit)
		case presto.WS:
			fmt.Println("WS")
		case presto.IDENT:
			fmt.Println("IDENT", lit)
		case presto.ASTERISK:
			fmt.Println("ASTERISK")
		case presto.COMMA:
			fmt.Println("COMMA")
		case presto.DOT:
			fmt.Println("DOT")
		case presto.SELECT:
			fmt.Println("SELECT")
		case presto.FROM:
			fmt.Println("FROM")
		}
	}
}
