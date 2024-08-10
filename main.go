package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)


const (
  type_sumbols = "letters"
  Reset = "\033[0m"
  // цвета, которые будут использоваться при выводе в консоль
  Error = "\033[31m"
  Info = "\033[36m"
)

var number_motion int = 1
var min_index_white_chess_figure uint64
var max_index_white_chess_figure uint64
var min_index_black_chess_figure uint64
var max_index_black_chess_figure uint64
var white_king, white_queen, white_rook, white_elephant, white_horse, white_pawns rune
var black_king, black_queen, black_rook, black_elephant, black_horse, black_pawns rune
var empty_cell rune = '_'
var clear map[string]func() //create a map for storing clear funcs


func init() {
    clear = make(map[string]func())
    clear["linux"] = func() { 
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}


func ClearConsole() {
    value, ok := clear[runtime.GOOS]
    if ok { 
        value()
    } else {
        panic("Your platform is unsupported!")
    }
}


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
func IsWhiteFigure(field [10][10]rune, x int, y int) bool {
  if uint64(field[y][x]) > min_index_white_chess_figure && uint64(field[y][x]) < max_index_white_chess_figure {
    return true
  }
  return false
}


// Проверяет что находиться в клетке с координатами (end_x, end_y)
func WhatIs(field [10][10]rune, start_x int, start_y int, end_x int, end_y int) string {
  if end_x < 2 || end_x > 9 || end_y < 2 || end_y > 9 {
    return "out_field"
  }

  if field[end_y][end_x] == empty_cell {
    return "empty"
  } else if IsWhiteFigure(field, start_x, start_y) == IsWhiteFigure(field, end_x, end_y) {
    return "ally"
  }
  return "enemy"
}


// Передвигает фигуру в конечную клетку
func DoMotion(field *[10][10]rune, start_x int, start_y int, end_x int, end_y int){
  field[end_y][end_x] = field[start_y][start_x]
  field[start_y][start_x] = empty_cell
}


// Проверяет может ли пешка сходить из (start_x; start_y) в клетку (end_x; end_y)
func IsValidPawnMotion(field [10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
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


// Проверяет может ли король сходить из (start_x; start_y) в клетку (end_x; end_y)
func IsValidKingMotion(field [10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
  if abs(int64(start_x) - int64(end_x)) > 1 || abs(int64(start_y) - int64(end_y)) > 1 {
    return false
  }
  if WhatIs(field, start_x, start_y, end_x, end_y) == "ally" || WhatIs(field, start_x, start_y, end_x, end_y) == "out_field" {
    return false
  }
  return true
}


// Проверяет может ли конь сходить из (start_x; start_y) в клетку (end_x; end_y)
func IsValidHorseMotion(field [10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
	if WhatIs(field, start_x, start_y, end_x, end_y) == "ally" || WhatIs(field, start_x, start_y, end_x, end_y) == "out_field" {
		return false
	}
	if abs(int64(start_x) - int64(end_x)) == 2 && abs(int64(start_y) - int64(end_y)) == 1 {
		return true
	} 
	if abs(int64(start_x) - int64(end_x)) == 1 && abs(int64(start_y) - int64(end_y)) == 2 {
		return true
	}
	return false
}


// Проверяет может ли тура сходить из (start_x; start_y) в клетку (end_x; end_y)
func IsValidRookMotion(field [10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
  if WhatIs(field, start_x, start_y, end_x, end_y) == "ally" || WhatIs(field, start_x, start_y, end_x, end_y) == "out_field" {
    return false
  }
  switch {
  case start_x == end_x && start_y > end_y:
    for i := start_y - 1; i > end_y; i-- {
      if field[i][start_x] != empty_cell {
        return false
      }
    }
  case start_x == end_x && start_y < end_y:
    for i := start_y + 1; i < end_y; i++ {
      if field[i][start_x] != empty_cell {
        return false
      }
    }
  case start_x > end_x && start_y == end_y:
    for i := start_x - 1; i > end_x; i-- {
      if field[start_y][i] != empty_cell {
        return false
      }
    }
  case start_x < end_x && start_y == end_y:
    for i := start_x + 1; i < end_x; i++ {
      if field[start_y][i] != empty_cell {
        return false
      }
    }
  default:
    return false
  }
  return true
}


// Проверяет может ли слон сходить из (start_x; start_y) в клетку (end_x; end_y)
func IsValidElephantMotion(field [10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
  if WhatIs(field, start_x, start_y, end_x, end_y) == "ally" || WhatIs(field, start_x, start_y, end_x, end_y) == "out_field" {
    return false
  }
  fmt.Println(WhatIs(field, start_x, start_y, end_x, end_y))
  if abs(int64(start_x) - int64(end_x)) != abs(int64(start_y) - int64(end_y)) {
    return false
  }

  var i int = 1
  switch {
  case start_x > end_x && start_y > end_y:
    // вверх влево
    for ; start_y - i > end_y; i++ {
      if field[start_y - i][start_x - i] != empty_cell {
        return false
      }
    }
  case start_x < end_x && start_y > end_y:
    // вверх вправо
    for ; start_y - i > end_y; i++ {
      if field[start_y - i][start_x + i] != empty_cell {
        return false
      }
    }
  case start_x > end_x && start_y < end_y:
    // вниз влево
    for ; start_y + i < end_x; i++ {
      if field[start_y + i][start_x - i] != empty_cell {
        return false
      }
    }
  case start_x < end_x && start_y < end_y:
    // вниз вправо
    for ; start_y + i < end_x; i++ {
      if field[start_y + i][start_x + i] != empty_cell {
        return false
      }
    }
  default:
    return false
  }
  return true
}


// Проверяет может ли королева сходить из (start_x; start_y) в клетку (end_x; end_y)
func IsValidQueenMotion(field [10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
  if start_x == end_x || start_y == end_y {
    return IsValidRookMotion(field, start_x, start_y, end_x, end_y)
  } else {
    return IsValidElephantMotion(field, start_x, start_y, end_x, end_y)
  }
}


// Распределяет действия в зависимости от фигуры, находящейся в клетке (x;y). Если ход был совершён, то возвращает true
func DistributionActions(field *[10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
  switch field[start_y][start_x] {
  case white_pawns:
    if IsValidPawnMotion(*field, start_x, start_y, end_x, end_y){
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }
  case black_pawns:
    if IsValidPawnMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }

  case white_king:
    if IsValidKingMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }
  case black_king:
    if IsValidKingMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }

	case white_horse:
    if IsValidHorseMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }
	case black_horse:
    if IsValidHorseMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }

  case white_rook:
    if IsValidRookMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }
  case black_rook:
    if IsValidRookMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }

  case white_elephant:
    if IsValidElephantMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }
  case black_elephant:
    if IsValidElephantMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }

  case white_queen:
    if IsValidQueenMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }
  case black_queen:
    if IsValidQueenMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }
  }
  return false
}


