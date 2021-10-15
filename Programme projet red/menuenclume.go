package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuEnclume() {
	Clear()
	image("enclume.PNG", []int{50, 13})
	var fer int = 0
	var midgard int = 0
	var bois int = 0
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Print("Voici les équipements que vous pouvez fabriquer:\n\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("1. Épee en fer")
	fmt.Println("Objets requis :\n- 2 Lingots de Fer\n- 1 Morceau de bois\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("2. Épee en acier de Midgard")
	fmt.Println("Objets requis :\n- 2 Lingots de Midgard\n- 1 Morceau de bois\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("3. Revenir à la Forge\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num int
	fmt.Scanf("%d\n", &num)

	switch num {
	case 1:
		if p.argent < 5 {
			fmt.Print("Vous n'avez pas assez d'argent revenez plus tard!")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEnclume()
		}
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Lingot de fer" {
				fer = fer + 1
			}
			if p.inventaire[i] == "Morceau de bois" {
				bois = bois + 1
			}

		}
		if fer >= 2 && bois >= 1 {
			p.RemoveInventory("Lingot de fer")
			p.RemoveInventory("Lingot de fer")
			p.RemoveInventory("Morceau de bois")
			p.transaction("Épee en fer", 5, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEnclume()
		} else {
			fmt.Println("Vous n'avez pas assez de matériaux, revenez plus tard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEnclume()
		}
	case 2:
		if p.argent < 5 {
			fmt.Print("Vous n'avez pas assez d'argent revenez plus tard!")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEnclume()
		}
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Lingot de Midgard" {
				fer = fer + 1
			}
			if p.inventaire[i] == "Morceau de bois" {
				bois = bois + 1
			}

		}
		if midgard >= 2 && bois >= 1 {
			p.RemoveInventory("Lingot de Midgard")
			p.RemoveInventory("Lingot de Midgard")
			p.RemoveInventory("Morceau de bois")
			p.transaction("Épee de Midgard", 10, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEnclume()
		} else {
			fmt.Println("Vous n'avez pas assez de matériaux, revenez plus tard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEnclume()
		}
	case 3:
		p.Forge()
	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.menuEnclume()
	}
}
