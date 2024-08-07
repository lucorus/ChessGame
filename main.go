package main

import (
	"fmt"
)


var min_index_white_chess_figure uint64
var max_index_white_chess_figure uint64
var min_index_black_chess_figure uint64
var max_index_black_chess_figure uint64
var white_king, white_queen, white_rook, white_elephant, white_horse, white_pawns rune
var black_king, black_queen, black_rook, black_elephant, black_horse, black_pawns rune
var empty_cell rune = '_'


// Устанавливает значения для шахматных фигур и индексы фигур каждой стороны
func SetSumbolsForFigures(type_sumbols string) {
	switch type_sumbols {
		case "sumbols":
			// В качестве шахматных фигур будут использоваться символы шахматных фигур из юникода
			min_index_white_chess_figure = 9818
			max_index_white_chess_figure = 9823
			min_index_black_chess_figure = 9812
			max_index_black_chess_figure = 9817
			// устанавливаем значки, которые будут у шахматных фигур
			white_king = '♚'
			white_queen = '♛'
			white_rook = '♜'
			white_elephant = '♝'
			white_horse = '♞'
			white_pawns = '♟'
			black_king = '♔'
			black_queen = '♕'
			black_rook = '♖'
			black_elephant = '♗'
			black_horse = '♘'
			black_pawns = '♙'
		case "letters":
			// В качестве шахматных фигур будут использоваться буквы
			min_index_white_chess_figure = 65
			max_index_white_chess_figure = 90
			min_index_black_chess_figure = 97
			max_index_black_chess_figure = 122
			// устанавливаем значки, которые будут у шахматных фигур
			white_king = 'K'
			white_queen = 'Q'
			white_rook = 'R'
			white_elephant = 'E'
			white_horse = 'H'
			white_pawns = 'P'
			black_king = 'k'
			black_queen = 'q'
			black_rook = 'r'
			black_elephant = 'e'
			black_horse = 'h'
			black_pawns = 'p'
		default:
			panic("Тип не найден")
	}
	/*if type_sumbols == "sumbols" {
		// В качестве шахматных фигур будут использоваться символы шахматных фигур из юникода
		min_index_white_chess_figure = 9818
		max_index_white_chess_figure = 9823
		min_index_black_chess_figure = 9812
		max_index_black_chess_figure = 9817
		// устанавливаем значки, которые будут у шахматных фигур
		white_king = '♚'
		white_queen = '♛'
		white_rook = '♜'
		white_elephant = '♝'
		white_horse = '♞'
		white_pawns = '♟'
		black_king = '♔'
		black_queen = '♕'
		black_rook = '♖'
		black_elephant = '♗'
		black_horse = '♘'
		black_pawns = '♙'
		return 
	} else if type_sumbols == "letters" {
		// В качестве шахматных фигур будут использоваться буквы
		min_index_white_chess_figure = 65
		max_index_white_chess_figure = 90
		min_index_black_chess_figure = 97
		max_index_black_chess_figure = 122
		// устанавливаем значки, которые будут у шахматных фигур
		white_king = 'K'
		white_queen = 'Q'
		white_rook = 'R'
		white_elephant = 'E'
		white_horse = 'H'
		white_pawns = 'P'
		black_king = 'k'
		black_queen = 'q'
		black_rook = 'r'
		black_elephant = 'e'
		black_horse = 'h'
		black_pawns = 'p'
		return 
	}
	panic("Тип не найден")*/
}



func PrintField(field [10][10]rune) {
	for _, line := range field {
			for _, item := range line {
				fmt.Printf("%c ", item)
			}
			fmt.Println()
	}
}


func main() {
	var type_sumbols string = "lettes"
	// fmt.Scan(&type_sumbols)
	SetSumbolsForFigures(type_sumbols)

	field := [10][10] rune {
		{'#', '|', '1', '2', '3', '4', '5', '6', '7', '8'},
    {'-', '+','-', '-', '-', '-', '-', '-', '-', '-'},
    {'1','|', black_rook, black_horse, black_elephant, black_queen, black_king, black_elephant, black_horse, black_rook},
    {'2','|', black_pawns, black_pawns, black_pawns, black_pawns, black_pawns, black_pawns, black_pawns, black_pawns},
    {'3','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'4','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'5','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'6','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'7','|', white_pawns, white_pawns, white_pawns, white_pawns, white_pawns, white_pawns, white_pawns, white_pawns},
    {'8','|', white_rook, white_horse, white_elephant, white_queen, white_king, white_elephant, white_horse, white_rook},
	}

	/*
	
		field := [10][10] rune {
				{'#', '|', '1', '2', '3', '4', '5', '6', '7', '8'},
       	{'-', '+','-', '-', '-', '-', '-', '-', '-', '-'},
        {'1','|', '♖', '♘', '♗', '♕', '♔', '♗', '♘', '♖'},
        {'2','|', '♙', '♙', '♙', '♙', '♙', '♙', '♙', '♙'},
        {'3','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'4','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'5','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'6','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'7','|', '♟', '♟', '♟', '♟', '♟', '♟', '♟', '♟'},
        {'8','|', '♜', '♞', '♝', '♛', '♚', '♝', '♞', '♜'},
	}

	*/

		/*field := [10][10] rune {
				{'#', '|', '1', '2', '3', '4', '5', '6', '7', '8'},
       	{'-', '+','-', '-', '-', '-', '-', '-', '-', '-'},
        {'1','|', '9814', '9816', '9815', '9813', '9812', '9815', '9816', '9814'},
        {'2','|', '9817', '9817', '9817', '9817', '9817', '9817', '9817', '9817'},
        {'3','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'4','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'5','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'6','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'7','|', '9823', '9823', '9823', '9823', '9823', '9823', '9823', '9823'},
        {'8','|', '9820', '9822', '9821', '9819', '9818', '9821', '9822', '9820'},
			}*/

				/*{'#', '|', '1', '2', '3', '4', '5', '6', '7', '8'},
       	{'-', '+','-', '-', '-', '-', '-', '-', '-', '-'},
        {'1','|', 'r', 'h', 'e', 'q', 'k', 'e', 'h', 'r'},
        {'2','|', 'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
        {'3','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'4','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'5','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'6','|', '_', '_', '_', '_', '_', '_', '_', '_'},
        {'7','|', 'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
        {'8','|', 'R', 'H', 'E', 'Q', 'K', 'E', 'H', 'R'}}*/

	for {
		PrintField(field)
		var x, y uint
		fmt.Scan(&x, &y)
		fmt.Println(x, y)
		fmt.Print("\033[H\033[2J")
	}
}
