package main

import (
	"fmt"
)

func (p *personnage) Armure() {
	if p.equipement.arme == "Épee en fer" {
		p.armedeg = 8 * 1 / 2
	}
	if p.equipement.arme == "Épee de Midgard" {
		p.armedeg = 8 * 3 / 2
	}
	if p.equipement.arme == "Hache Divine Rhitta" {
		p.armedeg = 8 * 6 / 2
	}
}

func (p *personnage) addArmure(s string) {
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == s && p.inventaire[i] == "Chapeau de l'aventurier" {
			if p.equipement.casque != "aucun" {
				p.addInventory(p.equipement.casque)
				p.pvm -= 15
				p.equipement.casque = "Chapeau de l'aventurier"
				p.RemoveInventory(p.equipement.casque)
				p.pvm += 10
			} else {
				p.equipement.casque = "Chapeau de l'aventurier"
				p.RemoveInventory("Chapeau de l'aventurier")
				p.pvm += 10
			}
			fmt.Println(p.equipement.casque, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Tunique de l'aventurier" {
			if p.equipement.armure != "aucun" {
				p.addInventory(p.equipement.armure)
				p.pvm -= 30
				p.equipement.armure = "Tunique de l'aventurier"
				p.RemoveInventory(p.equipement.armure)
				p.pvm += 25
			} else {
				p.equipement.armure = "Tunique de l'aventurier"
				p.RemoveInventory("Tunique de l'aventurier")
				p.pvm += 25
			}
			fmt.Println(p.equipement.armure, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Bottes de l'aventurier" {
			if p.equipement.bottes != "aucun" {
				p.addInventory(p.equipement.bottes)
				p.pvm -= 20
				p.equipement.bottes = "Bottes de l'aventurier"
				p.RemoveInventory(p.equipement.bottes)
				p.pvm += 15
			} else {
				p.equipement.bottes = "Bottes de l'aventurier"
				p.RemoveInventory("Bottes de l'aventurier")
				p.pvm += 15
			}
			fmt.Println(p.equipement.bottes, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Épee en fer" {
			if p.equipement.arme != "aucune" {
				p.addInventory(p.equipement.arme)
				p.equipement.arme = "Épee en fer"
				p.RemoveInventory(p.equipement.arme)
			} else {
				p.equipement.arme = "Épee en fer"
				p.RemoveInventory("Épee en fer")
			}
			fmt.Println(p.equipement.arme, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Épee de Midgard" {
			if p.equipement.arme != "aucune" {
				p.addInventory(p.equipement.arme)
				p.equipement.arme = "Épee de Midgard"
				p.RemoveInventory(p.equipement.arme)
			} else {
				p.equipement.arme = "Épee de Midgard"
				p.RemoveInventory("Épee de Midgard")
			}
			fmt.Println(p.equipement.arme, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Hache Divine Rhitta" {
			if p.equipement.arme != "aucune" {
				p.addInventory(p.equipement.arme)
				p.equipement.arme = "Hache Divine Rhitta"
				p.RemoveInventory(p.equipement.arme)
			} else {
				p.equipement.arme = "Hache Divine Rhitta"
				p.RemoveInventory("Hache Divine Rhitta")
			}
			fmt.Println(p.equipement.arme, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Casque en fer" {
			if p.equipement.casque != "aucun" {
				p.addInventory(p.equipement.casque)
				p.pvm -= 10
				p.equipement.casque = "Casque en fer"
				p.RemoveInventory(p.equipement.casque)
				p.pvm += 15
			} else {
				p.equipement.casque = "Casque en fer"
				p.RemoveInventory("Casque en fer")
				p.pvm += 15
			}
			fmt.Println(p.equipement.casque, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Armure en fer" {
			if p.equipement.armure != "aucun" {
				p.addInventory(p.equipement.armure)
				p.pvm -= 25
				p.equipement.armure = "Armure en fer"
				p.RemoveInventory(p.equipement.armure)
				p.pvm += 30
			} else {
				p.equipement.armure = "Armure en fer"
				p.RemoveInventory("Armure en fer")
				p.pvm += 30
			}
			fmt.Println(p.equipement.armure, "a bien été ajouté a votre équipement !")
		} else if p.inventaire[i] == s && p.inventaire[i] == "Bottes en fer" {
			if p.equipement.bottes != "aucun" {
				p.addInventory(p.equipement.bottes)
				p.pvm -= 15
				p.equipement.bottes = "Bottes en fer"
				p.RemoveInventory(p.equipement.bottes)
				p.pvm += 20
			} else {
				p.equipement.bottes = "Bottes en fer"
				p.RemoveInventory("Bottes en fer")
				p.pvm += 20
			}
			fmt.Println(p.equipement.bottes, "a bien été ajouté a votre équipement !")
		}
	}
}
