package main

import (
    "github.com/LukeOrth/go-maze/pkg/maze"
)

func main() {
    maze := maze.NewMaze(5000, 5000, 50)
    maze.MazeToSvg()
}