func main() {
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

  PrintField(field)
  for {
    var start_x, start_y, end_x, end_y int
  
    if number_motion % 2 == 1 {
      fmt.Println("Ход белых")
    } else {
      fmt.Println("Ход чёрных")
    }
    fmt.Printf("Введите координаты шахматной фигуры, %sкоторой будуте ходить%s \n", Info, Reset)

    fmt.Scan(&start_x, &start_y)
    if start_x > 8 || start_x < 1 || start_y > 8 || start_y < 1 || field[start_y+1][start_x+1] == empty_cell {
      fmt.Println(Error, "Ошибка введёных данных", Reset)
      continue
    }
    if (number_motion % 2 == 0 && IsWhiteFigure(field, start_x+1, start_y+1)) || (number_motion % 2 == 1 && IsWhiteFigure(field, start_x+1, start_y+1) == false) {
      fmt.Println(Error, "Сейчас не ваш ход", Reset)
      continue
    }

    fmt.Printf("Вы выбрали фигуру - %c \n", field[start_y+1][start_x+1])
    fmt.Printf("Введите координаты клетки, %sв которую собираетесь ходить%s \n", Info, Reset)
    fmt.Scan(&end_x, &end_y)

    if WhatIs(field, start_x+1, start_y+1, end_x+1, end_y+1) == "out_field" || WhatIs(field, start_x+1, start_y+1, end_x+1, end_y+1) == "ally" {
      fmt.Println(Error, "В эту клетку нельзя сходить", Reset)
      continue
    }
    if DistributionActions(&field, start_x+1, start_y+1, end_x+1, end_y+1) {
      // если игрок сделал ход, то передаём ход другому
      number_motion++
      ClearConsole()
      PrintField(field)
    } else {
      fmt.Println(Error, "Ошибка введёных данных", Reset)
    }
  }
}
