package main

const (
  type_sumbols = "letters"
  // цвета, которые будут использоваться при выводе в консоль
  Error = "\033[31m"
  Info = "\033[36m"
  Win = "\033[32m"
	empty_cell rune = '_'
)


// Устанавливает значения для шахматных фигур и индексы фигур каждой стороны
func SetSumbolsForFigures(type_sumbols string) {
  switch type_sumbols {
    case "sumbols":
      // В качестве шахматных фигур будут использоваться символы шахматных фигур из юникода
      min_index_white_chess_figure = 9818
      max_index_white_chess_figure = 9823
      // устанавливаем значки, которые будут у шахматных фигур
      WhiteKing = '♚'
      WhiteQueen = '♛'
      WhiteRook = '♜'
      WhiteElephant = '♝'
      WhiteHorse = '♞'
      WhitePawns = '♟'
      BlackKing = '♔'
      BlackQueen = '♕'
      BlackRook = '♖'
      BlackElephant = '♗'
      BlackHorse = '♘'
      BlackPawns = '♙'
    case "letters":
      // В качестве шахматных фигур будут использоваться буквы
      min_index_white_chess_figure = 65
      max_index_white_chess_figure = 90
      // устанавливаем значки, которые будут у шахматных фигур
      WhiteKing = 'K'
      WhiteQueen = 'Q'
      WhiteRook = 'R'
      WhiteElephant = 'E'
      WhiteHorse = 'H'
      WhitePawns = 'P'
      BlackKing = 'k'
      BlackQueen = 'q'
      BlackRook = 'r'
      BlackElephant = 'e'
      BlackHorse = 'h'
      BlackPawns = 'p'
    default:
      panic("Тип не найден")
  }
}
