package main

import (
	"fmt"
	"time"
)

func (p *personnage) chooseCombat() {
	Clear()
	image("combat.PNG", []int{50, 13})
	var num int
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Choisissez le monstre contre lequel vous voulez combattre :\n")
	fmt.Println("1. Combattez un", m1.nom, "(niveau 0 requis)\n")
	fmt.Println("2. Combattez ", m2.nom, "(niveau 3 requis)\n")
	fmt.Println("3. Combattez ", m3.nom, "(niveau 6 requis)\n")
	fmt.Println("4. Combattez ", m4.nom, "(niveau 9 requis)\n")
	if p.nom == "Escanor " {
		fmt.Println("5. Combattez ", m5.nom, "(cheetcode requis)\n")
		fmt.Println("6. Combattez ", m6.nom, "(cheetcode requis)\n")
		fmt.Println("7. Combattez ", m7.nom, "(cheetcode requis)\n")
		fmt.Println("8. Combattez ", m8.nom, "(cheetcode requis)\n")
	}
	fmt.Println("9. Revenir au menu\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		if p.niveau > 0 {
			for p.pva > 0 || m1.pva > 0 {
				p.trainingFight(&m1)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 2:
		if p.niveau >= 3 {
			for p.pva > 0 || m2.pva > 0 {
				p.trainingFight(&m2)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 3:
		if p.niveau >= 6 {
			for p.pva > 0 || m3.pva > 0 {
				p.trainingFight(&m3)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 4:
		if p.niveau >= 9 {
			for p.pva > 0 || m4.pva > 0 {
				p.trainingFight(&m4)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 5:
		if p.nom == "Escanor " {
			for p.pva > 0 || m5.pva > 0 {
				p.trainingFight(&m5)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 6:
		if p.nom == "Escanor " {
			for p.pva > 0 || m6.pva > 0 {
				p.trainingFight(&m6)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 7:
		if p.nom == "Escanor " {
			for p.pva > 0 || m7.pva > 0 {
				p.trainingFight(&m7)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 8:
		if p.nom == "Escanor " {
			for p.pva > 0 || m8.pva > 0 {
				p.trainingFight(&m8)
			}
		} else {
			fmt.Println("Vous n'avez pas le niveau requis ! Veuillez sélectionner un autre niveau")
			time.Sleep(3 * time.Second)
			Clear()
			p.chooseCombat()
		}
	case 9:
		p.BackMenu()
	default:
		fmt.Println("La commande entrée est inconnu ! Veuillez en retaper une valide !")
		time.Sleep(3 * time.Second)
		Clear()
		p.chooseCombat()
	}
}
