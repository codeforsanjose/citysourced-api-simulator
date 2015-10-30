package requests

import (
	"fmt"
	"testing"
	"github.com/davecgh/go-spew/spew"
)

func TestCreateMobileDevice(t *testing.T) {
	md := CreateMobileDevice_Type{
		ApiAuthKey: "1234567890",
		ApiRequestType: "CreateMobileDevice",	
		ApiRequestVersion: "1",
		DeviceId: "987654321",
		DeviceType: IPHONE,
		DeviceModel: "5S",
		DeviceNumber: "8183056052",
		
	}