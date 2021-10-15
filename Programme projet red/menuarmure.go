package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuArmure() {
	Clear()
	image("armure.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Quelle équipement voulez-vous équipé ? :\n")
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == "Chapeau de l'aventurier" {
			fmt.Println("Chapeau de l'aventurier |'Chapeau'\n")
		}
		if p.inventaire[i] == "Tunique de l'aventurier" {
			fmt.Println("Tunique de l'aventurier |'Tunique'\n")
		}
		if p.inventaire[i] == "Bottes de l'aventurier" {
			fmt.Println("Bottes de l'aventurier |'Bottes'\n")
		}
		if p.inventaire[i] == "Épee en fer" {
			fmt.Println("Épee en fer |'Fer'\n")
		}
		if p.inventaire[i] == "Épee de Midgard" {
			fmt.Println("Épee de Midgard |'Midgard'\n")
		}
		if p.inventaire[i] == "Hache Divine Rhitta" {
			fmt.Println("Hache Divine Rhitta |'Hache'\n")
		}
		if p.inventaire[i] == "Casque en fer" {
			fmt.Println("Casque en fer |'Casque'\n")
		}
		if p.inventaire[i] == "Armure en fer" {
			fmt.Println("Armure en fer |'Armure'\n")
		}
		if p.inventaire[i] == "Bottes en fer" {
			fmt.Println("Bottes en fer |'Botte'\n")
		}
	}
	fmt.Println("Retourner au menu ! | 'Retour'\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num string
	fmt.Scanf("%s\n", &num)
	switch num {
	case "Chapeau":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Chapeau de l'aventurier")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.casque == "Chapeau de l'aventurier" {
				fmt.Println("Vous avez déjà un Chapeau d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Chapeau de l'aventurier")
				p.Armure()
				casque++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Tunique":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Tunique de l'aventurier")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.armure == "Tunique de l'aventurier" {
				fmt.Println("Vous avez déjà une Tunique d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Tunique de l'aventurier")
				p.Armure()
				armure++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Bottes":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Bottes de l'aventurier")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.bottes == "Bottes de l'aventurier" {
				fmt.Println("Vous avez déjà une paire de Bottes d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Bottes de l'aventurier")
				p.Armure()
				bottes++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Fer":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Épee en fer")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.arme == "Épee en fer" {
				fmt.Println("Vous avez déjà une Épee en fer d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Épee en fer")
				p.Armure()
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Midgard":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Épee de Midgard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.arme == "Épee de Midgard" {
				fmt.Println("Vous avez déjà une Épee Midgard d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Épee de Midgard")
				p.Armure()
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Hache":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Hache Divine Rhitta")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.arme == "Hache Divine Rhitta" {
				fmt.Println("Vous avez déjà la Hache Divine Rhitta d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Hache Divine Rhitta")
				p.Armure()
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Casque":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Casque en fer")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.arme == "Casque en fer" {
				fmt.Println("Vous avez déjà le Casque en fer d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Casque en fer")
				p.Armure()
				c++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Armure":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Armure en fer")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.arme == "Armure en fer" {
				fmt.Println("Vous avez déjà l'Armure en fer d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Armure en fer")
				p.Armure()
				a++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Botte":
		if p.nom == "Escanor " {
			fmt.Println("Votre equipement à fondu à cause de la chaleur que votre corp dégage !")
			p.RemoveInventory("Bottes en fer")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArmure()
		} else {
			if p.equipement.arme == "Bottes en fer" {
				fmt.Println("Vous avez déjà les Bottes en fer d'équiper")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			} else {
				p.addArmure("Bottes en fer")
				p.Armure()
				b++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArmure()
			}
		}
	case "Retour":
		p.BackMenu()

	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.menuArmure()
	}
}
