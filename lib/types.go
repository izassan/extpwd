package lib

import "strings"

type ReplaceRule struct{
	ReplaceSource string
	ReplaceDestination string
}

func ParseOptionString(optStr string) ([]*ReplaceRule){
	var rules []*ReplaceRule
	ruleStrings := strings.Split(optStr, ";")
	for _, rule := range ruleStrings{
		src := strings.Split(rule, "=")[0]
		dest := strings.Split(rule, "=")[1]
		r := &ReplaceRule{
			ReplaceSource: src,
			ReplaceDestination: dest,
		}
		rules = append(rules, r)
	}
	return rules
}

func (r *ReplaceRule) ReplaceString(str string) string{
	return strings.Replace(str, r.ReplaceSource, r.ReplaceDestination, 1)
}
