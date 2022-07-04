package engine

import (
	"fmt"
	"strings"
	"unicode"
)

type State struct {
	CastlingRights int
	Rule50         int
	PliesFromNull  int
	EpSquare       Square
}

type Position struct {
	SideToMove         Color
	Board              [SquareNB]Piece
	ByTypeBB           [PieceTypeNB]Bitboard
	ByColorBB          [ColorNB]Bitboard
	PieceCount         [PieceNB]int
	CastlingRightsMask [SquareNB]int
	CastlingRookSquare [CastlingRightsNB]Square
	CastlingPath       [CastlingRightsNB]Bitboard
	State              State
}

func NewPosition(fen string) *Position {
	p := &Position{}

	str := strings.Fields(fen)

	// 1. Piece placement
	sq := SquareA8
	pieceToChar := " PNBRQK  pnbrqk"
	for _, r := range str[0] {
		switch {
		case unicode.IsDigit(r):
			sq += Square((r - '0') * rune(East))
		case r == '/':
			sq += Square(2 * South)
		default:
			if idx := strings.Index(pieceToChar, string(r)); idx != -1 {
				p.PutPiece(Piece(idx), sq)
				sq++
			}
		}
	}

	// 2. Active color
	p.SideToMove = White
	if str[1] == "b" {
		p.SideToMove = Black
	}

	// TODO: continue to parse the rest of the fen string

	return p
}

func (p Position) Fen() string {
	// TODO: implement
	return ""
}

func (p Position) Pieces(c Color, pt PieceType) Bitboard {
	// TODO: implement
	return 0
}

func (p Position) PiecesByType(pt PieceType) Bitboard {
	// TODO: implement
	return 0
}

func (p Position) PiecesByColor(c Color) Bitboard {
	// TODO: implement
	return 0
}

func (p Position) PieceOn(s Square) Piece {
	// TODO: implement
	return NoPiece
}

func (p Position) EpSquare() Square {
	return p.State.EpSquare
}

func (p *Position) PutPiece(pc Piece, s Square) {
	p.Board[s] = pc
	p.ByTypeBB[AllPieces] |= s.Bitboard()
	p.ByTypeBB[pc.Type()] |= s.Bitboard()
	p.ByColorBB[pc.Color()] |= s.Bitboard()
	p.PieceCount[pc]++
	p.PieceCount[NewPiece(Color(pc.Color()), AllPieces)]++
}

func (p *Position) RemovePiece(s Square) {
	pc := p.Board[s]
	p.ByTypeBB[AllPieces] ^= s.Bitboard()
	p.ByTypeBB[pc.Type()] ^= s.Bitboard()
	p.ByColorBB[pc.Color()] ^= s.Bitboard()
	p.Board[s] = NoPiece
	p.PieceCount[pc]--
	p.PieceCount[NewPiece(Color(pc.Color()), AllPieces)]--
}

func (p *Position) movePiece(fr, to Square) {
	// TODO: implement
}

func (p *Position) DoMove(m Move) {
	// TODO: implement
}

func (p *Position) UndooMove(m Move) {
	// TODO: implement
}

func (p *Position) DoNullMove(m Move) {
	// TODO: implement
}

func (p *Position) UndoNullMove(m Move) {
	// TODO: implement
}

func (p Position) String() string {
	s := "  +---+---+---+---+---+---+---+---+\n"
	for sq := SquareA8; sq.IsOK(); sq += -16 {
		s += fmt.Sprintf("%v ", sq.Rank())
		for i := 0; i < 8; i++ {
			s += fmt.Sprintf("| %v ", p.Board[sq])
			sq += 1
		}
		s += "|\n  +---+---+---+---+---+---+---+---+\n"
	}
	s += "    A   B   C   D   E   F   G   H\n"

	return s
}
