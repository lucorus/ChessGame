package main

import "fmt"

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
    for ; start_y + i < end_y; i++ {
      if field[start_y + i][start_x - i] != empty_cell {
        return false
      }
    }
  case start_x < end_x && start_y < end_y:
    // вниз вправо
    for ; start_y + i < end_y; i++ {
      fmt.Println(i, start_y, start_x, end_x, end_y)
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