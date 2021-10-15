package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuEtablie() {
	Clear()
	image("etablie.PNG", []int{50, 13})
	var cuir int = 0
	var plume int = 0
	var peau int = 0
	var fourrure int = 0
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Print("Voici les équipements que vous pouvez fabriquer:\n\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("1. Chapeau de l'aventurier")
	fmt.Println("Objets requis :\n- 1 Plume de Corbeau\n- 1 Cuir de Sanglier\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("2. Tunique de l'aventurier")
	fmt.Println("Objets requis :\n- 2 Fourrure de loup\n- 1 Peau de Troll\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("3. Bottes de l'aventurier")
	fmt.Println("Objets requis :\n- 1 Fourrure de loup\n- 1 Cuir de Sanglier\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("4. Revenir à la Forge ")
	var num int
	fmt.Scanf("%d\n", &num)

	switch num {
	case 1:
		if p.argent < 5 {
			fmt.Print("Vous n'avez pas assez d'argent revenez plus tard!")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		}
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Cuir de Sanglier" {
				cuir = cuir + 1
			}
			if p.inventaire[i] == "Plume de Corbeau" {
				plume = plume + 1
			}

		}
		if cuir >= 1 && plume >= 1 {
			p.transaction("Chapeau de l'aventurier", 5, p.inventaire)
			p.RemoveInventory("Plume de Corbeau")
			p.RemoveInventory("Cuir de Sanglier")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		} else {
			fmt.Println("Vous n'avez pas assez de matériaux, revenez plus tard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		}

	case 2:
		if p.argent < 5 {
			fmt.Print("Vous n'avez pas assez d'argent revenez plus tard !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		}
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Fourrure de Loup" {
				fourrure = fourrure + 1
			}
			if p.inventaire[i] == "Peau de Troll" {
				peau = peau + 1
			}

		}
		if peau >= 1 && fourrure >= 2 {
			p.transaction("Tunique de l'aventurier", 5, p.inventaire)
			p.RemoveInventory("Fourrure de Loup")
			p.RemoveInventory("Fourrure de Loup")
			p.RemoveInventory("Peau de Troll")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		} else {
			fmt.Println("Vous n'avez pas assez de matériaux, revenez plus tard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		}
	case 3:
		if p.argent < 5 {
			fmt.Print("Vous n'avez pas assez d'argent revenez plus tard !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		}
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Fourrure de Loup" {
				fourrure = fourrure + 1
			}
			if p.inventaire[i] == "Cuir de Sanglier" {
				cuir = cuir + 1
			}

		}
		if cuir >= 1 && fourrure >= 1 {
			p.transaction("Bottes de l'aventurier", 5, p.inventaire)
			p.RemoveInventory("Fourrure de Loup")
			p.RemoveInventory("Cuir de Sanglier")
			p.transaction("Bottes de l'aventurier", 5, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		} else {
			fmt.Println("Vous n'avez pas assez de matériaux, revenez plus tard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuEtablie()
		}
	case 4:
		p.Forge()
	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.menuEtablie()
	}
}
