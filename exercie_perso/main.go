package main

import "fmt"

const (
	Debutant = iota
	Intermediaire
	Avance
	Expert
)

func main() {
	joueurs := []struct {
		nom    string
		niveau int
	}{
		{"Alice", Expert},
		{"Bob", Avance},
		{"Clara", Intermediaire},
		{"David", Debutant},
	}

	for _, j := range joueurs {
		fmt.Printf("%s (niveau %d) a accès à : ", j.nom, j.niveau)

		switch j.niveau {
		case Expert:
			fmt.Print("[mode ranked] ")
			fallthrough
		case Avance:
			fmt.Print("[tournois] ")
			fallthrough
		case Intermediaire:
			fmt.Print("[classement] ")
			fallthrough
		case Debutant:
			fmt.Print("[partie normale]")
		}

		fmt.Println()
	}
}

/*
  ENONCE
  Contexte : Un jeu en ligne attribue des accès différents selon le niveau du joueur.

  1. Déclarez un groupe de constantes avec iota représentant 4 niveaux : Debutant, Intermediaire, Avance, Expert
  2. Créez une slice contenant au moins 4 joueurs, chacun avec un nom (string) et un niveau (int)
  3. Parcourez la slice avec une boucle for range et affichez pour chaque joueur ses accès
  4. Utilisez un switch avec fallthrough pour afficher les accès de façon cumulative :
     - Debutant      -> partie normale
     - Intermediaire -> classement + tout ce qui précède
     - Avance        -> tournois + tout ce qui précède
     - Expert        -> mode ranked + tout ce qui précède

*/
