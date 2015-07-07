package urlParser

import (
	"fmt"
	"regexp"
	"strings"
)

type Request struct {
	TableName string
	Fields    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//makes a new request given a string url
func ParseURL(url string) Request {
	r := new(Request)
	urlSections := strings.Split(url, "/")

	for _, section := range urlSections {
		//first check if match table name
		matchTableName, err := regexp.MatchString("^table=", section)
		check(err)

		//then check fields
		matchFields, err := regexp.MatchString("^fields=", section)
		check(err)

		if matchTableName {
			if r.TableName == "" {
				r.TableName = section[6:]
			} else {
				fmt.Printf("Error: multiple table name requests defined.")
			}
		} else if matchFields {
			if r.Fields == "" {
				r.Fields = section[7:]
			} else {
				fmt.Printf("Error: multiple fields defined.")
			}
		}

	}

	return *r
}

/*
func main() {
	s := "table=asn/fields=cat,dog/fields="
	r := ParseURL(s)

	fmt.Printf("%s\n", r.TableName)
	fmt.Printf("%s\n", r.Fields)
}
*/
