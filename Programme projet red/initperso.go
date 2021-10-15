package main

func (p *personnage) Init(nom string, classe string, niveau int, pvm int, pva int, inventaire []string, skill []string, argent int, equip []string, degat int, armedeg int, magie int, initiative int, experienceactu int, experiencemax int, mana int, manamax int, casque string, armure string, bottes string, arme string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.pvm = pvm
	p.pva = pva
	p.inventaire = inventaire
	p.skill = skill
	p.argent = argent
	p.mana = mana
	p.manamax = manamax
	p.degat = degat
	p.armedeg = armedeg
	p.magie = magie
	p.initiative = initiative
	p.experienceactu = experienceactu
	p.experiencemax = experiencemax
	p.equipement.casque = casque
	p.equipement.armure = armure
	p.equipement.bottes = bottes
	p.equipement.arme = arme
}
