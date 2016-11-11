package sql

import (
	"bufio"
	"io"
	"bytes"
	"strings"
	"fmt"
)

type SelectStatement struct {
Fields []string
TableName string
}
type GoifStatement struct {
	Fields []string
}
// Token represents a lexical token.
type Token int
//adding
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // fields, table_name

	// Misc characters
	ASTERISK // *
	COMMA    // ,
	ParenDer // (
	ParenIsq // )
	DosPunt // :


	// Keywords
	SELECT
	FROM
	Goif
)
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }
var eof = rune(0)
// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }
// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case '(':
		return ParenDer, string(ch)
	case ')':
		return ParenIsq, string(ch)
	case ':':
		return DosPunt, string(ch)
	}

	return ILLEGAL, string(ch)
}
// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}
// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "FROM":
		return FROM, buf.String()
	case "SELECT":
		return SELECT, buf.String()
	case "Goif":
		return SELECT, buf.String()

	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}
// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		    tok Token  // last read token
		    lit string // last read literal
		    n   int    // buffer size (max=1)
	    }
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}
// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}
func (p *Parser) Parse() (*SelectStatement, error) {

	stmt := &SelectStatement{}

	if tok, lit := p.scanIgnoreWhitespace(); tok != FROM{
		return nil, fmt.Errorf("found %q, expected SELECT", lit)
	}
	for {
		// Read a field.
		tok, lit := p.scanIgnoreWhitespace()
		if tok != IDENT && tok != ASTERISK {
			return nil, fmt.Errorf("found %q, expected field", lit)
		}
		stmt.Fields = append(stmt.Fields, lit)

		// If the next token is not a comma then break the loop.
		if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
			p.unscan()
			break
		}
	}
	// Next we should see the "FROM" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit
	return stmt, nil
}
func (p *Parser) Parses() (*GoifStatement, error) {

	stmts := &GoifStatement{}

	if tok, lit := p.scanIgnoreWhitespace(); tok != Goif{
		return nil, fmt.Errorf("found %q, expected Goif", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt{
		return nil, fmt.Errorf("found %q, expected :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParenDer{
		return nil, fmt.Errorf("found %q, expected (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParenIsq{
		return nil, fmt.Errorf("found %q, expected (", lit)
	}

	return stmts, nil
	}
