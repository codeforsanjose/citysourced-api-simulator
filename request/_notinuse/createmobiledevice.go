package requests

import (
	"encoding/xml"
    "fmt"
)


type CreateMobileDevice_Type struct {
	XMLName   xml.Name `xml:"CsRequest"`
	ApiAuthKey	string	 `xml:"ApiAuthKey"`
	ApiRequestType	string	 `xml:"ApiRequestType"`
	ApiRequestVersion	string	 `xml:"ApiRequestVersion"`
	DeviceId	string	 `xml:"DeviceId"`
	DeviceType	string	 `xml:"DeviceType"`
	DeviceModel	string	 `xml:"DeviceModel"`
	DeviceNumber	string	 `xml:"DeviceNumber"`
	DeviceToken	string	 `xml:"DeviceToken"`
}

