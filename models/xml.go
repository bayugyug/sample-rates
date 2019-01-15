package models

import "encoding/xml"

type XMLEnvelope struct {
	XMLName xml.Name    `xml:"Envelope"`
	Subject string      `xml:"subject,"`
	Sender  Sender      `xml:"Sender,"`
	Cubes   []CubeInfo1 `xml:"Cube>Cube"`
}

type Sender struct {
	Name string `xml:"name,"`
}

type CubeInfo1 struct {
	Time  string      `xml:"time,attr"`
	Cubes []CubeInfo2 `xml:"Cube"`
}

type CubeInfo2 struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
