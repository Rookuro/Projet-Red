package main

import (
	"fmt"
	"time"
)

func (p *personnage) chooseclass() {
	Clear()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Choisissez votre race :\n")
	fmt.Println("Humain\n")
	fmt.Println("Elfe\n")
	fmt.Println("Nain\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var choix string

	fmt.Scanf("%s\n", &choix)
	switch choix {

	case "Humain":
		p.classe = "Humain"
		p.pvm = 100
		p.pva = p.pvm / 2
		p.mana = 20
		p.manamax = 40
		p.niveau = 1
		p.skill = append(p.skill, "Coup de poing")
		p.initiative = r1.Intn(5)
		p.experiencemax = 2000
	case "Elfe":
		p.classe = "Elfe"
		p.pvm = 80
		p.pva = p.pvm / 2
		p.mana = 30
		p.manamax = 60
		p.niveau = 1
		p.skill = append(p.skill, "Coup de poing")
		p.initiative = r1.Intn(7)
		p.experiencemax = 2500
	case "Nain":
		p.classe = "Nain"
		p.pvm = 120
		p.pva = p.pvm / 2
		p.mana = 15
		p.manamax = 30
		p.niveau = 1
		p.skill = append(p.skill, "Coup de poing")
		p.initiative = r1.Intn(10)
		p.experiencemax = 3000

	default:
		fmt.Println("Merci de bien vouloir entrer une classe valide !")
		time.Sleep(5 * time.Second)
		p.chooseclass()
	}

}

func (p *personnage) charCreation() {
	//var choix string
	var na string
	var me string
	Clear()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("Choisissez votre nouveau nom :\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Scanf("%s %s", &na, &me)
	var tout = na + string(rune(32)) + me
	for i := 0; i < len(tout); i++ {
		var verif = []rune(tout)
		if MajMinNum2(verif[i]) == false {
			fmt.Println("Votre nom est incorrect, merci de bien vouloir entrer uniquement des majuscules et des minuscules !")
			time.Sleep(3 * time.Second)
			p.charCreation()
		}
	}
	p.nom = Capitalize(tout)
	if p.nom == "Escanor " {
		p.sunshine()
		fmt.Println("CHEAT CODE ACTIF !!!!!!!!!!!!")
		affiche(".\n")
		affiche("..\n")
		affiche("...\n")
		fmt.Println("Bien le bonjour monde si sympathique, \nj'ai entendu dire que le Roi des D??mons ??tait entrain d'attaquer cette splendide vall??e ! \nIl a d??ploy?? ses mis??rables sujets pour ann??antire ces terres ! \nMoi,", p.nom, "vais an??antir ses infames monstre en un seul coup !\nCes faibles ne m'arrivent point ?? ma hauteur !\nJe vais directement aller entamer une petite discussion avec le Roi des D??mons,\nCET ??TRE QUI M'EST BIEN INFERIEUR !")
		image("escanor.jpg", []int{128, 40})
	} else {
		p.chooseclass()
		Clear()
		fmt.Println("Vous ??tes un", p.classe, "et vous vous nommez", p.nom, "\nVous habitez dans la ville de Giromagny se situant dans les grandes plaines Quadro !\nAlors que vous viviez une vie paisible\nle Roi du Caf?? George Clooney d??cide d'attaquer la vall?? !\nNe pouvant pas rester les bras crois??s,\nvous d??cidez de menez une lutte acharn?? contre l'arm??e de Clooney !\nLa ville ayant de bonnes murailles, aucun monstres ne pouvaient entrer !\nLe seul moyen de les combattre est de sortir de la ville !\nBonne chance ?? vous jeune Guerrier !")
		time.Sleep(10 * time.Second)
		Clear()
		fmt.Println("Vous avez", p.pva, "/", p.pvm, "PV")
		fmt.Println("Votre niveau de d??part est", p.niveau)
		fmt.Println("Vous possedez le sort :", p.skill)
		fmt.Println("Votre initialtive est de", p.initiative)
		image("pnj.jpg", []int{75, 25})
	}
	time.Sleep(5 * time.Second)
	Clear()
	p1.menu()
}
