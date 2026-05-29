package main

import (
	"fmt"
	"strings"
)

const (
	MenuQuitter = iota
	MenuAjouter
	MenuChercher
	MenuSoldes
	MenuVendre
	MenuRapport
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

func (p Produit) String() string {
	return fmt.Sprintf("[%d] %s %s — %.2f€ | Stock: %d | Catégorie: %s",
		p.ID, p.Marque, p.Nom, p.Prix, p.Stock, p.Categorie)
}

type Catalogue struct {
	produits  []Produit
	prochainID int
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	c.prochainID++
	p.ID = c.prochainID
	c.produits = append(c.produits, p)
	return nil
}

func (c *Catalogue) TrouverParID(id int) (*Produit, error) {
	for i := range c.produits {
		if c.produits[i].ID == id {
			return &c.produits[i], nil
		}
	}
	return nil, fmt.Errorf("produit ID %d introuvable", id)
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultat []Produit
	for _, p := range c.produits {
		if strings.EqualFold(p.Categorie, cat) {
			resultat = append(resultat, p)
		}
	}
	return resultat
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nb := 0
	for i := range c.produits {
		if strings.EqualFold(c.produits[i].Categorie, categorie) && c.produits[i].Actif {
			c.produits[i].Prix -= c.produits[i].Prix * pct / 100
			nb++
		}
	}
	return nb
}

func (c *Catalogue) Vendre(id int, qte int) error {
	p, err := c.TrouverParID(id)
	if err != nil {
		return err
	}
	if p.Stock < qte {
		return fmt.Errorf("stock insuffisant : %d disponible(s), %d demandé(s)", p.Stock, qte)
	}
	p.Stock -= qte
	return nil
}

func (c Catalogue) Rapport() string {
	valeur := 0.0
	for _, p := range c.produits {
		valeur += p.Prix * float64(p.Stock)
	}
	return fmt.Sprintf("Catalogue : %d produits | Valeur totale du stock : %.2f€", len(c.produits), valeur)
}

func main() {
	cat := Catalogue{}

	produits := []Produit{
		{1, "iPhone 15", "Apple", 999.99, 10, "Smartphone", true},
		{2, "MacBook Pro", "Apple", 2499.99, 5, "Ordinateur", true},
		{3, "Galaxy S24", "Samsung", 849.99, 8, "Smartphone", true},
		{4, "ThinkPad X1", "Lenovo", 1599.99, 3, "Ordinateur", true},
		{5, "AirPods Pro", "Apple", 279.99, 15, "Audio", true},
	}
	for _, p := range produits {
		cat.AjouterProduit(p)
	}

	ajouter := func() {
		var stock int
		var nom, marque, categorie string
		var prix float64
		fmt.Print("Nom: ")
		fmt.Scan(&nom)
		fmt.Print("Marque: ")
		fmt.Scan(&marque)
		fmt.Print("Prix: ")
		fmt.Scan(&prix)
		fmt.Print("Stock: ")
		fmt.Scan(&stock)
		fmt.Print("Catégorie: ")
		fmt.Scan(&categorie)
		cat.AjouterProduit(Produit{Nom: nom, Marque: marque, Prix: prix, Stock: stock, Categorie: categorie, Actif: true})
		fmt.Printf("Produit ajouté avec l'ID %d.\n", cat.prochainID)
	}

	chercher := func() {
		fmt.Print("Rechercher par [1] ID  [2] Catégorie : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			var id int
			fmt.Print("ID : ")
			fmt.Scan(&id)
			p, err := cat.TrouverParID(id)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println(p)
			}
		case 2:
			var cat2 string
			fmt.Print("Catégorie : ")
			fmt.Scan(&cat2)
			produits := cat.TrouverParCategorie(cat2)
			if len(produits) == 0 {
				fmt.Println("Aucun produit dans cette catégorie.")
			} else {
				for _, p := range produits {
					fmt.Println(p)
				}
			}
		default:
			fmt.Println("Choix invalide.")
		}
	}

	soldes := func() {
		var categorie string
		var pct float64
		fmt.Print("Catégorie: ")
		fmt.Scan(&categorie)
		fmt.Print("Réduction (%): ")
		fmt.Scan(&pct)
		nb := cat.AppliquerReduction(categorie, pct)
		fmt.Printf("%d produit(s) mis en solde.\n", nb)
	}

	vendre := func() {
		var id, qte int
		fmt.Print("ID produit: ")
		fmt.Scan(&id)
		fmt.Print("Quantité: ")
		fmt.Scan(&qte)
		if err := cat.Vendre(id, qte); err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Vente effectuée.")
		}
	}

	rapport := func() {
		fmt.Println(cat.Rapport())
	}

	for {
		fmt.Printf("\n[%d] Ajouter  [%d] Chercher  [%d] Soldes  [%d] Vendre  [%d] Rapport  [%d] Quitter\n",
			MenuAjouter, MenuChercher, MenuSoldes, MenuVendre, MenuRapport, MenuQuitter)

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case MenuQuitter:
			fmt.Println("Au revoir !")
			return
		case MenuAjouter:
			ajouter()
		case MenuChercher:
			chercher()
		case MenuSoldes:
			soldes()
		case MenuVendre:
			vendre()
		case MenuRapport:
			rapport()
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
