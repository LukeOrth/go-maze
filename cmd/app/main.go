package main

import (

	"github.com/LukeOrth/go-maze/pkg/maze"
)

func main() {
    maze := maze.NewMaze(50, 50, 4)
    maze.Svg()
    //fmt.Printf("%s\n",maze.Json())
}
