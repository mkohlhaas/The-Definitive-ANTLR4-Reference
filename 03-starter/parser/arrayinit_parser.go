// Code generated from parser/ArrayInit.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ArrayInit

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ArrayInitParser struct {
	*antlr.BaseParser
}

var ArrayInitParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func arrayinitParserInit() {
	staticData := &ArrayInitParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "','", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"init", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 20, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 9, 8,
		0, 10, 0, 12, 0, 12, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 18, 8, 1, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 19, 0, 4, 1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 4, 5, 5,
		1, 0, 0, 5, 10, 3, 2, 1, 0, 6, 7, 5, 2, 0, 0, 7, 9, 3, 2, 1, 0, 8, 6, 1,
		0, 0, 0, 9, 12, 1, 0, 0, 0, 10, 8, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11,
		13, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 13, 14, 5, 3, 0, 0, 14, 1, 1, 0, 0,
		0, 15, 18, 3, 0, 0, 0, 16, 18, 5, 4, 0, 0, 17, 15, 1, 0, 0, 0, 17, 16,
		1, 0, 0, 0, 18, 3, 1, 0, 0, 0, 2, 10, 17,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ArrayInitParserInit initializes any static state used to implement ArrayInitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewArrayInitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ArrayInitParserInit() {
	staticData := &ArrayInitParserStaticData
	staticData.once.Do(arrayinitParserInit)
}

// NewArrayInitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewArrayInitParser(input antlr.TokenStream) *ArrayInitParser {
	ArrayInitParserInit()
	this := new(ArrayInitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ArrayInitParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ArrayInit.g4"

	return this
}

// ArrayInitParser tokens.
const (
	ArrayInitParserEOF  = antlr.TokenEOF
	ArrayInitParserT__0 = 1
	ArrayInitParserT__1 = 2
	ArrayInitParserT__2 = 3
	ArrayInitParserINT  = 4
	ArrayInitParserWS   = 5
)

// ArrayInitParser rules.
const (
	ArrayInitParserRULE_init  = 0
	ArrayInitParserRULE_value = 1
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArrayInitParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArrayInitParserRULE_init
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArrayInitParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *InitContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *ArrayInitParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ArrayInitParserRULE_init)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Match(ArrayInitParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(5)
		p.Value()
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArrayInitParserT__1 {
		{
			p.SetState(6)
			p.Match(ArrayInitParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(7)
			p.Value()
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(13)
		p.Match(ArrayInitParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Init() IInitContext
	INT() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArrayInitParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArrayInitParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArrayInitParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Init() IInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitContext)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ArrayInitParserINT, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ArrayInitParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ArrayInitParserRULE_value)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArrayInitParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Init()
		}

	case ArrayInitParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Match(ArrayInitParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
