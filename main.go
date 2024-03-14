package main

import (
	"example/policytree/attributes"
	"example/policytree/policytree"
	"fmt"
)

func main() {
	attributeLists := "|Manager|IT|Experience=5|Date = March 10, 2010|"
	policyTreeString := "(Manager AND IT) and (Experience in (1 - 10) and Date = March 11, 2025)"

	attribute := attributes.NewAttributeFromString(attributeLists);
	fmt.Println(attribute.ToString())
	
	policyTree := policytree.NewPolicyTree(policyTreeString);
	fmt.Println("\nMatch PolicyTree with Attribute: ", policyTree.EvaluatePolicyTree(*attribute));
}