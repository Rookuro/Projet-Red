package main

import (
	"fmt"
	"time"
)

func (p *personnage) poisonPot() {
	var degat int = 10 //variable degat qui enleve 10 de pv
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == "potion de poison" {
			p.RemoveInventory("potion de poison")
			for i := 0; i <= 2; i++ {
				time.Sleep(1 * time.Second)
				if p.pva > 0 {
					p.pva = p.pva - degat
					fmt.Println("Vous avez : ", p.pva, " PV")
					if p.pva < 10 {
						p.pva -= p.pva
						fmt.Println("Vous avez : ", p.pva, " PV")
					}
				}

			}
		}

	}
	p.dead()
}
