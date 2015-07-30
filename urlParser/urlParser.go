package urlParser

//package main

import (
	//	"errors"
	//"fmt"
	//	"regexp"
	"strings"
)

type Request struct {
	Type       string
	TableName  string
	Parameters []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//makes a new request given a string url
func ParseURL(url string) Request {
	r := Request{"", "", make([]string, 0)}

	url = strings.ToLower(url)

	//replace symbols
	url = strings.Replace(url, "%3c", "<", -1)
	url = strings.Replace(url, "%3e", ">", -1)

	urlSections := strings.Split(url, "/")

	if len(urlSections) > 0 {
		r.Type = urlSections[0]
	}

	if len(urlSections) > 1 {
		titleParamStr := urlSections[1]

		qMarkSplit := strings.Split(titleParamStr, "?")
		r.TableName = qMarkSplit[0]

		if len(qMarkSplit) > 1 {
			paramSplit := strings.Split(qMarkSplit[1], "&")
			for _, param := range paramSplit {
				r.Parameters = append(r.Parameters, param)
			}
		}
	}

	if len(urlSections) > 2 && urlSections[2] != "" {
		r.Parameters = append(r.Parameters, "id="+urlSections[2])
	}

	return r
}

//func main() {
//s := "url?param1=foo&param2=bar/3"
//r := ParseURL(s)
//fmt.Println(r)
//}
