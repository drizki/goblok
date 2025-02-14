package goblok

var ProfaneWords = []string{
	"ancuk", "ancok", "ajig", "anjay", "anjing", "anying", "anjir", "asu", "asyu", "babangus", "babi",
	"bacol", "bacot", "bagong", "bajingan", "balegug", "banci", "bangke", "bangsat", "bedebah", "bedegong",
	"bego", "belegug", "beloon", "bencong", "bloon", "blo'on", "bodoh", "boloho", "buduk", "budug", "celeng",
	"cibai", "cibay", "cocot", "cocote", "cok", "cokil", "colai", "colay", "coli", "colmek", "conge",
	"congean", "congek", "congor", "cuk", "cukima", "cukimai", "cukimay", "dancok", "entot", "entotan", "ewe",
	"ewean", "gelo", "genjik", "germo", "gigolo", "goblo", "goblog", "goblok", "hencet", "henceut", "heunceut",
	"homo", "idiot", "itil", "jancuk", "jancok", "jablay", "jalang", "jembut", "jiancok", "jilmek", "jurig",
	"kacung", "kampang", "kampret", "kampungan", "kehed", "kenthu", "kentot", "kentu", "keparat", "kimak",
	"kintil", "kirik", "kunyuk", "kurap", "konti", "kontol", "kopet", "koplok", "lacur", "lebok", "lonte",
	"maho", "meki", "memek", "monyet", "ndas", "ndasmu", "ngehe", "ngentot", "nggateli", "nyepong", "ngewe",
	"ngocok", "pante", "pantek", "patek", "pathek", "peju", "pejuh", "pecun", "pecundang", "pelacur", "pelakor",
	"peler", "pepek", "puki", "pukima", "pukimae", "pukimak", "pukimay", "sampah", "sepong", "sial", "sialan",
	"silit", "sinting", "sontoloyo", "tai", "taik", "tempek", "tempik", "tete", "tetek", "tiembokne", "titit",
	"toket", "tolol", "ublag", "udik", "wingkeng",
}


var LeetMap = map[rune]rune{
	'0': 'o', '1': 'i', '2': 'z', '3': 'e', '4': 'a', '5': 's', '6': 'g', '7': 't', '8': 'b', '9': 'g',
	'@': 'a', '!': 'i', '$': 's', '^': 'a', '*': 'x', '(': 'c', ')': 'd', '#': 'h', '&': 'n', '%': 'p',
	'+': 't', '=': 'e', '{': 'c', '}': 'd', '[': 'c', ']': 'd', '/': 'l', '\\': 'l', '|': 'l', '<': 'v',
	'>': 'w', '~': 'n', '`': 'e', '"': 'i', '\'': 'i', '?': 'q', '_': 'u', '-': 't', 'à': 'a', 'á': 'a',
	'â': 'a', 'ä': 'a', 'ã': 'a', 'å': 'a', 'ā': 'a', 'ă': 'a', 'ą': 'a', 'ç': 'c', 'ć': 'c', 'ĉ': 'c',
	'ċ': 'c', 'č': 'c', 'è': 'e', 'é': 'e', 'ê': 'e', 'ë': 'e', 'ē': 'e', 'ĕ': 'e', 'ė': 'e', 'ę': 'e',
	'ě': 'e', 'ì': 'i', 'í': 'i', 'î': 'i', 'ï': 'i', 'ī': 'i', 'ĭ': 'i', 'į': 'i', 'ñ': 'n', 'ń': 'n', 
	'ņ': 'n', 'ň': 'n', 'ò': 'o', 'ó': 'o', 'ô': 'o', 'ö': 'o', 'õ': 'o', 'ō': 'o', 'ŏ': 'o', 'ő': 'o',
	'ù': 'u', 'ú': 'u', 'û': 'u', 'ü': 'u', 'ū': 'u', 'ŭ': 'u', 'ů': 'u', 'ű': 'u', 'ų': 'u', 'ý': 'y', 
	'ÿ': 'y', 'ŷ': 'y', 'þ': 'p', 'ß': 's', 'æ': 'a', 'œ': 'o', 'ø': 'o', 'đ': 'd', 'ð': 'd', '€': 'e',
	'£': 'l', '¥': 'y', '₹': 'r', '₩': 'w', '₽': 'r', '₺': 'l', '₦': 'n', '₫': 'd',	'₱': 'p', 'µ': 'u', 
	'₿': 'b', '฿': 'b', '©': 'c', '®': 'r', '™': 't', '±': 'p', '÷': 'd', '≤': 'l', '≥': 'g', '√': 'r',
	'∞': 'i', 'π': 'p', '∆': 'd', '∑': 's', '∫': 'i', '≈': 'a', '≠': 'n', '∂': 'd', '°': 'd', '•': 'b',
	'¶': 'p', '§': 's', '«': 'l', '»': 'r', '„': 'q', '“': 'q', '”': 'q',
}