package main

func (m *Monstre) InitGoblin(nom string, pvm int, pva int, degat int, initiative int, gain int, gainxp int) {
	m.nom = nom
	m.pvm = pvm
	m.pva = pva
	m.degat = degat
	m.initiative = initiative
	m.gain = gain
	m.gainxp = gainxp
}
