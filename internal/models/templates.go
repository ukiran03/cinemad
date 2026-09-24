package models

// TemplateData will be the data used while rendering the template
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	IsAuthenticated bool
	CSRFToken       string
	Flash           string
	Warn            string
	Error           string
	CurrentYear     int
	API             string
	CSSVersion      string
	StripeSecretKey string
	StripePubKey    string
}
