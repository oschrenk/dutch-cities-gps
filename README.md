# README

Dutch cities and their GPS locations.

This repo contains the code, for a CSV file of the cities, please look into the releases.

## Background

I retrieved the list of the cities and their GPS locations via [](http://gps.herbschleb.net) in a [zip file](http://gps.herbschleb.net/compleet.zip) containing a xml file for each province in the Netherlands.

Unfortunately the original `Friesland.gpx` uses an unknown/broken file encoding that go breaks on, and [ecna](https://linux.die.net/man/1/enca) does not recognize. The Uppercase Circumflex A `Â` is encoded as `83` (hex) or `131` in binary. It's not Windows-1252, not ISO 8859-1, or ISO 8859-9. In the end I had to manually replace the code in the xml files.

The code in this repository parses the xml files (after pre-processing the encoding issue), normalizes the city and province names and creates a single CSV file.

## Correctness

The data has NOT being checked for correctness. If you find an issue or a better source for the data, please let me know.
