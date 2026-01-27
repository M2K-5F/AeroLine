package shared

type Airport struct {
	Code string
	City City
	Name string
}

var (
	Sheremetyevo      = Airport{Code: "SVO", City: Moscow, Name: "Шереметьево"}
	Domodedovo        = Airport{Code: "DME", City: Moscow, Name: "Домодедово"}
	Vnukovo           = Airport{Code: "VKO", City: Moscow, Name: "Внуково"}
	Pulkovo           = Airport{Code: "LED", City: SaintPetersburg, Name: "Пулково"}
	IstanbulNew       = Airport{Code: "IST", City: Istanbul, Name: "Стамбул"}
	SabihaGokcen      = Airport{Code: "SAW", City: Istanbul, Name: "Сабіха Гёкчен"}
	DubaiAirport      = Airport{Code: "DXB", City: Dubai, Name: "Дубай"}
	Suvarnabhumi      = Airport{Code: "BKK", City: Bangkok, Name: "Суварнабхуми"}
	DonMueang         = Airport{Code: "DMK", City: Bangkok, Name: "Донмыанг"}
	AntalyaAirport    = Airport{Code: "AYT", City: Antalya, Name: "Анталья"}
	Zvartnots         = Airport{Code: "EVN", City: Yerevan, Name: "Звартноц"}
	TbilisiAirport    = Airport{Code: "TBS", City: Tbilisi, Name: "Тбилиси"}
	AlmatyAirport     = Airport{Code: "ALA", City: Almaty, Name: "Алматы"}
	Manas             = Airport{Code: "FRU", City: Bishkek, Name: "Манас"}
	SimferopolAirport = Airport{Code: "SIP", City: Simferopol, Name: "Симферополь"}
	SochiAirport      = Airport{Code: "AER", City: Sochi, Name: "Сочи"}
	KrasnodarAirport  = Airport{Code: "KRR", City: Krasnodar, Name: "Краснодар"}
	Platov            = Airport{Code: "ROV", City: RostovOnDon, Name: "Платов"}
	Koltsovo          = Airport{Code: "SVX", City: Yekaterinburg, Name: "Кольцово"}
	KazanAirport      = Airport{Code: "KZN", City: Kazan, Name: "Казань"}
	UfaAirport        = Airport{Code: "UFA", City: Ufa, Name: "Уфа"}
	Strigino          = Airport{Code: "GOJ", City: NizhnyNovgorod, Name: "Стригино"}
	Khrabrovo         = Airport{Code: "KGD", City: Kaliningrad, Name: "Храброво"}
	Tolmachevo        = Airport{Code: "OVB", City: Novosibirsk, Name: "Толмачёво"}
	CharlesDeGaulle   = Airport{Code: "CDG", City: Paris, Name: "Шарль-де-Голль"}
	Orly              = Airport{Code: "ORY", City: Paris, Name: "Орли"}
	FrankfurtAirport  = Airport{Code: "FRA", City: Frankfurt, Name: "Франкфурт"}
	Heathrow          = Airport{Code: "LHR", City: London, Name: "Хитроу"}
	Gatwick           = Airport{Code: "LGW", City: London, Name: "Гатвик"}
	JFK               = Airport{Code: "JFK", City: NewYork, Name: "Джон Кеннеди"}
	Newark            = Airport{Code: "EWR", City: NewYork, Name: "Ньюарк"}
	PhuketAirport     = Airport{Code: "HKT", City: Phuket, Name: "Пхукет"}
	MaleAirport       = Airport{Code: "MLE", City: Male, Name: "Мале"}
	HannoverAirport   = Airport{Code: "HAJ", City: Hannover, Name: "Ганновер"}
	VaclavHavel       = Airport{Code: "PRG", City: Prague, Name: "Вацлав Гавел"}
	ViennaAirport     = Airport{Code: "VIE", City: Vienna, Name: "Вена"}
	ZurichAirport     = Airport{Code: "ZRH", City: Zurich, Name: "Цюрих"}
)
