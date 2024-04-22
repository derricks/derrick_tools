/*
Copyright © 2022 Derrick Schneider derrick.schneider@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

var spellingBeeCmd = &cobra.Command{
	Use:   "spellingbee",
	Short: "Quiz Spelling Bee word sets",
	Run:   quizSpellingBee,
}

var spellingBeeSets = [][]string{
	{"FAIR", "FRIAR", "AFFAIR", "RIFFRAFF", "RAFFIA"},
	{"LATHE", "ATHLETE", "LETHAL", "HEALTH", "TELEHEALTH"},
	{"ACORN", "CORONA", "RANCOR", "RACCOON", "NARCO"},
	{"HEAD", "AHEAD", "HEADED", "DEADHEAD", "DEADHEADED"},
	{"TRIAL", "RATTAIL", "LARIAT", "TRAIL", "ATRIAL"},
	{"RATIO", "TRAITOR", "ORATORIO", "TRATTORIA"},
	{"FOAL", "ALOOF", "OFFAL", "LOAF", "LOOFA", "FALLOFF"},
	{"ELBOW", "BELLOW", "BELOW", "WOBBLE", "BOWEL"},
	{"DINE", "DENIED", "INDEED", "DINED", "INDIE"},
	{"LINGO", "OGLING", "GOOGLING", "GOGGLING", "LOGGING", "OILING", "LOGIN"},
	{"DICE", "DICED", "DEICE", "DEICED", "DECIDE", "DECIDED", "DEICIDE", "ICED"},
	{"INGOT", "TOGGING", "TOTING", "TOOTING", "IGNITION", "INTONING", "TONING", "NOTING"},
	{"ACHED", "HEADACHE", "CACHED"},
	{"BINGO", "GIBBON", "BONING", "BOING", "BOOING"},
	{"NOTICE", "CONCEIT", "INNOCENT", "CONTINENT", "TECTONIC", "CONNECTION", "CONTENTION", "INCONTINENT"},
	{"GENIE", "EGGING", "GENII", "ENGINE"},
	{"PIPET", "PIPETTE", "PETIT", "TIPPET", "PETITE"},
	{"CART", "TRACT", "CATARACT", "ATTRACT", "CARAT"},
	{"TACIT", "TACTIC", "CACTI", "ATTIC"},
	{"AWING", "GNAWING", "AWNING", "WAGING", "WAGGING", "WIGWAG", "WIGWAGGING", "WANING"},
	{"CLEAN", "ENLACE", "LANCE", "CANCEL", "NACELLE"},
	{"OWING", "WOOING", "GOWNING", "OWNING", "WOWING"},
	{"GALE", "EAGLE", "LEGAL", "ALLEGE", "GAGGLE", "ALGAE"},
	{"MEAN", "NAME", "ENEMA", "MANE", "AMEN"},
	{"TEAM", "TEAMMATE", "MATTE", "MEAT", "META", "MATE"},
	{"TRAP", "APPARAT", "RAPT", "APART", "PART", "TARP", "RATTRAP"},
	{"PEND", "PENDED", "DEPEND", "DEPENDED", "DEEPEN", "DEEPENED", "PENNED"},
	{"PACE", "APACE", "PEACE", "CAPE"},
	{"PEAL", "APPLE", "PALE", "LAPEL", "LEAP", "APPEAL", "PAELLA", "APPELLEE", "PLEA"},
	{"HOLE", "HELLO", "HELLHOLE"},
	{"PLEAD", "PLEADED", "DAPPLED", "PADDLE", "PEDAL", "LEAPED", "PALED", "PALLED", "LAPPED", "DAPPLE", "PEALED", "APPEALED", "APPALLED"},
	{"PECAN", "PENANCE", "PANACEA", "CANAPE"},
	{"ICON", "COIN", "IONIC", "CONIC", "ICONIC"},
	{"TEEN", "TENT", "TENET", "ENTENTE"},
	{"ATOM", "MOAT", "TOMATO"},
	{"NICE", "NIECE"},
	{"ACID", "ACIDIC", "CICADA"},
	{"INTEL", "LENTIL", "LINTEL", "LINNET", "INLET", "ENTITLE", "LENIENT"},
	{"LAIC", "LAICAL", "LILAC", "ILIAC"},
	{"DANCE", "CANED", "CANNED", "DECADENCE", "DANCED", "CADENCE"},
	{"LAMA", "LLAMA", "MAMMAL", "MALL"},
	{"ABLE", "BALE", "LABEL", "BABBLE", "BABEL"},
	{"BILE", "BIBLE", "BELIE", "LIBEL", "LIBELEE"},
	{"PAINT", "PATINA", "PINATA", "INAPT"},
	{"LOOT", "LOTTO", "TOOL", "TOLL"},
	{"PALL", "APPALL", "PAPAL", "PALAPA"},
	{"VIAL", "AVAIL", "VILLA"},
	{"ANVIL", "VILLAIN", "VANILLA", "VANILLIN"},
	{"LUNA", "ULNA", "ANNUL", "ANNUAL", "LUNULA"},
	{"GAVE", "AGAVE", "GAVAGE"},
	{"OUGHT", "OUTTHOUGHT", "TOUGH", "THOUGHT", "THOUGH"},
	{"RANT", "RATTAN", "TARTAN", "TANTRA", "TANTARA", "ARRANT"},
	{"GAIN", "AGAIN", "AGING", "GAINING", "NAGGING", "ANGINA", "GAGGING", "GANGING"},
	{"DOGE", "GEODE", "DODGE", "DOGGED", "DODGED"},
	{"DUNE", "DUNNED", "UNNEEDED", "NUDE", "ENDUED", "ENDUE", "DENUDE", "DENUDED", "UNDUE"},
	{"DUEL", "DELUDED", "ELUDE", "DULLED", "ELUDED", "DELUDE", "LULLED"},
	{"APNEA", "PANE", "NEAP", "PAEAN", "NAPE"},
	{"TOME", "TOTEM", "EMOTE", "MOTE", "MOTE"},
	{"OWNED", "WOODEN", "ENDOWED", "ENDOW", "DOWNED"},
	{"LEWD", "WELLED", "WELDED", "WELD", "DWELLED", "DWELL"},
	{"TACTICIAN", "CANTINA", "INCANT", "INTACT", "ANTIC", "TITANIC", "TANNIC"},
	{"INTENT", "NINETEEN", "TINE", "NITE"},
	{"OPTED", "DEPOT", "TOPPED", "POTTED", "TOPED"},
	{"PAIN", "APIAN", "PANINI"},
	{"PANED", "PANNED", "DEADPAN", "DEADPANNED", "NAPPED", "APPEND", "APPENDED"},
	{"PINNED", "PINED", "PINNIPED", "NIPPED"},
	{"CHAT", "CATCH", "ATTACH", "HATCH", "THATCH", "TACH"},
	{"CODE", "COOED", "DECODED", "DECODE", "COED", "DECO", "CODED"},
	{"ALONG", "LAGOON", "GALLON", "NONAGONAL", "ANALOG", "LONAGAN"},
	{"GLOAT", "GALOOT", "GLOTTAL"},
	{"IDLY", "DIDDLY", "IDYL", "IDYLL", "DILLY"},
	{"AIRMAN", "MARINA", "MARINARA"},
	{"DEBATE", "DEBATED", "ABATED", "BATED", "TABBED", "DEADBEAT", "ABETTED", "BATTED"},
	{"MAIN", "ANIMA", "MANIA", "MINIMA"},
	{"MODE", "DEMOED", "MODEM", "DOOMED", "DOMED", "DOME", "DEMO", "MOOED"},
	{"OMIT", "TOMTIT"},
	{"THEE", "TEETH", "TEETHE"},
	{"INCLINE", "LENIENCE"},
	{"LANE", "LEAN", "ANNEAL", "ELAN"},
	{"LEAVE", "VALE", "LAVE", "VALVE", "VEAL"},
	{"ROAD", "ARDOR", "DORADO"},
	{"ACTOR", "CARROT", "TRACTOR"},
	{"LORD", "DROLL", "DOLOR", "DROOL"},
	{"ORATOR", "ROTATOR", "TARO", "AORTA", "TAROT"},
	{"HATE", "HEAT", "HEATH", "THETA"},
	{"ENVIED", "DEVEIN", "DEVEINED", "ENDIVE", "DIVIDEND", "VEINED", "DIVINE", "DIVINED"},
	{"RANDO", "ADORN", "RADON"},
	{"MAYOR", "MORAY", "ARMORY"},
	{"ROAM", "ARMOR", "AROMA"},
	{"CUED", "DEDUCED", "DEUCE", "DEDUCE", "EDUCE", "EDUCED"},
	{"TABLET", "TABLE", "BLEAT", "ABETTAL", "BALLET", "BATTLE", "BEATABLE", "ABLATE", "BELATE", "EATABLE"},
	{"MARTIN", "MARTIAN", "MARTINI", "TAMARIN", "TRIMARAN"},
	{"CLEAT", "CATTLE", "ECLAT", "ACETAL"},
	{"LADLING", "LADING", "DIALING", "ADDLING"},
	{"MONEY", "YEOMEN"},
	{"GNOME", "GENOME"},
	{"GONE", "EGGNOG"},
	{"THEY", "EYETEETH"},
	{"MADE", "MADAME", "DAME", "DAMMED", "EDEMA", "EDAMAME", "MEAD"},
	{"LEMMA", "MALE", "LAME", "MEAL"},
	{"MEDIA", "MAIMED", "AIMED", "DIADEM"},
	{"FELLA", "FALAFEL", "FLEA", "LEAF"},
	{"ELFIN", "FELINE", "LIFELINE"},
	{"FALL", "ALFALFA"},
	{"FEAT", "FATE", "FETA", "TAFFETA"},
	{"LEAFLET", "FLATFEET", "FELLATE", "FETAL"},
	{"FETE", "EFFETE", "FEET", "TEFF"},
	{"FLAT", "FATAL"},
	{"BEING", "BEGGING", "BEGIN", "BINGEING", "BENIGN", "BINGE", "BEGINNING", "EBBING"},
	{"TIDIED", "DIET", "TIDE", "TIED", "EDITED", "TIDED", "DIETED"},
	{"MODEL", "LOOMED", "MODELED", "MOLDED"},
	{"IGLOO", "GIGOLO"},
	{"POINT", "PINTO", "PINOT", "PITON", "OPTION", "POTION"},
	{"POINTING", "TOPING", "OPTING", "POTTING", "TOPPING"},
	{"INTO", "NOTION"},
	{"PAPPY", "YAPPY", "PAPAYA"},
	{"MAHATMA", "MATH"},
	{"LOCAL", "COAL", "COLA", "CLOACA", "CALLALOO"},
	{"NONLOCAL", "CANOLA", "COLCANNON"},
	{"AUNT", "TAUNT", "TUNA"},
	{"AFRO", "FORA", "FARRO"},
	{"TONE", "NOTE", "NONET", "TENON", "TONNE"},
	{"CALLED", "LACED", "DECAL", "CLADE"},
	{"GIGGING", "GINNING", "INNING"},
	{"TACTICAL", "LACTIC", "CATTAIL", "ITALIC"},
	{"LAIC", "LILAC", "CILIA", "ILIAC", "LAICAL"},
	{"CHICA", "CHIA", "CHAI"},
	{"CHIT", "ITCH", "HITCH"},
	{"ACME", "CAME", "MECCA", "MACE"},
	{"ENACT", "CETACEAN", "ACCENT", "CANTEEN"},
	{"TALLIT", "ALIT", "TAIL", "ATILT", "TALI"},
	{"PARROT", "RAPPORT", "RAPTOR", "PARATROOP", "TAPROOT"},
	{"DROP", "DROOP", "PROD"},
	{"TROOP", "PORT", "TORPOR"},
	{"IDLED", "DELI", "LIED", "ELIDED", "IDLE", "LIDDED", "ELIDE", "DIDDLE", "DIDDLED"},
	{"DIED", "EDDIED"},
	{"EKED", "DEKED", "DEKE"},
	{"FOWL", "FOLLOW", "WOLF", "FLOW"},
	{"FARROW", "FOOFARAW"},
	{"TRAIN", "IRRITANT"},
	{"RABBI", "BRIAR"},
	{"BARBARIAN", "BRAIN"},
	{"HITTING", "THING", "NIGHT", "HINTING", "THINNING", "TITHING"},
	{"CLINGING", "INCLINING", "CLING"},
	{"LOAN", "LLANO"},
	{"MANOR", "MAROON"},
	{"BILLET", "BELITTLE"},
	{"TOILET", "TOILE", "TOILETTE"},
	{"COLLECT", "COLLET", "OCELOT"},
	{"COTE", "OCTET"},
	{"BELT", "BETEL", "BEETLE"},
	{"ARMY", "MARRY", "MAMMARY"},
	{"BEAM", "AMEBA", "AMEBAE"},
	{"ANIME", "MEANIE", "ANEMIA"},
	{"CAIMAN", "MINICAM", "MANIAC", "MANIC"},
	{"ANEMIC", "CINEMA", "ICEMAN"},
	{"MINCE", "EMINENCE", "IMMINENCE"},
	{"GAIT", "AGITA", "TAIGA"},
	{"GIRT", "GRIT", "TRIG"},
	{"TACET", "ACETATE"},
	{"POLO", "LOLLOP", "POOL", "LOOP", "POLL", "PLOP"},
	{"ONLY", "LOONY", "NYLON"},
	{"EELY", "YELL"},
	{"PLOY", "POLYP"},
	{"ROWAN", "NARROW"},
	{"WRATH", "THWART", "ATHWART"},
	{"WROTH", "THROW", "WORTH"},
	{"TROT", "TORT", "ROTOR", "ROOT", "TORO"},
	{"NARRATOR", "NOTATOR", "ANNOTATOR", "NONART"},
	{"HORA", "HOORAH", "HOAR"},
	{"TENTING", "IGNITE", "TEEING"},
	{"TALC", "CATCALL"},
	{"LATCH", "CATCHALL"},
	{"HALL", "HALAL"},
	{"CLAN", "CANAL"},
	{"NATAL", "LANTANA"},
	{"GOING", "ONGOING", "GONGING", "NOGGIN"},
	{"ZONING", "IONIZING", "OOZING"},
	{"ACTUAL", "TACTUAL"},
	{"DOZING", "IODIZING"},
	{"GOOD", "DOGGO"},
	{"INDIGO", "DINGO", "DOING", "DOGGING", "DODGING", "NODDING", "DONNING"},
	{"BAKE", "KEBAB", "BEAK"},
	{"ABED", "BADE", "BEAD", "BEADED", "DABBED", "BAAED"},
	{"BAKED", "BEAKED"},
	{"EBBED", "BEDDED"},
	{"BEAN", "BANE"},
	{"NABBED", "BEANED", "BANDED", "BANNED"},
	{"BAND", "BANDANA", "BANDANNA"},
	{"HARM", "HARAM"},
	{"EMIT", "TIME", "ITEM", "MITE"},
	{"TUTTED", "ETUDE", "DUET", "DUETTED"},
	{"PITIED", "TIPPED", "TEPID", "PEPTIDE"},
	{"EMITTED", "TIMED", "DEMITTED", "DEMIT"},
	{"PUTTED", "DEPUTED", "DEPUTE"},
	{"TEMP", "TEMPT"},
	{"METE", "TEEM", "MEET"},
	{"METED", "TEEMED"},
	{"TUTU", "TUTEE", "TUTTI"},
	{"FIFED", "DEFIED", "EDIFIED", "DEIFIED"},
	{"FELID", "FIELD", "FIDDLED", "FIDDLE", "FLIED", "FILLED", "FILED", "DEFILE", "DEFILED"},
	{"COUTH", "TOUCH"},
	{"MOUNT", "MUTTON"},
	{"HOOT", "TOOTH"},
	{"FELL", "FLEE", "FEEL"},
	{"COUNT", "COCONUT"},
	{"ICILY", "CYCLIC"},
	{"COLONIAL", "LACONIC", "CANONICAL", "CONICAL", "CANNOLI", "OILCAN"},
	{"FEINT", "FIFTEEN", "FINITE", "INFINITE"},
	{"EFFICIENT", "INFECT", "INEFFICIENT"},
	{"BINDI", "BIND"},
	{"BALLAD", "BALD"},
	{"WING", "WINGING", "WIGGING", "WINNING", "WINING"},
	{"WARNING", "WARRING"},
	{"GAWPING", "PAWING"},
	{"WRINGING", "WRING", "WIRING"},
	{"LOYALLY", "ALLOY", "LOYAL"},
	{"HAPPEN", "PEAHEN"},
	{"HEADLAND", "HANDLE", "HANDHELD", "HANDLED"},
	{"HANDED", "HENNAED"},
	{"HAPPENED", "DAPHNE"},
	{"COIL", "LOCI", "COLIC"},
	{"VENULE", "UNLEVEL"},
	{"VENUE", "UNEVEN"},
	{"LOBO", "BOLO", "BOLL", "BLOB"},
	{"BOOBOO", "BOOB"},
	{"LOBE", "BOBBLE"},
	{"HINT", "NINTH", "THIN"},
	{"THANE", "NEATH", "ETHANE"},
	{"HEATED", "HATED", "HATTED", "DEATH"},
	{"THEN", "TENTH"},
	{"THEM", "METH", "THEME"},
	{"MOMENT", "MEMENTO", "MONTE", "MONOTONE"},
	{"MOTTO", "MOOT"},
	{"MENTEE", "TENEMENT"},
	{"BILLABLE", "LIABLE", "LABILE", "BAILEE"},
	{"OUCH", "COUCH"},
	{"OGLE", "GOOGLE", "LOGE", "GOGGLE"},
	{"LEONINE", "ONLINE"},
	{"BEELINE", "NIBBLE"},
	{"ENTITLEMENT", "TIMELINE", "LINIMENT"},
	{"NETTLE", "LENT"},
	{"METTLE", "MELT"},
	{"NETTED", "DETENTE", "TENTED", "TENT", "DENT", "DENTED"},
	{"FLATFOOT", "FOOTFALL", "ALOFT", "FLOAT", "AFLOAT"},
	{"NEEDY", "DYNE", "DENY", "YENNED"},
	{"BLEEP", "PLEB", "PLEEB", "PEBBLE"},
	{"CAPABLE", "PEACEABLE"},
	{"PLACE", "PALACE"},
	{"CALLABLE", "CABLE"},
	{"CARPI", "PRIAPIC"},
	{"AGENT", "TANGENT", "TEENAGE", "NEGATE"},
	{"ORDAIN", "ANDROID", "ANDIRON"},
	{"DONOR", "RONDO"},
	{"CAPTAIN", "CATNIP"},
	{"CANNA", "CANCAN"},
	{"CANTATA", "CANT"},
	{"PANCETTA", "ACCEPTANCE"},
	{"CANDID", "CANID", "INDICIA", "INDICA"},
	{"DACHA", "CHAD"},
	{"RIBBING", "BRIBING", "BRINING", "BRINGING", "BRING"},
	{"BRAGGING", "BARRAGING", "BARGING", "BARING", "GRABBING", "BARGAIN", "GARBING", "BARGAINING"},
	{"BANGING", "BANNING", "BAAING", "BAGGING", "GABBING"},
	{"EPODE", "DOPE", "DOPED", "POPPED", "POOPED"},
	{"NODE", "ODEON", "DONE", "DONNED", "NODDED"},
	{"BEFELL", "FEEBLE"},
	{"BINNING", "BINGING", "GIBING"},
	{"PINTAIL", "PLIANT", "PLANTAIN", "PLAINT"},
	{"INITIALLY", "LITANY", "NATALITY", "NATTILY"},
	{"FELT", "LEFT", "FLEET", "FETTLE"},
	{"VIED", "DIVE", "IVIED", "DIVED", "DIVVIED", "DIVIDED", "DIVIDE"},
	{"HOBO", "BOHO", "BOOHOO"},
	{"BOON", "NOOB", "BONBON", "BONOBO"},
	{"BONNY", "BONY"},
	{"TONIC", "CONCOCTION"},
	{"OCTILLION", "COTILLION"},
	{"CONFIT", "NONFICTION", "FICTION"},
	{"AHOY", "YAHOO"},
	{"HALO", "ALOHA"},
	{"GAMINE", "ENIGMA", "IMAGINE", "MEANING"},
	{"MANAGE", "MANGE"},
	{"LIEGE", "GIGGLE"},
	{"NIGGLE", "NEGLIGEE", "LEGGING", "GELLING"},
	{"GAMIN", "MANNING", "NAMING", "GAMING", "IMAGINING", "IMAGING", "AIMING", "MAIMING"},
	{"NIGGLING", "GIGGLING", "GILLING", "LINING"},
	{"HINGING", "NIGH"},
	{"CHINNING", "CINCHING", "INCHING"},
	{"APPETITE", "PEPITA"},
	{"PITAPAT", "PITA"},
	{"PENITENT", "TENPIN", "INEPT"},
	{"PATENTEE", "PATENT", "PENNANT"},
	{"ACCOUNTANT", "ACCOUNT", "TOUCAN"},
	{"APEMAN", "PENMAN", "APEMEN"},
	{"CONCOCT", "COTTON"},
	{"ATOP", "POTATO"},
	{"CANNOT", "CANTON"},
	{"COHO", "HOOCH"},
	{"COYLY", "COOLLY"},
	{"ROCOCO", "CROC"},
	{"HYPO", "HOPPY"},
	{"ATONE", "ANNOTATE", "NEONATE", "NOTATE"},
	{"POPPET", "POET", "TOPE"},
	{"POTENTATE", "PANETTONE"},
	{"OPPONENT", "POTENT"},
	{"VALUE", "VULVAE", "UVULAE"},
	{"GAYLY", "LALLYGAG", "LAGGY"},
	{"GALL", "ALGAL", "GALA", "ALGA"},
	{"VULVA", "UVULA"},
	{"CORN", "CROON"},
	{"CAIRN", "CRANIA", "ARNICA", "ARANCINI"},
	{"OCARINA", "CARRION"},
	{"MONOTONED", "ODDMENT"},
	{"CHANCE", "ENHANCE"},
	{"CHANGE", "GANACHE"},
	{"PALATIAL", "PLAIT", "TILAPIA"},
	{"DOLED", "DOLE", "LOLLED", "DOLLED", "DOODLE", "DOODLED", "LODE"},
	{"ROLLBAR", "LABOR"},
	{"PEPPING", "PEEPING", "PEEING", "PIGPEN"},
	{"PIKE", "KEPI"},
	{"KEEN", "KNEE"},
	{"BAOBAB", "BOBA"},
	{"COBRA", "BARBACOA", "CAROB"},
	{"BLACK", "BLACKBALL", "CALLBACK"},
	{"BLAB", "BALL"},
	{"DOWNWIND", "WINDOW", "WOODWIND"},
	{"DAILY", "DILLYDALLY"},
	{"LADY", "DALLY"},
	{"DYAD", "DADDY"},
	{"EFFACE", "FACE", "CAFE"},
	{"AGAPE", "PAGE", "GAPE"},
	{"POGO", "GOOP"},
	{"CUTTLE", "LETTUCE", "CUTLET"},
	{"AFFECT", "FACET"},
	{"FANNED", "DEAFEN", "DEAFENED"},
	{"PHOTO", "HOTPOT"},
	{"TOPIC", "OPTIC", "OCTOPI", "PICOT"},
	{"PHOTOPIC", "PHOTIC"},
	{"CHAIR", "ARCHAIC"},
	{"EGAD", "ADAGE", "AGED", "GADDED", "GAGGED", "GAGED"},
	{"GATED", "GADGET", "TAGGED"},
	{"NOIR", "NORI", "IRON"},
	{"TWIN", "NITWIT"},
	{"TRAFFIC", "AIRCRAFT"},
	{"TARIFF", "FRITTATA"},
	{"ARID", "RADII", "RAID"},
	{"ACRID", "CARDIAC", "ARCADIA"},
	{"INITIATE", "INNATE"},
	{"ARRIVAL", "RIVAL", "VIRAL"},
	{"LIAR", "RAIL", "RIAL", "LIRA", "ARIL"},
	{"FRIARY", "FAIRY"},
	{"CAVEAT", "VACATE"},
	{"CADET", "ACTED"},
	{"DEAD", "ADDED"},
	{"DATE", "TATTED", "DATED"},
	{"ACCEDED", "ACCEDED", "ACED", "DECADE"},
	{"FUROR", "FOUR", "FROUFROU"},
	{"FROG", "FORGO"},
	{"EXACT", "EXACTA"},
	{"ADDITION", "DONATION"},
	{"DOMINO", "DOMINION"},
	{"LIGHT", "HIGHLIGHT"},
	{"BANC", "CABANA"},
	{"WARD", "DRAW", "AWARD"},
	{"BANDEAU", "UNBANNED"},
	{"PINE", "NINEPIN"},
	{"ROCK", "CORK", "CROCK"},
	{"CRAM", "MARACA"},
	{"BAFFLE", "AFFABLE"},
	{"FAIL", "FLAIL", "FILIAL"},
	{"DAMN", "ADMAN", "MADMAN"},
	{"MANY", "MYNA"},
	{"MICROMINI", "MICRON", "MORONIC", "OMICRON"},
	{"NORM", "MORON", "MORN"},
	{"DELETE", "DELT", "DELETED"},
	{"GATING", "TANNING", "GIANT", "INITIATING", "TAGGING", "TATTING"},
	{"IPECAC", "APIECE"},
	{"PIED", "PIPED", "DIPPING"},
	{"RACIAL", "RAILCAR"},
	{"BOGGLING", "GOBLIN", "GOBBLING", "OBLIGING"},
	{"NIBBLING", "BLING"},
	{"EIDETIC", "EDICT", "CITED", "DECEIT", "DIETETIC"},
	{"ENABLE", "BALEEN", "BEANBALL"},
	{"ANGEL", "ANGLE", "GLEAN"},
	{"BAGGLE", "BEAGLE", "BAGEL"},
	{"NUDGE", "NUDGED", "DENGUE", "GUNNED"},
	{"EIGHTEEN", "EIGHTEENTH", "TIGHTEN", "TIGHTENING", "HEIGHTEN", "HEIGHTENING", "NIGHTIE"},
	{"VENAL", "LEAVEN", "NAVEL"},
	{"MINDED", "MINED", "DENIM"},
}

func quizSpellingBee(cmd *cobra.Command, args []string) {
	wordSet := spellingBeeSets[rand.Intn(len(spellingBeeSets))]
	word := wordSet[rand.Intn(len(wordSet))]
	inputSet := responseFromPrompt(promptAndResponse{fmt.Sprintf("What are other Spelling Bee words for %s (separate by commas)?", word), ""})

	enteredWords := strings.Split(inputSet, ",")

	error := false
	// verify that entered words doesn't have entries not in the list
	for _, enteredWord := range enteredWords {
		if !isStringInSlice(enteredWord, wordSet) {
			error = true
			fmt.Printf("%s is not in the list of words for %s\n", enteredWord, word)
		}
	}

	// and now verify that the list isn't missing any words
	for _, validWord := range wordSet {
		if !isStringInSlice(validWord, enteredWords) && validWord != word {
			error = true
			fmt.Printf("You missed %s\n", validWord)
		}
	}

	if !error {
		fmt.Println("You got them all!")
	}
}

func init() {
	memoryquizCmd.AddCommand(spellingBeeCmd)
}
