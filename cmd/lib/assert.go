package lib

import "testing"

func Assert(t *testing.T, condition bool, message ...string) {
	if !condition {
		if len(message) > 0 {
			t.Errorf("Assertion failed: %s", message)
		} else {
			t.Errorf("Assertion Failed")
		}
	}
}
