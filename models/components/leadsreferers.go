// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package components

type LeadsReferers struct {
	// The name of the referer. If unknown, this will be `(direct)`
	Referer string `json:"referer"`
	// The number of leads from this referer
	Leads float64 `json:"leads"`
}

func (o *LeadsReferers) GetReferer() string {
	if o == nil {
		return ""
	}
	return o.Referer
}

func (o *LeadsReferers) GetLeads() float64 {
	if o == nil {
		return 0.0
	}
	return o.Leads
}
