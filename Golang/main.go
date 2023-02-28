package main

import (
    "fmt"
)

func main() {
    var n int
	for n<4{ fmt.Println("Combien d'allumettes voulez-vous utiliser ?")
    fmt.Scanln(&n) }
    // Vérifie que le nombre d'allumettes est supérieur ou égal à 4
    // if n < 4 {
    //     fmt.Println("Le nombre d'allumettes doit être supérieur ou égal à 4")
    //     return
    // }

    // Boucle de jeu
    for n > 0 {
        fmt.Printf("Il reste %d allumettes. Combien d'allumettes prenez-vous (1-3) ? ", n)
        var choix int
        fmt.Scanln(&choix)

        // Vérifie que le choix est valide
        if choix < 1 || choix > 3 {
            fmt.Println("Choix invalide, veuillez entrer un nombre entre 1 et 3")
            continue
        }else if choix>n{
			fmt.Println("reduisez votre votre choix")
			continue
		}

        // Retire les allumettes choisies
        n =n- choix

        // Vérifie si le jeu est terminé
        if n == 0 {
            fmt.Println("Vous avez perdu !")
            break
        }

        // Laisse l'autre joueur jouer
        fmt.Println("C'est à l'autre joueur de jouer.")
    }
}
























/*package main

import (
	"fmt"
	"math/rand"
)

func jeuxAllumettes() {
	RandomIntegerwithinRange := rand.Intn(3-1) + 1 // range is min to max

	var (
		allumette uint
		joueur    uint
		prise	uint
		machine   = uint(RandomIntegerwithinRange)
	)

	for allumette < 4 {
		fmt.Print("Entrez le nombres d'allumettes de depart (Doit etre superieur a 4)")
		fmt.Scan(&allumette)
	}
	print("Le jeu va commencer...\n")
	for i:=allumette; i > 0; i=allumette {
		for prise<1 || prise>3{fmt.Print("Combien d'allumettes prenez vous? (Entre 1 et 3)")
		fmt.Scan(&prise)}
		joueur=prise
		allumette=allumette-joueur
		if allumette != 0 {
			fmt.Print("Il en reste", allumette, "\n C'est au tour de la machine\n La machine joue...\n")
			allumette = allumette - machine
			if allumette != 0 {
				fmt.Print("Il en reste", allumette, "\n")
				
			} else if allumette == 0 {
				fmt.Print("La machine a pris la derniere allumette vous avez donc gagne\n")
				break
			}

		} else if allumette == 0 {
			fmt.Print("Vous avez pris la derniere allumette\n Vous avez perdu\n")
			break
		}
	}	
}		

func main() {

	jeuxAllumettes()

}

*/
