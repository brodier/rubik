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

func NewTransform(from Rubik, to Rubik) Transform {
	return InitialTransform(from).Reverse().Then(InitialTransform(to))
}

func InitialTransform(to Rubik) Transform {
	st := make(Transform, 0)
	for i, v := range to[:nbCorner] {
		me := MoveElem{i, v % nbCorner, v - (v % nbCorner)}
		if me.from != me.to || me.flip != 0 {
			st = append(st, me)
		}
	}
	for i, v := range to[bordersIdxOffset:centersIdxOffset] {
		me := MoveElem{i - bordersIdxOffset, v % nbBorder, v - (v % nbBorder)}
		if me.from != me.to || me.flip != 0 {
			st = append(st, me)
		}
	}
	for i, v := range to[centersIdxOffset:] {
		me := MoveElem{i - centersIdxOffset, v, 0}
		if me.from != me.to || me.flip != 0 {
			st = append(st, me)
		}
	}
	return st
}

func (t Transform) Reverse() Transform {
	rt := make(Transform, len(t))
	for i, me := range t {
		rt[i] = MoveElem{me.to, me.from, reverseFlip(me.flip)}
	}
	return rt
}

func reverseFlip(flip int) int {
	switch flip {
	case 8:
		return 16
	case 16:
		return 8
	default:
		return flip
	}
}

func (t Transform) Then(next Transform) Transform {
	combineTransform := next // TODO: not yet implemented
	return combineTransform
}
