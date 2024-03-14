package main

import (
	"example/policytree/attributes"
	"example/policytree/policytree"
	"fmt"
)

func main() {
	attributeLists := "|Manager|IT|Experience=1|Date=March 10,2010|"
	policyTreeString := "(Manager AND IT) AND (Experience >= 3 or Date = March 2-10,2010)"

	attribute := attributes.NewAttributeFromString(attributeLists);
	fmt.Println(attribute.ToString())
	
	policyTree := policytree.NewPolicyTree(policyTreeString);
	fmt.Println("\nMatch PolicyTree with Attribute: ", policyTree.EvaluatePolicyTree(*attribute));
}