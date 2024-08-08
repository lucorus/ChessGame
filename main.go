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


func abs(num int64) int64 {
	if num < 0 {
		return num / (-1)
	}
	return num
}

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
}


func PrintField(field [10][10]rune) {
	for _, line := range field {
			for _, item := range line {
				fmt.Printf("%c ", item)
			}
			fmt.Println()
	}
}


// Возвращает true, если фигура, находящаяся по указанному индексу, - фигура белых
func IsWhiteFigure(field [10][10]rune, x uint, y uint) bool {
	if uint64(field[y][x]) > min_index_white_chess_figure && uint64(field[y][x]) < max_index_white_chess_figure {
		return true
	}
	return false
}


// Проверяет что находиться в клетке с координатами (end_x, end_y)
func WhatIs(field [10][10]rune, start_x uint, start_y uint, end_x uint, end_y uint) string {
	if end_x < 2 || end_x > 9 || end_y < 2 || end_y > 9 {
		fmt.Println("OUT_FIELD")
		return "out_field"
	}

	if field[end_y][end_x] == empty_cell {
		fmt.Println("EMPTY")
		return "empty"
	} else if IsWhiteFigure(field, start_x, start_y) == IsWhiteFigure(field, end_x, end_y) {
		fmt.Println("ALLY")
		return "ally"
	}
	fmt.Println("ENEMY")
	return "enemy"
}


// Передвигает фигуру в конечную клетку
func DoMotion(field [10][10]rune, start_x uint, start_y uint, end_x uint, end_y uint) [10][10]rune {
	field[end_y][end_x] = field[start_y][start_x]
	field[start_y][start_x] = empty_cell
	return field
}


// Проверяет может ли пешка сходит из (start_x; start_y) в клетку (end_x; end_y)
func IsValidPawnMotion(field [10][10]rune, start_x uint, start_y uint, end_x uint, end_y uint) bool {
	if (start_x != end_x && WhatIs(field, start_x, start_y, end_x, end_y) != "enemy") || (start_x == end_x && WhatIs(field, start_x, start_y, end_x, end_y) != "empty") {
		// проверяем валидность движения по x
		return false
	}
	if IsWhiteFigure(field, start_x, start_y) {
		if start_y <= end_y || (start_y - end_y > 2 && start_y == 8) || (start_y - end_y != 1 && start_y != 8) {
			// проверяем валидность движения по y
			return false
		}
	} else {
		if start_y >= end_y || (end_y - start_y > 2 && start_y == 3) || (end_y - start_y != 1 && start_y != 3) {
			// проверяем валидность движения по y
			return false
		}
	}
	return true
}


// Распределяет действия в зависимости от фигуры, находящейся в клетке (x;y)
func DistributionActions(field [10][10]rune, start_x uint, start_y uint, end_x uint, end_y uint) [10][10]rune {
	
	switch field[start_y][start_x] {
	case white_pawns:
		if IsValidPawnMotion(field, start_x, start_y, end_x, end_y){
			fmt.Println("Access")
			field = DoMotion(field, start_x, start_y, end_x, end_y)
		}
	case black_pawns:
		if IsValidPawnMotion(field, start_x, start_y, end_x, end_y) {
			field = DoMotion(field, start_x, start_y, end_x, end_y)
		}
	default:
		fmt.Println("Фигура не найдена")
	}
	return field
}


func main() {
	var type_sumbols string = "letters"
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
    {'8','|', white_rook, white_horse, white_elephant, white_queen, white_king, white_elephant, white_horse, white_rook}}

	for {
		PrintField(field)
		var start_x, start_y, end_x, end_y uint
		fmt.Println("Введите координаты шахматной фигуры, которой будете ходить")
		fmt.Scan(&start_x, &start_y)
		fmt.Println("Введите координаты клетки, в которую собираетесь сходить")
		fmt.Scan(&end_x, &end_y)
		field = DistributionActions(field, start_x+1, start_y+1, end_x+1, end_y+1)
		// fmt.Print("\033[H\033[2J")
	}
}
