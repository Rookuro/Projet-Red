package main

import (
	"fmt"
	"time"
)

func (p *personnage) menuArme() {
	Clear()
	image("arme.PNG", []int{50, 13})
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("1. Épee en fer | 10 pièces d'or\n")
	fmt.Println("2. Épee de Midgard | 20 pièces d'or\n")
	fmt.Println("3. Hache Divine Rhitta | 200 pièces d'or\n")
	fmt.Println("4. Casque en fer | 15 pièces d'or\n")
	fmt.Println("5. Armure en fer | 20 pièces d'or\n")
	fmt.Println("6. Bottes en fer | 10 pièces d'or\n")
	fmt.Println("7.  Revenir au marchand.\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	var num int
	fmt.Scanf("%d\n", &num)
	var choose string
	switch num {
	case 1:
		if epeefer >= 1 {
			fmt.Println("Vous possedez deja cet item !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArme()
		} else {
			fmt.Println("Voulez vous acheter cet objet | 'acheter'")
			fmt.Scanf("%s", &choose)
			switch choose {
			case "acheter":
				p.transaction("Épee en fer", 10, p.inventaire)
				epeefer++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			default:
				fmt.Println("Veuillez rechoisir l'article")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			}
		}
	case 2:
		if epeemidgard >= 1 {
			fmt.Println("Vous possedez deja cet item !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArme()
		} else {
			fmt.Println("Voulez vous acheter cet objet | 'acheter'")
			fmt.Scanf("%s", &choose)
			switch choose {
			case "acheter":
				p.transaction("Épee de Midgard", 20, p.inventaire)
				epeemidgard++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			default:
				fmt.Println("Veuillez rechoisir l'article")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			}
		}
	case 3:
		if hache >= 1 {
			fmt.Println("Vous possedez deja cet item !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArme()
		} else {
			fmt.Println("Voulez vous acheter cet objet | 'acheter'")
			fmt.Scanf("%s", &choose)
			switch choose {
			case "acheter":
				p.transaction("Hache Divine Rhitta", 200, p.inventaire)
				hache++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			default:
				fmt.Println("Veuillez rechoisir l'article")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			}
		}
	case 4:
		if c >= 1 {
			fmt.Println("Vous possedez deja cet item !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArme()
		} else {
			fmt.Println("Voulez vous acheter cet objet | 'acheter'")
			fmt.Scanf("%s", &choose)
			switch choose {
			case "acheter":
				p.transaction("Casque en fer", 15, p.inventaire)
				c++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			default:
				fmt.Println("Veuillez rechoisir l'article")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			}
		}
	case 5:
		if a >= 1 {
			fmt.Println("Vous possedez deja cet item !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArme()
		} else {
			fmt.Println("Voulez vous acheter cet objet | 'acheter'")
			fmt.Scanf("%s", &choose)
			switch choose {
			case "acheter":
				p.transaction("Armure en fer", 20, p.inventaire)
				a++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			default:
				fmt.Println("Veuillez rechoisir l'article")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			}
		}
	case 6:
		if b >= 1 {
			fmt.Println("Vous possedez deja cet item !")
			time.Sleep(3 * time.Second)
			Clear()
			p.menuArme()
		} else {
			fmt.Println("Voulez vous acheter cet objet | 'acheter'")
			fmt.Scanf("%s", &choose)
			switch choose {
			case "acheter":
				p.transaction("Bottes en fer", 10, p.inventaire)
				b++
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			default:
				fmt.Println("Veuillez rechoisir l'article")
				time.Sleep(3 * time.Second)
				Clear()
				p.menuArme()
			}
		}
	case 7:
		Clear()
		p.marchand()
	default:
		fmt.Println("La commande n'a pas été reconnu veuillez la retaper !")
		time.Sleep(3 * time.Second)
		Clear()
		p.menuArme()
	}
}
