package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Pokémon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {

	http.HandleFunc("/pokemon/{id}", func(w http.ResponseWriter, r *http.Request) {
		id_str := r.PathValue("id")
		id, err := strconv.Atoi(id_str)
		if err != nil || id < 0 {
			http.Error(w, "Invalid id, not an int", http.StatusBadRequest)
			return
		}

		// TODO: Rate limit http requests
		resp, err := http.Get(`https://pokeapi.co/api/v2/pokemon-species/` + id_str)
		if err != nil {
			http.Error(w, "Failed to fetch upstream data from pokeapi", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read upstream data from pokeapi", http.StatusInternalServerError)
			return
		}
		var apiPokémon APIPokémon
		err = json.Unmarshal(body, &apiPokémon)
		if err != nil {
			http.Error(w, "Failed to parse upstream data from pokeapi", http.StatusInternalServerError)
			return
		}

		pokemon := Pokémon{
			ID:   apiPokémon.ID,
			Name: apiPokémon.Name,
		}

		for _, flavorText := range apiPokémon.FlavorTextEntries {
			if flavorText.Language.Name == "en" {
				pokemon.Description = flavorText.FlavorText
				break
			}
		}

		err = json.NewEncoder(w).Encode(pokemon)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Println(err)
			return
		}

	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
