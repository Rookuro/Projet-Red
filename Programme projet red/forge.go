package main

import (
	"fmt"
	"time"
)

func (p *personnage) Forge() {
	Clear()
	image("forge.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Print("Vous devez choisir oú aller dans la forge\n\n")
	fmt.Println("1. Aller vers la Fonderie\n")
	fmt.Println("2. Aller vers l'Établie\n")
	fmt.Println("3. Aller vers l'Enclume\n")
	fmt.Println("4. Revenir au menu\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num int
	fmt.Scanf("%d\n", &num)
	switch num {
	case 1:
		fmt.Println("Vous êtes dans la Fonderie,\nVous pouvez y faire fondre vos minerai !")
		p.menuFonderie()
	case 2:
		fmt.Println("Vous êtes dans l'Établie,\nVous pouvew y fabriquer toute sorte d'équipement !")
		p.menuEtablie()
	case 3:
		fmt.Println("Vous êtes devant l'enclume,\nVous pouvew y fabriquer toute sorte d'armes !")
		p.menuEnclume()
	case 4:
		p.BackMenu()
	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.Forge()
	}
}
