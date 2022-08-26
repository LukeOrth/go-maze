package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/LukeOrth/go-maze/maze"
	"github.com/gorilla/mux"
)

func generateMazeGET(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    q := r.URL.Query()
    columns, _ := strconv.Atoi(q.Get("columns"))
    rows, _ := strconv.Atoi(q.Get("rows"))
    scale, _ := strconv.Atoi(q.Get("scale"))
    maze := maze.NewMaze(columns, rows, scale)
    json.NewEncoder(w).Encode(maze)
}

func mazeToPngPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var maze maze.Maze
	_ = json.NewDecoder(r.Body).Decode(&maze)
    maze.Png(w)
}

func mazeToSvg(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/svg")
    q := r.URL.Query()
    columns, _ := strconv.Atoi(q.Get("columns"))
    rows, _ := strconv.Atoi(q.Get("rows"))
    scale, _ := strconv.Atoi(q.Get("scale"))
    maze := maze.NewMaze(columns, rows, scale)
    maze.Svg(w)
}

func Run() {
    r := mux.NewRouter()

    r.HandleFunc("/maze", generateMazeGET).Methods("GET")
    r.HandleFunc("/maze/svg", mazeToSvg).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", r))
}
