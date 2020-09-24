package models

type RuleApplicability struct {
	CodeName string `json:"codeName"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
