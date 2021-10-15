package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuFonderie() {
	Clear()
	image("fonderie.PNG", []int{50, 13})
	var fer int = 0
	var midgard int = 0
	var charbon int = 0
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Print("Bienvenue dans la Fonderie, ici vous pouvez faire fondre vos minerais !\n\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("1. Minerai de fer en Lingot de fer")
	fmt.Println("Objet requis : \n- 2 Minerai de fer\n- 1 Charbon\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("2. Minerai de Midgard en Lingot de Midgard")
	fmt.Println("Objet requis : \n- 2 Minerai de Midgard\n- 1 Charbon\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("3. Retourner à la Forge")
	var num int
	fmt.Scanf("%d\n", &num)
	switch num {
	case 1:
		if p.argent < 5 {
			fmt.Print("Vous n'avez pas assez d'argent revenez plus tard!")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuFonderie()
		}
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Minerai de fer" {
				fer = fer + 1
			}
			if p.inventaire[i] == "Charbon de bois" {
				charbon = charbon + 1
			}
		}
		if fer >= 2 && charbon >= 1 {
			p.RemoveInventory("Minerai de fer")
			p.RemoveInventory("Minerai de fer")
			p.RemoveInventory("Charbon de bois")
			p.transaction("Lingot de fer", 5, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuFonderie()
		} else {
			fmt.Println("Vous n'avez pas assez de matériaux, revenez plus tard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuFonderie()
		}
	case 2:
		if p.argent < 5 {
			fmt.Print("Vous n'avez pas assez d'argent revenez plus tard!")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuFonderie()
		}
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Minerai de Midgard" {
				midgard = midgard + 1
			}
			if p.inventaire[i] == "Charbon de bois" {
				charbon = charbon + 1
			}
		}
		if midgard >= 2 && charbon >= 1 {
			p.RemoveInventory("Minerai de Midgard")
			p.RemoveInventory("Minerai de Midgard")
			p.RemoveInventory("Charbon de bois")
			p.transaction("Lingot de Midgard", 10, p.inventaire)
			time.Sleep(3 * time.Second)
			Clear()
			p.menuFonderie()
		} else {
			fmt.Println("Vous n'avez pas assez de matériaux, revenez plus tard")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuFonderie()
		}
	case 3:
		p.Forge()
	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.menuFonderie()
	}
}
