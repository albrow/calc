package token

import (
	"fmt"
	"io"
)

type Buffer struct {
	tokens []Token
	pos    int
}

func NewBuffer(tokens []Token) *Buffer {
	return &Buffer{
		tokens: tokens,
		pos:    0,
	}
}

func (b Buffer) Len() int {
	return len(b.tokens)
}

func (b Buffer) Pos() int {
	return b.pos
}

func (b *Buffer) Seek(pos int) error {
	if pos < 0 || pos > len(b.tokens)-1 {
		return fmt.Errorf("Seek out of range: %d", pos)
	}
	b.pos = pos
	return nil
}

func (b *Buffer) MustSeek(pos int) {
	if err := b.Seek(pos); err != nil {
		panic(err)
	}
}

func (b *Buffer) Read() (Token, error) {
	if b.pos > len(b.tokens)-1 {
		return Token{}, io.EOF
	}
	token := b.tokens[b.pos]
	b.pos++
	return token, nil
}
