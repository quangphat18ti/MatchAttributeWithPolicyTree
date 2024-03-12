package attributes

import (
	"bytes"
	"fmt"
	"strings"
)

// Attribute is a struct that holds a list of attributes
type Attribute struct {
	List map[string]int `json "list"`
}

func (a *Attribute) ToString() string {
	b := new(bytes.Buffer);

	for key, value := range a.List {
			fmt.Fprintf(b, "%s = %d\n", key, value);
	}
	return b.String();
}

func NewAttribute() *Attribute {
	return &Attribute{};
}

func parseValue(s string) int {
	// check if s is an int
	var valueInt int;
	_, err := fmt.Sscanf(s, "%d", &valueInt);
	if(err == nil) {
		return valueInt;
	}

	// check if s is a date
	date := NewDateFromString(s);
	if date != nil {
		return date.DateToInt();
	}
	return -1;
}

func parseAttribute(s string) (string, int) {
	splitSlice := strings.Split(s, "=");
	field := splitSlice[0];
	field = strings.TrimSpace(field);
	fmt.Println("field = ", field);

	value := 0;
	if(len(splitSlice) > 1) {
		valueString := splitSlice[1];
		value = parseValue(valueString);
	}

	fmt.Println("value = ", value);
	return field, value;
}

func NewAttributeFromString(s string) *Attribute{
	fmt.Println("Attributes: ", s);

	splitSlice := strings.Split(s, "|");
	splitSlice = splitSlice[1: len(splitSlice) - 1]; // remove the first and last element

	list := make(map[string]int);
	for _, v := range splitSlice {
		fmt.Println("attribute string = ", v);
		field, value := parseAttribute(v);
		list[field] = value;
		fmt.Println();
	}
	return &Attribute{
		List: list,
	};
}

func (a *Attribute) SetAttribute(lable string, value int) {
	a.List[lable] = value;
}

func (a Attribute) GetAttribute(lable string) int {
	return a.List[lable];
}