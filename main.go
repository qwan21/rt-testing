package main

//Calculate ...
func Calculate(product *Product, conditions []Condition) (*Offer, error) {
	if product == nil || conditions == nil {
		return nil, nil
	}

	if len(product.Components) == 0 || len(conditions) == 0 {
		return &Offer{}, nil
	}

	type Index struct {
		ComponentInd int
		PriceInd     int
	}

	type RuleWithIndex struct {
		Identify Index
		Rule     RuleApplicability
	}

	rulesWithInd := make(map[RuleWithIndex]bool)

	for ic, component := range product.Components {
		for ip, price := range component.Prices {
			for _, rule := range price.RuleApplicabilities {
				rulesWithInd[RuleWithIndex{Identify: Index{ComponentInd: ic, PriceInd: ip}, Rule: rule}] = false
			}
		}
	}

	for key := range rulesWithInd {
		for _, condition := range conditions {
			if condition.RuleName == key.Rule.CodeName {
				if condition.Value == key.Rule.Value {
					if product.Components[key.Identify.ComponentInd].IsMain {
						rulesWithInd[key] = true
					} else {
						rulesWithInd[key] = true
					}
				}
			}
		}
	}

	//rules group by indexes
	type RulesByIndex struct {
		Identify Index
		RuleList []RuleApplicability
	}
	var rulesByIndexArr []RulesByIndex
	var flag bool

	for key, v := range rulesWithInd {
		if v {
			flag = false
			for ind, v := range rulesByIndexArr {

				if v.Identify == key.Identify {
					rulesByIndexArr[ind].RuleList = append(rulesByIndexArr[ind].RuleList, key.Rule)
					flag = true
				}
			}

			var rules []RuleApplicability

			if !flag {
				rules = append(rules, key.Rule)
				rulesByIndexArr = append(rulesByIndexArr, RulesByIndex{Identify: key.Identify, RuleList: rules})
			}
		}
	}

	//select correct rules
	var validRulesByIndexArr []RulesByIndex
	for _, v := range rulesByIndexArr {
		if len(product.Components[v.Identify.ComponentInd].Prices[v.Identify.PriceInd].RuleApplicabilities) == len(v.RuleList) {
			validRulesByIndexArr = append(validRulesByIndexArr, v)

		}
	}

	if len(validRulesByIndexArr) == 0 {
		return nil, nil
	}

	type ValidRulesByComponent struct {
		Index           int
		RuleListByIndex []RulesByIndex
	}

	var validRulesByComponent []ValidRulesByComponent
	flag = false
	for _, v := range validRulesByIndexArr {
		flag = false

		for ind, val := range validRulesByComponent {
			if v.Identify.ComponentInd == val.Index {
				rulesByIndex := RulesByIndex{Identify: v.Identify, RuleList: v.RuleList}
				validRulesByComponent[ind].RuleListByIndex = append(validRulesByComponent[ind].RuleListByIndex, rulesByIndex)
				flag = true
			}
		}

		if !flag {
			rulesByIndex := []RulesByIndex{{Identify: v.Identify, RuleList: v.RuleList}}

			tmp := ValidRulesByComponent{Index: v.Identify.ComponentInd, RuleListByIndex: rulesByIndex}
			validRulesByComponent = append(validRulesByComponent, tmp)
		}
	}

	//result former
	var result Offer
	result.Name = product.Name

	var isFoundMain bool

	for _, ruleByComponent := range validRulesByComponent {
		var maxDiscount float64

		for _, v := range ruleByComponent.RuleListByIndex {
			if product.Components[v.Identify.ComponentInd].Prices[v.Identify.PriceInd].PriceType == PriceTypeDiscount {
				if maxDiscount < product.Components[v.Identify.ComponentInd].Prices[v.Identify.PriceInd].Cost {
					maxDiscount = product.Components[v.Identify.ComponentInd].Prices[v.Identify.PriceInd].Cost
				}
			}
		}

		for _, v := range ruleByComponent.RuleListByIndex {
			if product.Components[v.Identify.ComponentInd].Prices[v.Identify.PriceInd].PriceType == PriceTypeCost {
				if product.Components[v.Identify.ComponentInd].IsMain {
					isFoundMain = true
				}

				prices := []Price{{
					Cost:                product.Components[v.Identify.ComponentInd].Prices[v.Identify.PriceInd].Cost * (1 - maxDiscount/100),
					PriceType:           PriceTypeCost,
					RuleApplicabilities: v.RuleList,
				}}

				components := Component{
					Name:   product.Components[v.Identify.ComponentInd].Name,
					IsMain: product.Components[v.Identify.ComponentInd].IsMain,
					Prices: prices,
				}
				result.Components = append(result.Components, components)
				result.TotalCost.Cost += product.Components[v.Identify.ComponentInd].Prices[v.Identify.PriceInd].Cost * (1 - maxDiscount/100)
			}

		}

	}

	if !isFoundMain {
		return nil, nil
	}

	return &result, nil

}

func main() {
}
