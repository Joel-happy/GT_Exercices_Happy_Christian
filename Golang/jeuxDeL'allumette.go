package main

import (
	"fmt"
)

func jeuxAllumettes() {
	var (
		allumette int
		choix     int
	)

	for allumette < 4 {
		fmt.Println("Combien d'allumettes voulez-vous utiliser ?")
		fmt.Scanln(&allumette)
	}
	// jeu
	for allumette > 0 {
		fmt.Printf("Il reste %d allumettes. Combien d'allumettes prenez-vous (1-3) ? ", allumette)
		fmt.Scanln(&choix)
		// Vérifie que le choix est valide
		if choix < 1 || choix > 3 {
			fmt.Println("Choix invalide, veuillez entrer un nombre entre 1 et 3")
			continue
		} else if choix > allumette {
			fmt.Println("reduisez votre votre choix")
			continue
		}
		// Comptage d'allumettes
		allumette = allumette - choix
		// condition de fin
		if allumette == 0 {
			fmt.Println("Vous avez perdu !")
			break
		}
		// tour de l'autre joueur
		fmt.Println("C'est à l'autre joueur de jouer.")
	}
}

func main() {
	jeuxAllumettes()
}
