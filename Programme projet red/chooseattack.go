package main

import (
	"fmt"
	"time"
)

func (m *Monstre) chooseattack(p *personnage) { // modif
	fmt.Println("Quelle sort voulez-vous utiliser ?")
	fmt.Println("Vous avez", p.mana, "/", p.manamax)
	if p.nom == "Escanor " {
		fmt.Println("utiliser Epéee Divine ESCANOR !!!!!! avec 'Seikin' One Shot", m.nom)
	} else {
		fmt.Println("utiliser le sort 'Poing' fait 8 de déggat + % de dégat de l'arme")
	} //-----------------------------------
	fmt.Println("utiliser le sort 'Feu' coûte 10 mana et fait 18 de dégat")
	fmt.Println("utiliser le sort 'Eau' coûte 12 mana et fait 22 de dégat")
	var sort string
	fmt.Scanf("%s\n", &sort)
	//-----------------------------------------------------------------
	if sort != "Poing" && sort != "Feu" && sort != "Seikin" && sort != "Eau" {
		fmt.Println("La commande est invalide veuillez la retaper !")
		time.Sleep(5 * time.Second)
		Clear()
		m.chooseattack(p)
	}
	//---------------------------------------------------------------------------
	if p.nom == "Escanor " {
		if sort == "Seikin" {
			if m.pva > p.degat {
				m.pva = m.pva - p.degat
				fmt.Println(p.nom, "inflige", p.degat, "de dégat à", m.nom)
				fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
			} else if m.pva <= p.degat {
				m.pva -= m.pva
				fmt.Println(p.nom, "inflige", p.degat, "de dégat à", m.nom)
				fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
				fmt.Println("Vous avez battu le monstre !")
				p.Experience(m)
			}
		}
	} else {
		if sort == "Poing" {
			p.degat = 8
			p.degat += p.armedeg
			if m.pva > p.degat {
				m.pva = m.pva - p.degat
				fmt.Println(p.nom, "inflige", p.degat, "de dégat à", m.nom)
				fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
			} else if m.pva <= p.degat {
				m.pva -= m.pva
				fmt.Println(p.nom, "inflige", p.degat, "de dégat à", m.nom)
				fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
				fmt.Println("Vous avez battu le monstre !")
				p.Experience(m)
			}
		}
	}
	if sort == "Feu" {
		p.magie = 18
		for i, _ := range p.skill {
			if p.skill[i] == "Boule de feu" {
				if p.mana < 10 {
					fmt.Println("Vous ne pouvez pas executer le sort car vous ne possedez pas assez de mana !")
					m.chooseattack(p)
				} else {
					p.mana -= 10
					if m.pva > p.magie {
						m.pva = m.pva - p.magie
						fmt.Println(p.nom, "inflige", p.magie, "de dégat à", m.nom)
						fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
					} else if m.pva <= p.magie {
						m.pva -= m.pva
						fmt.Println(p.nom, "inflige", p.magie, "de dégat à", m.nom)
						fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
						fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
						fmt.Println("Vous avez battu le monstre !")
						p.Experience(m)
					}
				}
			} else {
				fmt.Println("Vous ne possedez pas le sort !")
				time.Sleep(5 * time.Second)
				Clear()
				m.chooseattack(p)
			}
		}
	}
	if sort == "Eau" {
		p.magie = 22
		for i, _ := range p.skill {
			if p.skill[i] == "Ras de marée" {
				if p.mana < 12 {
					fmt.Println("Vous ne pouvez pas executer le sort car vous ne possedez pas assez de mana !")
					m.chooseattack(p)
				} else {
					p.mana -= 12
					if m.pva > p.magie {
						m.pva = m.pva - p.magie
						fmt.Println(p.nom, "inflige", p.magie, "de dégat à", m.nom)
						fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
					} else if m.pva <= p.magie {
						m.pva -= m.pva
						fmt.Println(p.nom, "inflige", p.magie, "de dégat à", m.nom)
						fmt.Println(m.nom, "à", m.pva, "de vie sur", m.pvm)
						fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
						fmt.Println("Vous avez battu le monstre !")
						p.Experience(m)
					}
				}
			} else {
				fmt.Println("Vous ne possedez pas le sort !")
				time.Sleep(5 * time.Second)
				Clear()
				m.chooseattack(p)
			}
		}
	}
}
