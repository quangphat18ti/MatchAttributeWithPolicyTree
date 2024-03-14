package attributes

import (
	"fmt"
	"log"
	"time"

	"github.com/rickb777/date"
)

type MyDate struct{
	day, year int 
	month     string
	d         date.Date
}

var longMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func monthStringToMonth(month string) time.Month{
	for index, value := range longMonthNames {
		if (month == value){
			return time.Month(index + 1);
		} 
	}

	// TODO: Do After
	log.Fatal("month String is wrong format");
	return time.January;
}

func NewDate(day int, month string, year int) *MyDate{
	// @TODO: check valid date
	d := date.New(year, monthStringToMonth(month), day);
	return &MyDate{
		day: day,
		month: month,
		year: year,
		d: d,
	};
}

// format string: month day, year
func NewDateFromString(s string) *MyDate{
	var month string;
	var day, year int;

	_, err := fmt.Sscanf(s, "%s %d, %d", &month, &day, &year);
	if (err != nil) {
		log.Fatal("Date Format Error!");
		return nil;
	}

	return &MyDate{
		day: day, 
		month: month,
		year: year,
		d: date.New(year, monthStringToMonth(month), day),
	};
}

// format string: month day1 - day2, year
func DateRangeFromString(s string) (*MyDate, *MyDate){
	return nil, nil;
} 

func (d MyDate) DateToInt() int {
	return int(d.d.DaysSinceEpoch());
}

func (d MyDate) ToString() string {
	return fmt.Sprintf("%s %d, %d", d.month, d.day, d.year);
}