package model

type HiscoreSkill struct {
	Name       string
	Rank       int
	Level      int
	Experience int
}

type HiscoreActivity struct {
	Name  string
	Rank  int
	Score int
}

type Hiscore struct {
	Username   string
	Skills     []HiscoreSkill
	Activities []HiscoreActivity
}

const (
	HiscoresApiUrl      = "https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws"
	HiscoresFriendlyUrl = "https://secure.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws"
)

const (
	MinigamesEmoji = "<:minigames:530206797215301632>"
)

const (
	Overall               = 0
	Attack                = 1
	Defence               = 2
	Strength              = 3
	Hitpoints             = 4
	Ranged                = 5
	Prayer                = 6
	Magic                 = 7
	Cooking               = 8
	Woodcutting           = 9
	Fletching             = 10
	Fishing               = 11
	Firemaking            = 12
	Crafting              = 13
	Smithing              = 14
	Mining                = 15
	Herblore              = 16
	Agility               = 17
	Thieving              = 18
	Slayer                = 19
	Farming               = 20
	Runecraft             = 21
	Hunter                = 22
	Construction          = 23
	League                = 24
	Bhhunter              = 25
	Bhrogue               = 26
	Cluesall              = 27
	Cluesbeginner         = 28
	Clueseasy             = 29
	Cluesmedium           = 30
	Clueshard             = 31
	Clueselite            = 32
	Cluesmaster           = 33
	Lmsrank               = 34
	Sire                  = 35
	Hydra                 = 36
	Barrows               = 37
	Bryophyta             = 38
	Callisto              = 39
	Cerberus              = 40
	Xeric                 = 41
	Xericchallenge        = 42
	Chaoselemental        = 43
	Chaosfanatic          = 44
	Zilyana               = 45
	Corporealbeast        = 46
	Crazyarchaeologist    = 47
	Dkprime               = 48
	Dkrex                 = 49
	Dksupreme             = 50
	Derangedarchaeologist = 51
	Graardor              = 52
	Mole                  = 53
	Grotesqueguardians    = 54
	Hespori               = 55
	Kalphitequeen         = 56
	Kingblackdragon       = 57
	Kraken                = 58
	Kreearra              = 59
	Kril                  = 60
	Mimic                 = 61
	Nightmare             = 62
	Obor                  = 63
	Sarachnis             = 64
	Scorpia               = 65
	Skotizo               = 66
	Gauntlet              = 67
	Corruptedgauntlet     = 68
	Theatreofblood        = 69
	Thermonucleardevil    = 70
	Tzkalzuk              = 71
	Tztokjad              = 72
	Venenatis             = 73
	Vetion                = 74
	Vorkath               = 75
	Wintertodt            = 76
	Zulcano               = 77
	Zulrah                = 78
	Offset                = League
)

