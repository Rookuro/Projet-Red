package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuObjet() {
	Clear()
	image("objet.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Objet de catégorie Armure :\n")
	fmt.Println("1. Fourrure de Loup | 4 pièces d'or\n")
	fmt.Println("2. Peau de Troll | 8 pièces d'or\n")
	fmt.Println("3. Cuir de Sanglier | 4 pièces d'or\n")
	fmt.Println("4. Plume de Corbeau | 2 pièces d'or\n")
	fmt.Println("Objet de catégorie Arme :\n")
	fmt.Println("5. Minerai de fer | 10 pièces d'or\n")
	fmt.Println("6. Minerai de Midgard | 26 pièces d'or\n")
	fmt.Println("7. Charbon de bois | 2 pièces d'or\n")
	fmt.Println("8. Morceau de bois | 2 pièces d'or\n")
	fmt.Println("9. Revenir au marchand\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num int
	fmt.Scanf("%d\n", &num)
	var choose string
	switch num {
	case 1:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Fourrure de Loup", 4, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Fourrure de Loup", 4, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 2:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Peau de Troll", 8, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Peau de Troll", 8, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 3:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Cuir de Sanglier", 4, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Cuir de Sanglier", 4, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 4:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Plume de Corbeau", 2, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Plume de Corbeau", 2, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 5:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Minerai de fer", 10, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Minerai de fer", 10, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 6:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Minerai de Midgard", 26, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Minerai de Midgard", 26, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 7:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Charbon de bois", 2, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Charbon de bois", 2, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 8:
		fmt.Println("Voulez vous acheter ou vendre cet objet")
		fmt.Scanf("%s", &choose)
		switch choose {
		case "acheter":
			p.transaction("Morceau de bois", 2, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		case "vendre":
			p.SellRemove("Morceau de bois", 2, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		default:
			fmt.Println("Veuillez rechoisir l'article")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuObjet()
		}
	case 9:
		p.marchand()
	default:
		fmt.Println("Veuillez rechoisir l'article")
		time.Sleep(3 * time.Second)
		p.menuObjet()
	}
}
