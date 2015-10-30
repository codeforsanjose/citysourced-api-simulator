package request

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func showData(data CreateThreeOneOne_Type) {
	fmt.Printf("ApiAuthKey: %s\n", data.ApiAuthKey)
	fmt.Printf("ApiRequestType: %s\n", data.ApiRequestType)
	fmt.Printf("ApiRequestVersion: %s\n", data.ApiRequestVersion)
	fmt.Printf("DateCreated: %s\n", data.DateCreated)
	fmt.Printf("DeviceType: %s\n", data.DeviceType)
	fmt.Printf("DeviceModel: %s\n", data.DeviceModel)
	fmt.Printf("DeviceId: %s\n", data.DeviceId)
	fmt.Printf("RequestType: %s\n", data.RequestType)
	fmt.Printf("RequestTypeId: %s\n", data.RequestTypeId)
	fmt.Printf("Latitude: %s\n", data.Latitude)
	fmt.Printf("Longitude: %s\n", data.Longitude)
	fmt.Printf("Directionality: %s\n", data.Directionality)
	fmt.Printf("Description: %s\n", data.Description)
	fmt.Printf("AuthorNameFirst: %s\n", data.AuthorNameFirst)
	fmt.Printf("AuthorNameLast: %s\n", data.AuthorNameLast)
	fmt.Printf("AuthorEmail: %s\n", data.AuthorEmail)
	fmt.Printf("AuthorTelephone: %s\n", data.AuthorTelephone)
	fmt.Printf("AuthorIsAnonymous: %s\n", data.AuthorIsAnonymous)
}

func TestCreateThreeOneOne(t *testing.T) {
	input := `
		<CsRequest>
		<ApiAuthKey>a0124852109248523948z</ApiAuthKey>
		<ApiRequestType>CreateThreeOneOne</ApiRequestType>
		<ApiRequestVersion>1</ApiRequestVersion>
		<DateCreated>2015-05-20T13:45:30</DateCreated>
		<DeviceType>IPHONE</DeviceType>
		<DeviceModel>6</DeviceModel>
		<DeviceId>987654321</DeviceId>
		<RequestType>Graffiti Removal</RequestType>
		<RequestTypeId>10</RequestTypeId>
		<Latitude>34.0632809</Latitude>
		<Longitude>-118.445211</Longitude>
		<Directionality>25 N NW</Directionality>
		<Description>This is my description.</Description>
		<AuthorNameFirst>James</AuthorNameFirst>
		<AuthorNameLast>Haskell</AuthorNameLast>
		<AuthorEmail>xyz@xxx.com</AuthorEmail>
		<AuthorTelephone>555-555-5555</AuthorTelephone>
		<AuthorIsAnonymous>False</AuthorIsAnonymous>
		<KeyValuePairs>
			<KeyValuePair Key="Key1">Value1</KeyValuePair>
			<KeyValuePair Key="Key2">Value2</KeyValuePair>
		</KeyValuePairs>
		</CsRequest>
	`
	data, err := NewCreateThreeOneOne(input)
	if err != nil {
		t.Errorf("Error unmashalling test package: %s", err)
	}

	fmt.Println(spew.Sdump(data))
}
