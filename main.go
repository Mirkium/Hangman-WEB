package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Choix(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		languageStr := r.FormValue("language")
		language, err := strconv.Atoi(languageStr)
		if err != nil {
			fmt.Println("Erreur de conversion de la langue en entier.")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if language == 2 {
			ReadFileContent("langue/English.txt")
		}

		difficulty := r.FormValue("difficulty")

		fmt.Printf("Langue sélectionnée : %d\n", language)
		fmt.Printf("Difficulté sélectionnée : %s\n", difficulty)

		switch difficulty {
		case "easy":
			http.ServeFile(w, r, "Hangman_facile.html")
		case "normal":
			http.ServeFile(w, r, "Hangman_normal.html")
		case "hard":
			http.ServeFile(w, r, "Hangman_difficile.html")
		default:
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		return
	}

	http.ServeFile(w, r, "page_menu.html")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "page_menu.html")
	})

	http.HandleFunc("/submit", Choix)

	http.HandleFunc("/Hangman_facile.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Hangman_facile.html")
	})

	http.HandleFunc("/Hangman_normal.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Hangman_normal.html")
	})

	http.HandleFunc("/Hangman_difficile.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Hangman_difficile.html")
	})

	RootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(RootDoc + "/asset/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	http.ListenAndServe("localhost: 8080", nil)
}

var test []string

func ReadFileContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	}
	fmt.Println("Contenu du fichier :")
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		test = append(test, lines[i])
	}
	mot_random := rand.Intn(len(test))
	return test[mot_random]
}
