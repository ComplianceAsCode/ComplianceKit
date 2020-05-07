package v1

import (
	"encoding/xml"
	"fmt"
)

type Benchmark struct {
	XMLName  xml.Name  `xml:"Benchmark" json:"-"`
	Name     string    `xml:"-" json:"benchmark"`
	Profiles []Profile `xml:"Profile" json:"profiles"`
	Groups   []Group   `xml:"Group" json:"group"`
}

type Profile struct {
	XMLName      xml.Name       `xml:"Profile" json:"-"`
	Id           string         `xml:"id,attr" json:"id"`
	Title        string         `xml:"title" json:"title"`
	Description  string         `xml:"description" json:"description"`
	Selects      []Select       `xml:"select" json:"select"`
	RefineValues []RefinedValue `xml:"refine-value,omitempty" json:"refine-value,omitempty"`
}

type Select struct {
	XMLName  xml.Name `xml:"select" json:"-"`
	Idref    string   `xml:"idref,attr" json:"idref"`
	Selected bool     `xml:"selected,attr" json:"selected"`
}

type RefinedValue struct {
	XMLName  xml.Name `xml:"refine-value" json:"-"`
	Idref    string   `xml:"idref,attr" json:"idref"`
	Selector string   `xml:"selector,attr" json:"selector"`
}

type Group struct {
	XMLName     xml.Name `xml:"Group" json:"-"`
	Id          string   `xml:"id,attr" json:"id"`
	Title       string   `xml:"title" json:"title"`
	Description string   `xml:"description" json:"description"`
	Rules       []Rule   `xml:"Rule" json:"rule,omitempty"`
	Groups      []Group  `xml:"Group" json:"group,omitempty"`
}

type Rule struct {
	XMLName     xml.Name    `xml:"Rule" json:"-"`
	Id          string      `xml:"id,attr" json:"id"`
	Title       string      `xml:"title" json:"title"`
	Description string      `xml:"description" json:"description"`
	References  []Reference `xml:"reference" json:"references"`
	Rationale   string      `xml:"rationale" json:"rationale"`
	Idents      []Ident     `xml:"ident" json:"ident,omitempty"`
}

type Reference struct {
	XMLName xml.Name `xml:"reference" json:"-"`
	Catalog string   `xml:"href,attr" json:"catalog"`
	Control string   `xml:",chardata" json:"control"`
}

type Ident struct {
	XMLName xml.Name `xml:"ident" json:"-"`
	System  string   `xml:"system" json:"system"`
}

func (b Benchmark) GetProfile(profile string) Profile {
	for _, p := range b.Profiles {
		if p.Id == profile {
			return p
		}
	}
	return b.Profiles[0]
}

func (b Benchmark) GetProfileRules(profile string) []Select {
	for _, p := range b.Profiles {
		if p.Id == profile {
			return p.Selects
		}
	}
	return b.Profiles[0].Selects
}

func (b Benchmark) GetGroups() []Group {
	for _, g := range b.Groups {

		fmt.Println(g.Groups)
	}
	return b.Groups
}

func (b Benchmark) GetGroup(group string) Group {
	for _, g := range b.Groups {
		if g.Id == group {
			return g
		}
	}
	return b.Groups[0]
}

func (b Benchmark) GetRules() []Group {
	for _, g := range b.Groups {
		fmt.Println("Group: ", g.Id)
		for _, r := range g.GetRules() {
			fmt.Println("Rule: ", r.Id)
		}
		for _, sg := range g.Groups {
			for _, r := range sg.Rules {
				fmt.Println("Rule: ", r.Id)
			}
		}
	}
	return b.Groups
}
