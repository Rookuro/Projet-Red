package main

import "fmt"

func (m *Monstre) goblinPattern(p *personnage) {
	//for i := 0; i <= p.pva; i++
	if p.pva > m.degat {
		if tour%3 != 0 {
			p.pva = p.pva - m.degat
			fmt.Println(m.nom, "inflige", m.degat, "de dégat à", p.nom)
			fmt.Println("Le joueur à", p.pva, "de vie sur", p.pvm)
		} else {
			p.pva -= m.degat * 2
			fmt.Println(m.nom, "inflige", m.degat*2, "de dégat à", p.nom)
			fmt.Println("Le joueur à", p.pva, "de vie sur", p.pvm)
		}
	} else if p.pva <= m.degat {
		if tour%3 != 0 {
			p.pva -= p.pva
			fmt.Println(m.nom, "inflige", m.degat, "de dégat à", p.nom)
			fmt.Println("Le joueur à", p.pva, "de vie sur", p.pvm)
		} else {
			p.pva -= p.pva
			fmt.Println(m.nom, "inflige", m.degat*2, "de dégat à", p.nom)
			fmt.Println("Le joueur à", p.pva, "de vie sur", p.pvm)
		}
	}
}
