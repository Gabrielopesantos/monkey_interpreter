package lexer

import (
	"fmt"
	"github.com/gabrielopesantos/monkey_interpreter/token"
	"unicode"
)

type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch rune // current char under examination
}
 func New(input string) *Lexer {

 	l := &Lexer{
 		input: input,
 	}

 	l.readChar()

 	return l
 }

 func (l *Lexer) readChar() {

 	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}

	l.position = l.readPosition
	l.readPosition += 1
 }

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
				tok = newToken(token.BANG, l.ch)
			}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.getIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok

		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok

		} else {
			fmt.Printf("%q\n %v", l.ch, l.position)
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) getIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return  l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}
func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return rune(l.input[l.readPosition])
	}
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{
		Type: tokenType,
		Literal: string(ch),
	}
}

func isLetter(ch rune) bool {
	//return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
	return unicode.IsLetter(ch) || unicode.IsSymbol(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
