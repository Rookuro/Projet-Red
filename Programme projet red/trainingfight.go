package main

import (
	"fmt"
	"time"
)

func (p *personnage) trainingFight(m *Monstre) { // modif
	Clear()
	tour = 0
	m.pva = m.pvm
	if m.nom == m1.nom {
		image("goblin.jpg", []int{50, 70})
	} else if m.nom == m2.nom {
		image("thanos.jpg", []int{125, 50})
	} else if m.nom == m3.nom {
		image("maximus.jpg", []int{100, 25})
	} else if m.nom == m4.nom {
		image("george2.jpg", []int{100, 40})
	} else if m.nom == m5.nom {
		image("estarossa.jpg", []int{100, 25})
	} else if m.nom == m6.nom {
		image("meliodas.jpg", []int{100, 25})
	} else if m.nom == m7.nom {
		image("zeldris.jpg", []int{100, 40})
	} else if m.nom == m8.nom {
		image("roi.jpg", []int{100, 25})
	}
	fmt.Println("Le combat d'entrainement a commencé !")
	fmt.Println("Votre initialtive est de", p.initiative)
	fmt.Println("L'initialtive du", m.nom, "est de", m.initiative)

	if m.initiative > p.initiative {
		fmt.Println(m.nom, "commence en premier")
	} else {
		fmt.Println("Vous commencez en premier")
	}
	time.Sleep(3 * time.Second)
	Clear()
	if m.pva == 0 {
		fmt.Println(p.nom, "a gagné le combat !")
		fmt.Println("Vous avez gagné", m.gain, "d'argent")
		p.argent = p.argent + m.gain
		fmt.Println("Vous avez maintenant", p.argent, "d'argent")
		fmt.Println("Le combat est terminé !")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		if m4.pva == 0 {
			time.Sleep(5 * time.Second)
			Clear()
			credit()
		} else {
			time.Sleep(5 * time.Second)
			Clear()
			p.BackMenu()
		}
		m.pva = m.pvm
	}
	if p.pva == 0 {
		fmt.Println(m.nom, "a gagné le combat !")
		fmt.Println("Le combat est terminé !")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		time.Sleep(5 * time.Second)
		Clear()
		p.dead()
	}
	for 0 < m.pva && 0 < p.pva {
		time.Sleep(3 * time.Second)
		Clear()
		tour = tour + 1
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		fmt.Println("Vous etes au tour n°", tour)
		fmt.Println(m.nom, "est à", m.pva, "/", m.pvm, "PV")
		fmt.Println(p.nom, "est à", p.pva, "/", p.pvm, "PV")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		if m.initiative > p.initiative {
			if m.pva == 0 {
				fmt.Println(p.nom, "a gagné le combat !")
				fmt.Println("Vous avez gagné", m.gain, "d'argent")
				p.argent = p.argent + m.gain
				fmt.Println("Vous avez maintenant", p.argent, "d'argent")
				fmt.Println("Le combat est terminé !")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
				if m4.pva == 0 {
					time.Sleep(5 * time.Second)
					Clear()
					credit()
				} else {
					time.Sleep(5 * time.Second)
					Clear()
					p.BackMenu()
					m.pva = m.pvm
				}
			}
			m.goblinPattern(p)
		} else {
			if p.pva == 0 {
				fmt.Println(m.nom, "a gagné le combat !")
				fmt.Println("Le combat est terminé !")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
				time.Sleep(5 * time.Second)
				Clear()
				p.dead()
			}
			m.charturn(p)
		}
		if m.initiative > p.initiative {
			if p.pva == 0 {
				fmt.Println(m.nom, "a gagné le combat !")
				fmt.Println("Le combat est terminé !")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
				time.Sleep(5 * time.Second)
				Clear()
				p.dead()
			}
			m.charturn(p)
		} else {
			if m.pva == 0 {
				fmt.Println(p.nom, "a gagné le combat !")
				fmt.Println("Vous avez gagné", m.gain, "d'argent")
				p.argent = p.argent + m.gain
				fmt.Println("Vous avez maintenant", p.argent, "d'argent")
				fmt.Println("Le combat est terminé !")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
				if m4.pva == 0 {
					time.Sleep(5 * time.Second)
					Clear()
					credit()
				} else {
					time.Sleep(5 * time.Second)
					Clear()
					p.BackMenu()
					m.pva = m.pvm
				}
			}
			m.goblinPattern(p)
		}
	}
	if m.pva == 0 {
		fmt.Println(p.nom, "a gagné le combat !")
		fmt.Println("Vous avez gagné", m.gain, "d'argent")
		p.argent = p.argent + m.gain
		fmt.Println("Vous avez maintenant", p.argent, "d'argent")
		fmt.Println("Le combat est terminé !")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		if m4.pva == 0 {
			time.Sleep(5 * time.Second)
			Clear()
			credit()
		} else {
			time.Sleep(5 * time.Second)
			Clear()
			p.BackMenu()
			m.pva = m.pvm
		}
	}
	if p.pva == 0 {
		fmt.Println(m.nom, "a gagné le combat !")
		fmt.Println("Le combat est terminé !")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		time.Sleep(5 * time.Second)
		Clear()
		p.dead()
	}
	p.BackMenu()
	tour = 0
	m.pva = m.pvm
}
