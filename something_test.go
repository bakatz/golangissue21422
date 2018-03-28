package golangissue21422

import "testing"

func TestSomethingReturnsTrueWhenTrueIsPassed(t *testing.T) {
	if !Something(true) {
		t.Fail()
	}
}

func TestSomethingReturnsFalseWhenFalseIsPassed(t *testing.T) {
	if Something(false) {
		t.Fail()
	}

	if 1 == 2 {
		t.Fail()
	}
}
