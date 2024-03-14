package main

import (
	"example/policytree/attributes"
	"example/policytree/policytree"
	"fmt"
)

func EvaluatePolicyTree(attributeStr string, policyTreeStr string) bool {
	attribute := attributes.NewAttributeFromString(attributeStr);
	policyTree := policytree.NewPolicyTree(policyTreeStr);
	return policyTree.EvaluatePolicyTree(*attribute);
}

func main() {
	attributeLists := "|Date = March 11, 2010|Manager|IT|Experience=10|"
	policyTreeString := "((Manager) AND IT) and ((Experience <= 10 or Experience > 100) and Date = March 10 -12, 2010)"
	

	fmt.Println(EvaluatePolicyTree(attributeLists, policyTreeString));
}