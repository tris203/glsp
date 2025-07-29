package protocol

import (
	"strings"
	"unicode/utf8"
)

func (r Range) IndexesIn(content string) (int, int) {
	return r.Start.IndexIn(content), r.End.IndexIn(content)
}

func (p Position) IndexIn(content string) int {
	// This code is modified from the gopls implementation found:
	// https://cs.opensource.google/go/x/tools/+/refs/tags/v0.1.5:internal/span/utf16.go;l=70

	// In accordance with the LSP Spec:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocuments
	// self.Character represents utf-16 code units, not bytes and so we need to
	// convert utf-16 code units to a byte offset.

	// Find the byte offset for the line
	index := 0
	for row := uint32(0); row < p.Line; row++ {
		content_ := content[index:]
		if next := strings.Index(content_, "\n"); next != -1 {
			index += next + 1
		} else {
			return 0
		}
	}

	// The index represents the byte offset from the beginning of the line
	// count self.Character utf-16 code units from the index byte offset.

	byteOffset := index
	remains := content[index:]
	chr := int(p.Character)

	for count := 1; count <= chr; count++ {

		if len(remains) <= 0 {
			// char goes past content
			// this a error
			return 0
		}

		r, w := utf8.DecodeRuneInString(remains)
		if r == '\n' {
			// Per the LSP spec:
			//
			// > If the character value is greater than the line length it
			// > defaults back to the line length.
			break
		}

		remains = remains[w:]
		if r >= 0x10000 {
			// a two point rune
			count++
			// if we finished in a two point rune, do not advance past the first
			if count > chr {
				break
			}
		}
		byteOffset += w

	}

	return byteOffset
}

func (p Position) EndOfLineIn(content string) Position {
	index := p.IndexIn(content)
	content_ := content[index:]
	if eol := strings.Index(content_, "\n"); eol != -1 {
		return Position{
			Line:      p.Line,
			Character: p.Character + uint32(eol),
		}
	} else {
		return p
	}
}
