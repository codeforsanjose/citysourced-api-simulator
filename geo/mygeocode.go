package geo

import (
	"CitySourcedAPI/logs"
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var (
	log = logs.Log
)

func GetLatLng(addr string) (float64, float64, error) {
	req := &Request{
		Address:  addr,
		Provider: GOOGLE,
	}
	resp, err := req.Lookup(nil)
	// log.Debug(">>>Found: %s\n", resp.Found)
	// log.Debug(">>>Response:\n%s\n", spew.Sdump(resp))
	if err != nil || resp.Status != "OK" {
		return 0.0, 0.0, errors.New(fmt.Sprintf("Unable to determine GeoLoc for %q", addr))
	}
	p := resp.GoogleResponse.Results[0].Geometry.Location
	return p.Lat, p.Lng, nil
}

func GetAddress(lat, lng float64) (string, error) {
	loc := Point{lat, lng}
	req := &Request{
		Location: &loc,
		Provider: GOOGLE,
	}
	resp, err := req.Lookup(nil)
	log.Debug(">>>Found: %s\n", resp.Found)
	log.Debug(">>>Response:\n%s\n", spew.Sdump(resp))
	if err != nil || resp.Status != "OK" {
		return "", errors.New(fmt.Sprintf("Unable to determine GeoLoc for %v | %v", lat, lng))
	}
	return resp.Found, nil
}
