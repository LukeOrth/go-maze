package main

import (
    "fmt"

	"github.com/LukeOrth/go-maze/pkg/maze"
)

func main() {
    maze := maze.NewMaze(10, 10, 4)
    //maze.Svg()
    fmt.Printf("%s\n",maze.Json())
}
