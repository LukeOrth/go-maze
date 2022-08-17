package main

import "fmt"

type Cell struct {
    X       int
    Y       int
    Top     bool
    Right   bool
    Bottom  bool
    Left    bool
    Visited bool
    Current bool
}

type Maze struct {
    Cells       []*Cell
    CellSize    int
    Cols        int
    Rows        int
}

func main() {
    maze := InitMaze(5, 3, 1)
    fmt.Println(maze)
}

func InitMaze(cols int, rows int, cellSize int) *Maze {
    maze := &Maze{ Cells: make([]*Cell, cols * rows), CellSize: cellSize, Cols: cols, Rows: rows  }

    for y := 0; y < rows; y++ {
        for x := 0; x < cols; x++ {
            maze.Cells[cols * y + x] = &Cell{ X: x, Y: y, Top: true, Right: true, Bottom: true, Left: true, Visited: false, Current: false }
        }
    }
    return maze
}

func cellAt(maze *Maze, x int, y int) *Cell {
    return maze.Cells[maze.Cols * y + x]
}

/*
func MazeToSvg(maze *Maze) {

}
*/
