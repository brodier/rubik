package rubik

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Rubik []int

const nbCorner = 8
const cornerSize = 3
const nbBorder = 12
const borderSize = 2
const nbCenter = 6
const bordersIdxOffset = 8
const centersIdxOffset = 20

var loadingSequence []int
var cornersMap map[string]int
var bordersMap map[string]int
var centersMap map[string]int

var unloadingSequence []int
var reverseCornersMap []string
var reverseBordersMap []string
var reverseCentersMap []string

//             00 01 02
//             03 04 05
//             06 07 08
//
//   09 10 11  12 13 14  15 16 17  18 19 20
//   21 22 23  24 25 26  27 28 29  30 31 32
//   33 34 35  36 37 38  39 40 41  42 43 44
//
//             45 46 47
//             48 49 50
//             51 52 53

func init() {

	loadingSequence = make([]int, 54)
	unloadingSequence = make([]int, 54)
	copy(loadingSequence, []int{
		6, 11, 12, 14, 15, 8, // Corner 0, 1
		9, 0, 20, 2, 17, 18, // Corner 2, 3
		45, 36, 35, 47, 39, 38, // Corner 4, 5
		33, 44, 51, 53, 42, 41, // Corner 6, 7
		7, 13, 16, 5, 1, 19, 10, 3, // Border 0, 1, 2, 3
		24, 23, 26, 27, 30, 29, 32, 21, // Border 4, 5, 6, 7
		46, 37, 40, 50, 52, 43, 34, 48, // Border 8, 9,10,11
		25, 28, 31, 22, 4, 49}) // Center 0 to 5
	for k, v := range loadingSequence {
		unloadingSequence[v] = k
	}
	cornersSeed := []string{"ULF", "FRU", "LUB", "URB", "DFL", "DRF", "LBD", "DBR"}
	cornersMap = make(map[string]int, 24)
	reverseCornersMap = make([]string, 24)
	for k, v := range cornersSeed {
		var flip string
		var i int
		bs := []byte(v)
		i = k
		cornersMap[v] = i
		reverseCornersMap[i] = v
		i += 8
		flip = string([]byte{bs[1], bs[2], bs[0]})
		cornersMap[flip] = i
		reverseCornersMap[i] = flip
		i += 8
		flip = string([]byte{bs[2], bs[0], bs[1]})
		cornersMap[flip] = i
		reverseCornersMap[i] = flip
	}
	bordersSeed := []string{
		"UF", "RU", "UB", "LU",
		"FL", "FR", "BR", "BL",
		"DF", "RD", "DB", "LD"}
	bordersMap = make(map[string]int, 24)
	reverseBordersMap = make([]string, 24)
	for k, v := range bordersSeed {
		bs := []byte(v)
		bordersMap[v] = k
		reverseBordersMap[k] = v
		flip := string([]byte{bs[1], bs[0]})
		bordersMap[flip] = k + 12
		reverseBordersMap[k+12] = flip
	}
	centersMap = make(map[string]int)
	reverseCentersMap = []string{"F", "R", "B", "L", "U", "D"}
	for k, v := range reverseCentersMap {
		centersMap[v] = k
	}
	log.Printf("corner map : %v\n", cornersMap)
	log.Printf("borders map : %v\n", bordersMap)
	log.Printf("center map : %v\n", centersMap)

}

func NewRubik(r io.Reader) *Rubik {
	bytes := make([]byte, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		block := strings.Split(line, " ")
		for _, elem := range block {
			if len(elem) > 0 {
				for _, v := range []byte(elem) {
					bytes = append(bytes, v)
				}
			}
		}
	}
	elems := make([]byte, 54)
	for k, v := range loadingSequence {
		elems[k] = bytes[v]
	}
	rubik := make(Rubik, nbCorner+nbBorder+nbCenter)
	corners := rubik[:8]
	for idCorner := 0; idCorner < nbCorner; idCorner++ {
		i := idCorner * cornerSize
		cornerString := string([]byte{elems[i], elems[i+1], elems[i+2]})
		corners[idCorner] = cornersMap[cornerString]
		log.Printf("Loading corner %v with %v = %v\n", idCorner, cornerString, corners[idCorner])
	}
	borders := rubik[8:20]
	for idBorder := 0; idBorder < nbBorder; idBorder++ {
		i := (nbCorner * cornerSize) + (idBorder * borderSize)
		borderString := string([]byte{elems[i], elems[i+1]})
		borders[idBorder] = bordersMap[borderString]
		log.Printf("Loading border %v with %v = %v\n", idBorder, borderString, borders[idBorder])
	}
	centers := rubik[20:]
	for idCenter := 0; idCenter < nbCenter; idCenter++ {
		i := (nbCorner * cornerSize) + (nbBorder * borderSize) + idCenter
		centers[idCenter] = centersMap[string([]byte{elems[i]})]
		log.Printf("Loading center %v : %v\n", idCenter, centers[idCenter])
	}

	return &rubik
}

func (r *Rubik) corner(id int) int {
	return (*r)[id]
}
func (r *Rubik) border(id int) int {
	return (*r)[id+bordersIdxOffset]
}
func (r *Rubik) center(id int) int {
	return (*r)[id+centersIdxOffset]
}

func (r *Rubik) unload() []string {
	ul := make([]byte, 54)
	for idCorner := 0; idCorner < nbCorner; idCorner++ {
		s := reverseCornersMap[r.corner(idCorner)]
		i := idCorner * cornerSize
		ul[loadingSequence[i]] = s[0]
		ul[loadingSequence[i+1]] = s[1]
		ul[loadingSequence[i+2]] = s[2]
	}
	for idBorder := 0; idBorder < nbBorder; idBorder++ {
		s := reverseBordersMap[r.border(idBorder)]
		i := nbCorner*cornerSize + idBorder*borderSize
		ul[loadingSequence[i]] = s[0]
		ul[loadingSequence[i+1]] = s[1]
	}
	for idCenter := 0; idCenter < nbCenter; idCenter++ {
		i := nbCorner*cornerSize + nbBorder*borderSize + idCenter
		ul[loadingSequence[i]] = reverseCentersMap[r.center(idCenter)][0]
	}
	ulStringSize := 3
	unload := make([]string, 18)
	for i := 0; i < ulStringSize*len(unload); i += ulStringSize {
		unload[i/ulStringSize] = string([]byte{ul[i], ul[i+1], ul[i+2]})
	}
	return unload
}

func irc(corner int) int {
	return rc(rc(corner))
}

func rc(corner int) int {
	corner += nbCorner
	if corner >= nbCorner*3 {
		corner -= nbCorner * 3
	}
	return corner
}

func fb(border int) int {
	border += nbBorder
	if border >= nbBorder*2 {
		border -= nbBorder * 2
	}
	return border
}

func (rubik *Rubik) Display(out io.Writer) {
	r := rubik.unload()
	fmt.Fprintf(out, "    %s\n    %s\n    %s\n", r[0], r[1], r[2])
	fmt.Fprintf(out, "\n%s %s %s %s\n", r[3], r[4], r[5], r[6])
	fmt.Fprintf(out, "%s %s %s %s\n", r[7], r[8], r[9], r[10])
	fmt.Fprintf(out, "%s %s %s %s\n\n", r[11], r[12], r[13], r[14])
	fmt.Fprintf(out, "    %s\n    %s\n    %s\n", r[15], r[16], r[17])
}

func (rubik *Rubik) MoveUpDirect() {
	*rubik = U.Apply(*rubik)
}
