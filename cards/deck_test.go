package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 36 {
		t.Errorf("Expected deck length of 20, but got %v",  len(d))
	}
}