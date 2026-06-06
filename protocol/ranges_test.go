package protocol

import "testing"

func TestPositionIndexInASCII(t *testing.T) {
	content := "alpha\nbeta\ngamma"
	position := Position{Line: 1, Character: 2}

	if got := position.IndexIn(content); got != len("alpha\nbe") {
		t.Fatalf("Expected byte index 8, got %d", got)
	}
}

func TestPositionIndexInUTF16SurrogatePair(t *testing.T) {
	content := "a😀b"

	beforeEmoji := Position{Line: 0, Character: 1}.IndexIn(content)
	if beforeEmoji != len("a") {
		t.Fatalf("Expected index before emoji to be 1, got %d", beforeEmoji)
	}

	afterEmoji := Position{Line: 0, Character: 3}.IndexIn(content)
	if afterEmoji != len("a😀") {
		t.Fatalf("Expected index after emoji to be %d, got %d", len("a😀"), afterEmoji)
	}

	midEmoji := Position{Line: 0, Character: 2}.IndexIn(content)
	if midEmoji != len("a") {
		t.Fatalf("Expected index inside surrogate pair to stay before emoji, got %d", midEmoji)
	}
}

func TestPositionIndexInClampsCharacterPastLineEnd(t *testing.T) {
	content := "ab\ncd"
	position := Position{Line: 0, Character: 99}

	if got := position.IndexIn(content); got != len("ab") {
		t.Fatalf("Expected character past line end to clamp to line end, got %d", got)
	}
}

func TestPositionIndexInUnknownLineReturnsZero(t *testing.T) {
	content := "ab\ncd"
	position := Position{Line: 4, Character: 0}

	if got := position.IndexIn(content); got != 0 {
		t.Fatalf("Expected unknown line to return 0, got %d", got)
	}
}

func TestRangeIndexesIn(t *testing.T) {
	content := "alpha\nbeta\ngamma"
	range_ := Range{
		Start: Position{Line: 1, Character: 1},
		End:   Position{Line: 1, Character: 3},
	}

	start, end := range_.IndexesIn(content)
	if start != len("alpha\nb") || end != len("alpha\nbet") {
		t.Fatalf("Expected range indexes for 'et', got start=%d end=%d", start, end)
	}
}

func TestPositionEndOfLineIn(t *testing.T) {
	content := "alpha\nbeta\ngamma"
	position := Position{Line: 1, Character: 1}

	end := position.EndOfLineIn(content)
	if end != (Position{Line: 1, Character: 4}) {
		t.Fatalf("Expected end of line at character 4, got %#v", end)
	}
}
