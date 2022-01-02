package loop_test

import (
	"testing"
)
func TestLoop(t *testing.T)  {
	
	var n int = 0
    for n < 3 {
		t.Log(n)
		n = n + 1
	}

	
}
