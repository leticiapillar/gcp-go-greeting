package main

import "testing"

func TestMainSuccess(t *testing.T) {

	expected := "<b>Code.education Rocks!</b>"

	result := greeting("Code.education Rocks!")

	if result != expected {
		t.Errorf("Function sum failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("Function sum success")
	}
}
