// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package components

type Security struct {
	Token *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *Security) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
