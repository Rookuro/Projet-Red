package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuMagie() {
	Clear()
	image("sort.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Faites votre choix :\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Différents sorts :\n")
	fmt.Println("1. Livre de sort : Boule de feu || 30 pièces\n")
	fmt.Println("2. Livre de sort : Ras de marée || 40 pièces\n")
	fmt.Println("3. Revenir au marchand\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num int
	fmt.Scanf("%d\n", &num)
	//switch case permettant de choisir la catégorie
	switch num {
	case 1:
		if sort >= 1 {
			fmt.Println("Vous possedez deja ce sort !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuMagie()
		} else {
			p.transaction("Livre de sort : Boule de feu", 30, p.inventaire)
			sort++
			time.Sleep(3 * time.Second)
			Clear()
			p.menuMagie()
		}
	case 2:
		if sort2 >= 1 {
			fmt.Println("Vous possedez deja ce sort !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuMagie()
		} else {
			p.transaction("Livre de sort : Ras de marée", 40, p.inventaire)
			sort2++
			time.Sleep(3 * time.Second)
			Clear()
			p.menuMagie()
		}
	case 3:
		p.marchand()
	default:
		fmt.Println("La commande n'a pas été reconnu veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.menuMagie()
	}
}
