package main

import (
	"example/policytree/attributes"
	"example/policytree/policytree"
	"fmt"
)

func main() {
	attributeLists := "|Manager|IT|Experience=5|Date = December 20, 2015|"
	policyTreeString := "(Manager AND IT) AND (Experience >= 3 OR Date > March 1, 2010)"
	
	attribute := attributes.NewAttributeFromString(attributeLists);
	fmt.Println(attribute.ToString())
	
	policyTree := policytree.NewPolicyTree(policyTreeString);
	fmt.Println(policyTree.EvaluatePolicyTree(*attribute));
}