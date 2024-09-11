// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
)

// AnalyticsCitiesCountry - The 2-letter country code of the city: https://d.to/geo
type AnalyticsCitiesCountry string

const (
	AnalyticsCitiesCountryAf AnalyticsCitiesCountry = "AF"
	AnalyticsCitiesCountryAl AnalyticsCitiesCountry = "AL"
	AnalyticsCitiesCountryDz AnalyticsCitiesCountry = "DZ"
	AnalyticsCitiesCountryAs AnalyticsCitiesCountry = "AS"
	AnalyticsCitiesCountryAd AnalyticsCitiesCountry = "AD"
	AnalyticsCitiesCountryAo AnalyticsCitiesCountry = "AO"
	AnalyticsCitiesCountryAi AnalyticsCitiesCountry = "AI"
	AnalyticsCitiesCountryAq AnalyticsCitiesCountry = "AQ"
	AnalyticsCitiesCountryAg AnalyticsCitiesCountry = "AG"
	AnalyticsCitiesCountryAr AnalyticsCitiesCountry = "AR"
	AnalyticsCitiesCountryAm AnalyticsCitiesCountry = "AM"
	AnalyticsCitiesCountryAw AnalyticsCitiesCountry = "AW"
	AnalyticsCitiesCountryAu AnalyticsCitiesCountry = "AU"
	AnalyticsCitiesCountryAt AnalyticsCitiesCountry = "AT"
	AnalyticsCitiesCountryAz AnalyticsCitiesCountry = "AZ"
	AnalyticsCitiesCountryBs AnalyticsCitiesCountry = "BS"
	AnalyticsCitiesCountryBh AnalyticsCitiesCountry = "BH"
	AnalyticsCitiesCountryBd AnalyticsCitiesCountry = "BD"
	AnalyticsCitiesCountryBb AnalyticsCitiesCountry = "BB"
	AnalyticsCitiesCountryBy AnalyticsCitiesCountry = "BY"
	AnalyticsCitiesCountryBe AnalyticsCitiesCountry = "BE"
	AnalyticsCitiesCountryBz AnalyticsCitiesCountry = "BZ"
	AnalyticsCitiesCountryBj AnalyticsCitiesCountry = "BJ"
	AnalyticsCitiesCountryBm AnalyticsCitiesCountry = "BM"
	AnalyticsCitiesCountryBt AnalyticsCitiesCountry = "BT"
	AnalyticsCitiesCountryBo AnalyticsCitiesCountry = "BO"
	AnalyticsCitiesCountryBa AnalyticsCitiesCountry = "BA"
	AnalyticsCitiesCountryBw AnalyticsCitiesCountry = "BW"
	AnalyticsCitiesCountryBv AnalyticsCitiesCountry = "BV"
	AnalyticsCitiesCountryBr AnalyticsCitiesCountry = "BR"
	AnalyticsCitiesCountryIo AnalyticsCitiesCountry = "IO"
	AnalyticsCitiesCountryBn AnalyticsCitiesCountry = "BN"
	AnalyticsCitiesCountryBg AnalyticsCitiesCountry = "BG"
	AnalyticsCitiesCountryBf AnalyticsCitiesCountry = "BF"
	AnalyticsCitiesCountryBi AnalyticsCitiesCountry = "BI"
	AnalyticsCitiesCountryKh AnalyticsCitiesCountry = "KH"
	AnalyticsCitiesCountryCm AnalyticsCitiesCountry = "CM"
	AnalyticsCitiesCountryCa AnalyticsCitiesCountry = "CA"
	AnalyticsCitiesCountryCv AnalyticsCitiesCountry = "CV"
	AnalyticsCitiesCountryKy AnalyticsCitiesCountry = "KY"
	AnalyticsCitiesCountryCf AnalyticsCitiesCountry = "CF"
	AnalyticsCitiesCountryTd AnalyticsCitiesCountry = "TD"
	AnalyticsCitiesCountryCl AnalyticsCitiesCountry = "CL"
	AnalyticsCitiesCountryCn AnalyticsCitiesCountry = "CN"
	AnalyticsCitiesCountryCx AnalyticsCitiesCountry = "CX"
	AnalyticsCitiesCountryCc AnalyticsCitiesCountry = "CC"
	AnalyticsCitiesCountryCo AnalyticsCitiesCountry = "CO"
	AnalyticsCitiesCountryKm AnalyticsCitiesCountry = "KM"
	AnalyticsCitiesCountryCg AnalyticsCitiesCountry = "CG"
	AnalyticsCitiesCountryCd AnalyticsCitiesCountry = "CD"
	AnalyticsCitiesCountryCk AnalyticsCitiesCountry = "CK"
	AnalyticsCitiesCountryCr AnalyticsCitiesCountry = "CR"
	AnalyticsCitiesCountryCi AnalyticsCitiesCountry = "CI"
	AnalyticsCitiesCountryHr AnalyticsCitiesCountry = "HR"
	AnalyticsCitiesCountryCu AnalyticsCitiesCountry = "CU"
	AnalyticsCitiesCountryCy AnalyticsCitiesCountry = "CY"
	AnalyticsCitiesCountryCz AnalyticsCitiesCountry = "CZ"
	AnalyticsCitiesCountryDk AnalyticsCitiesCountry = "DK"
	AnalyticsCitiesCountryDj AnalyticsCitiesCountry = "DJ"
	AnalyticsCitiesCountryDm AnalyticsCitiesCountry = "DM"
	AnalyticsCitiesCountryDo AnalyticsCitiesCountry = "DO"
	AnalyticsCitiesCountryEc AnalyticsCitiesCountry = "EC"
	AnalyticsCitiesCountryEg AnalyticsCitiesCountry = "EG"
	AnalyticsCitiesCountrySv AnalyticsCitiesCountry = "SV"
	AnalyticsCitiesCountryGq AnalyticsCitiesCountry = "GQ"
	AnalyticsCitiesCountryEr AnalyticsCitiesCountry = "ER"
	AnalyticsCitiesCountryEe AnalyticsCitiesCountry = "EE"
	AnalyticsCitiesCountryEt AnalyticsCitiesCountry = "ET"
	AnalyticsCitiesCountryFk AnalyticsCitiesCountry = "FK"
	AnalyticsCitiesCountryFo AnalyticsCitiesCountry = "FO"
	AnalyticsCitiesCountryFj AnalyticsCitiesCountry = "FJ"
	AnalyticsCitiesCountryFi AnalyticsCitiesCountry = "FI"
	AnalyticsCitiesCountryFr AnalyticsCitiesCountry = "FR"
	AnalyticsCitiesCountryGf AnalyticsCitiesCountry = "GF"
	AnalyticsCitiesCountryPf AnalyticsCitiesCountry = "PF"
	AnalyticsCitiesCountryTf AnalyticsCitiesCountry = "TF"
	AnalyticsCitiesCountryGa AnalyticsCitiesCountry = "GA"
	AnalyticsCitiesCountryGm AnalyticsCitiesCountry = "GM"
	AnalyticsCitiesCountryGe AnalyticsCitiesCountry = "GE"
	AnalyticsCitiesCountryDe AnalyticsCitiesCountry = "DE"
	AnalyticsCitiesCountryGh AnalyticsCitiesCountry = "GH"
	AnalyticsCitiesCountryGi AnalyticsCitiesCountry = "GI"
	AnalyticsCitiesCountryGr AnalyticsCitiesCountry = "GR"
	AnalyticsCitiesCountryGl AnalyticsCitiesCountry = "GL"
	AnalyticsCitiesCountryGd AnalyticsCitiesCountry = "GD"
	AnalyticsCitiesCountryGp AnalyticsCitiesCountry = "GP"
	AnalyticsCitiesCountryGu AnalyticsCitiesCountry = "GU"
	AnalyticsCitiesCountryGt AnalyticsCitiesCountry = "GT"
	AnalyticsCitiesCountryGn AnalyticsCitiesCountry = "GN"
	AnalyticsCitiesCountryGw AnalyticsCitiesCountry = "GW"
	AnalyticsCitiesCountryGy AnalyticsCitiesCountry = "GY"
	AnalyticsCitiesCountryHt AnalyticsCitiesCountry = "HT"
	AnalyticsCitiesCountryHm AnalyticsCitiesCountry = "HM"
	AnalyticsCitiesCountryVa AnalyticsCitiesCountry = "VA"
	AnalyticsCitiesCountryHn AnalyticsCitiesCountry = "HN"
	AnalyticsCitiesCountryHk AnalyticsCitiesCountry = "HK"
	AnalyticsCitiesCountryHu AnalyticsCitiesCountry = "HU"
	AnalyticsCitiesCountryIs AnalyticsCitiesCountry = "IS"
	AnalyticsCitiesCountryIn AnalyticsCitiesCountry = "IN"
	AnalyticsCitiesCountryID AnalyticsCitiesCountry = "ID"
	AnalyticsCitiesCountryIr AnalyticsCitiesCountry = "IR"
	AnalyticsCitiesCountryIq AnalyticsCitiesCountry = "IQ"
	AnalyticsCitiesCountryIe AnalyticsCitiesCountry = "IE"
	AnalyticsCitiesCountryIl AnalyticsCitiesCountry = "IL"
	AnalyticsCitiesCountryIt AnalyticsCitiesCountry = "IT"
	AnalyticsCitiesCountryJm AnalyticsCitiesCountry = "JM"
	AnalyticsCitiesCountryJp AnalyticsCitiesCountry = "JP"
	AnalyticsCitiesCountryJo AnalyticsCitiesCountry = "JO"
	AnalyticsCitiesCountryKz AnalyticsCitiesCountry = "KZ"
	AnalyticsCitiesCountryKe AnalyticsCitiesCountry = "KE"
	AnalyticsCitiesCountryKi AnalyticsCitiesCountry = "KI"
	AnalyticsCitiesCountryKp AnalyticsCitiesCountry = "KP"
	AnalyticsCitiesCountryKr AnalyticsCitiesCountry = "KR"
	AnalyticsCitiesCountryKw AnalyticsCitiesCountry = "KW"
	AnalyticsCitiesCountryKg AnalyticsCitiesCountry = "KG"
	AnalyticsCitiesCountryLa AnalyticsCitiesCountry = "LA"
	AnalyticsCitiesCountryLv AnalyticsCitiesCountry = "LV"
	AnalyticsCitiesCountryLb AnalyticsCitiesCountry = "LB"
	AnalyticsCitiesCountryLs AnalyticsCitiesCountry = "LS"
	AnalyticsCitiesCountryLr AnalyticsCitiesCountry = "LR"
	AnalyticsCitiesCountryLy AnalyticsCitiesCountry = "LY"
	AnalyticsCitiesCountryLi AnalyticsCitiesCountry = "LI"
	AnalyticsCitiesCountryLt AnalyticsCitiesCountry = "LT"
	AnalyticsCitiesCountryLu AnalyticsCitiesCountry = "LU"
	AnalyticsCitiesCountryMo AnalyticsCitiesCountry = "MO"
	AnalyticsCitiesCountryMg AnalyticsCitiesCountry = "MG"
	AnalyticsCitiesCountryMw AnalyticsCitiesCountry = "MW"
	AnalyticsCitiesCountryMy AnalyticsCitiesCountry = "MY"
	AnalyticsCitiesCountryMv AnalyticsCitiesCountry = "MV"
	AnalyticsCitiesCountryMl AnalyticsCitiesCountry = "ML"
	AnalyticsCitiesCountryMt AnalyticsCitiesCountry = "MT"
	AnalyticsCitiesCountryMh AnalyticsCitiesCountry = "MH"
	AnalyticsCitiesCountryMq AnalyticsCitiesCountry = "MQ"
	AnalyticsCitiesCountryMr AnalyticsCitiesCountry = "MR"
	AnalyticsCitiesCountryMu AnalyticsCitiesCountry = "MU"
	AnalyticsCitiesCountryYt AnalyticsCitiesCountry = "YT"
	AnalyticsCitiesCountryMx AnalyticsCitiesCountry = "MX"
	AnalyticsCitiesCountryFm AnalyticsCitiesCountry = "FM"
	AnalyticsCitiesCountryMd AnalyticsCitiesCountry = "MD"
	AnalyticsCitiesCountryMc AnalyticsCitiesCountry = "MC"
	AnalyticsCitiesCountryMn AnalyticsCitiesCountry = "MN"
	AnalyticsCitiesCountryMs AnalyticsCitiesCountry = "MS"
	AnalyticsCitiesCountryMa AnalyticsCitiesCountry = "MA"
	AnalyticsCitiesCountryMz AnalyticsCitiesCountry = "MZ"
	AnalyticsCitiesCountryMm AnalyticsCitiesCountry = "MM"
	AnalyticsCitiesCountryNa AnalyticsCitiesCountry = "NA"
	AnalyticsCitiesCountryNr AnalyticsCitiesCountry = "NR"
	AnalyticsCitiesCountryNp AnalyticsCitiesCountry = "NP"
	AnalyticsCitiesCountryNl AnalyticsCitiesCountry = "NL"
	AnalyticsCitiesCountryNc AnalyticsCitiesCountry = "NC"
	AnalyticsCitiesCountryNz AnalyticsCitiesCountry = "NZ"
	AnalyticsCitiesCountryNi AnalyticsCitiesCountry = "NI"
	AnalyticsCitiesCountryNe AnalyticsCitiesCountry = "NE"
	AnalyticsCitiesCountryNg AnalyticsCitiesCountry = "NG"
	AnalyticsCitiesCountryNu AnalyticsCitiesCountry = "NU"
	AnalyticsCitiesCountryNf AnalyticsCitiesCountry = "NF"
	AnalyticsCitiesCountryMk AnalyticsCitiesCountry = "MK"
	AnalyticsCitiesCountryMp AnalyticsCitiesCountry = "MP"
	AnalyticsCitiesCountryNo AnalyticsCitiesCountry = "NO"
	AnalyticsCitiesCountryOm AnalyticsCitiesCountry = "OM"
	AnalyticsCitiesCountryPk AnalyticsCitiesCountry = "PK"
	AnalyticsCitiesCountryPw AnalyticsCitiesCountry = "PW"
	AnalyticsCitiesCountryPs AnalyticsCitiesCountry = "PS"
	AnalyticsCitiesCountryPa AnalyticsCitiesCountry = "PA"
	AnalyticsCitiesCountryPg AnalyticsCitiesCountry = "PG"
	AnalyticsCitiesCountryPy AnalyticsCitiesCountry = "PY"
	AnalyticsCitiesCountryPe AnalyticsCitiesCountry = "PE"
	AnalyticsCitiesCountryPh AnalyticsCitiesCountry = "PH"
	AnalyticsCitiesCountryPn AnalyticsCitiesCountry = "PN"
	AnalyticsCitiesCountryPl AnalyticsCitiesCountry = "PL"
	AnalyticsCitiesCountryPt AnalyticsCitiesCountry = "PT"
	AnalyticsCitiesCountryPr AnalyticsCitiesCountry = "PR"
	AnalyticsCitiesCountryQa AnalyticsCitiesCountry = "QA"
	AnalyticsCitiesCountryRe AnalyticsCitiesCountry = "RE"
	AnalyticsCitiesCountryRo AnalyticsCitiesCountry = "RO"
	AnalyticsCitiesCountryRu AnalyticsCitiesCountry = "RU"
	AnalyticsCitiesCountryRw AnalyticsCitiesCountry = "RW"
	AnalyticsCitiesCountrySh AnalyticsCitiesCountry = "SH"
	AnalyticsCitiesCountryKn AnalyticsCitiesCountry = "KN"
	AnalyticsCitiesCountryLc AnalyticsCitiesCountry = "LC"
	AnalyticsCitiesCountryPm AnalyticsCitiesCountry = "PM"
	AnalyticsCitiesCountryVc AnalyticsCitiesCountry = "VC"
	AnalyticsCitiesCountryWs AnalyticsCitiesCountry = "WS"
	AnalyticsCitiesCountrySm AnalyticsCitiesCountry = "SM"
	AnalyticsCitiesCountrySt AnalyticsCitiesCountry = "ST"
	AnalyticsCitiesCountrySa AnalyticsCitiesCountry = "SA"
	AnalyticsCitiesCountrySn AnalyticsCitiesCountry = "SN"
	AnalyticsCitiesCountrySc AnalyticsCitiesCountry = "SC"
	AnalyticsCitiesCountrySl AnalyticsCitiesCountry = "SL"
	AnalyticsCitiesCountrySg AnalyticsCitiesCountry = "SG"
	AnalyticsCitiesCountrySk AnalyticsCitiesCountry = "SK"
	AnalyticsCitiesCountrySi AnalyticsCitiesCountry = "SI"
	AnalyticsCitiesCountrySb AnalyticsCitiesCountry = "SB"
	AnalyticsCitiesCountrySo AnalyticsCitiesCountry = "SO"
	AnalyticsCitiesCountryZa AnalyticsCitiesCountry = "ZA"
	AnalyticsCitiesCountryGs AnalyticsCitiesCountry = "GS"
	AnalyticsCitiesCountryEs AnalyticsCitiesCountry = "ES"
	AnalyticsCitiesCountryLk AnalyticsCitiesCountry = "LK"
	AnalyticsCitiesCountrySd AnalyticsCitiesCountry = "SD"
	AnalyticsCitiesCountrySr AnalyticsCitiesCountry = "SR"
	AnalyticsCitiesCountrySj AnalyticsCitiesCountry = "SJ"
	AnalyticsCitiesCountrySz AnalyticsCitiesCountry = "SZ"
	AnalyticsCitiesCountrySe AnalyticsCitiesCountry = "SE"
	AnalyticsCitiesCountryCh AnalyticsCitiesCountry = "CH"
	AnalyticsCitiesCountrySy AnalyticsCitiesCountry = "SY"
	AnalyticsCitiesCountryTw AnalyticsCitiesCountry = "TW"
	AnalyticsCitiesCountryTj AnalyticsCitiesCountry = "TJ"
	AnalyticsCitiesCountryTz AnalyticsCitiesCountry = "TZ"
	AnalyticsCitiesCountryTh AnalyticsCitiesCountry = "TH"
	AnalyticsCitiesCountryTl AnalyticsCitiesCountry = "TL"
	AnalyticsCitiesCountryTg AnalyticsCitiesCountry = "TG"
	AnalyticsCitiesCountryTk AnalyticsCitiesCountry = "TK"
	AnalyticsCitiesCountryTo AnalyticsCitiesCountry = "TO"
	AnalyticsCitiesCountryTt AnalyticsCitiesCountry = "TT"
	AnalyticsCitiesCountryTn AnalyticsCitiesCountry = "TN"
	AnalyticsCitiesCountryTr AnalyticsCitiesCountry = "TR"
	AnalyticsCitiesCountryTm AnalyticsCitiesCountry = "TM"
	AnalyticsCitiesCountryTc AnalyticsCitiesCountry = "TC"
	AnalyticsCitiesCountryTv AnalyticsCitiesCountry = "TV"
	AnalyticsCitiesCountryUg AnalyticsCitiesCountry = "UG"
	AnalyticsCitiesCountryUa AnalyticsCitiesCountry = "UA"
	AnalyticsCitiesCountryAe AnalyticsCitiesCountry = "AE"
	AnalyticsCitiesCountryGb AnalyticsCitiesCountry = "GB"
	AnalyticsCitiesCountryUs AnalyticsCitiesCountry = "US"
	AnalyticsCitiesCountryUm AnalyticsCitiesCountry = "UM"
	AnalyticsCitiesCountryUy AnalyticsCitiesCountry = "UY"
	AnalyticsCitiesCountryUz AnalyticsCitiesCountry = "UZ"
	AnalyticsCitiesCountryVu AnalyticsCitiesCountry = "VU"
	AnalyticsCitiesCountryVe AnalyticsCitiesCountry = "VE"
	AnalyticsCitiesCountryVn AnalyticsCitiesCountry = "VN"
	AnalyticsCitiesCountryVg AnalyticsCitiesCountry = "VG"
	AnalyticsCitiesCountryVi AnalyticsCitiesCountry = "VI"
	AnalyticsCitiesCountryWf AnalyticsCitiesCountry = "WF"
	AnalyticsCitiesCountryEh AnalyticsCitiesCountry = "EH"
	AnalyticsCitiesCountryYe AnalyticsCitiesCountry = "YE"
	AnalyticsCitiesCountryZm AnalyticsCitiesCountry = "ZM"
	AnalyticsCitiesCountryZw AnalyticsCitiesCountry = "ZW"
	AnalyticsCitiesCountryAx AnalyticsCitiesCountry = "AX"
	AnalyticsCitiesCountryBq AnalyticsCitiesCountry = "BQ"
	AnalyticsCitiesCountryCw AnalyticsCitiesCountry = "CW"
	AnalyticsCitiesCountryGg AnalyticsCitiesCountry = "GG"
	AnalyticsCitiesCountryIm AnalyticsCitiesCountry = "IM"
	AnalyticsCitiesCountryJe AnalyticsCitiesCountry = "JE"
	AnalyticsCitiesCountryMe AnalyticsCitiesCountry = "ME"
	AnalyticsCitiesCountryBl AnalyticsCitiesCountry = "BL"
	AnalyticsCitiesCountryMf AnalyticsCitiesCountry = "MF"
	AnalyticsCitiesCountryRs AnalyticsCitiesCountry = "RS"
	AnalyticsCitiesCountrySx AnalyticsCitiesCountry = "SX"
	AnalyticsCitiesCountrySs AnalyticsCitiesCountry = "SS"
	AnalyticsCitiesCountryXk AnalyticsCitiesCountry = "XK"
)

