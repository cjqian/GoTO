package urlParser

import (
	"errors"
	//"fmt"
	"regexp"
	"strings"
)

type Request struct {
	TableName        string
	Id               string
	Parameters       map[string]string
	UpdateParameters map[string]string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//makes a new request given a string url
func ParseURL(url string) Request {
	r := Request{"", "", map[string]string{}, map[string]string{}}

	url = strings.ToLower(url)
	urlSections := strings.Split(url, "/")

	for _, section := range urlSections {
		//first check if match table name
		if matchTableName, err := regexp.MatchString("^table=", section); matchTableName && err == nil {
			if r.TableName == "" {
				r.TableName = section[6:]
			} else {
				err := errors.New("Error: multiple table name requests defined.")
				check(err)
			}

		} else if matchUpdate, err := regexp.MatchString("^new:", section); matchUpdate && err == nil {
			ua := strings.Split(section[4:], ",")
			for _, param := range ua {
				pa := strings.Split(param, "=")
				r.UpdateParameters[pa[0]] = pa[1]
			}
		} else if matchEquiv, err := regexp.MatchString("^.*=", section); matchEquiv && err == nil {
			ea := strings.Split(section, "=")
			r.Parameters[ea[0]] = ea[1]
		}
	}
	return r
}

/*
func main() {
	s := "table=asn/fields=cat,dog/fields="
	r := ParseURL(s)

	fmt.Printf("%s\n", r.TableName)
	fmt.Printf("%s\n", r.Fields)
}
*/
