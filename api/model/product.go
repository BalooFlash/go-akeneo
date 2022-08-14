package model

type Product struct {
	Identifier string   `json:"identifier"`
	Enabled    bool     `json:"enabled,omitempty"`
	Family     string   `json:"family,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Groups     []string `json:"groups,omitempty"`
	Parent     *string  `json:"parent,omitempty"`
	Values     Values   `json:"values"`
}

type Values map[string][]Attribute

type LinkedData struct {
	Attribute string            `json:"attribute"`
	Code      string            `json:"code"`
	Labels    map[string]string `json:"labels"`
}

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Prices []Price

type Attribute interface {
	AttributeType()
}

type TextAttribute struct {
	Data   string  `json:"data"`
	Locale *string `json:"locale"`
	Scope  *string `json:"scope"`
}

type SimpleSelectAttribute struct {
	Data       string     `json:"data"`
	Locale     *string    `json:"locale"`
	Scope      *string    `json:"scope"`
	LinkedData LinkedData `json:"linked_data"`
}

type PriceAttribute struct {
	Locale *string `json:"locale"`
	Scope  *string `json:"scope"`
	Data   Prices  `json:"data"`
}

func (a TextAttribute) AttributeType()         {}
func (a SimpleSelectAttribute) AttributeType() {}
func (a PriceAttribute) AttributeType()        {}
