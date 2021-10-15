package main

import (
	"fmt"
	"os"
	"time"

	"github.com/01-edu/z01"
)

func MajMinNum2(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || a == rune(32) {
		return true
	}
	return false
}

func MajMinNum(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	res := []rune(s)
	lettre := true
	for i := 0; i < len(s); i++ {
		var maj bool = res[i] >= 'A' && res[i] <= 'Z'
		var min bool = res[i] >= 'a' && res[i] <= 'z'
		if MajMinNum(res[i]) == true && lettre {
			if min {
				res[i] = res[i] - 32
			}
			lettre = false
		} else if maj {
			res[i] = res[i] + 32
		} else if !MajMinNum(res[i]) {
			lettre = true
		}
	}
	return string(res)
}

func affiche(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
		time.Sleep(25 * time.Millisecond)
	}
}

func (p *personnage) BackMenu() {
	fmt.Print("Vous êtes revenu au menu !\n")
	p.menu()
}

func (p *personnage) RemoveInventory(s string) {
	var i int
	for i = 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == s {
			break
		}
	}
	if i < len(p.inventaire) {
		p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
		return
	}
}

func (p *personnage) RemoveInventorySkill(s string) {
	var i int
	for i = 0; i < len(p.skill); i++ {
		if p.skill[i] == s {
			break
		}
	}
	if i < len(p.skill) {
		p.skill = append(p.skill[:i], p.skill[i+1:]...)
		return
	}
}

func (p *personnage) addInventory(s string) []string {
	if len(p.inventaire) < LimitInv {
		p.inventaire = append(p.inventaire, s)
	} else {
		p.OverWeight()
	}
	return p.inventaire
}

func (p *personnage) SpellBook(s string, a string) {
	p.skill = append(p.skill, s)
	p.RemoveInventory(a)
}

func (p *personnage) dead() int { // modif
	if p.pva == 0 {
		fmt.Println("Vous êtes MORT --> YOU ARE DEAD !!!!!")
		fmt.Println("Vous êtes réscussité par JEAN REMIS ! Vous regagnez 50 PV")
		p.pva += p.pvm / 2
		p.mana = p.manamax / 2
		p.BackMenu()
	}
	return p.pva
}

func (p *personnage) OverWeight() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.inventaire) == LimitInv {
			fmt.Println("Vous avez atteint la limite d'inventaire.")
			fmt.Println("Vous pouvez acheter ")
		}
	}
}

func credit() {
	image("escanorfin.jpg", []int{100, 50})
	image("rokuroo.jpg", []int{100, 50})
	fmt.Println("Bravo vous avez battu", m4.nom, "!\nVous avez éliminé la terrible menace qu'il représenté pour la vallé !\nVous êtes devenu le héro de la ville et pour vous récompenser,\nle Dieu JEAN REMIS vous à renvoyer dans votre monde d'origine !\nMerci d'avoir jouer à notre jeu !\nCréateurs : Alexis Giromany et Thomas Quadro")
	os.Exit(0)
}

func (p *personnage) SellRemove(obj string, price int, tab []string) {
	if len(tab) == 0 {
		fmt.Println("Ben alors, à ce qui parait on est un peu pauvre ? On essaie de m'arnaquer ! Casse toi d'la sale Gueux !")
	}
	for i, _ := range tab {
		if obj == tab[i] {
			price = price / 2
			p.argent = p.argent + price
			fmt.Println("Vous avez vendu", obj, "pour", price, "piéces d'or")
			fmt.Println("Vous possédez maintenant", p.argent, "d'or !")
			p.RemoveInventory(obj)
			break
		}
		if i == len(tab)-1 {
			fmt.Println("Ben alors, à ce qui parait on est un peu pauvre ? On essaie de m'arnaquer ! Casse toi d'la sale Gueux !")
		}
	}
}

func (p *personnage) transaction(obj string, price int, inv []string) {
	if len(p.inventaire) == LimitInv {
		fmt.Println("Vous avez atteint la limite d'inventaire.")
	} else {
		if p.argent >= price {
			p.argent = p.argent - price
			fmt.Println("Votre transaction est de", price, "piéces d'or")
			fmt.Println("Vous possédez maintenant", p.argent, "d'or !")
			p.addInventory(obj)
		} else {
			fmt.Println("Ben alors, à ce qui parait on est un peu pauvre ? Reviens quand tu seras aussi riche que Jeff Bezos !")
		}
	}
}

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
