package main

import "fmt"

func (p *personnage) afficheInventory(t []string) {
	for i, mot := range t {
		fmt.Println("Objet", i+1, ":", mot)
	}
}

func (p *personnage) afficheSkill(t []string) {
	for i, mot := range t {
		fmt.Println("Skill", i+1, ":", mot)
	}
}

func (p *personnage) DisplayInfo() {
	Clear()
	image("info.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Nom:", p.nom)
	fmt.Println("Race :", p.classe)
	fmt.Println("Niveau :", p.niveau)
	fmt.Println("EXP :", p.experienceactu)
	fmt.Println("EXP Max :", p.experiencemax)
	fmt.Println("Points de vie MAX :", p.pvm)
	fmt.Println("Points de vie ACTUEL :", p.pva)
	fmt.Println("L'inventaire de", p.nom, "contient :")
	p.afficheInventory(p.inventaire)
	fmt.Println("Taille de l'inventaire actuellement :", LimitInv)
	fmt.Println("Les skills de", p.nom, "sont :")
	p.afficheSkill(p.skill)
	fmt.Println("Vous possedez", p.argent, "de piéces d'or !")
	fmt.Println("Mana ACTUEL :", p.mana, "mana")
	fmt.Println("Mana NAX :", p.manamax, "manamax")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Vous êtes équipé de :")
	fmt.Println("Casque :", p.equipement.casque)
	fmt.Println("Armure :", p.equipement.armure)
	fmt.Println("Bottes :", p.equipement.bottes)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Vous êtes équipé d'une :")
	fmt.Println("Arme :", p.equipement.arme, "")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Revenir au menu 'Retour'!")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	var num string
	fmt.Scanf("%s\n", &num)
	switch num {
	case "Retour": //On choisi la potion que l'on veut utiliser
		p.BackMenu()
	default:
		fmt.Println("La commande n'a pas été comprise")
		p.DisplayInfo()
	}
}
