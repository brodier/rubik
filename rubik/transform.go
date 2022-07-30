package rubik

type MoveElem struct {
	from int
	to   int
	flip int
}

type Transform []MoveElem

var U Transform

func init() {
	U = Transform{
		0: MoveElem{from: 0, to: 1, flip: 8},
		1: MoveElem{from: 1, to: 3, flip: 16},
		2: MoveElem{from: 3, to: 2, flip: 16},
		3: MoveElem{from: 2, to: 0, flip: 8},
		4: MoveElem{from: 8, to: 9, flip: 12},
		5: MoveElem{from: 9, to: 10, flip: 12},
		6: MoveElem{from: 10, to: 11, flip: 12},
		7: MoveElem{from: 11, to: 8, flip: 12},
	}
}

func (t Transform) Apply(r Rubik) (tr Rubik) {
	tr = make(Rubik, 26)
	copy(tr, r)
	for _, me := range t {
		tr[me.to] = flip(r[me.from], me.flip)
	}
	return
}

func flip(orientedElem int, flip int) int {
	if orientedElem += flip; orientedElem >= 24 {
		orientedElem %= 24
	}
	return orientedElem
}
