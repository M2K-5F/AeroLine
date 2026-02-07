package shared

type Country string

const (
	Russia        Country = "Россия"
	Turkey        Country = "Турция"
	UAE           Country = "ОАЭ"
	Thailand      Country = "Таиланд"
	Armenia       Country = "Армения"
	Georgia       Country = "Грузия"
	Kazakhstan    Country = "Казахстан"
	Kyrgyzstan    Country = "Киргизия"
	France        Country = "Франция"
	Germany       Country = "Германия"
	UnitedKingdom Country = "Великобритания"
	USA           Country = "США"
	Maldives      Country = "Мальдивы"
	CzechRepublic Country = "Чехия"
	Austria       Country = "Австрия"
	Switzerland   Country = "Швейцария"
)

type City struct {
	Code    string
	Name    string
	Country Country
}

var (
	Moscow          = City{Code: "MOW", Name: "Москва", Country: Russia}
	SaintPetersburg = City{Code: "LED", Name: "Санкт-Петербург", Country: Russia}
	Istanbul        = City{Code: "IST", Name: "Стамбул", Country: Turkey}
	Dubai           = City{Code: "DXB", Name: "Дубай", Country: UAE}
	Bangkok         = City{Code: "BKK", Name: "Бангкок", Country: Thailand}
	Antalya         = City{Code: "AYT", Name: "Анталья", Country: Turkey}
	Yerevan         = City{Code: "EVN", Name: "Ереван", Country: Armenia}
	Tbilisi         = City{Code: "TBS", Name: "Тбилиси", Country: Georgia}
	Almaty          = City{Code: "ALA", Name: "Алматы", Country: Kazakhstan}
	Bishkek         = City{Code: "FRU", Name: "Бишкек", Country: Kyrgyzstan}
	Simferopol      = City{Code: "SIP", Name: "Симферополь", Country: Russia}
	Sochi           = City{Code: "AER", Name: "Сочи", Country: Russia}
	Krasnodar       = City{Code: "KRR", Name: "Краснодар", Country: Russia}
	RostovOnDon     = City{Code: "ROV", Name: "Ростов-на-Дону", Country: Russia}
	Yekaterinburg   = City{Code: "SVX", Name: "Екатеринбург", Country: Russia}
	Kazan           = City{Code: "KZN", Name: "Казань", Country: Russia}
	Ufa             = City{Code: "UFA", Name: "Уфа", Country: Russia}
	NizhnyNovgorod  = City{Code: "GOJ", Name: "Нижний Новгород", Country: Russia}
	Kaliningrad     = City{Code: "KGD", Name: "Калининград", Country: Russia}
	Novosibirsk     = City{Code: "OVB", Name: "Новосибирск", Country: Russia}
	Paris           = City{Code: "CDG", Name: "Париж", Country: France}
	Frankfurt       = City{Code: "FRA", Name: "Франкфурт", Country: Germany}
	London          = City{Code: "LHR", Name: "Лондон", Country: UnitedKingdom}
	NewYork         = City{Code: "JFK", Name: "Нью-Йорк", Country: USA}
	Phuket          = City{Code: "HKT", Name: "Пхукет", Country: Thailand}
	Male            = City{Code: "MLE", Name: "Мале", Country: Maldives}
	Hannover        = City{Code: "HAJ", Name: "Ганновер", Country: Germany}
	Prague          = City{Code: "PRG", Name: "Прага", Country: CzechRepublic}
	Vienna          = City{Code: "VIE", Name: "Вена", Country: Austria}
	Zurich          = City{Code: "ZRH", Name: "Цюрих", Country: Switzerland}
)

var CityList = []City{
	Moscow,
	SaintPetersburg,
	Istanbul,
	Dubai,
	Bangkok,
	Antalya,
	Yerevan,
	Tbilisi,
	Almaty,
	Bishkek,
	Simferopol,
	Sochi,
	Krasnodar,
	RostovOnDon,
	Yekaterinburg,
	Kazan,
	Ufa,
	NizhnyNovgorod,
	Kaliningrad,
	Novosibirsk,
	Paris,
	Frankfurt,
	London,
	NewYork,
	Phuket,
	Male,
	Hannover,
	Prague,
	Vienna,
	Zurich,
}
