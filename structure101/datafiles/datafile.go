package datafiles

import (
	"encoding/xml"
)

type Module struct {
	XMLName xml.Name `xml:"module"`
	Name    string   `xml:"name,attr"`
	Id      int      `xml:"id,attr"`
	Type    string   `xml:"type,attr"`
}

type DataFile struct {
	XMLName xml.Name `xml:"data"`
	Flavor  string   `xml:"flavor,attr"`
	Modules []Module
}
