package main

import (
	"fmt"
	"net/http"
	"strconv"

	"distro-hub/internal/distro"
	"distro-hub/view"
)

func filterDistros(category string, distros []distro.Distro) []distro.Distro {
	if category == "" {
		return distros
	}
	newDistros := []distro.Distro{}

	for _, distro := range distros {
		for _, cat := range distro.Categories {
			if cat == category {
				newDistros = append(newDistros, distro)
				break
			}
		}
	}

	return newDistros
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./view/public/"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", fs))

	var distros, err = distro.LoadDistrosFromFile("./distros.json")
	if err != nil {
		fmt.Println(err)
	}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")

		filteredDistros := filterDistros(category, distros)
		indexProps := map[string]any{
			"Category": category,
			"Distros":  filteredDistros,
		}

		if category != "" {
			if r.Header.Get("HX-Request") == "true" {
				err := view.ChildRender(w, "index", "category", indexProps)
				if err != nil {
					fmt.Println(err)
				}
				return
			}
			err := view.Render(w, "index", indexProps)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err := view.Render(w, "index", indexProps)
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("GET /packages", func(w http.ResponseWriter, r *http.Request) {
		view.Render(w, "packages", nil)
	})

	mux.HandleFunc("GET /distros/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		ID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		if ID < 0 || ID > len(distros) {
			http.Error(w, "Distro not found", http.StatusNotFound)
			return
		}
		var distro distro.Distro
		for i := 0; i < len(distros); i++ {
			if distros[i].ID == ID {
				distro = distros[i]
			}
		}

		view.Render(w, "distro", distro)
	})

	http.ListenAndServe("0.0.0.0:6969", mux)
}
