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
    query := r.URL.Query()
    columns, _ := strconv.Atoi(query.Get("columns"))
    rows, _ := strconv.Atoi(query.Get("rows"))
    scale, _ := strconv.Atoi(query.Get("scale"))

    maze := maze.NewMaze(columns, rows, scale)
    json.NewEncoder(w).Encode(maze)
}

func Run() {
    r := mux.NewRouter()

    r.HandleFunc("/maze", generateMazeGET).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", r))
}
