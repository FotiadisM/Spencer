package engine

import (
	"fmt"
	"strings"
	"unicode"
)

type State struct {
	castlingRights CastlingRights
	rule50         int
	pliesFromNull  int
	epSquare       Square

	checkersBB    Bitboard
	kingBlockers  [ColorNB]Bitboard
	pinners       [ColorNB]Bitboard
	checkSquares  [PieceTypeNB]Bitboard
	capturedPiece Piece
	repetition    int

	previous *State
}

type Position struct {
	sideToMove         Color
	board              [SquareNB]Piece
	byTypeBB           [PieceTypeNB]Bitboard
	byColorBB          [ColorNB]Bitboard
	pieceCount         [PieceNB]int
	castlingRightsMask [SquareNB]int
	castlingRookSquare [CastlingRightsNB]Square
	castlingPath       [CastlingRightsNB]Bitboard
	state              *State
}

func NewPosition(fen string) *Position {
	p := &Position{
		state: &State{},
	}

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
	p.sideToMove = White
	if str[1] == "b" {
		p.sideToMove = Black
	}

	// TODO: continue to parse the rest of the fen string

	return p
}

func (p Position) Fen() string {
	// TODO: implement
	return ""
}

// Position representation

func (p Position) Pieces(c Color, pt PieceType) Bitboard {
	return p.byColorBB[c] & p.byTypeBB[pt]
}

func (p Position) PiecesByType(pt PieceType) Bitboard {
	return p.byTypeBB[pt]
}

func (p Position) PiecesByColor(c Color) Bitboard {
	return p.byColorBB[c]
}

func (p Position) PieceOn(s Square) Piece {
	return p.board[s]
}

func (p Position) EpSquare() Square {
	return p.state.epSquare
}

func (p *Position) PutPiece(pc Piece, s Square) {
	p.board[s] = pc
	p.byTypeBB[AllPieces] |= s.Bitboard()
	p.byTypeBB[pc.Type()] |= s.Bitboard()
	p.byColorBB[pc.Color()] |= s.Bitboard()
	p.pieceCount[pc]++
	p.pieceCount[NewPiece(Color(pc.Color()), AllPieces)]++
}

func (p *Position) RemovePiece(s Square) {
	pc := p.board[s]
	p.byTypeBB[AllPieces] ^= s.Bitboard()
	p.byTypeBB[pc.Type()] ^= s.Bitboard()
	p.byColorBB[pc.Color()] ^= s.Bitboard()
	p.board[s] = NoPiece
	p.pieceCount[pc]--
	p.pieceCount[NewPiece(Color(pc.Color()), AllPieces)]--
}

func (p *Position) movePiece(fr, to Square) {
	pc := p.board[fr]
	fromTo := fr.Bitboard() | to.Bitboard()
	p.byTypeBB[AllPieces] ^= fromTo
	p.byTypeBB[pc.Type()] ^= fromTo
	p.byColorBB[pc.Color()] ^= fromTo
	p.board[fr] = NoPiece
	p.board[to] = pc
}

// Castling

func (p Position) CastlingRights(c Color) CastlingRights {
	return CastlingRights(c) & p.state.castlingRights
}

func (p Position) CanCastle(cr CastlingRights) bool {
	if p.state.castlingRights&cr == 0 {
		return false
	}
	return true
}

func (p Position) CastlingImpeded(cr CastlingRights) bool {
	if p.byTypeBB[AllPieces]&p.castlingPath[cr] == 0 {
		return false
	}
	return true
}

func (p Position) CastlingRookSquare(cr CastlingRights) Square {
	return p.castlingRookSquare[cr]
}

// Checking

func (p Position) Checkers() Bitboard {
	return p.state.checkersBB
}

func (p Position) KingBlockers(c Color) Bitboard {
	return p.state.kingBlockers[c]
}

func (p Position) CheckSquares(pt PieceType) Bitboard {
	return p.state.checkSquares[pt]
}

func (p Position) Pinners(c Color) Bitboard {
	return p.state.pinners[c]
}

// Attacks to/from a given square
func (p Position) AttackersTo(s Square) Bitboard {
	return p.attackersTo(s, p.byTypeBB[AllPieces])
}

func (p Position) attackersTo(s Square, occupied Bitboard) Bitboard {
	// TODO: implement
	return 0
}

// Properties of Moves

func (p Position) IsMoveLegal(m Move) bool {
	// TODO: implement
	return false
}

func (p Position) IsMovePseudoLegal(m Move) bool {
	// TODO: implement
	return false
}

func (p Position) IsMoveCapture(m Move) bool {
	// TODO: implement
	return false
}

func (p Position) GivesCheck(m Move) bool {
	// TODO: implement
	return false
}

func (p Position) MovedPiece(m Move) Piece {
	// TODO: implement
	return 0
}

func (p Position) CapturedPiece(m Move) Piece {
	// TODO: implement
	return 0
}

// Doing and undoing moves

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
			s += fmt.Sprintf("| %v ", p.board[sq])
			sq += 1
		}
		s += "|\n  +---+---+---+---+---+---+---+---+\n"
	}
	s += "    A   B   C   D   E   F   G   H\n"

	return s
}
