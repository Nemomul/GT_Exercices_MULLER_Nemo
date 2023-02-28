package main

import (
	"fmt"
)

func choix_num_allumette() int {
	var num_allumette int
	fmt.Println("Choisissez un nombre d'allumette superieur ou égal à 4 : ")
	fmt.Scan(&num_allumette)
	for num_allumette < 4 {
		fmt.Println("Choisissez un nombre d'allumette superieur ou égal à 4 : ")
		fmt.Scan(&num_allumette)
	}
	return num_allumette
}

func affiche_allumette(num_allumette int) {
	fmt.Println("Il reste ", num_allumette, " allumettes.")
}

func choix_num_allumette_down() int {
	var num_allumette_down int
	fmt.Println("Choisissez un nombre d'allumette a retirer (entre 1 et 3) : ")
	fmt.Scan(&num_allumette_down)
	for num_allumette_down < 1 || num_allumette_down > 3 {
		fmt.Println("Choisissez un nombre d'allumette a retirer (entre 1 et 3) : ")
		fmt.Scan(&num_allumette_down)
	}
	return num_allumette_down
}

func defaite(num_allumette int) bool {
	if num_allumette == 0 {
		fmt.Println("Vous avez perdu ! Dommage.")
		return true
	}
	return false
}

func compteur_allumette(num_allumette int, num_allumette_down int) int {
	num_allumette -= num_allumette_down
	if num_allumette < 0 {
		num_allumette = 0
	}
	return num_allumette
}

func joue() {
	num_allumette := choix_num_allumette()
	for !defaite(num_allumette) {
		affiche_allumette(num_allumette)
		num_allumette_down := choix_num_allumette_down()
		num_allumette = compteur_allumette(num_allumette, num_allumette_down)
	}
}

func main() {
	joue()
}
