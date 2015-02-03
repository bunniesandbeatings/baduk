package datafiles

import (
	"encoding/xml"
)

type Module struct {
	XMLName xml.Name `xml:"module"`
	Name    string   `xml:"name,attr"`
	Id      string   `xml:"id,attr"`
	Type    string   `xml:"type,attr"`
}

type Dependency struct {
	XMLName xml.Name `xml:"dependency"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Type    string   `xml:"type,attr"`
}

type DataFile struct {
	XMLName      xml.Name `xml:"data"`
	Flavor       string   `xml:"flavor,attr"`
	Modules      []Module
	Dependencies []Dependency
}

func NewDataFile(flavorName string) *DataFile{
	return &DataFile{
		Flavor:       flavorName,
		Modules:      []Module{},
		Dependencies: []Dependency{},
	}
}