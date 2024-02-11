package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/oschrenk/dutch-cities-gps/internal"
)

func main() {
	gpxDirPtr := flag.String("gpx", "./assets", "input path to directory containg gpx files")
	outputPtr := flag.String("csv", "./.release/cities.csv", "output path to csv file")
	flag.Parse()

	if _, err := os.Stat(*gpxDirPtr); os.IsNotExist(err) {
		log.Fatal(*gpxDirPtr, "does not exist")
	}

	var cities []internal.City
	for _, path := range internal.Find(*gpxDirPtr, ".gpx") {
		fmt.Println("Reading", path)
		data, err := internal.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		points, err := internal.Parse(data)
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range points {
			cities = append(cities, p.ToCity())
		}
	}
	sort.SliceStable(cities, func(i, j int) bool {
		return cities[i].Name < cities[j].Name
	})
	internal.Write(cities, *outputPtr)

}
