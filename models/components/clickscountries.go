// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Country - The 2-letter ISO 3166-1 country code for the country associated with the location of the user. Learn more: https://d.to/geo
type Country string

const (
	CountryAf Country = "AF"
	CountryAl Country = "AL"
	CountryDz Country = "DZ"
	CountryAs Country = "AS"
	CountryAd Country = "AD"
	CountryAo Country = "AO"
	CountryAi Country = "AI"
	CountryAq Country = "AQ"
	CountryAg Country = "AG"
	CountryAr Country = "AR"
	CountryAm Country = "AM"
	CountryAw Country = "AW"
	CountryAu Country = "AU"
	CountryAt Country = "AT"
	CountryAz Country = "AZ"
	CountryBs Country = "BS"
	CountryBh Country = "BH"
	CountryBd Country = "BD"
	CountryBb Country = "BB"
	CountryBy Country = "BY"
	CountryBe Country = "BE"
	CountryBz Country = "BZ"
	CountryBj Country = "BJ"
	CountryBm Country = "BM"
	CountryBt Country = "BT"
	CountryBo Country = "BO"
	CountryBa Country = "BA"
	CountryBw Country = "BW"
	CountryBv Country = "BV"
	CountryBr Country = "BR"
	CountryIo Country = "IO"
	CountryBn Country = "BN"
	CountryBg Country = "BG"
	CountryBf Country = "BF"
	CountryBi Country = "BI"
	CountryKh Country = "KH"
	CountryCm Country = "CM"
	CountryCa Country = "CA"
	CountryCv Country = "CV"
	CountryKy Country = "KY"
	CountryCf Country = "CF"
	CountryTd Country = "TD"
	CountryCl Country = "CL"
	CountryCn Country = "CN"
	CountryCx Country = "CX"
	CountryCc Country = "CC"
	CountryCo Country = "CO"
	CountryKm Country = "KM"
	CountryCg Country = "CG"
	CountryCd Country = "CD"
	CountryCk Country = "CK"
	CountryCr Country = "CR"
	CountryCi Country = "CI"
	CountryHr Country = "HR"
	CountryCu Country = "CU"
	CountryCy Country = "CY"
	CountryCz Country = "CZ"
	CountryDk Country = "DK"
	CountryDj Country = "DJ"
	CountryDm Country = "DM"
	CountryDo Country = "DO"
	CountryEc Country = "EC"
	CountryEg Country = "EG"
	CountrySv Country = "SV"
	CountryGq Country = "GQ"
	CountryEr Country = "ER"
	CountryEe Country = "EE"
	CountryEt Country = "ET"
	CountryFk Country = "FK"
	CountryFo Country = "FO"
	CountryFj Country = "FJ"
	CountryFi Country = "FI"
	CountryFr Country = "FR"
	CountryGf Country = "GF"
	CountryPf Country = "PF"
	CountryTf Country = "TF"
	CountryGa Country = "GA"
	CountryGm Country = "GM"
	CountryGe Country = "GE"
	CountryDe Country = "DE"
	CountryGh Country = "GH"
	CountryGi Country = "GI"
	CountryGr Country = "GR"
	CountryGl Country = "GL"
	CountryGd Country = "GD"
	CountryGp Country = "GP"
	CountryGu Country = "GU"
	CountryGt Country = "GT"
	CountryGn Country = "GN"
	CountryGw Country = "GW"
	CountryGy Country = "GY"
	CountryHt Country = "HT"
	CountryHm Country = "HM"
	CountryVa Country = "VA"
	CountryHn Country = "HN"
	CountryHk Country = "HK"
	CountryHu Country = "HU"
	CountryIs Country = "IS"
	CountryIn Country = "IN"
	CountryID Country = "ID"
	CountryIr Country = "IR"
	CountryIq Country = "IQ"
	CountryIe Country = "IE"
	CountryIl Country = "IL"
	CountryIt Country = "IT"
	CountryJm Country = "JM"
	CountryJp Country = "JP"
	CountryJo Country = "JO"
	CountryKz Country = "KZ"
	CountryKe Country = "KE"
	CountryKi Country = "KI"
	CountryKp Country = "KP"
	CountryKr Country = "KR"
	CountryKw Country = "KW"
	CountryKg Country = "KG"
	CountryLa Country = "LA"
	CountryLv Country = "LV"
	CountryLb Country = "LB"
	CountryLs Country = "LS"
	CountryLr Country = "LR"
	CountryLy Country = "LY"
	CountryLi Country = "LI"
	CountryLt Country = "LT"
	CountryLu Country = "LU"
	CountryMo Country = "MO"
	CountryMg Country = "MG"
	CountryMw Country = "MW"
	CountryMy Country = "MY"
	CountryMv Country = "MV"
	CountryMl Country = "ML"
	CountryMt Country = "MT"
	CountryMh Country = "MH"
	CountryMq Country = "MQ"
	CountryMr Country = "MR"
	CountryMu Country = "MU"
	CountryYt Country = "YT"
	CountryMx Country = "MX"
	CountryFm Country = "FM"
	CountryMd Country = "MD"
	CountryMc Country = "MC"
	CountryMn Country = "MN"
	CountryMs Country = "MS"
	CountryMa Country = "MA"
	CountryMz Country = "MZ"
	CountryMm Country = "MM"
	CountryNa Country = "NA"
	CountryNr Country = "NR"
	CountryNp Country = "NP"
	CountryNl Country = "NL"
	CountryNc Country = "NC"
	CountryNz Country = "NZ"
	CountryNi Country = "NI"
	CountryNe Country = "NE"
	CountryNg Country = "NG"
	CountryNu Country = "NU"
	CountryNf Country = "NF"
	CountryMk Country = "MK"
	CountryMp Country = "MP"
	CountryNo Country = "NO"
	CountryOm Country = "OM"
	CountryPk Country = "PK"
	CountryPw Country = "PW"
	CountryPs Country = "PS"
	CountryPa Country = "PA"
	CountryPg Country = "PG"
	CountryPy Country = "PY"
	CountryPe Country = "PE"
	CountryPh Country = "PH"
	CountryPn Country = "PN"
	CountryPl Country = "PL"
	CountryPt Country = "PT"
	CountryPr Country = "PR"
	CountryQa Country = "QA"
	CountryRe Country = "RE"
	CountryRo Country = "RO"
	CountryRu Country = "RU"
	CountryRw Country = "RW"
	CountrySh Country = "SH"
	CountryKn Country = "KN"
	CountryLc Country = "LC"
	CountryPm Country = "PM"
	CountryVc Country = "VC"
	CountryWs Country = "WS"
	CountrySm Country = "SM"
	CountrySt Country = "ST"
	CountrySa Country = "SA"
	CountrySn Country = "SN"
	CountrySc Country = "SC"
	CountrySl Country = "SL"
	CountrySg Country = "SG"
	CountrySk Country = "SK"
	CountrySi Country = "SI"
	CountrySb Country = "SB"
	CountrySo Country = "SO"
	CountryZa Country = "ZA"
	CountryGs Country = "GS"
	CountryEs Country = "ES"
	CountryLk Country = "LK"
	CountrySd Country = "SD"
	CountrySr Country = "SR"
	CountrySj Country = "SJ"
	CountrySz Country = "SZ"
	CountrySe Country = "SE"
	CountryCh Country = "CH"
	CountrySy Country = "SY"
	CountryTw Country = "TW"
	CountryTj Country = "TJ"
	CountryTz Country = "TZ"
	CountryTh Country = "TH"
	CountryTl Country = "TL"
	CountryTg Country = "TG"
	CountryTk Country = "TK"
	CountryTo Country = "TO"
	CountryTt Country = "TT"
	CountryTn Country = "TN"
	CountryTr Country = "TR"
	CountryTm Country = "TM"
	CountryTc Country = "TC"
	CountryTv Country = "TV"
	CountryUg Country = "UG"
	CountryUa Country = "UA"
	CountryAe Country = "AE"
	CountryGb Country = "GB"
	CountryUs Country = "US"
	CountryUm Country = "UM"
	CountryUy Country = "UY"
	CountryUz Country = "UZ"
	CountryVu Country = "VU"
	CountryVe Country = "VE"
	CountryVn Country = "VN"
	CountryVg Country = "VG"
	CountryVi Country = "VI"
	CountryWf Country = "WF"
	CountryEh Country = "EH"
	CountryYe Country = "YE"
	CountryZm Country = "ZM"
	CountryZw Country = "ZW"
	CountryAx Country = "AX"
	CountryBq Country = "BQ"
	CountryCw Country = "CW"
	CountryGg Country = "GG"
	CountryIm Country = "IM"
	CountryJe Country = "JE"
	CountryMe Country = "ME"
	CountryBl Country = "BL"
	CountryMf Country = "MF"
	CountryRs Country = "RS"
	CountrySx Country = "SX"
	CountrySs Country = "SS"
	CountryXk Country = "XK"
)

func (e Country) ToPointer() *Country {
	return &e
}
func (e *Country) UnmarshalJSON(data []byte) error {
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
		*e = Country(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Country: %v", v)
	}
}

type ClicksCountries struct {
	// The 2-letter ISO 3166-1 country code for the country associated with the location of the user. Learn more: https://d.to/geo
	Country Country `json:"country"`
	// The number of clicks from this country
	Clicks float64 `json:"clicks"`
}

func (o *ClicksCountries) GetCountry() Country {
	if o == nil {
		return Country("")
	}
	return o.Country
}

func (o *ClicksCountries) GetClicks() float64 {
	if o == nil {
		return 0.0
	}
	return o.Clicks
}
