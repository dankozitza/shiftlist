package shiftlist

import (
	"bytes"
	"fmt"
)

type ShiftList struct {
	truelist   []interface{}
	MaxIndex   int
	NumEntries int
	oldest     int
	newest     int
	Full       bool
}

func New(max int) *ShiftList {
	if max < 0 {
		max = 100
	}
	return &ShiftList{make([]interface{}, max+1), max, 0, 0, -1, false}
}

func (shl *ShiftList) Add(val interface{}) {

	if shl.NumEntries <= shl.MaxIndex {
		shl.NumEntries++
	}

	shl.newest++

	if shl.newest > shl.MaxIndex {
		shl.newest = 0
		shl.Full = true
	}

	if shl.Full {

		shl.oldest++
		if shl.oldest > shl.MaxIndex {
			shl.oldest = 0
		}
	}
	shl.truelist[shl.newest] = val
}

func (shl *ShiftList) Get(i int) interface{} {
	return shl.truelist[shl.shiftindex(i)]
}

func (shl *ShiftList) shiftindex(i int) int {
	var n int

	n = i + shl.oldest

	if n > shl.MaxIndex {
		n = n - shl.MaxIndex - 1
	}

	if n > shl.MaxIndex || 0 > shl.MaxIndex {
		panic("shiftlist.ShiftList: error index out of range!: " + string(n))
	}

	return n
}

func (shl *ShiftList) GetArray() []interface{} {

	a := make([]interface{}, shl.NumEntries)

	for i := 0; i < shl.NumEntries; i++ {
		a[i] = shl.Get(i)
	}

	return a
}

func (shl *ShiftList) String() string {

	var buff bytes.Buffer
	fmt.Fprint(&buff, shl.GetArray())

	return buff.String()
}
