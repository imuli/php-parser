package php7

import (
	"io"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/scanner"
)

func (lval *yySymType) Token(t *scanner.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	*scanner.Lexer
	path            string
	currentTokenID  int
	nextTokenID     int
	currentToken    *scanner.Token
	nextToken       *scanner.Token
	nextLval        *yySymType
	positionBuilder *parser.PositionBuilder
	errors          []*errors.Error
	rootNode        node.Node
	comments        parser.Comments
	positions       parser.Positions
}

// NewParser creates and returns new Parser
func NewParser(src io.Reader, path string) *Parser {
	lexer := scanner.NewLexer(src, path)

	return &Parser{
		lexer,
		path,
		0,
		0,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	}
}

// Lex proxy to lexer Lex
func (l *Parser) Lex(lval *yySymType) int {
	if l.nextLval == nil {
		l.nextLval = &yySymType{}
		l.nextTokenID = l.Lexer.Lex(l.nextLval)
		l.nextToken = l.nextLval.token
	}

	l.currentTokenID = l.nextTokenID
	l.currentToken = l.nextToken

	l.nextTokenID = l.Lexer.Lex(l.nextLval)
	l.nextToken = l.nextLval.token

	lval.token = l.currentToken
	return l.currentTokenID
}

func (l *Parser) Error(msg string) {
	l.errors = append(l.errors, errors.NewError(msg, l.currentToken))
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse() int {
	yyDebug = 0
	yyErrorVerbose = true

	// init
	l.nextLval = nil
	l.errors = nil
	l.rootNode = nil
	l.comments = parser.Comments{}
	l.positions = parser.Positions{}
	l.positionBuilder = &parser.PositionBuilder{
		Positions: &l.positions,
	}

	// parse

	return yyParse(l)
}

func (l *Parser) listGetFirstNodeComments(list []node.Node) []*comment.Comment {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return l.comments[node]
}

// GetPath return path to file
func (l *Parser) GetPath() string {
	return l.path
}

// GetRootNode returns root node
func (l *Parser) GetRootNode() node.Node {
	return l.rootNode
}

// GetErrors returns errors list
func (l *Parser) GetErrors() []*errors.Error {
	return l.errors
}

// GetComments returns comments list
func (l *Parser) GetComments() parser.Comments {
	return l.comments
}

// GetPositions returns positions list
func (l *Parser) GetPositions() parser.Positions {
	return l.positions
}
