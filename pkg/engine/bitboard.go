package engine

type Bitboard uint64

const (
	AllSquares  Bitboard = ^Bitboard(0)
	DarkSquares Bitboard = 0xAA55AA55AA55AA55

	FileABB Bitboard = 0x0101010101010101
	FileBBB          = FileABB << 1
	FileCBB          = FileABB << 2
	FileDBB          = FileABB << 3
	FileEBB          = FileABB << 4
	FileFBB          = FileABB << 5
	FileGBB          = FileABB << 6
	FileHBB          = FileABB << 7

	Rank1BB Bitboard = 0xFF
	Rank2BB          = Rank1BB << (8 * 1)
	Rank3BB          = Rank1BB << (8 * 2)
	Rank4BB          = Rank1BB << (8 * 3)
	Rank5BB          = Rank1BB << (8 * 4)
	Rank6BB          = Rank1BB << (8 * 5)
	Rank7BB          = Rank1BB << (8 * 6)
	Rank8BB          = Rank1BB << (8 * 7)

	QueenSideBB   = FileABB | FileBBB | FileCBB | FileDBB
	KingSideBB    = FileEBB | FileFBB | FileGBB | FileHBB
	CenterFilesBB = FileCBB | FileDBB | FileEBB | FileFBB
	Center        = (FileDBB | FileEBB) & (Rank4BB | Rank5BB)
)

var (
	PopCount16     [1 << 16]uint8
	SquareDistance [SquareNB][SquareNB]uint8

	SquareBB      [SquareNB]Bitboard
	LineBB        [SquareNB][SquareNB]Bitboard
	BetweenBB     [SquareNB][SquareNB]Bitboard
	PseudoAttacks [PieceTypeNB][SquareNB]Bitboard
	PawnAttacks   [ColorNB][SquareNB]Bitboard
)

func (b Bitboard) MoreThanOne() bool {
	return b&(b-1) != 0
}

func (b Bitboard) Shift(d Direction) Bitboard {
	switch d {
	case North:
		return b << 8
	case South:
		return b >> 8
	case North + North:
		return b << 16
	case South + South:
		return b >> 16
	case East:
		return (b & ^FileHBB) << 1
	case West:
		return (b & ^FileABB) >> 1
	case North + East:
		return (b & ^FileHBB) << 9
	case North + West:
		return (b & ^FileABB) << 7
	case South + East:
		return (b & ^FileHBB) >> 7
	case South + West:
		return (b & ^FileABB) >> 9
	default:
		panic("Direction is not applicable")
	}
}

func PawnAttacksBB(c Color, b Bitboard) Bitboard {
	if c == White {
		return b.Shift(NorthWest) | b.Shift(NorthEast)
	}
	return b.Shift(SouthWest) | b.Shift(SouthEast)
}

type Magic struct {
	Mask    Bitboard
	Magic   Bitboard
	Attacks []Bitboard
	Shift   uint
}

func (m Magic) Inxex(occupied Bitboard) uint {
	// WARN: 64bit only
	return uint(((occupied & m.Mask) * m.Magic) >> m.Shift)
}

var (
	RookMagics   [SquareNB]Magic
	BishopMagics [SquareNB]Magic
)
