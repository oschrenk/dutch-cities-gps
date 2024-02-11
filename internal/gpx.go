package internal

import (
	"encoding/xml"
)

type Gpx struct {
	XMLName xml.Name `xml:"gpx"`
	Points  []Wpt    `xml:"wpt"`
}

type Wpt struct {
	XMLName xml.Name `xml:"wpt"`
	Name    string   `xml:"name"`
	Cmt     string   `xml:"cmt"`
	Lat     float32  `xml:"lat,attr"`
	Lon     float32  `xml:"lon,attr"`
}

func Parse(data []byte) ([]Wpt, error) {
	var gpx Gpx
	if err := xml.Unmarshal(data, &gpx); err != nil {
		return nil, err
	}

	return gpx.Points, nil
}
