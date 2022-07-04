package engine

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
}

func NewPosition(fen string) Position {
	return Position{}
}
