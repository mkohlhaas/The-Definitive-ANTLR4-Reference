package main

import (
	"fmt"
	"os"

	"hello/parser"
	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseHelloListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
  fmt.Println("Entering Every Rule")
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewHelloLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewHelloParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
  tree := p.R()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
