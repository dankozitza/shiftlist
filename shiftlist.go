package shiftlist

import (
	//"fmt"
	//"github.com/dankozitza/seestack"
)

type ShiftList struct {
	truelist []interface{}
	MaxIndex int
	NumEntries int
	oldest int
	newest int
	Full bool
}

func New(max int) ShiftList {
	if (max <= 0) {
		max = 100
	}
	shl := ShiftList{make([]interface{}, max + 1), max, 0, 0, -1, false}
	return shl
}

func (shl *ShiftList) Add(val interface{}) {

	shl.newest++

	if (shl.NumEntries <= shl.MaxIndex) {
		shl.NumEntries++
	}

	if shl.newest > shl.MaxIndex {
		shl.newest = 0
		shl.Full = true
	}
	if (shl.Full) {
		shl.oldest += 1
		if shl.oldest > shl.MaxIndex {
			shl.oldest = 0
		}
	}

	//fmt.Println(seestack.Short(),
	//	"shl.newest:", shl.newest, "shl.MaxIndex:", shl.MaxIndex)
	//fmt.Println(seestack.Short(), "shl:")
	//fmt.Println(shl)


	shl.truelist[shl.newest] = val
}

func (shl *ShiftList) Get(i int) interface{} {
	return shl.truelist[shl.shiftindex(i)]
}

func (shl *ShiftList) shiftindex(i int) int {
	var n int

	n = i + shl.oldest

	if (n > shl.MaxIndex) {
		n = n - shl.MaxIndex - 1
	}

	if (n > shl.MaxIndex || 0 > shl.MaxIndex) {
		panic("ShiftList: error index out of range!: " + string(n))
	}

	return n
}
