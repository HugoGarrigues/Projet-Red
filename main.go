package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	menu()
}

// initialisation du JEU a l'aide du menu //

func name() string {
	var name string
	fmt.Println("        Vous avez choisi de créer un personnage")
	fmt.Println("   C'est un choix des plus important car c'est le nom de")
	fmt.Println("votre personnage avec qui vous allez vivre de nombreuses aventures")
	fmt.Println("")
	fmt.Println("       Alors quel sera le nom de votre personnage ?")
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		name = line
	}
	for {
		if verifname(name) == false {
			fmt.Println("Vous utilisez des caractères non autorisés")
			fmt.Println("Veuillez entrer un autre nom")
			fmt.Scanln(&name)
		} else {
			break
		}
	}
	if name[0] >= 'a' && name[0] <= 'z' {
		name = string(name[0]-32) + name[1:]
	}
	for i := 1; i < len(name); i++ {
		if name[i] >= 'A' && name[i] <= 'Z' {
			name = name[:i] + string(name[i]+32) + name[i+1:]
		}
	}
	fmt.Println("___________________________________________")
	fmt.Println("|                                          ")
	fmt.Println("|     Vous avez choisi le nom de", name, " ")
	fmt.Println("|__________________________________________")
	return name
}

func menu() {
	var personnage Personnage
	var choix string
	fmt.Println("______________________________________")
	fmt.Println("|                                     |")
	fmt.Println("|    BIENVENUE A VOUS DANS CE JEU     |")
	fmt.Println("|   Dans ce jeu vous aurez uniquement |")
	fmt.Println("|    des choix a faire pour réaliser  |")
	fmt.Println("|         votre folle aventure        |")
	fmt.Println("|                                     |")
	fmt.Println("| Et viens la votre premier choix :   |")
	fmt.Println("|                                     |")
	fmt.Println("|        1. Créer un personnage       |")
	fmt.Println("|                                     |")
	fmt.Println("|        2. Quitter le jeu            |")
	fmt.Println("|_____________________________________|")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		name := name()
		personnage = chooseclass(name)
		time.Sleep(1 * time.Second)
		fmt.Println("______________________________________")
		fmt.Println("|                                     |")
		fmt.Println("|              BRAVO !                |")
		fmt.Println("|                                     |")
		fmt.Println("|    Vous avez finis la création      |")
		fmt.Println("|        de votre personnage          |")
		fmt.Println("|                                     |")
		fmt.Println("|           Bonne chance a            |")
		fmt.Println("|             vous dans               |")
		fmt.Println("|     l'aventure qui vous attends     |")
		fmt.Println("|                                     |")
		fmt.Println("|         Amusez-vous bien !          |")
		fmt.Println("|_____________________________________|")
		time.Sleep(1 * time.Second)
	case "2":
		fmt.Println("Vous avez quitté le jeu")
		os.Exit(0)
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		menu()
	}
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("_____________________________________")
		fmt.Println("|               MENU                 |")
		fmt.Println("|                                    |")
		fmt.Println("|      1. Informations Personnage    |")
		fmt.Println("|                                    |")
		fmt.Println("|      2. Acceder a l'inventaire     |")
		fmt.Println("|                                    |")
		fmt.Println("|      3. Parler au marchand         |")
		fmt.Println("|                                    |")
		fmt.Println("|      4. Parler au forgeron         |")
		fmt.Println("|                                    |")
		fmt.Println("|      5. Combattre un monstre       |")
		fmt.Println("|                                    |")
		fmt.Println("|      6. Qui sont ils ?             |")
		fmt.Println("|                                    |")
		fmt.Println("|      7. Quitter le jeu             |")
		fmt.Println("|____________________________________|")
		fmt.Println("")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			personnage.displayinfo()
		case "2":
			personnage.accessInventory()
		case "3":
			personnage.marchand()
		case "4":
			personnage.forgeron()
		case "5":
			personnage.menuCombat()
		case "6":
			quiSontIls()
		case "7":
			fmt.Println("Vous avez quitté le jeu")
			os.Exit(0)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide")
		}
	}
}

