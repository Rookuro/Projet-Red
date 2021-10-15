package main

import (
	"fmt"
)

func main() {
	m1.InitGoblin("Goblelin d'entrainement", 40, 40, 5, 5, 10+r1.Intn(5), 1500)
	m2.InitGoblin("Josh Brolin allias Thanos", 60, 60, 8, 7, 20+r1.Intn(10), 2000)
	m3.InitGoblin("Russell Crowe allias Maximus", 80, 80, 11, 9, 30+r1.Intn(15), 2500)
	m4.InitGoblin("Le Roi du Café, George Clooney", 100, 100, 15, 12, 50+r1.Intn(30), 3000)
	m5.InitGoblin("Estarossa de la Charité", 110, 110, 25, 15, 550+r1.Intn(330), 4000)
	m6.InitGoblin("Méliodas Mode Assault", 120, 120, 40, 18, 700+r1.Intn(400), 5000)
	m7.InitGoblin("Zeldris de la Piété", 130, 130, 50, 20, 800+r1.Intn(500), 6000)
	m8.InitGoblin("Le Roi des Démons", 200, 200, 1000, 50, 1000+r1.Intn(800), 10000)
	fmt.Println("Vous avez été invoqué dans ce monde pour combattre la menace que représente l'armée de George Clooney !")
	p1.Init(p.nom, p.classe, p.pvm, p.pva, p.niveau, []string{}, p.skill, 151515, []string{}, p.degat, p.armedeg, p.magie, 1, 0, p.experiencemax, p.mana, p.manamax, "aucun", "aucun", "aucun", "aucune")
	p1.charCreation()

	for {
		if p1.pva <= 0 {
			p1.dead()
		} else {
			p1.menu()
		}
		p1.Armure()
	}

}
