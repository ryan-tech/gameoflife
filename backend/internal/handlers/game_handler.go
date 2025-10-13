package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/ryan-tech/gameoflife/internal/models"
	"log"
)


func Step(w http.ResponseWriter, r *http.Request) {
	// read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	// unmarshal the body into the game state
	var gridBody models.RequestBody
	err = json.Unmarshal(body, &gridBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	printGrid(gridBody)
	response := map[string]string{"message": "Game state updated"}
	json.NewEncoder(w).Encode(response)

}

func printGrid(gridBody models.RequestBody) {
	output := "Printing Grid:\n"
	for _, row := range gridBody.Grid {
		output += fmt.Sprintf("%d", row.Row) + ":\t" + " "	
		for _, cell := range row.Cells {
			output += fmt.Sprintf("%t", cell.Alive) + " "
			output += "\t"
		}
		output += "\n"
	}
	log.Println(output)
}

	// // get the game state from the request
	// gameState := r.Context().Value("gameState").(GameState)
	// // step the game
	// gameState.Step()
	// // return the game state
	// json.NewEncoder(w).Encode(gameState)