func GetHiscoreName(index int) (name string) {

	switch index {
	case Overall:
		return "Overall"
	case Attack:
		return "Attack"
	case Defence:
		return "Defence"
	case Strength:
		return "Strength"
	case Hitpoints:
		return "Hitpoints"
	case Ranged:
		return "Ranged"
	case Prayer:
		return "Prayer"
	case Magic:
		return "Magic"
	case Cooking:
		return "Cooking"
	case Woodcutting:
		return "Woodcutting"
	case Fletching:
		return "Fletching"
	case Fishing:
		return "Fishing"
	case Firemaking:
		return "Firemaking"
	case Crafting:
		return "Crafting"
	case Smithing:
		return "Smithing"
	case Mining:
		return "Mining"
	case Herblore:
		return "Herblore"
	case Agility:
		return "Agility"
	case Thieving:
		return "Thieving"
	case Slayer:
		return "Slayer"
	case Farming:
		return "Farming"
	case Runecraft:
		return "Runecraft"
	case Hunter:
		return "Hunter"
	case Construction:
		return "Construction"
	case Bhhunter:
		return "BH/Hunter"
	case Bhrogue:
		return "BH/Rogue"
	case Cluesall:
		return "Clues/All"
	case Cluesbeginner:
		return "Clues/Beginner"
	case Clueseasy:
		return "Clues/Easy"
	case Cluesmedium:
		return "Clues/Medium"
	case Clueshard:
		return "Clues/Hard"
	case Clueselite:
		return "Clues/Elite"
	case Cluesmaster:
		return "Clues/Master"
	case Lmsrank:
		return "LMS/Rank"
	case Sire:
		return "Abyssal Sire"
	case Hydra:
		return "Alchemical Hydra"
	case Barrows:
		return "Barrows Chests"
	case Bryophyta:
		return "Bryophyta"
	case Callisto:
		return "Callisto"
	case Cerberus:
		return "Cerberus"
	case Xeric:
		return "Chambers of Xeric"
	case Xericchallenge:
		return "Chambers of Xeric: Challenge Mode"
	case Chaoselemental:
		return "Chaos Elemental"
	case Chaosfanatic:
		return "Chaos Fanatic"
	case Zilyana:
		return "Commander Zilyana"
	case Corporealbeast:
		return "Corporeal Beast"
	case Crazyarchaeologist:
		return "Crazy Archaeologist"
	case Dkprime:
		return "Dagannoth Prime"
	case Dkrex:
		return "Dagannoth Rex"
	case Dksupreme:
		return "Dagannoth Supreme"
	case Derangedarchaeologist:
		return "Deranged Archaeologist"
	case Graardor:
		return "General Graardor"
	case Mole:
		return "Giant Mole"
	case Grotesqueguardians:
		return "Grotesque Guardians"
	case Hespori:
		return "Hespori"
	case Kalphitequeen:
		return "Kalphite Queen"
	case Kingblackdragon:
		return "King Black Dragon"
	case Kraken:
		return "Kraken"
	case Kreearra:
		return "Kree'arra"
	case Kril:
		return "K'ril Tsutsaroth"
	case Mimic:
		return "Mimic"
	case Nightmare:
		return "Nightmare"
	case Obor:
		return "Obor"
	case Sarachnis:
		return "Sarachnis"
	case Scorpia:
		return "Scorpia"
	case Skotizo:
		return "Skotizo"
	case Gauntlet:
		return "The Gauntlet"
	case Corruptedgauntlet:
		return "The Corrupted Gauntlet"
	case Theatreofblood:
		return "Theatre of Blood"
	case Thermonucleardevil:
		return "Thermonuclear Smoke Devil"
	case Tzkalzuk:
		return "TzKal-Zuk"
	case Tztokjad:
		return "TzTok-Jad"
	case Venenatis:
		return "Venenatis"
	case Vetion:
		return "Vet'ion"
	case Vorkath:
		return "Vorkath"
	case Wintertodt:
		return "Wintertodt"
	case Zulcano:
		return "Zulcano"
	case Zulrah:
		return "Zulrah"

	default:
		return "Error"
	}

}

func GetHiscoreEmoji(index int) (emoji string) {

	switch index {
	case Overall:
		return "<:stats:529107862316908564>"
	case Attack:
		return "<:Attack:529105287664369674>"
	case Defence:
		return "<:Defence:529105287773421568>"
	case Strength:
		return "<:Strength:529105288020754452>"
	case Hitpoints:
		return "<:Hitpoints:529105289316663307>"
	case Ranged:
		return "<:Ranged:529105287534346246>"
	case Prayer:
		return "<:Prayer:529105287857307658>"
	case Magic:
		return "<:Magic:529105287861501989>"
	case Cooking:
		return "<:Cooking:529105287706312704>"
	case Woodcutting:
		return "<:Woodcutting:529105287681015831>"
	case Fletching:
		return "<:Fletching:529105287852982272>"
	case Fishing:
		return "<:Fishing:529105287878017051>"
	case Firemaking:
		return "<:Firemaking:529105287790067742>"
	case Crafting:
		return "<:Crafting:529105287727284224>"
	case Smithing:
		return "<:Smithing:529105287798325284>"
	case Mining:
		return "<:Mining:529105287819427850>"
	case Herblore:
		return "<:Herblore:529105287966097462>"
	case Agility:
		return "<:Agility:529105287718895616>"
	case Thieving:
		return "<:Thieving:529105287882342411>"
	case Slayer:
		return "<:Slayer:529105287488208898>"
	case Farming:
		return "<:Farming:529105287521501185>"
	case Runecraft:
		return "<:runecrafting:529105287806976001>"
	case Hunter:
		return "<:Hunter:529105287601455119>"
	case Construction:
		return "<:Construction:529105287651786753>"
	case Bhhunter:
		return "<:bh_hunter:529105998980448256>"
	case Bhrogue:
		return "<:bh_rogue:529105998892236868>"
	case Cluesall:
		return "<:cluescroll:529106218753720320>"
	case Cluesbeginner:
		return "<:cluescroll:529106218753720320>"
	case Clueseasy:
		return "<:cluescroll:529106218753720320>"
	case Cluesmedium:
		return "<:cluescroll:529106218753720320>"
	case Clueshard:
		return "<:cluescroll:529106218753720320>"
	case Clueselite:
		return "<:cluescroll:529106218753720320>"
	case Cluesmaster:
		return "<:cluescroll:529106218753720320>"
	case Lmsrank:
		return "<:lms:529108371811467264>"
	default:
		return ""
	}
}