func (e AnalyticsCitiesCountry) ToPointer() *AnalyticsCitiesCountry {
	return &e
}
func (e *AnalyticsCitiesCountry) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AF":
		fallthrough
	case "AL":
		fallthrough
	case "DZ":
		fallthrough
	case "AS":
		fallthrough
	case "AD":
		fallthrough
	case "AO":
		fallthrough
	case "AI":
		fallthrough
	case "AQ":
		fallthrough
	case "AG":
		fallthrough
	case "AR":
		fallthrough
	case "AM":
		fallthrough
	case "AW":
		fallthrough
	case "AU":
		fallthrough
	case "AT":
		fallthrough
	case "AZ":
		fallthrough
	case "BS":
		fallthrough
	case "BH":
		fallthrough
	case "BD":
		fallthrough
	case "BB":
		fallthrough
	case "BY":
		fallthrough
	case "BE":
		fallthrough
	case "BZ":
		fallthrough
	case "BJ":
		fallthrough
	case "BM":
		fallthrough
	case "BT":
		fallthrough
	case "BO":
		fallthrough
	case "BA":
		fallthrough
	case "BW":
		fallthrough
	case "BV":
		fallthrough
	case "BR":
		fallthrough
	case "IO":
		fallthrough
	case "BN":
		fallthrough
	case "BG":
		fallthrough
	case "BF":
		fallthrough
	case "BI":
		fallthrough
	case "KH":
		fallthrough
	case "CM":
		fallthrough
	case "CA":
		fallthrough
	case "CV":
		fallthrough
	case "KY":
		fallthrough
	case "CF":
		fallthrough
	case "TD":
		fallthrough
	case "CL":
		fallthrough
	case "CN":
		fallthrough
	case "CX":
		fallthrough
	case "CC":
		fallthrough
	case "CO":
		fallthrough
	case "KM":
		fallthrough
	case "CG":
		fallthrough
	case "CD":
		fallthrough
	case "CK":
		fallthrough
	case "CR":
		fallthrough
	case "CI":
		fallthrough
	case "HR":
		fallthrough
	case "CU":
		fallthrough
	case "CY":
		fallthrough
	case "CZ":
		fallthrough
	case "DK":
		fallthrough
	case "DJ":
		fallthrough
	case "DM":
		fallthrough
	case "DO":
		fallthrough
	case "EC":
		fallthrough
	case "EG":
		fallthrough
	case "SV":
		fallthrough
	case "GQ":
		fallthrough
	case "ER":
		fallthrough
	case "EE":
		fallthrough
	case "ET":
		fallthrough
	case "FK":
		fallthrough
	case "FO":
		fallthrough
	case "FJ":
		fallthrough
	case "FI":
		fallthrough
	case "FR":
		fallthrough
	case "GF":
		fallthrough
	case "PF":
		fallthrough
	case "TF":
		fallthrough
	case "GA":
		fallthrough
	case "GM":
		fallthrough
	case "GE":
		fallthrough
	case "DE":
		fallthrough
	case "GH":
		fallthrough
	case "GI":
		fallthrough
	case "GR":
		fallthrough
	case "GL":
		fallthrough
	case "GD":
		fallthrough
	case "GP":
		fallthrough
	case "GU":
		fallthrough
	case "GT":
		fallthrough
	case "GN":
		fallthrough
	case "GW":
		fallthrough
	case "GY":
		fallthrough
	case "HT":
		fallthrough
	case "HM":
		fallthrough
	case "VA":
		fallthrough
	case "HN":
		fallthrough
	case "HK":
		fallthrough
	case "HU":
		fallthrough
	case "IS":
		fallthrough
	case "IN":
		fallthrough
	case "ID":
		fallthrough
	case "IR":
		fallthrough
	case "IQ":
		fallthrough
	case "IE":
		fallthrough
	case "IL":
		fallthrough
	case "IT":
		fallthrough
	case "JM":
		fallthrough
	case "JP":
		fallthrough
	case "JO":
		fallthrough
	case "KZ":
		fallthrough
	case "KE":
		fallthrough
	case "KI":
		fallthrough
	case "KP":
		fallthrough
	case "KR":
		fallthrough
	case "KW":
		fallthrough
	case "KG":
		fallthrough
	case "LA":
		fallthrough
	case "LV":
		fallthrough
	case "LB":
		fallthrough
	case "LS":
		fallthrough
	case "LR":
		fallthrough
	case "LY":
		fallthrough
	case "LI":
		fallthrough
	case "LT":
		fallthrough
	case "LU":
		fallthrough
	case "MO":
		fallthrough
	case "MG":
		fallthrough
	case "MW":
		fallthrough
	case "MY":
		fallthrough
	case "MV":
		fallthrough
	case "ML":
		fallthrough
	case "MT":
		fallthrough
	case "MH":
		fallthrough
	case "MQ":
		fallthrough
	case "MR":
		fallthrough
	case "MU":
		fallthrough
	case "YT":
		fallthrough
	case "MX":
		fallthrough
	case "FM":
		fallthrough
	case "MD":
		fallthrough
	case "MC":
		fallthrough
	case "MN":
		fallthrough
	case "MS":
		fallthrough
	case "MA":
		fallthrough
	case "MZ":
		fallthrough
	case "MM":
		fallthrough
	case "NA":
		fallthrough
	case "NR":
		fallthrough
	case "NP":
		fallthrough
	case "NL":
		fallthrough
	case "NC":
		fallthrough
	case "NZ":
		fallthrough
	case "NI":
		fallthrough
	case "NE":
		fallthrough
	case "NG":
		fallthrough
	case "NU":
		fallthrough
	case "NF":
		fallthrough
	case "MK":
		fallthrough
	case "MP":
		fallthrough
	case "NO":
		fallthrough
	case "OM":
		fallthrough
	case "PK":
		fallthrough
	case "PW":
		fallthrough
	case "PS":
		fallthrough
	case "PA":
		fallthrough
	case "PG":
		fallthrough
	case "PY":
		fallthrough
	case "PE":
		fallthrough
	case "PH":
		fallthrough
	case "PN":
		fallthrough
	case "PL":
		fallthrough
	case "PT":
		fallthrough
	case "PR":
		fallthrough
	case "QA":
		fallthrough
	case "RE":
		fallthrough
	case "RO":
		fallthrough
	case "RU":
		fallthrough
	case "RW":
		fallthrough
	case "SH":
		fallthrough
	case "KN":
		fallthrough
	case "LC":
		fallthrough
	case "PM":
		fallthrough
	case "VC":
		fallthrough
	case "WS":
		fallthrough
	case "SM":
		fallthrough
	case "ST":
		fallthrough
	case "SA":
		fallthrough
	case "SN":
		fallthrough
	case "SC":
		fallthrough
	case "SL":
		fallthrough
	case "SG":
		fallthrough
	case "SK":
		fallthrough
	case "SI":
		fallthrough
	case "SB":
		fallthrough
	case "SO":
		fallthrough
	case "ZA":
		fallthrough
	case "GS":
		fallthrough
	case "ES":
		fallthrough
	case "LK":
		fallthrough
	case "SD":
		fallthrough
	case "SR":
		fallthrough
	case "SJ":
		fallthrough
	case "SZ":
		fallthrough
	case "SE":
		fallthrough
	case "CH":
		fallthrough
	case "SY":
		fallthrough
	case "TW":
		fallthrough
	case "TJ":
		fallthrough
	case "TZ":
		fallthrough
	case "TH":
		fallthrough
	case "TL":
		fallthrough
	case "TG":
		fallthrough
	case "TK":
		fallthrough
	case "TO":
		fallthrough
	case "TT":
		fallthrough
	case "TN":
		fallthrough
	case "TR":
		fallthrough
	case "TM":
		fallthrough
	case "TC":
		fallthrough
	case "TV":
		fallthrough
	case "UG":
		fallthrough
	case "UA":
		fallthrough
	case "AE":
		fallthrough
	case "GB":
		fallthrough
	case "US":
		fallthrough
	case "UM":
		fallthrough
	case "UY":
		fallthrough
	case "UZ":
		fallthrough
	case "VU":
		fallthrough
	case "VE":
		fallthrough
	case "VN":
		fallthrough
	case "VG":
		fallthrough
	case "VI":
		fallthrough
	case "WF":
		fallthrough
	case "EH":
		fallthrough
	case "YE":
		fallthrough
	case "ZM":
		fallthrough
	case "ZW":
		fallthrough
	case "AX":
		fallthrough
	case "BQ":
		fallthrough
	case "CW":
		fallthrough
	case "GG":
		fallthrough
	case "IM":
		fallthrough
	case "JE":
		fallthrough
	case "ME":
		fallthrough
	case "BL":
		fallthrough
	case "MF":
		fallthrough
	case "RS":
		fallthrough
	case "SX":
		fallthrough
	case "SS":
		fallthrough
	case "XK":
		*e = AnalyticsCitiesCountry(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnalyticsCitiesCountry: %v", v)
	}
}

