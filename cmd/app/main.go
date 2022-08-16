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
    maze := InitMaze(5, 5, 1)
    fmt.Println(maze)
}

func InitMaze(cols int, rows int, cellSize int) *Maze {
    maze := &Maze{ Cells: make([]*Cell, cols * rows), CellSize: cellSize, Cols: cols, Rows: rows  }
   
    for x := 0; x < cols; x++ {
        for y := 0; y < cols; y++ {
            var cell Cell
            cell = *maze.Cells[0]
            fmt.Println(cell)
            cell.X = x
            cell.Y = y
            cell.Top = true
            cell.Right = true
            cell.Bottom = true
            cell.Left = true
            cell.Visited = false
            cell.Current = false
        }
    }
    return maze
}
