package main

import (
	"fmt"
	"time"
)

func (p *personnage) marchand() {
	/* var potion string = "potion"
	var prixpotion int = 0*/
	Clear()
	image("marchand.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Print("1. Boutique de Potion\n\n")
	fmt.Print("2. Boutique d'Arme\n\n")
	fmt.Print("3. Boutique d'Objet\n\n")
	fmt.Print("4. Boutique de sort\n\n")
	fmt.Print("5. Augmentation Inventaire\n\n")
	fmt.Print("6. Retour au menu\n\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num int
	fmt.Scanf("%d\n", &num)
	/*Fonction qui permet d'afficher l'erreur
	// _, err := fmt.Scanf("%d\n", &num)
	// if err != nil {
	// 	log.Fatal(err)
	// }*/
	switch num {
	case 1:
		p.menuPotion()
	case 2:
		p.menuArme()
	case 3:
		p.menuObjet()
	case 4:
		p.menuMagie()
	case 5:
		p.upgradeInventorySlot()
	case 6:
		p.BackMenu()
	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.marchand()
	}
}
