// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ClickEvents struct {
	Timestamp string   `json:"timestamp"`
	ClickID   string   `json:"click_id"`
	LinkID    string   `json:"link_id"`
	Domain    string   `json:"domain"`
	Key       string   `json:"key"`
	URL       string   `json:"url"`
	Continent *string  `json:"continent"`
	Country   *string  `json:"country"`
	City      *string  `json:"city"`
	Device    *string  `json:"device"`
	Browser   *string  `json:"browser"`
	Os        *string  `json:"os"`
	Referer   *string  `json:"referer"`
	IP        *string  `json:"ip"`
	Qr        *float64 `json:"qr"`
}

func (o *ClickEvents) GetTimestamp() string {
	if o == nil {
		return ""
	}
	return o.Timestamp
}

func (o *ClickEvents) GetClickID() string {
	if o == nil {
		return ""
	}
	return o.ClickID
}

func (o *ClickEvents) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

func (o *ClickEvents) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *ClickEvents) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *ClickEvents) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *ClickEvents) GetContinent() *string {
	if o == nil {
		return nil
	}
	return o.Continent
}

func (o *ClickEvents) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *ClickEvents) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *ClickEvents) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *ClickEvents) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *ClickEvents) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *ClickEvents) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *ClickEvents) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *ClickEvents) GetQr() *float64 {
	if o == nil {
		return nil
	}
	return o.Qr
}
