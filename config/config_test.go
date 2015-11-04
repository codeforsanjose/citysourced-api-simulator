package config_test

import (
	"CitySourcedAPI/config"
	"fmt"
	"testing"
)

func TestReadConfigInvalidPath(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestReadConfigInvalidPath <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Load Config
	if err := config.Init("../configxxx.json"); err == nil {
		t.Errorf("Attempting to load an invalid file should have caused an error", err)
	}
	fmt.Println("   (Should have just received a CRIT error 'Failed to open...')")
}

func TestReadConfigInvalidJSON(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestReadConfigInvalidJSON <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Load Config
	if err := config.Init("tests/config_faulty.json"); err == nil {
		t.Errorf("Attempting to load a faulty file should have caused an error", err)
	}
	fmt.Println("    (Should have just received a CRIT error 'Invalid JSON...')")
}

func TestReadConfig(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestReadConfig <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Load Config
	if err := config.Init("../config.json"); err != nil {
		t.Errorf("Error %v occurred when loading the config.", err)
	}
	fmt.Printf("%v", config.C.Display())
	if config.C.Loaded != true {
		t.Errorf("System configuration is not marked as loaded.")
	}

	//  Test Auth()
	ac := "1234567890"
	if a := config.Auth(ac); !a {
		t.Errorf("Auth() failed.")
	}

	ac = "1111"
	if a := config.Auth(ac); a {
		t.Errorf("Auth() passed erroneously for: %q", ac)
	}

}

func TestRepeatReadConfig(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestRepeatReadConfig <<<<<<<<<<<<<<<<<<<<<<<<<<")
	if err := config.Init("../config.json"); err == nil {
		t.Errorf("Duplicate calls to config.Init() should have resulted in a warning")
	}
	fmt.Println("   (Should have just received a WARN error 'Duplicate calls...')")

}
