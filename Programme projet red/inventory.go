package main

import (
	"fmt"
	"time"
)

func (p *personnage) accesInventory() {
	Clear()
	image("inventaire.PNG", []int{50, 13})
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
		time.Sleep(3 * time.Second)
		Clear()
		p.BackMenu()
	} else {
		for i := 0; i < len(p.inventaire); i++ {
			fmt.Println("Item", i+1, " : ", p.inventaire[i])
		}
		p.searchInventory(p.inventaire)
	}
}

func (p *personnage) upgradeInventorySlot() {
	Clear()
	image("upgrade.PNG", []int{50, 25})
	var num int
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Print("Voulez-vous augmenter votre inventaire pour 30 pièces d'or ?\n")
	fmt.Print("1. Oui\n")
	fmt.Print("2. Revenir au marchand !\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Scanf("%d\n", &num)
	switch num {
	case 1:
		if LimitInv < 40 {
			LimitInv += 10
			fmt.Println("Vous avez augmentez votre inventaire de 10 slot.")
			fmt.Println("Vous avez maintenant", LimitInv, "de place !")
			p.argent -= 30
			fmt.Println("Il vous reste", p.argent, "d'argent")
			time.Sleep(3 * time.Second)
			Clear()
			p.marchand()
		} else if LimitInv >= 40 {
			fmt.Println("Vous avez atteint la limite maximale de l'inventaire .")
			time.Sleep(3 * time.Second)
			Clear()
			p.marchand()
		}
	case 2:
		p.marchand()
	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.upgradeInventorySlot()
	}
}

func (p *personnage) searchInventory(tab []string) {
	var obj string
	inventaire := tab
	if inventaire != nil {
		fmt.Println("Quel objet souhaitez vous utilisez \nPour utiliser la potion de vie, tapez 'vie' \nPour utiliser la potion de poison, tapez 'poison'\nPour utiliser le Livre de sort : Boule de feu tapez 'feu'\nPour utiliser le Livre de sort : Ras de marée tapez 'eau'\nPour utiliser la potion de mana, tapez 'mana'\nPour revenir au menu tapez 'Retour'")
		fmt.Println("Entrez le nom de votre potion:")
		fmt.Scanf("%s\n", &obj)
		switch obj {
		case "vie":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "potion de soin" {
					if p.pva == p.pvm {
						fmt.Println("Vous êtes foulles LIIFFE")
						time.Sleep(3 * time.Second)
						Clear()
						p.accesInventory()
						p.searchInventory(p.inventaire)
					} else {
						p.takePot()
						fmt.Print("Vous avez utiliser la potion\n")
						time.Sleep(3 * time.Second)
						Clear()
						p.BackMenu()
					}
				}
			}
			fmt.Println("Vous ne possedez pas l'objet")
			time.Sleep(3 * time.Second)
			Clear()
			p.accesInventory()
			p.searchInventory(p.inventaire)
		case "poison":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "potion de poison" {
					p.poisonPot()
					fmt.Print("Vous avez utiliser la potion\n")
					time.Sleep(3 * time.Second)
					Clear()
					p.BackMenu()
				}
			}
			fmt.Println("Vous ne possedez pas l'objet")
			time.Sleep(3 * time.Second)
			Clear()
			p.accesInventory()
			p.searchInventory(p.inventaire)

		case "feu":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "Livre de sort : Boule de feu" {
					p.SpellBook("Boule de feu", "Livre de sort : Boule de feu")
					fmt.Println("Vous possédez désormais le sort Boule de feu")
					time.Sleep(3 * time.Second)
					Clear()
					p.BackMenu()
				}
			}
			fmt.Println("Vous ne possedez pas l'objet")
			time.Sleep(3 * time.Second)
			Clear()
			p.accesInventory()
			p.searchInventory(p.inventaire)
		case "eau":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "Livre de sort : Ras de marée" {
					p.SpellBook("Ras de marée", "Livre de sort : Ras de marée")
					fmt.Println("Vous possédez désormais le sort Ras de marée")
					time.Sleep(3 * time.Second)
					Clear()
					p.BackMenu()
				}
			}
			fmt.Println("Vous ne possedez pas l'objet")
			time.Sleep(3 * time.Second)
			Clear()
			p.accesInventory()
			p.searchInventory(p.inventaire)
		case "mana":
			for i := 0; i < len(inventaire); i++ {
				if inventaire[i] == "potion de mana" {
					if p.mana == p.manamax {
						fmt.Println("Votre mana est au max")
						time.Sleep(3 * time.Second)
						Clear()
						p.accesInventory()
						p.searchInventory(p.inventaire)
					} else {
						p.takeMana()
						fmt.Print("Vous avez utiliser la potion\n")
						time.Sleep(3 * time.Second)
						Clear()
						p.BackMenu()
					}
				} else {
					fmt.Println("Vous ne possedez pas l'objet")
					time.Sleep(3 * time.Second)
					Clear()
					p.accesInventory()
					p.searchInventory(p.inventaire)
				}
			}

		case "Retour":
			fmt.Println("Vous revenez au menu !")
			p.BackMenu()
		default:
			fmt.Println("Cet item n'est pas disponible")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("La commande est invalide veuillez la retaper !")
			time.Sleep(3 * time.Second)
			Clear()
			p.accesInventory()
			p.searchInventory(p.inventaire)
		}
	}
}
