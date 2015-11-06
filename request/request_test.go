package request

import (
	"CitySourcedAPI/config"
	"CitySourcedAPI/data"
	"fmt"
	"testing"
	// "github.com/davecgh/go-spew/spew"
)

var (
	inCreateThreeOneOne = `
		<?xml version="1.0" encoding="utf-8" ?>
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
	inGetReportsByAddress = `
	<?xml version="1.0" encoding="utf-8" ?>
	<CsRequest>
		<ApiAuthKey>{INSERT YOUR AUTH KEY HERE!}</ApiAuthKey>
		<ApiRequestType>GetReportsByAddress</ApiRequestType>
		<ApiRequestVersion>1</ApiRequestVersion>
		<Address>200 E. Santa Clara St., San Jose, CA</Address>
		<Radius>500</Radius>
		<MaxResults></MaxResults>
		<IncludeDetails>False</IncludeDetails>
		<DateRangeStart>2015-05-20T13:45:30</DateRangeStart>
		<DateRangeEnd>2015-05-20T13:45:30</DateRangeEnd>
		<CurrentStatus></CurrentStatus>
	</CsRequest>
	`
	inGetReportsByAddress2 = `
	<?xml version="1.0" encoding="utf-8" ?>
	<CsRequest>
		<ApiAuthKey>{INSERT YOUR AUTH KEY HERE!}</ApiAuthKey>
		<ApiRequestType>GetReportsByAddress</ApiRequestType>
		<ApiRequestVersion>1</ApiRequestVersion>
		<Address>200 E. Santa Clara St., San Jose, CA</Address>
		<Radius>10000</Radius>
		<MaxResults></MaxResults>
		<IncludeDetails></IncludeDetails>
		<DateRangeStart></DateRangeStart>
		<DateRangeEnd></DateRangeEnd>
		<CurrentStatus></CurrentStatus>
	</CsRequest>
	`
)

func init() {
	if err := config.Init("../config.json"); err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
	}

	if err := data.Init("../data.json"); err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
	}

}

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
	fmt.Println("\n>>>>>>>>>>>>>>>>>>> TestCreateThreeOneOne <<<<<<<<<<<<<<<<<<<<<<<<<<")
	input := inCreateThreeOneOne
	data, _ := Process(input)
	fmt.Printf("[ProcessRequest] returned: %q\n", data)
}

func TestGetReportsByAddress(t *testing.T) {
	fmt.Println("\n>>>>>>>>>>>>>>>>>>> TestGetReportsByAddress <<<<<<<<<<<<<<<<<<<<<<<<<<")
	input := inGetReportsByAddress
	data, _ := Process(input)
	fmt.Printf("[ProcessRequest] returned: %q\n", data)

	input = inGetReportsByAddress2
	data, _ = Process(input)
	fmt.Printf("[ProcessRequest] returned: %q\n", data)
}
