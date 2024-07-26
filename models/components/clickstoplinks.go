// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package components

type ClicksTopLinks struct {
	// The unique ID of the short link
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Link string `json:"link"`
	// The unique ID of the short link
	ID string `json:"id"`
	// The domain of the short link
	Domain string `json:"domain"`
	// The key of the short link
	Key string `json:"key"`
	// The short link URL
	ShortLink string `json:"shortLink"`
	// The destination URL of the short link
	URL string `json:"url"`
	// The creation timestamp of the short link
	CreatedAt string `json:"createdAt"`
	// The number of clicks from this link
	Clicks float64 `json:"clicks"`
}

func (o *ClicksTopLinks) GetLink() string {
	if o == nil {
		return ""
	}
	return o.Link
}

func (o *ClicksTopLinks) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ClicksTopLinks) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *ClicksTopLinks) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *ClicksTopLinks) GetShortLink() string {
	if o == nil {
		return ""
	}
	return o.ShortLink
}

func (o *ClicksTopLinks) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *ClicksTopLinks) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *ClicksTopLinks) GetClicks() float64 {
	if o == nil {
		return 0.0
	}
	return o.Clicks
}
