package urlParser

import (
	"errors"
	"regexp"
	"strings"
)

type Request struct {
	TableName string
	Id        string
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

		matchId, err := regexp.MatchString("^id=", section)
		check(err)

		if matchTableName {
			if r.TableName == "" {
				r.TableName = section[6:]
			} else {
				err := errors.New("Error: multiple table name requests defined.")
				check(err)
			}
		} else if matchId {
			if r.Id == "" {
				r.Id = section[3:]
			} else {
				err := errors.New("Error: multiple IDs defined.")
				check(err)
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
