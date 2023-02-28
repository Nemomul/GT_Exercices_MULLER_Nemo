package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func lireTexte(nom_fichier string) (string, error) {
	fichier, err := os.Open(nom_fichier)
	if err != nil {
		return "", err
	}
	defer fichier.Close()
	contenu, err := ioutil.ReadAll(fichier)
	if err != nil {
		return "", err
	}
	texte := string(contenu)
	return texte, nil
}

func lancementLireTexte() {
	texte, err := lireTexte("document.txt")
	if err != nil {
		panic(err)
	}
	println(texte)
}

func AjouterTexte(nom_fichier string, texte_a_ajouter string) error {
	fichier, err := os.OpenFile(nom_fichier, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fichier.Close()
	_, err = fichier.WriteString(texte_a_ajouter)
	if err != nil {
		return err
	}
	return nil
}

func lancementAjouterTexte() {
	err := AjouterTexte("document.txt", "Texte ajouté")
	if err != nil {
		panic(err)
	}
}

func SupprimerTexte(nom_fichier string) error {
	fichier, err := os.OpenFile(nom_fichier, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fichier.Close()
	return nil
}

func lancementSupprimerTexte() {
	err := SupprimerTexte("document.txt")
	if err != nil {
		panic(err)
	}
}

func RemplacerTexte(nom_fichier string, texte_a_remplacer string) error {
	err := SupprimerTexte(nom_fichier)
	if err != nil {
		return err
	}
	err = AjouterTexte(nom_fichier, texte_a_remplacer)
	if err != nil {
		return err
	}
	return nil
}

func lancementRemplacerTexte() {
	err := RemplacerTexte("document.txt", "Texte remplacé")
	if err != nil {
		panic(err)
	}
}

func menu() {
	println("1 - Lire le contenu d'un fichier")
	println("2 - Ajouter du texte")
	println("3 - Supprimer le contenu d'un fichier")
	println("4 - Remplacer le contenu d'un fichier")
	println("5 - Quitter")
}

func lancementMenu() {
	menu()
	var choix int
	println("Entrez votre choix")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		lancementLireTexte()
		lancementMenu()
	case 2:
		lancementAjouterTexte()
		lancementMenu()
	case 3:
		lancementSupprimerTexte()
		lancementMenu()
	case 4:
		lancementRemplacerTexte()
		lancementMenu()
	case 5:
		println("Au revoir")
		os.Exit(0)
	default:
		println("Choix invalide")
	}
}

func main() {
	lancementMenu()
}
