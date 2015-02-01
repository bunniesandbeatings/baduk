package datafiles

import (
	"encoding/xml"
)

type Module struct {
	XMLName xml.Name `xml:"module"`
	Name    string   `xml:",attr"`
	Id      int      `xml:",attr"`
	Type    string   `xml:",attr"`
}

type DataFile struct {
	XMLName xml.Name `xml:"data"`
	Flavor  string   `xml:",attr"`
	Modules []Module
}
