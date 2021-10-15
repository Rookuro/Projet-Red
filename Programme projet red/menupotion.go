package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuPotion() {
	Clear()
	image("potion.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("1. Potion de soin | 3 pèces d'or\n")
	fmt.Println("2. Potion de poison | 6 pièces d'or\n")
	fmt.Println("3. Potion de mana | 5 pièces d'or\n")
	fmt.Println("4. Revenir au marchand\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num int
	var choose string
	fmt.Scanf("%d\n", &num)
	switch num {
	case 1:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("potion de soin", 3, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		case "vendre":
			p.SellRemove("potion de soin", 3, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		}
	case 2:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("potion de poison", 6, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		case "vendre":
			p.SellRemove("potion de poison", 6, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		}
	case 3:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("potion de mana", 5, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		case "vendre":
			p.SellRemove("potion de mana", 5, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuPotion()
		}
	case 4:
		Clear()
		p.marchand()
	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.menuPotion()
	}
}