// Création du personnage //

type Personnage struct {
	nom           string
	classe        string
	niveau        int
	exp           int
	expmax        int
	pvmax         int
	pv            int
	Pa            int
	inventaire    []string
	skill         []string
	gold          int
	equipement    *Equipement
	initiative    int
	mana          int
	manamax       int
	upgrade       int
	inventairemax int
}

type Equipement struct {
	casque   string
	plastron string
	bottes   string
}

func verifname(name string) bool {
	if name == "" {
		return false
	} else {
		for _, v := range name {
			if v >= '0' && v <= '9' {
				return false
			}
		}
		return true
	}
}

func chooseclass(name string) Personnage {
	var classe string
	var personnage Personnage
	fmt.Println("______________________________________")
	fmt.Println("|                                     |")
	fmt.Println("|     Choisissez votre classe :       |")
	fmt.Println("|                                     |")
	fmt.Println("|            1. Samurai               |")
	fmt.Println("|                                     |")
	fmt.Println("|            2. Assassin              |")
	fmt.Println("|                                     |")
	fmt.Println("|            3. Indigène              |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&classe)
	switch classe {
	case "1":
		fmt.Println("___________________________________________")
		fmt.Println("|                                          |")
		fmt.Println("|     Vous avez choisi la classe Samurai   |")
		fmt.Println("|                                          |")
		fmt.Println("|   Cette classe est la plus polyvalente   |")
		fmt.Println("|__________________________________________|")
		personnage = Personnage{name, "Samurai", 1, 0, 80, 120, 60, 10, []string{"Potion de vie", "Potion de vie", "Potion de vie"}, []string{"Coup de Katana"}, 200, &Equipement{}, 3, 40, 80, 0, 10}
	case "2":
		fmt.Println("___________________________________________")
		fmt.Println("|                                          |")
		fmt.Println("|    Vous avez choisi la classe Assassin   |")
		fmt.Println("|                                          |")
		fmt.Println("|     Cette classe est la plus Agile       |")
		fmt.Println("|__________________________________________|")
		personnage = Personnage{name, "Assassin", 1, 0, 80, 70, 35, 18, []string{"Potion de vie", "Potion de vie", "Potion de poison"}, []string{"Coup de poignard"}, 300, &Equipement{}, 8, 40, 80, 0, 10}
	case "3":
		fmt.Println("___________________________________________")
		fmt.Println("|                                          |")
		fmt.Println("|   Vous avez choisi la classe Indigène    |")
		fmt.Println("|                                          |")
		fmt.Println("|  Cette classe est la plus faible         |")
		fmt.Println("|__________________________________________|")
		personnage = Personnage{name, "Indigène", 1, 0, 80, 50, 25, 6, []string{}, []string{"Coup de Poing"}, 50, &Equipement{}, 1, 25, 50, 0, 10}
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
	return personnage
}

// Fonctions personnage //

func (p *Personnage) displayinfo() {
	fmt.Println("_________________________________________")
	fmt.Println("Nom: ", p.nom)
	fmt.Println("Classe: ", p.classe)
	fmt.Println("Niveau: ", p.niveau)
	fmt.Println("Expérience: ", p.exp, "/", p.expmax)
	fmt.Println("PV :", p.pv, "/", p.pvmax)
	fmt.Println("PA: ", p.Pa)
	fmt.Println("Inventaire: ", p.inventaire)
	fmt.Println("Skill: ", p.skill)
	fmt.Println("Gold: ", p.gold)
	fmt.Println("Equipement: ", p.equipement)
	fmt.Println("Initiative: ", p.initiative)
	fmt.Println("Mana: ", p.mana, "/", p.manamax)
	fmt.Println("_________________________________________")
	fmt.Println("")
}

func (p *Personnage) dead() {
	if p.pv <= 0 {
		fmt.Println("_________________________________________")
		fmt.Println("|									     |")
		fmt.Println("|									     |")
		fmt.Println("|          Vous êtes mort !             |")
		fmt.Println("|									     |")
		fmt.Println("|_______________________________________|")
		p.pv = p.pvmax / 2
		fmt.Println("Vous avez  ressuscité et vous avez maintenant", p.pv, "pv")
	}
}

// Fonctions inventaire //

func (p *Personnage) accessInventory() {
	var choix string
	fmt.Println("______________________________________")
	fmt.Println("|            INVENTAIRE               |")
	fmt.Println("|                                     |")
	fmt.Println("|     1. Voir l'inventaire            |")
	fmt.Println("|     2. Boire une potion de vie      |")
	fmt.Println("|     3. Boire une potion de mana     |")
	fmt.Println("|     4. Boire une potion de poison   |")
	fmt.Println("|     5. Apprendre un sort            |")
	fmt.Println("|     6. Equiper un item              |")
	fmt.Println("|     7. Déséquiper un item           |")
	fmt.Println("|     8.Augmenter inventaire | 10 gold|")
	fmt.Println("|     9.Retourner au menu             |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		p.displayinventory()
	case "2":
		p.drinkPotion("Potion de vie")
	case "3":
		p.drinkPotion("Potion de mana")
	case "4":
		p.drinkPotion("Potion de poison")
	case "5":
		p.spellBook()
	case "6":
		p.equipItem("")
	case "7":
		p.unequipItem("")
	case "8":
		p.upgradeInventory()
	case "9":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) addInventory(item string) {
	if len(p.inventaire) < p.inventairemax {
		p.inventaire = append(p.inventaire, item)
	} else {
		fmt.Println("Votre inventaire est plein")
	}
}

func (p *Personnage) removeInventory(item string) {
	for i := range p.inventaire {
		if p.inventaire[i] == item {
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
			break
		}
	}
}

func (p *Personnage) inInventory(item string) bool {
	for _, v := range p.inventaire {
		if v == item {
			return true
		}
	}
	return false
}

func (p *Personnage) upgradeInventory() {
	if p.upgrade < 3 {
		if p.gold >= 30 {
			p.inventairemax = p.inventairemax + 10
			p.gold = p.gold - 30
			fmt.Println("Vous avez augmenté votre inventaire de 10 slots pour 30 gold")
			p.upgrade = p.upgrade + 1
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	} else {
		fmt.Println("Vous avez utilisé toutes vos upgrades pour l'inventaire")
	}
}

func (p *Personnage) displayinventory() {
	fmt.Println("______________________________________")
	fmt.Println("|            INVENTAIRE               |")
	fmt.Println("|_____________________________________|")
	for i := range p.inventaire {
		fmt.Println("|", p.inventaire[i])
	}
	fmt.Println("|_____________________________________|")
	fmt.Println("")
}

// Fonctions Potions //

func (p *Personnage) drinkPotion(potion string) {
	if p.inInventory(potion) {
		if potion == "Potion de vie" {
			potionVie(p)
			p.removeInventory(potion)
		} else if potion == "Potion de mana" {
			potionMana(p)
			p.removeInventory(potion)
		} else if potion == "Potion de poison" {
			potionPoison(p)
			p.removeInventory(potion)
		}
	} else {
		fmt.Println("Vous n'avez pas de", potion, "dans votre inventaire")
		fmt.Println("Vous pouvez en acheter dans le magasin de potions")
	}
}

func potionVie(p *Personnage) {
	if p.pv < p.pvmax {
		p.pv += 50
		if p.pv > p.pvmax {
			p.pv = p.pvmax
		}
		fmt.Println("Vous avez bu une potion de vie et vous avez maintenant", p.pv, "pv")
	} else {
		fmt.Println("Vous avez déjà le maximum de pv")
		p.addInventory("Potion de vie")
	}
}

func potionMana(p *Personnage) {
	if p.mana < p.manamax {
		p.mana += 20
		if p.mana > p.manamax {
			p.mana = p.manamax
		}
		fmt.Println("Vous avez bu une potion de mana et vous avez maintenant", p.mana, "mana")
	} else {
		fmt.Println("Vous avez déjà le maximum de mana")
		p.addInventory("Potion de mana")
	}
}

func potionPoison(p *Personnage) {
	if p.pv > 0 {
		for i := 0; i < 3; i++ {
			p.pv -= 10
			fmt.Println("Vous avez bu une potion de poison et vous avez maintenant", p.pv, "pv")
			time.Sleep(1 * time.Second)
		}
		if p.pv <= 0 {
			p.pv = 0
			p.dead()

		}
	} else {
		p.dead()
	}
}

func (p *Personnage) buyPotion() {
	var choix string
	fmt.Println("______________________________________")
	fmt.Println("|                Potions              |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|                                     |")
	fmt.Println("|    1. Potion de vie     | 5 or      |")
	fmt.Println("|    2. Potion de mana    | 8 or      |")
	fmt.Println("|    3. Potion de poison  | 10 or     |")
	fmt.Println("|    4. Boutique principal            |")
	fmt.Println("|    5. Menu                          |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		if p.gold >= 5 {
			p.addInventory("Potion de vie")
			p.gold -= 5
			fmt.Println("Vous avez acheté une potion de vie")
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "2":
		if p.gold >= 8 {
			p.addInventory("Potion de mana")
			p.gold -= 8
			fmt.Println("Vous avez acheté une potion de mana")
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "3":
		if p.gold >= 10 {
			p.addInventory("Potion de poison")
			p.gold -= 10
			fmt.Println("Vous avez acheté une potion de poison")
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "4":
		p.marchand()
	case "5":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

// Fonctions Sorts //

func (p *Personnage) spellBook() {
	var choix string
	fmt.Println("______________________________________")
	fmt.Println("|                Sorts                |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|                                     |")
	fmt.Println("|    1. Boule de feu                  |")
	fmt.Println("|    2. Menu                          |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		if p.inSpell("Boule de feu") {
			fmt.Println("Vous avez déjà appris ce sort")
		} else {
			if p.inInventory("Boule de feu") {
				p.addSpell("Boule de feu")
				fmt.Println("Vous avez appris le sort Boule de feu")
				p.removeInventory("Boule de feu")
			} else {
				fmt.Println("Vous n'avez pas de boule de feu dans votre inventaire")
				fmt.Println("Vous pouvez acheter ce sort à la boutique")
			}
		}
	case "2":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) addSpell(spell string) {
	p.skill = append(p.skill, spell)
}

func (p *Personnage) buySpell(spell string) {
	var choix string
	fmt.Println("______________________________________")
	fmt.Println("|               Sorts                 |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|                                     |")
	fmt.Println("|    1. Boule de feu      | 25 or     |")
	fmt.Println("|    2. Boutique                      |")
	fmt.Println("|    3. Menu                          |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scan(&choix)
	switch choix {
	case "1":
		if p.gold >= 25 {
			if !p.inInventory("Boule de feu") {
				p.addInventory("Boule de feu")
				p.gold -= 25
				fmt.Println("Vous avez acheté le sort Boule de feu")
			} else {
				fmt.Println("Vous avez déjà acheté ce sort")
			}
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "2":
		p.marchand()
	case "3":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) inSpell(spell string) bool {
	for _, s := range p.skill {
		if s == spell {
			return true
		}
	}
	return false
}

func (p *Personnage) coupDePoing(g *Monstre) {
	g.pv -= p.Pa
	fmt.Println("Vous avez infligé 10 dégâts à", g.nom)
	fmt.Println(g.nom, "a", g.pv, "pv sur", g.pvmax)
}

func (p *Personnage) bouleDeFeu(g *Monstre) {
	if p.inSpell("Boule de feu") {
		if p.mana >= 10 {
			p.mana -= 10
			fmt.Println("Vous avez utilisé le sort Boule de feu")
			g.pv -= 20
			fmt.Println("Vous avez infligé 20 dégâts à", g.nom)
			fmt.Println(g.nom, "a", g.pv, "pv sur", g.pvmax)
		} else {
			fmt.Println("Vous n'avez pas assez de mana")
			p.charTurn(g)
		}
	} else {
		fmt.Println("Vous n'avez pas appris ce sort")
		p.charTurn(g)
	}
}

// Fonctions Marchand //

func (p *Personnage) marchand() {
	fmt.Println("______________________________________")
	fmt.Println("|            BOUTIQUE                 |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|          1. Acheter                 |")
	fmt.Println("|          2. Vendre                  |")
	fmt.Println("|          3. Retourner au menu       |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")

	var choix string
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		p.buy()
	case "2":
		p.sell()
	case "3":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) buy() {
	fmt.Println("______________________________________")
	fmt.Println("|            BOUTIQUE                 |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|          1. Potions                 |")
	fmt.Println("|          2. Sorts                   |")
	fmt.Println("|          3. Items Fabrications      |")
	fmt.Println("|          4. Retourner au menu       |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")

	var choix string
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		p.buyPotion()
	case "2":
		p.buySpell("")
	case "3":
		p.buyItemForgeron()
	case "4":
		fmt.Println("Vous retournez au menu")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) sell() {
	fmt.Println("______________________________________")
	fmt.Println("|            Vente D'objets           |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|                                     |")
	fmt.Println("|    1. Potion de vie     | 5 or      |")
	fmt.Println("|    2. Potion de mana    | 5 or      |")
	fmt.Println("|    3. Potion de poison  | 5 or      |")
	fmt.Println("|    4. Boule de Feu      | 10 or     |")
	fmt.Println("|    5. Menu                          |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	var choix string
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		if p.inInventory("Potion de vie") {
			p.gold += 5
			p.removeInventory("Potion de vie")
			fmt.Println("Vous avez vendu une potion de vie")
		} else {
			fmt.Println("Vous n'avez pas de potion de vie dans votre inventaire")
		}
	case "2":
		if p.inInventory("Potion de mana") {
			p.gold += 5
			p.removeInventory("Potion de mana")
			fmt.Println("Vous avez vendu une potion de mana")
		} else {
			fmt.Println("Vous n'avez pas de potion de mana dans votre inventaire")
		}
	case "3":
		if p.inInventory("Potion de poison") {
			p.gold += 5
			p.removeInventory("Potion de poison")
			fmt.Println("Vous avez vendu une potion de poison")
		} else {
			fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire")
		}
	case "4":
		if p.inInventory("Boule de Feu") {
			p.gold += 10
			p.removeInventory("Boule de Feu")
			fmt.Println("Vous avez vendu Boule de Feu")
		} else {
			fmt.Println("Vous n'avez pas de Boule de Feu dans votre inventaire")
		}
	case "5":
		p.marchand()
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

// Fonctions Forgerons / items forgerons //

func (p *Personnage) buyItemForgeron() {
	fmt.Println("______________________________________")
	fmt.Println("|         Outils Fabrications         |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|                                     |")
	fmt.Println("|    1. Fourrure de loup  | 3 or     |")
	fmt.Println("|    2. Peau de Troll     | 4 or     |")
	fmt.Println("|    3. cuir de sanglier  | 2 or     |")
	fmt.Println("|    4. Plume de corbeau  | 2 or     |")
	fmt.Println("|    5. Boutique Principal            |")
	fmt.Println("|    6. Menu                          |")
	fmt.Println("|_____________________________________|")
	var choix string
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		if p.gold >= 3 {
			if !p.inInventory("Fourrure de loup") {
				p.addInventory("Fourrure de loup")
				p.gold -= 3
				fmt.Println("Vous avez acheté une fourrure de loup")
			} else {
				fmt.Println("Vous avez déjà acheté cette fourrure")
			}
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "2":
		if p.gold >= 4 {
			if !p.inInventory("Peau de Troll") {
				p.addInventory("Peau de Troll")
				p.gold -= 4
				fmt.Println("Vous avez acheté une Peau de Troll")
			} else {
				fmt.Println("Vous avez déjà acheté cette Peau de Troll")
			}
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "3":
		if p.gold >= 2 {
			if !p.inInventory("Cuir de sanglier") {
				p.addInventory("Cuir de sanglier")
				p.gold -= 2
				fmt.Println("Vous avez acheté du cuir de sanglier")
			} else {
				fmt.Println("Vous avez déjà acheté ce cuir de sanglier")
			}
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "4":
		if p.gold >= 2 {
			if !p.inInventory("Plume de corbeau") {
				p.addInventory("Plume de corbeau")
				p.gold -= 2
				fmt.Println("Vous avez acheté une Plume de corbeau")
			} else {
				fmt.Println("Vous avez déjà acheté cette Plume de corbeau")
			}
		} else {
			fmt.Println("Vous n'avez pas assez d'or")
		}
	case "5":
		p.marchand()
	case "6":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) forgeron() {
	fmt.Println("______________________________________")
	fmt.Println("|              Forgeron               |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|     1. Forger un item               |")
	fmt.Println("|     2. Menu                         |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	var choix string
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		p.forger()
	case "2":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) forger() {
	fmt.Println("______________________________________")
	fmt.Println("|               Forge                 |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|                                     |")
	fmt.Println("|    1. Forger un Casque  | 20 or     |")
	fmt.Println("|    2. Forger un plastron| 20 or     |")
	fmt.Println("|    3. Forger des bottes | 20 or     |")
	fmt.Println("|    4. Menu                          |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	var choix string
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		p.forgerCasque()
	case "2":
		p.forgerPlastron()
	case "3":
		p.forgerBottes()
	case "4":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) forgerCasque() {
	if p.inInventory("Plume de corbeau") && p.inInventory("Peau de Troll") && p.gold >= 20 {
		p.addInventory("Casque en cuir")
		p.removeInventory("Plume de corbeau")
		p.removeInventory("Peau de Troll")
		p.gold -= 20
		fmt.Println("Vous avez forger un casque en cuir")
	} else {
		fmt.Println("Vous avez besoin de 1 Plume de corbeau, 1 Peau de Troll et 20 or")
		fmt.Println("Vous pouvez acheter ces items dans la boutique du forgeron")
	}
}

func (p *Personnage) forgerPlastron() {
	if p.inInventory("Fourrure de loup") && p.inInventory("Peau de Troll") && p.gold >= 20 {
		p.addInventory("Plastron en cuir")
		p.removeInventory("Fourrure de loup")
		p.removeInventory("Peau de Troll")
		p.gold -= 20
		fmt.Println("Vous avez forger un plastron en cuir")
	} else {
		fmt.Println("Vous avez besoin de 1 Fourrure de loup, 1 Peau de Troll et 20 or")
		fmt.Println("Vous pouvez acheter ces items dans la boutique du forgeron")
	}
}

func (p *Personnage) forgerBottes() {
	if p.inInventory("Cuir de sanglier") && p.inInventory("Fourrure de loup") && p.gold >= 20 {
		p.addInventory("Bottes en cuir")
		p.removeInventory("Cuir de sanglier")
		p.removeInventory("Fourrure de loup")
		p.gold -= 20
		fmt.Println("Vous avez forger des bottes en cuir")
	} else {
		fmt.Println("Vous avez besoin d'1 cuir de sanglier, 1 Peau de Troll et 20 or")
		fmt.Println("Vous pouvez acheter ces items dans la boutique du forgeron")
	}
}

// Items //

func (p *Personnage) equipItem(item string) {
	var choix string
	fmt.Println("______________________________________")
	fmt.Println("|            EQUIPEMENT               |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|        1. Equiper un casque         |")
	fmt.Println("|        2. Equiper un plastron       |")
	fmt.Println("|        3. Equiper des bottes        |")
	fmt.Println("|        4. Retourner au menu         |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		p.equipCasque()
	case "2":
		p.equipPlastron()
	case "3":
		p.equipBottes()
	case "4":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) equipCasque() {
	if p.inInventory("Casque en cuir") {
		p.removeInventory("Casque en cuir")
		p.equipement.casque = "Casque en cuir"
		fmt.Println("Vous avez équipé un casque en cuir")
		p.pv += 5
		p.pvmax += 5
	} else {
		fmt.Println("Vous n'avez pas de casque en cuir")
		fmt.Println("Vous pouvez en forger dans la boutique du forgeron")
	}
}

func (p *Personnage) equipPlastron() {
	if p.inInventory("Plastron en cuir") {
		p.removeInventory("Plastron en cuir")
		p.equipement.plastron = "Plastron en cuir"
		fmt.Println("Vous avez équipé un plastron en cuir")
		p.pv += 10
		p.pvmax += 10
	} else {
		fmt.Println("Vous n'avez pas de plastron en cuir")
		fmt.Println("Vous pouvez en forger dans la boutique du forgeron")
	}
}

func (p *Personnage) equipBottes() {
	if p.inInventory("Bottes en cuir") {
		p.removeInventory("Bottes en cuir")
		p.equipement.bottes = "Bottes en cuir"
		fmt.Println("Vous avez équipé des bottes en cuir")
		p.pv += 5
		p.pvmax += 5
	} else {
		fmt.Println("Vous n'avez pas de bottes en cuir")
		fmt.Println("Vous pouvez en forger dans la boutique du forgeron")
	}
}

func (p *Personnage) unequipItem(item string) {
	var choix string
	fmt.Println("______________________________________")
	fmt.Println("|            EQUIPEMENT               |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|        1. Désequiper un casque      |")
	fmt.Println("|        2. Désequiper un plastron    |")
	fmt.Println("|        3. Désequiper des bottes     |")
	fmt.Println("|        4. Retourner au menu         |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		p.unequipCasque()
	case "2":
		p.unequipPlastron()
	case "3":
		p.unequipBottes()
	case "4":
		fmt.Println("Retour au menu principal")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func (p *Personnage) unequipCasque() {
	if p.equipement.casque != "" {
		p.addInventory(p.equipement.casque)
		p.equipement.casque = ""
		fmt.Println("Vous avez déséquipé un casque")
		p.pv -= 5
		p.pvmax -= 5
	} else {
		fmt.Println("Vous n'avez pas de casque")
	}
}

func (p *Personnage) unequipPlastron() {
	if p.equipement.plastron != "" {
		p.addInventory(p.equipement.plastron)
		p.equipement.plastron = ""
		fmt.Println("Vous avez déséquipé un plastron")
		p.pv -= 10
		p.pvmax -= 10
	} else {
		fmt.Println("Vous n'avez pas de plastron")
	}
}

func (p *Personnage) unequipBottes() {
	if p.equipement.bottes != "" {
		p.addInventory(p.equipement.bottes)
		p.equipement.bottes = ""
		fmt.Println("Vous avez déséquipé des bottes")
		p.pv -= 5
		p.pvmax -= 5
	} else {
		fmt.Println("Vous n'avez pas de bottes")
	}
}

// fonction de combat //

type Monstre struct {
	nom        string
	pvmax      int
	pv         int
	pa         int
	initiative int
}

func InitGoblin() Monstre {
	var goblin Monstre
	goblin.nom = "Goblin"
	goblin.pvmax = 50
	goblin.pv = 50
	goblin.pa = 5
	goblin.initiative = 5
	return goblin
}

func (g *Monstre) goblinPattern(p *Personnage) {
	var tour int
	tour++
	if tour%3 == 0 {
		p.pv -= g.pa * 2
		fmt.Println("Le goblin vous a infligé", g.pa*2, "dégâts")
	} else {
		p.pv -= g.pa
		fmt.Println("Le goblin vous a infligé", g.pa, "dégâts")
	}
}

func (p *Personnage) charTurn(g *Monstre) {
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|             Menu Combat             |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|            1 - Attaquer             |")
	fmt.Println("|            2 - Sorts                |")
	fmt.Println("|            3 - Objets               |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Println("Statistiques du personnage :")
	fmt.Println("Pv :", p.pv, "/", p.pvmax)
	fmt.Println("Mana :", p.mana, "/", p.manamax)
	fmt.Println("Statistiques du monstre :")
	fmt.Println("Pv :", g.pv, "/", g.pvmax)
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		p.coupDePoing(g)
	case 2:
		p.bouleDeFeu(g)
	case 3:
		p.useItem()
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func fight(p *Personnage, g *Monstre) {
	for p.pv > 0 && g.pv > 0 {
		if p.initiative >= g.initiative {
			p.charTurn(g)
			g.goblinPattern(p)
		} else {
			g.goblinPattern(p)
			p.charTurn(g)
		}
	}
	if p.pv <= 0 {
		p.dead()
	} else if g.pv <= 0 {
		p.win()
	}
}

func (p *Personnage) win() {
	fmt.Println("________________________________")
	fmt.Println("|                               |")
	fmt.Println("|                               |")
	fmt.Println("|                               |")
	fmt.Println("|           VICTOIRE !          |")
	fmt.Println("|                               |")
	fmt.Println("|                               |")
	fmt.Println("|_______________________________|")
	fmt.Println("")
	fmt.Println("Vous avez gagné 20 points d'expérience")
	p.exp += 20
	p.expUp()
}

func (p *Personnage) trainingfight() {
	goblin := InitGoblin()
	fight(p, &goblin)
}

// fonction de lvl //

func (p *Personnage) lvlUp() {
	p.niveau++
	p.pvmax += 10
	p.pv = p.pvmax
	p.Pa += 2
	p.initiative += 2
	fmt.Println("Vous avez gagné un niveau !")
	fmt.Println("Vous êtes maintenant niveau", p.niveau)
}

func (p *Personnage) expUp() {
	if p.exp >= p.expmax {
		p.exp -= p.expmax
		p.expmax += 20
		p.lvlUp()
	}
}

func (p *Personnage) useItem() {
	var choix int
	fmt.Println("Menu")
	for i := 0; i < len(p.inventaire); i++ {
		fmt.Println(i+1, "-", p.inventaire[i])
	}
	fmt.Println("0 - Retour")
	fmt.Scanln(&choix)
	if choix != 0 {
		if p.inventaire[choix-1] == "Potion de vie" {
			potionVie(p)
		} else if p.inventaire[choix-1] == "Potion de mana" {
			potionMana(p)
		} else {
			fmt.Println("Vous ne pouvez pas utiliser cet objet")
		}
	}
}

func (p *Personnage) menuCombat() {
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|             Menu Combat             |")
	fmt.Println("|_____________________________________|")
	fmt.Println("|       1 - Combat d'entrainement     |")
	fmt.Println("|       2 - Infos personnage          |")
	fmt.Println("|       3 - Infos Monstre             |")
	fmt.Println("|       4 - Inventaire                |")
	fmt.Println("|       5 - Menu                      |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		p.trainingfight()
	case 2:
		p.displayinfo()
	case 3:
		displayinfoMonstre()
	case 4:
		p.displayinventory()
	case 5:
		fmt.Println("Retour au menu")
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
	}
}

func displayinfoMonstre() {
	goblin := InitGoblin()
	fmt.Println("Nom :", goblin.nom)
	fmt.Println("PV :", goblin.pv, "/", goblin.pvmax)
	fmt.Println("PA :", goblin.pa)
	fmt.Println("Initiative :", goblin.initiative)
}

func quiSontIls() {
	fmt.Println("_________________________________________________")
	fmt.Println("|                 Qui sont ils ?                |")
	fmt.Println("|_______________________________________________|")
	fmt.Println("|      Partie 2 : ABBA                          |")
	fmt.Println("|      Partie 3 : Steven Spielberg              |")
	fmt.Println("|_______________________________________________|")
	fmt.Println("")
}
