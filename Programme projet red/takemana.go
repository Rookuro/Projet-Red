package main

import "fmt"

func (p *personnage) takeMana() {
	inventaire := p.inventaire
	for i := 0; i < len(inventaire); i++ {
		if inventaire[i] == "potion de mana" {
			if p.mana == p.manamax {
				p.mana = p.manamax
				fmt.Println("Vous êtes full mana")
			}
			if p.mana > p.manamax/2 {
				p.mana = p.mana + (p.manamax - p.mana)
				fmt.Println("Vous avez maintenant ", p.mana)
				p.RemoveInventory("potion de soin")
			}
			if p.mana <= p.manamax/2 {
				p.mana = p.mana + 20
				fmt.Println("Vous avez maintenant ", p.mana)
				p.RemoveInventory("potion de soin")

			}
		}
	}
}
