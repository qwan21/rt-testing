package models

type Price struct {
	Cost                float64             `json:"cost"`
	PriceType           string              `json:"priceType,omitempty"`
	RuleApplicabilities []RuleApplicability `json:"ruleApplicabilities,omitempty"`
}
