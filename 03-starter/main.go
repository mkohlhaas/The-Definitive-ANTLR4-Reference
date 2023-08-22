package main

import (
	"fmt"
	"os"
	"strconv"

	"starter/parser"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseArrayInitListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterInit(ctx *parser.InitContext) {
	fmt.Print("\"")
}

func (this *TreeShapeListener) ExitInit(ctx *parser.InitContext) {
	fmt.Print("\"")
}

func (this *TreeShapeListener) EnterValue(ctx *parser.ValueContext) {
	strValue := ctx.GetText()
	intValue, _ := strconv.Atoi(strValue)
	fmt.Printf("\\u%04x", intValue)
}

// func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Print("Entering EveryRule: ")
// 	fmt.Println(ctx.GetText())
// }

func main() {
	if len(os.Args) > 1 {
		input, _ := antlr.NewFileStream(os.Args[1])
		lexer := parser.NewArrayInitLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewArrayInitParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		tree := p.Init()
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
		fmt.Println()
	} else {
		fmt.Println("Please provide a file name.")
	}
}
