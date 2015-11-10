package response

import (
	"CitySourcedAPI/config"
	"CitySourcedAPI/data"
	"fmt"
	"testing"
	// "github.com/davecgh/go-spew/spew"
)

func init() {
	if err := config.Init("../config.json"); err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
	}

	if err := data.Init("../data.json"); err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
	}

}

func TestNewResponseReports(t *testing.T) {
	fmt.Println("\n>>>>>>>>>>>>>>>>>>> TestNewResponseReports <<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("\n!!!CREATE SOME TEST CASES!!!!\n\n")
}
