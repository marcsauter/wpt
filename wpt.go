package wpt

import (
	"encoding/xml"
	"io/ioutil"
	"math"
)

type Waypoint struct {
	XMLName   xml.Name `xml:"wpt"`
	Latitude  float64  `xml:"lat,attr"`
	Longitude float64  `xml:"lon,attr"`
	Elevation int      `xml:"ele"`
	Name      string   `xml:"name"`
}

func NewWaypoints(filename string) (*GPX, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	g := GPX{}
	err = xml.Unmarshal(data, &g)
	if err != nil {
		return nil, err
	}
	return &g, err
}

func (w *Waypoint) Distance(lat, lon float64) float64 {
	lat1 := lat * math.Pi / 180
	lon1 := lon * math.Pi / 180
	lat2 := w.Latitude * math.Pi / 180
	lon2 := w.Longitude * math.Pi / 180
	return 6378.388 * math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1)) * 1000
}

type GPX struct {
	XMLName   xml.Name   `xml:"gpx"`
	Version   string     `xml:"version,attr"`
	Creator   string     `xml:"creator,attr"`
	Waypoints []Waypoint `xml:"wpt"`
}

func (g *GPX) Find(lat, lon float64) (string, int) {
	nearest := g.Waypoints[0]
	distance := g.Waypoints[0].Distance(lat, lon)
	for _, wp := range g.Waypoints[1:] {
		d := wp.Distance(lat, lon)
		if d < distance {
			nearest = wp
			distance = d
		}
	}
	return nearest.Name, int(distance)
}
