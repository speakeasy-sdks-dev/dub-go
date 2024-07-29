// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ClicksTimeseries struct {
	// The starting timestamp of the interval
	Start string `json:"start"`
	// The number of clicks in the interval
	Clicks float64 `json:"clicks"`
}

func (o *ClicksTimeseries) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}

func (o *ClicksTimeseries) GetClicks() float64 {
	if o == nil {
		return 0.0
	}
	return o.Clicks
}
