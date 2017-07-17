package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// +build ignore

const url = "http://www.isbn-international.org/export_rangemessage.xml"
const header = `package isbn

type Range struct {
	Prefix string
	Agency string
	Ranges [][2]string
}

var Ranges = map[string]Range{`
const footer = `
}

`

type Group struct {
	Prefix string
	Agency string
	Rules  []Rule `xml:">Rule"`
}

type Rule struct {
	Range  string
	Length string
}

func main() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	dec := xml.NewDecoder(res.Body)

	var groups []Group
	for {
		t, err := dec.Token()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if t == nil {
			break
		}
		switch elem := t.(type) {
		case xml.StartElement:
			if elem.Name.Local == "Group" {
				var g Group
				dec.DecodeElement(&g, &elem)
				groups = append(groups, g)
			}
		}
	}

	fmt.Println(header)
	for _, g := range groups {
		fmt.Printf("%q: Range{\n", g.Prefix[4:])
		fmt.Printf("\tPrefix: %q,\n", g.Prefix[:3])
		fmt.Printf("\tAgency: %q,\n", g.Agency)
		fmt.Printf("\tRanges: [][2]string{")
		for _, r := range g.Rules {
			if r.Length == "0" {
				continue
			}
			l, _ := strconv.Atoi(r.Length)
			ranges := strings.Split(r.Range, "-")
			fmt.Printf("{%q,%q},", ranges[0][:l], ranges[1][:l])
		}
		fmt.Printf("\t},")
		fmt.Printf("\t},\n")

	}
	fmt.Println(footer)
}
