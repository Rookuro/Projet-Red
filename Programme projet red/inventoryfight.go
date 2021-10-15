package main

import (
	"fmt"
	"time"
)

func (p *personnage) accesInventoryFight() {
	soin := 0
	poison := 0
	mana := 0
	var yn string
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Print("Vous pouvez utiliser les objets suivants:\n")
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == "potion de soin" {
			soin = soin + 1
		}

		if p.inventaire[i] == "potion de poison" {
			poison = poison + 1
		}

		if p.inventaire[i] == "potion de mana" {
			mana = mana + 1
		}

	}
	fmt.Println(soin, "x Potion de vie")
	fmt.Println(poison, "x Potion de poison")
	fmt.Println(mana, "x Potion de mana")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Voulez vous utiliser un objet de l'inventaire ?")
	fmt.Scanf("%s\n", &yn)
	switch yn {
	case "oui":
		p.searchInventoryFight(p.inventaire)
	case "non":
		m.charturn(p)
	default:
		fmt.Println("La commande entrée est inconnu, merci d'entrer une commande valide !")
		time.Sleep(5 * time.Second)
		Clear()
		p.accesInventoryFight()

	}
}

func (p *personnage) searchInventoryFight(tab []string) {
	var obj string
	inventaire := tab
	if inventaire != nil {
		fmt.Println("Quel objet souhaitez vous utilisez \nPour utiliser la potion de vie, tapez 'vie' \nPour utiliser la potion de poison, tapez 'poison'\nPour utiliser le Livre de sort : Boule de feu tapez 'feu'\nPour utiliser le Livre de sort : Ras de marée tapez 'eau'\nPour utiliser la potion de mana, tapez 'mana'\nPour revenir au menu tapez 'Retour'")
		fmt.Println("Entrez le nom de votre potion:")
		fmt.Scanf("%s\n", &obj)
		switch obj {
		case "vie":
			for i := 0; i < len(p.inventaire); i++ {
				if p.inventaire[i] == "potion de soin" {
					if p.pva == p.pvm {
						fmt.Println("Vous êtes foulles LIIFFE")
						time.Sleep(3 * time.Second)
						fmt.Println("\033[2J")
						p.searchInventoryFight(p.inventaire)
					} else {
						p.takePot()
						fmt.Print("Vous avez utiliser la potion\n")
						return
					}
				}
			}
			fmt.Println("Vous ne possedez pas l'objet")
			time.Sleep(3 * time.Second)
			Clear()
			p.searchInventoryFight(p.inventaire)
		case "poison":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "potion de poison" {
					p.poisonPot()
					fmt.Print("Vous avez utiliser la potion\n")
					return
				}
			}
			fmt.Println("Vous ne possedez pas l'objet")
			time.Sleep(3 * time.Second)
			Clear()
			p.searchInventoryFight(p.inventaire)

		case "feu":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "Livre de sort : Boule de feu" {
					p.SpellBook("Boule de feu", "Livre de sort : Boule de feu")
					return
				} else {
					fmt.Println("Vous ne possedez pas l'objet")
					time.Sleep(3 * time.Second)
					Clear()
					p.searchInventoryFight(p.inventaire)
				}
			}
		case "eau":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "Livre de sort : Boule de feu" {
					p.SpellBook("Ras de marée", "Livre de sort : Ras de marée")
					return
				} else {
					fmt.Println("Vous ne possedez pas l'objet")
					time.Sleep(3 * time.Second)
					Clear()
					p.searchInventoryFight(p.inventaire)
				}
			}
		case "mana":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "potion de mana" {
					if p.mana == p.manamax {
						fmt.Println("Votre mana est au max")
						time.Sleep(3 * time.Second)
						Clear()
						p.searchInventoryFight(p.inventaire)
					} else {
						p.takeMana()
						fmt.Print("Vous avez utiliser la potion\n")
						return
					}
				} else {
					fmt.Println("Vous ne possedez pas l'objet")
					time.Sleep(3 * time.Second)
					Clear()
					p.searchInventoryFight(p.inventaire)
				}
			}

		case "retour":
			fmt.Println("Vous revenez au combat !")
			m.charturn(p)
		default:
			fmt.Println("Cet item n'est pas disponible")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("La commande est invalide veuillez la retaper !")
			time.Sleep(3 * time.Second)
			Clear()
			p.searchInventoryFight(p.inventaire)
		}
	}
}
