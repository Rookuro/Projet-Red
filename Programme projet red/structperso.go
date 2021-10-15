package main

type personnage struct {
	nom            string
	classe         string
	niveau         int
	pvm            int
	pva            int
	inventaire     []string
	skill          []string
	argent         int
	degat          int
	armedeg        int
	magie          int
	initiative     int
	experienceactu int
	experiencemax  int
	mana           int
	manamax        int
	equipement     Equipement
}

type Equipement struct {
	casque string
	armure string
	bottes string
	arme   string
}
