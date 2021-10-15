package main

import (
	"fmt"
	"os"
)

//Menu principal
func (p *personnage) menu() {
	Clear()
	image("menu.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Print("Faites votre choix :\n\n")
	fmt.Print("1. Regarder mes informations :\n\n")
	fmt.Print("2. Regarder dans son inventaire :\n\n")
	fmt.Print("3. Aller au Marchand :\n\n")
	fmt.Print("4. Aller à la Forge :\n\n")
	fmt.Print("5. Aller à l'Armurie :\n\n")
	fmt.Print("6. Aller au terrain d'Entrainement !\n\n")
	fmt.Print("7. Quitter ce monde !\n\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var chi int
	fmt.Scanf("%d\n", &chi)
	fmt.Println("Votre choix est :", chi)
	switch chi {
	case 1:
		p.DisplayInfo()
	case 2:
		p.accesInventory()
	case 3:
		p.marchand()
	case 4:
		p.Forge()
	case 5:
		p.menuArmure()
	case 6:
		//fmt.Println("Le combat d'entrainement a commencé !")
		p.chooseCombat()

	case 7:
		fmt.Println("Vous avez quitté le menu")
		os.Exit(0)
	}
}
