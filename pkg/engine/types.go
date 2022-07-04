package engine

type Color int

const (
	White Color = iota
	Black
	ColorNB Color = 2
)

type MoveType int

const (
	Normal    MoveType = 0
	Promotion MoveType = 1 << 14
	EnPassant MoveType = 2 << 14
	Castling  MoveType = 3 << 14
)

// A move needs 16 bits to be stored
//
// bit  0- 5: destination square (from 0 to 63)
// bit  6-11: origin square (from 0 to 63)
// bit 12-13: promotion piece type - 2 (from KNIGHT-2 to QUEEN-2)
// bit 14-15: special move flag: promotion (1), en passant (2), castling (3)
// NOTE: en passant bit is set only when a pawn can be captured
//
// Special cases are MOVE_NONE and MOVE_NULL. We can sneak these in because in
// any normal move destination square is always different from origin square
// while MOVE_NONE and MOVE_NULL have the same origin and destination square.
type Move int

const (
	MoveNone Move = 0
	MoveNull Move = 65
)

func NewSimpleMove(fr, to Square) Move {
	return Move((int(fr) << 6) + int(to))
}

func NewMove(fr, to Square, pt PieceType, mt MoveType) Move {
	return Move(int(mt) + (int(pt-Knight) << 12) + (int(fr) << 6) + int(to))
}

func (m Move) IsOK() bool {
	return m.FromSquare() != m.ToSquare() // catch MoveNull and MoveNone
}

func (m Move) FromSquare() Square {
	return Square((int(m) >> 6) & 0x3F)
}

func (m Move) ToSquare() Square {
	return Square(m & 0x3F)
}

func (m Move) TypeOf() MoveType {
	return MoveType(int(m) & (3 << 14))
}

func (m Move) PromotionType() PieceType {
	return PieceType(((int(m) >> 12) & 3) + int(Knight))
}

type CastlingRights int

const (
	NoCastling CastlingRights = 0
	WhiteOO    CastlingRights = 1
	WhiteOOO                  = WhiteOO << 1
	BlackOO                   = WhiteOO << 2
	BlackOOO                  = WhiteOO << 3

	KingSide      = WhiteOO | BlackOO
	QueenSide     = WhiteOOO | BlackOOO
	WhiteCastling = WhiteOO | WhiteOOO
	BlackCastling = BlackOO | BlackOOO
	AnyCastling   = WhiteCastling | BlackCastling

	CastlingRightsNB CastlingRights = 16
)

type PieceType int

const (
	NoPieceType PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King

	AllPieces   PieceType = 0
	PieceTypeNB PieceType = 8
)

type Piece int

const (
	NoPiece Piece = iota
	WPawn
	WKnight
	WBishop
	WRook
	WQueen
	WKing

	BPawn Piece = iota + 2
	BKnight
	BBishop
	BRook
	BQueen
	BKing

	PieceNB Piece = 16
)

func NewPiece(c Color, pt PieceType) Piece {
	return Piece((int(c) << 3) + int(pt))
}

func (pc Piece) TypeOf() PieceType {
	return PieceType(pc & 7)
}

func (pc Piece) ColorOf() PieceType {
	if pc == NoPiece {
		panic("Piece is of type NoPiece")
	}
	return PieceType(pc >> 3)
}

func (pc Piece) SwapColor() Piece {
	return pc ^ 8
}

type Direction int

const (
	North Direction = 8
	East  Direction = 1
	South           = -North
	West            = -East

	NorthEast = North + East
	NorthWest = North + West
	SouthEast = South + East
	SouthWest = South + West
)

func PawnPush(c Color) Direction {
	if c == White {
		return North
	}
	return South
}

type File int

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	FileNB
)

func (f File) Bitboard() Bitboard {
	return FileABB << f
}

type Rank int

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	RankNB
)

func (r Rank) RelativeRank(c Color) Rank {
	return Rank(int(r) ^ (int(c) * 7))
}

func (r Rank) Bitboard() Bitboard {
	return Rank1BB << (8 * r)
}

type Square int

const (
	SquareA1 Square = iota
	SquareB1
	SquareC1
	SquareD1
	SquareE1
	SquareF1
	SquareG1
	SquareH1

	SquareA2
	SquareB2
	SquareC2
	SquareD2
	SquareE2
	SquareF2
	SquareG2
	SquareH2

	SquareA3
	SquareB3
	SquareC3
	SquareD3
	SquareE3
	SquareF3
	SquareG3
	SquareH3

	SquareA4
	SquareB4
	SquareC4
	SquareD4
	SquareE4
	SquareF4
	SquareG4
	SquareH4

	SquareA5
	SquareB5
	SquareC5
	SquareD5
	SquareE5
	SquareF5
	SquareG5
	SquareH5

	SquareA6
	SquareB6
	SquareC6
	SquareD6
	SquareE6
	SquareF6
	SquareG6
	SquareH6

	SquareA7
	SquareB7
	SquareC7
	SquareD7
	SquareE7
	SquareF7
	SquareG7
	SquareH7

	SquareA8
	SquareB8
	SquareC8
	SquareD8
	SquareE8
	SquareF8
	SquareG8
	SquareH8

	SquareZero Square = 0
	SquareNB   Square = 64
)

func NewSquare(f File, r Rank) Square {
	return Square((int(r) << 3) + int(f))
}

func (s Square) IsOK() bool {
	return s >= SquareA1 && s <= SquareH8
}

func (s Square) Bitboard() Bitboard {
	if !s.IsOK() {
		panic("Suqare out of range")
	}
	return SquareBB[s]
}

func (s Square) FlipRank() Square {
	return s ^ SquareA8
}

func (s Square) FlipFile() Square {
	return s ^ SquareH1
}

func (s Square) GetRank() Rank {
	return Rank(int(s) >> 3)
}

func (s Square) GetFile() File {
	return File(int(s) & 7)
}

func (s Square) RelativeSquare(c Color) Square {
	return Square(int(s) & (int(c) * 56))
}

func (s Square) RelativeRank(c Color) Rank {
	return s.GetRank().RelativeRank(c)
}

func (s Square) RankBB() Bitboard {
	return s.GetRank().Bitboard()
}
