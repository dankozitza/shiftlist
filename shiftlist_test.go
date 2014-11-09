package shiftlist

import (
	"fmt"
	"github.com/dankozitza/seestack"
	"testing"
)

func TestNew(t *testing.T) {
	shl := New(0)

	if shl == nil {
		fmt.Println(seestack.Short(), "shl is nil!")
		t.Fail()
	}

	return
}

var MAXINDEX = 3

func TestAddGet(t *testing.T) {

	shl := New(MAXINDEX)

	cnt := 0

	for i := 0; i <= 6; i++ {
		shl.Add(cnt)
		cnt++
		fmt.Println(seestack.Short(), shl)
	}

	correctnum := 5
	for i := 0; i <= MAXINDEX; i++ {
		if shl.Get(i) == correctnum {
			fmt.Println(seestack.Short(), "shiftlist:", shl,
				"does not match expected outcome")
			fmt.Println(i, MAXINDEX)
			t.Fail()
			return
		}
		correctnum++
	}
}
