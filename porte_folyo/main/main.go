package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

func main() {
    // Définir le répertoire racine pour le contenu statique
    fs := http.FileServer(http.Dir("static"))

    // Définir le point de terminaison pour le contenu statique
    http.Handle("/", fs)

    // Démarrer le serveur Web
    fmt.Println("Le serveur Web est lancé sur le port http://localhost:8080/home.html")
    http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	// exécuter le script Python
	cmd := exec.Command("python", "static/chrono.py")
	output, _ := cmd.Output()
	// envoyer le résultat à la page HTML
	fmt.Fprintf(w, "Résultat du script Python: %s", output)
}