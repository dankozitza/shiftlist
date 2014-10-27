package shiftlist

type Shiftlist struct {
	truelist []interface{}
	MaxIndex int
	NumEntries int
	oldest int
	newest int
	Full bool
}

func New(max int) Shiftlist {
	shl := Shiftlist{make([]interface{}, max + 1), max, 0, 0, -1, false}
	shl.MaxIndex = max
	return shl
}

func (shl *Shiftlist) Add(val interface{}) {

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

	shl.truelist[shl.newest] = val
}

func (shl *Shiftlist) Get(i int) interface{} {
	return shl.truelist[shl.shiftindex(i)]
}

func (shl *Shiftlist) shiftindex(i int) int {
	var n int

	n = i + shl.oldest

	if (n > shl.MaxIndex) {
		n = n - shl.MaxIndex - 1
	}

	if (n > shl.MaxIndex || 0 > shl.MaxIndex) {
		panic("Shiftlist: error index out of range!: " + string(n))
	}

	return n
}
