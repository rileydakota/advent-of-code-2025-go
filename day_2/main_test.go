package main

import "testing"

func TestInvalidId(t *testing.T) {
	invalidIds := [...]string{
		"11",
		"22",
		"99",
		"1188511885",
		"222222",
		"446446",
		"38593859",
	}

	for _, id := range invalidIds {
		if isIdValid(id) {
			t.Errorf("expected %s to be valid", id)
		}
	}

	validIds := [...]string{
		"115",
		"1188511890",
	}

	for _, id := range validIds {
		if !isIdValid(id) {
			t.Errorf("expected %s to be valid", id)
		}
	}
}
