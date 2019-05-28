package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected length 9, but got: %v", len(d))
	}

	firstSuite := "Ace of suit 1"
	if d[0] != firstSuite {
		t.Errorf("Expected '%v', but got: %v", firstSuite, d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	os.Remove(testFileName)
	d := newDeck()
	d.saveToFile(testFileName)

	loadedFile := newDeckFromFile(testFileName)

	if len(loadedFile) != 9 {
		t.Errorf("Expected length of file was 9, but got: %v", len(loadedFile))
	}

	os.Remove(testFileName)
}
