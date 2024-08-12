package main

import (
	"fmt"
	"os"
	"time"
)

// Возвращает кол-во очков, которые даёт переданная фигура
func PointsForChessFigure(figure rune) int {
  switch figure {
    case WhitePawns:
      return 1
    case BlackPawns:
      return 1

    case WhiteRook:
      return 5
    case BlackRook:
      return 5

    case WhiteHorse:
      return 3
    case BlackHorse:
      return 3

    case WhiteElephant:
      return 3
    case BlackElephant:
      return 3

    case WhiteQueen:
      return 9
    case BlackQueen:
      return 9

    case WhiteKing:
      return 100
    case BlackKing:
      return 100
    }
  return 0
}


func PrintField(field [10][10]rune) {
  fmt.Printf("Ход номер: %s%d%s", Info, number_motion, Reset)
  if number_motion % 2 == 1 {
    fmt.Printf("\t%s(ход белых)%s\n", Info, Reset)
  } else {
    fmt.Printf("\t %s(ход чёрных)%s\n", Info, Reset)
  }

  white_figures_points := 0
  black_figures_points := 0

  for i := 0; i < len(field); i++ {
    for j := 0; j < len(field[i]); j++ {
      if IsWhiteFigure(field, j, i) {
        white_figures_points += PointsForChessFigure(field[i][j])
      } else {
        black_figures_points += PointsForChessFigure(field[i][j])
      }
      fmt.Printf("%c ", field[i][j])
    }
    fmt.Println()
  }
  fmt.Printf("Кол-во баллов у белых: %s%d%s\n", Info, white_figures_points-100, Reset)
  fmt.Printf("Кол-во баллов у чёрных: %s%d%s\n\n", Info, black_figures_points-100, Reset)
  if white_figures_points < 100 {
    fmt.Printf("%sКоманда чёрных выйграла!!!%s", Win, Reset)
    time.Sleep(10 * time.Second)
    os.Exit(0)
  }
  if black_figures_points < 100 {
    fmt.Printf("%sКоманда белых выйграла!!!%s", Win, Reset)
    time.Sleep(10 * time.Second)
    os.Exit(0)
  }
}


// Когда игрок доходит до конца поля пешкой, он может сменить пешку на другую фигуру
func ChangeFigure(field *[10][10]rune, x int, y int) {
  figure_changed := true
  fmt.Printf("%sПоздравляем, вы дошли до конца поля!%s \nВыберите шахматную фигуру: %s(h - конь, e - слон, r - тура, q - королева)%s\n", Win, Reset, Info, Reset)
  for ; figure_changed; {
    var figure string
    fmt.Scan(&figure)
    switch figure {
    case "h":
      if IsWhiteFigure(*field, x+1, y+1) {
        field[y+1][x+1] = WhiteHorse
      } else {
        field[y+1][x+1] = BlackHorse
      }
      figure_changed = false
    case "e":
      if IsWhiteFigure(*field, x+1, y+1) {
        field[y+1][x+1] = WhiteElephant
      } else {
        field[y+1][x+1] = BlackElephant
      }
      figure_changed = false
    case "r":
      if IsWhiteFigure(*field, x+1, y+1) {
        field[y+1][x+1] = WhiteRook
      } else {
        field[y+1][x+1] = BlackRook
      }
      figure_changed = false
    case "q":
      if IsWhiteFigure(*field, x+1, y+1) {
        field[y+1][x+1] = WhiteQueen
      } else {
        field[y+1][x+1] = BlackQueen
      }
      figure_changed = false
    default:
      fmt.Printf("%sВы ввели некорректные данные!%s\n", Error, Reset)
    }
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


// Распределяет действия в зависимости от фигуры, находящейся в клетке (x;y). Если ход был совершён, то возвращает true
func DistributionActions(field *[10][10]rune, start_x int, start_y int, end_x int, end_y int) bool {
  switch field[start_y][start_x] {
  case WhitePawns:
    if IsValidPawnMotion(*field, start_x, start_y, end_x, end_y){
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }
  case BlackPawns:
    if IsValidPawnMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }

  case WhiteKing:
    if IsValidKingMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }
  case BlackKing:
    if IsValidKingMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }

	case WhiteHorse:
    if IsValidHorseMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }
	case BlackHorse:
    if IsValidHorseMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
			return true
    }

  case WhiteRook:
    if IsValidRookMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }
  case BlackRook:
    if IsValidRookMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }

  case WhiteElephant:
    if IsValidElephantMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }
  case BlackElephant:
    if IsValidElephantMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }

  case WhiteQueen:
    if IsValidQueenMotion(*field, start_x, start_y, end_x, end_y) {
      DoMotion(field, start_x, start_y, end_x, end_y)
      return true
    }
  case BlackQueen:
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
    {'1','|', BlackRook, BlackHorse, BlackElephant, BlackQueen, BlackKing, BlackElephant, BlackHorse, BlackRook},
    {'2','|', BlackPawns, BlackPawns, BlackPawns, BlackPawns, BlackPawns, BlackPawns, BlackPawns, BlackPawns},
    {'3','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'4','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'5','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'6','|', empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell, empty_cell},
    {'7','|', WhitePawns, WhitePawns, WhitePawns, WhitePawns, WhitePawns, WhitePawns, WhitePawns, WhitePawns},
    {'8','|', WhiteRook, WhiteHorse, WhiteElephant, WhiteQueen, WhiteKing, WhiteElephant, WhiteHorse, WhiteRook}}

  PrintField(field)
  for {
    var start_x, start_y, end_x, end_y int
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
      if (field[end_y+1][end_x+1] == WhitePawns && end_y == 1) || (field[end_y+1][end_x+1] == BlackPawns && end_y == 8) {
        // игрок дошёл до конца поля пешкой
        ChangeFigure(&field, end_x, end_y)
      }
      number_motion++
      ClearConsole()
      PrintField(field)
    } else {
      fmt.Println(Error, "Ошибка введёных данных", Reset)
    }
  }
}
