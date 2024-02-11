package internal

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type City struct {
	Name     string
	Province string
	Lat      float32
	Lon      float32
}

func (c *City) Header() []string {
	header := []string{"name", "province", "lat", "lon"}
	return header
}

func titleCase(s string) string {
	return cases.Title(language.English, cases.NoLower).String(strings.ToLower(s))
}

func findBetween(s string) string {
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(":") + `(.*?)` + regexp.QuoteMeta(","))
	matches := rx.FindAllStringSubmatch(s, -1)
	return strings.TrimSpace(matches[0][1])
}

// PROV: DRENTHE,
func parseProvince(s string) string {
	between := findBetween(s)
	titled := cases.Title(language.English, cases.NoLower).String(strings.ToLower(between))

	return titled
}

func (c *City) ToRow() []string {
	row := []string{c.Name, c.Province, fmt.Sprintf("%f", c.Lat), fmt.Sprintf("%f", c.Lon)}
	return row
}

func (p *Wpt) ToCity() City {
	name := titleCase(p.Name)
	province := parseProvince(p.Cmt)
	city := City{}
	city.Name = name
	city.Province = province
	city.Lat = p.Lat
	city.Lon = p.Lon

	return city
}
