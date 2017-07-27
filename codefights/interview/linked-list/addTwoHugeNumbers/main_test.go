package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	if addTwoNumbers("9", "10") != "91" {
		t.Error("Expected 19")
	}
	if addTwoNumbers("1999", "3") != "2002" {
		t.Error("Expected 2002")
	}
}

func TestFillLeadingZeros(t *testing.T) {
	if fillLeadingZeros("9") != "0009" {
		t.Error("Expected 0009")
	}
}

func TestTrimZeros(t *testing.T) {
	if trimZeros("00029") != "29" {
		t.Error("Expected 29")
	}
}
