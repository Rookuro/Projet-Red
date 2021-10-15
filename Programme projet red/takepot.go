package main

import "fmt"

func (p *personnage) takePot() {
	inventaire := p.inventaire
	for i := 0; i < len(inventaire); i++ {
		if inventaire[i] == "potion de soin" {
			if p.pva == p.pvm {
				p.pva = p.pvm
				fmt.Println("Vous êtes full life")
				break
			} else if p.pva > p.pvm/2 {
				p.pva = p.pva + (p.pvm - p.pva)
				fmt.Println("Vous avez maintenant ", p.pva)
				p.RemoveInventory("potion de soin")
				break
			} else if p.pva <= p.pvm/2 {
				p.pva = p.pva + 50
				if p.pva >= p.pvm {
					p.pva = p.pvm
				}
				fmt.Println("Vous avez maintenant ", p.pva)
				p.RemoveInventory("potion de soin")
				break
			}
		}
	}
}
