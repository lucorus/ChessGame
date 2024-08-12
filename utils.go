package main

import (
	"os"
	"os/exec"
	"runtime"
)

const (
  Reset = "\033[0m"
)

var number_motion int = 1
var min_index_white_chess_figure uint64
var max_index_white_chess_figure uint64

var WhiteKing, WhiteQueen, WhiteRook, WhiteElephant, WhiteHorse, WhitePawns rune
var BlackKing, BlackQueen, BlackRook, BlackElephant, BlackHorse, BlackPawns rune
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
