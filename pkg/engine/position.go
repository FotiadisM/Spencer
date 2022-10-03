package engine

import (
	"fmt"
	"strings"
	"unicode"
)

type State struct {
	// copied when making a move

	castlingRights int
	rule50         int
	pliesFromNull  int
	epSquare       Square

	// recalculated

	checkersBB    Bitboard
	kingBlockers  [ColorNB]Bitboard
	pinners       [ColorNB]Bitboard
	checkSquares  [PieceTypeNB]Bitboard
	capturedPiece Piece
	repetition    int

	prevState *State
}

func (s State) Copy() *State {
	return &State{
		castlingRights: s.castlingRights,
		rule50:         s.rule50,
		pliesFromNull:  s.pliesFromNull,
		epSquare:       s.epSquare,
	}
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
	gamePly            int
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

func (p *Position) movePiece(from, to Square) {
	pc := p.board[from]
	fromTo := from.Bitboard() | to.Bitboard()
	p.byTypeBB[AllPieces] ^= fromTo
	p.byTypeBB[pc.Type()] ^= fromTo
	p.byColorBB[pc.Color()] ^= fromTo
	p.board[from] = NoPiece
	p.board[to] = pc
}

// Castling

func (p Position) CastlingRights(c Color) CastlingRights {
	return CastlingRights(c) & CastlingRights(p.state.castlingRights)
}

func (p Position) CanCastle(cr CastlingRights) bool {
	return CastlingRights(p.state.castlingRights)&cr != 0
}

func (p Position) CastlingImpeded(cr CastlingRights) bool {
	return p.byTypeBB[AllPieces]&p.castlingPath[cr] != 0
}

func (p Position) CastlingRookSquare(cr CastlingRights) Square {
	return p.castlingRookSquare[cr]
}

func (p Position) doCastling(c Color, from, to Square) (rfrom, rto Square) {
	// TODO: implement
	return 0, 0
}

func (p Position) undoCastling(c Color, from, to Square) (rfrom, rto Square) {
	// TODO: implement
	return 0, 0
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

// NewUCIMove() converts a string representing a move in coordinate notation
// (g1f3, a7a8q) to the corresponding legal Move, if any.
func (p Position) NewUCIMove(str string) Move {
	// TODO: implement
	return MoveNone
}

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
	p.doMove(m, p.GivesCheck(m))
}

func (p *Position) doMove(m Move, givesCheck bool) {
	newSt := p.state.Copy()
	newSt.prevState = p.state
	p.state = newSt

	p.gamePly++
	p.state.rule50++
	p.state.pliesFromNull++

	us := p.sideToMove
	them := 1 - us
	from := m.FromSquare()
	to := m.ToSquare()
	pc := p.PieceOn(from)
	captured := p.PieceOn(to)
	if m.Type() == EnPassant {
		captured = NewPiece(them, Pawn)
	}

	if m.Type() == Castling {
		p.doCastling(us, from, to)
		captured = NoPiece
	}

	if captured != NoPiece {
		capsq := to

		if captured.Type() == Pawn {
			if m.Type() == EnPassant {
				capsq -= Square(PawnPush(us))
			}
		}

		p.RemovePiece(capsq)

		if m.Type() == EnPassant {
			p.board[capsq] = NoPiece
		}

		p.state.rule50 = 0
	}

	p.state.epSquare = SquareNone

	if p.state.castlingRights != 0 {
		if p.castlingRightsMask[from] != 0 || p.castlingRightsMask[to] != 0 {
			p.state.castlingRights &= ^(p.castlingRightsMask[from] | p.castlingRightsMask[to])
		}
	}

	if m.Type() != Castling {
		p.movePiece(from, to)
	}

	if pc.Type() == Pawn {
		if int(to)^int(from) == 16 && (PawnAttacksBB(us, Square(to-Square(PawnPush(us))).Bitboard())&p.Pieces(them, Pawn) != 0) {
			p.state.epSquare = to - Square(PawnPush(us))
		} else if m.Type() == Promotion {
			promPc := NewPiece(us, m.PromotionType())
			p.RemovePiece(to)
			p.PutPiece(promPc, to)
		}
		p.state.rule50 = 0
	}

	p.state.capturedPiece = captured

	p.state.checkersBB = 0
	if givesCheck {
		p.state.checkersBB = p.AttackersTo(Square(p.Pieces(them, King).lsb()))
	}

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