type AnalyticsCities struct {
	// The name of the city
	City string `json:"city"`
	// The 2-letter country code of the city: https://d.to/geo
	Country AnalyticsCitiesCountry `json:"country"`
	// The number of clicks from this city
	Clicks *float64 `default:"0" json:"clicks"`
	// The number of leads from this city
	Leads *float64 `default:"0" json:"leads"`
	// The number of sales from this city
	Sales *float64 `default:"0" json:"sales"`
	// The total amount of sales from this city, in cents
	SaleAmount *float64 `default:"0" json:"saleAmount"`
}

func (a AnalyticsCities) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AnalyticsCities) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AnalyticsCities) GetCity() string {
	if o == nil {
		return ""
	}
	return o.City
}

func (o *AnalyticsCities) GetCountry() AnalyticsCitiesCountry {
	if o == nil {
		return AnalyticsCitiesCountry("")
	}
	return o.Country
}

func (o *AnalyticsCities) GetClicks() *float64 {
	if o == nil {
		return nil
	}
	return o.Clicks
}

func (o *AnalyticsCities) GetLeads() *float64 {
	if o == nil {
		return nil
	}
	return o.Leads
}

func (o *AnalyticsCities) GetSales() *float64 {
	if o == nil {
		return nil
	}
	return o.Sales
}

func (o *AnalyticsCities) GetSaleAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.SaleAmount
}
