package country

import (
	"fmt"
	"strings"
)

func NormalizeCode(key string) (string, error) {
	key = strings.ToUpper(strings.TrimSpace(key))
	_, ok := names[key]

	if !ok {
		return "", fmt.Errorf("unknown country code: %q", key)
	}

	return key, nil
}

var names = map[string]string{
	"AR": "Argentina",
	"BO": "Bolivia",
	"BR": "Brazil",
	"CA": "Canada",
	"CL": "Chile",
	"CO": "Colombia",
	"CR": "Costa Rica",
	"CU": "Cuba",
	"DO": "Dominican Republic",
	"EC": "Ecuador",
	"SV": "El Salvador",
	"GT": "Guatemala",
	"HN": "Honduras",
	"HT": "Haiti",
	"JM": "Jamaica",
	"MX": "Mexico",
	"NI": "Nicaragua",
	"PA": "Panama",
	"PY": "Paraguay",
	"PE": "Peru",
	"PR": "Puerto Rico",
	"UY": "Uruguay",
	"US": "United States",
	"VE": "Venezuela",
	"AT": "Austria",
	"BE": "Belgium",
	"BG": "Bulgaria",
	"BY": "Belarus",
	"CH": "Switzerland",
	"CZ": "Czechia",
	"DE": "Germany",
	"DK": "Denmark",
	"ES": "Spain",
	"FI": "Finland",
	"FR": "France",
	"GB": "United Kingdom",
	"GR": "Greece",
	"HR": "Croatia",
	"HU": "Hungary",
	"IE": "Ireland",
	"IT": "Italy",
	"NL": "Netherlands",
	"NO": "Norway",
	"PL": "Poland",
	"PT": "Portugal",
	"RO": "Romania",
	"RU": "Russia",
	"SE": "Sweden",
	"UA": "Ukraine",
	"AE": "United Arab Emirates",
	"AF": "Afghanistan",
	"CN": "China",
	"ID": "Indonesia",
	"IL": "Israel",
	"IN": "India",
	"IQ": "Iraq",
	"IR": "Iran",
	"JP": "Japan",
	"KP": "North Korea",
	"KR": "South Korea",
	"MY": "Malaysia",
	"PH": "Philippines",
	"PK": "Pakistan",
	"SA": "Saudi Arabia",
	"SG": "Singapore",
	"SY": "Syria",
	"TH": "Thailand",
	"TR": "Turkey",
	"VN": "Vietnam",

	"DZ": "Algeria",
	"EG": "Egypt",
	"ET": "Ethiopia",
	"GH": "Ghana",
	"KE": "Kenya",
	"MA": "Morocco",
	"NG": "Nigeria",
	"TN": "Tunisia",
	"ZA": "South Africa",

	"AU": "Australia",
	"NZ": "New Zealand",
}
