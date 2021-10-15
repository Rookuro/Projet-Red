package main

import "fmt"

func (p *personnage) Experience(m *Monstre) {
	if m.pva == 0 {
		p.experienceactu += m.gainxp
	}
	if p.experienceactu >= p.experiencemax {
		surplus := (p.experienceactu - p.experiencemax)
		if surplus <= 0 {
			p.experienceactu = 0
			p.niveau += 1
			p.experiencemax = p.experiencemax + 250
		} else if surplus > 0 {
			p.experienceactu = surplus
			p.niveau += 1
			p.experiencemax = p.experiencemax + 250
		}
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Vous êtes au niveau", p.niveau, ".\n")
	fmt.Println("Vous avez gagné", m.gainxp, "XP\n")
	fmt.Println("Vous avez maintenant", p.experienceactu, "/", p.experiencemax, "expériences.\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
}
