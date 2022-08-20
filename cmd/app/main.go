package main

import (
    "github.com/LukeOrth/go-maze/pkg/maze"
)

func main() {
    maze := maze.NewMaze(1000, 1000, 2)
    maze.MazeToSvg()
    //maze.MazeToSvg()
}
