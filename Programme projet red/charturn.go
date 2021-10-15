package main

import (
	"fmt"
	"time"
)

func (m *Monstre) charturn(p *personnage) { // modif
	var num int
	//for i := 0; i < m.pva; i++
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Quelle est l'action que vous voulez réaliser ?:\n")
	fmt.Println("1. Attaquer\n")
	fmt.Println("2. Inventaire\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Scanf("%d\n", &num)
	switch num {
	case 1:
		m.chooseattack(p)
	case 2:
		p.accesInventoryFight()

	default:
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(5 * time.Second)
		Clear()
		m.charturn(p)
	}
}
