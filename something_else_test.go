package golangissue21422

import "testing"

func TestSomethingElseReturnsTrueWhenTrueIsPassed(t *testing.T) {
	if SomethingElse(true) {
		t.Fail()
	}
}

func TestSomethingElseReturnsFalseWhenFalseIsPassed(t *testing.T) {
	if !SomethingElse(false) {
		t.Fail()
	}
}
