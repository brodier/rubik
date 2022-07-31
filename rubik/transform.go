package rubik

const Corner = 0 // Sommet Corner
const Border = 1 // Border
const Center = 2 // Center

type MoveElem struct {
	piece int
	from  int
	to    int
	flip  int
}

type Transform []MoveElem

var U Transform

func init() {
	U = Transform{
		0: MoveElem{piece: Corner, from: 0, to: 1, flip: 8},
		1: MoveElem{piece: Corner, from: 1, to: 3, flip: 16},
		2: MoveElem{piece: Corner, from: 3, to: 2, flip: 16},
		3: MoveElem{piece: Corner, from: 2, to: 0, flip: 8},
		4: MoveElem{piece: Border, from: 0, to: 1, flip: 12},
		5: MoveElem{piece: Border, from: 1, to: 2, flip: 12},
		6: MoveElem{piece: Border, from: 2, to: 3, flip: 12},
		7: MoveElem{piece: Border, from: 3, to: 0, flip: 12},
	}
}

func (t Transform) Apply(r Rubik) (tr Rubik) {
	tr = make(Rubik, 26)
	copy(tr, r)
	for _, me := range t {
		offset := 0
		switch me.piece {
		case Border:
			offset += bordersIdxOffset
		case Center:
			offset += centersIdxOffset
		}
		tr[me.to+offset] = flip(r[me.from+offset], me.flip)
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
	for i, v := range to[:bordersIdxOffset] {
		cid := v % nbCorner
		me := MoveElem{Corner, cid, i, v - cid}
		if me.from != me.to || me.flip != 0 {
			st = append(st, me)
		}
	}
	for i, v := range to[bordersIdxOffset:centersIdxOffset] {
		bId := v % nbBorder
		me := MoveElem{Border, bId, i, v - bId}
		if me.from != me.to || me.flip != 0 {
			st = append(st, me)
		}
	}
	for i, v := range to[centersIdxOffset:] {
		me := MoveElem{Center, v, i, 0}
		if me.from != me.to || me.flip != 0 {
			st = append(st, me)
		}
	}
	return st
}

func (t Transform) Reverse() Transform {
	rt := make(Transform, len(t))
	for i, me := range t {
		rt[i] = MoveElem{me.piece, me.to, me.from, reverseFlip(me.flip)}
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
