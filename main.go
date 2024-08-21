package main

import (
	"fmt"
	"net/http"

	"groupieTracker/roots"
)

func main() {
	css := http.FileServer(http.Dir("style")) // Serve static files from "Assets" directory
	http.Handle("/style/", http.StripPrefix("/style/", css))

	http.HandleFunc("/", roots.HandleMainPage)
	http.HandleFunc("/details", roots.HandleDetailsPage)

	fmt.Println("\x1b[92m" + " -----------------------------------------------")
	fmt.Println(" |  " + "\033[1m" + "im working, port :" + "\x1b[91m" + " http://localhost:8080" + "\x1b[0m" + "\x1b[92m" + "   |")
	fmt.Println(" -----------------------------------------------" + "\x1b[0m")
	http.ListenAndServe(":8080", nil)
}
