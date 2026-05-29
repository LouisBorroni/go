package main

import "fmt"

type Personne struct {
	Prenom, Nom string
	Age         int
	Email       string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans, %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue, Ville, CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("--- Employé ---\n%s\nPoste : %s\nSalaire : %.2f€\nAdresse : %s",
		e.Presentation(), e.Poste, e.Salaire, e.Format())
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (et Etudiant) MentionObtenue() string {
	switch {
	case et.Moyenne >= 16:
		return "Très Bien"
	case et.Moyenne >= 14:
		return "Bien"
	case et.Moyenne >= 12:
		return "Assez Bien"
	default:
		return "Passable"
	}
}

func (et Etudiant) FicheEtudiant() string {
	return fmt.Sprintf("--- Étudiant ---\n%s\nPromo : %s\nMoyenne : %.2f — %s",
		et.Presentation(), et.Promo, et.Moyenne, et.MentionObtenue())
}

func main() {
	employes := []Employe{
		{Personne{"Alice", "Martin", 32, "alice@entreprise.fr"}, Adresse{"12 rue de la Paix", "Paris", "75001"}, "Développeuse", 3500},
		{Personne{"Marc", "Dupont", 45, "marc@entreprise.fr"}, Adresse{"8 avenue Foch", "Lyon", "69006"}, "Chef de projet", 4800},
	}
	etudiants := []Etudiant{
		{Personne{"Laura", "Petit", 20, "laura@ecole.fr"}, "BUT Info 2025", 17.5},
		{Personne{"Tom", "Bernard", 21, "tom@ecole.fr"}, "BUT Info 2025", 11.8},
	}

	employes = append(employes, Employe{
		Personne: Personne{"Sara", "Lopes", 28, "sara@entreprise.fr"},
		Adresse:  Adresse{"3 rue Voltaire", "Bordeaux", "33000"},
		Poste:    "Designer",
		Salaire:  3100,
	})

	for i := range employes {
		employes[i].AugmenterSalaire(10)
	}

	for _, e := range employes {
		fmt.Println(e.FicheEmploye())
		fmt.Println()
	}

	for _, et := range etudiants {
		fmt.Println(et.FicheEtudiant())
		fmt.Println()
	}

	annuaire := make(map[string]string)
	for _, e := range employes {
		annuaire[e.Email] = e.NomComplet()
	}
	for _, et := range etudiants {
		annuaire[et.Email] = et.NomComplet()
	}

	recherches := []string{"alice@entreprise.fr", "inconnu@mail.fr"}
	for _, email := range recherches {
		if nom, ok := annuaire[email]; ok {
			fmt.Printf("Contact trouvé : %s (%s)\n", nom, email)
		} else {
			fmt.Printf("Contact introuvable : %s\n", email)
		}
	}
}
