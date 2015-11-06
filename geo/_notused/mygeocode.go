package geocode

import (
	"fmt"
	"github.com/paulmach/go.geo"
) 

func GetBounds(p *Point, radius float64) (*Point, *Point, error) {
	var center = geo.NewPointFromLatLng(p.Lat, p.Lng)
	fmt.Printf("Center: %#v\n", center)
	
	b := geo.NewGeoBoundAroundPoint(center, radius)
	gne := b.NorthEast()
	gsw := b.SouthWest()

	ne := Point{gne.Lat(), gne.Lng()}
	sw := Point{gsw.Lat(), gsw.Lng()}
	return &ne, &sw, nil 
}