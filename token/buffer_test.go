package token

import (
	"io"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestRead(t *testing.T) {
	tokens := []Token{
		{
			Class: Number,
			Value: "2",
		},
		{
			Class: Add,
			Value: "+",
		},
		{
			Class: Number,
			Value: "2",
		},
	}
	buf := NewBuffer(tokens)
	for i := 0; i < len(tokens); i++ {
		token, err := buf.Read()
		if err != nil {
			t.Error(err)
		}
		if token != tokens[i] {
			t.Errorf(
				"Expected: %s\n  but got: %s",
				spew.Sdump(tokens[i]),
				spew.Sdump(token),
			)
		}
	}
	_, err := buf.Read()
	if err != io.EOF {
		t.Errorf("Expected io.EOF but got: %v", err)
	}
}

func TestSeek(t *testing.T) {
	tokens := []Token{
		{
			Class: Number,
			Value: "2",
		},
	}
	buf := NewBuffer(tokens)
	_, err := buf.Read()
	if err != nil {
		t.Error(err)
	}
	if err := buf.Seek(0); err != nil {
		t.Error(err)
	}
	token, err := buf.Read()
	if err != nil {
		t.Error(err)
	}
	if token != tokens[0] {
		t.Errorf(
			"Expected: %s\n  but got: %s",
			spew.Sdump(tokens[0]),
			spew.Sdump(token),
		)
	}
}

func TestPos(t *testing.T) {
	tokens := []Token{
		{
			Class: Number,
			Value: "2",
		},
		{
			Class: Add,
			Value: "+",
		},
		{
			Class: Number,
			Value: "2",
		},
	}
	buf := NewBuffer(tokens)
	if pos := buf.Pos(); pos != 0 {
		t.Errorf("Expected pos to be 0 but got: %d", pos)
	}
	if _, err := buf.Read(); err != nil {
		t.Error(err)
	}
	if pos := buf.Pos(); pos != 1 {
		t.Errorf("Expected pos to be 1 but got: %d", pos)
	}
	if err := buf.Seek(2); err != nil {
		t.Error(err)
	}
	if pos := buf.Pos(); pos != 2 {
		t.Errorf("Expected pos to be 2 but got: %d", pos)
	}
	if err := buf.Seek(-1); err == nil ||
		!strings.Contains(err.Error(), "out of range") {
		t.Errorf("Expected out of rage error but got: %v", err)
	}
	if err := buf.Seek(len(tokens)); err == nil ||
		!strings.Contains(err.Error(), "out of range") {
		t.Errorf("Expected out of rage error but got: %v", err)
	}
}
