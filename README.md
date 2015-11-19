## Description
This program simulates the CitySourced back-end API.  It follows the CitySourced API as closely as possible (there are some differences!).  
The Simulator reads the initial data set from the "data.json" file.  Once running, data additions and modifications (e.g. new reports created by calls to "CreateThreeOneOne") are appended to the initial set.

### CitySourced Docs
This is based on the limited CitySource docs available in their [Knowledge Base](https://citysourced.zendesk.com/home).  

## Operation
To run, simply copy and unzip the appropriate program file and the two JSON files (config.json and data.json) to a suitable directory.  You may need to set permissions on Linux/MacOS ("chmod 777 CSSimulator"). If you're running on MacOS or Linux, start the program in the standard *nix way: "./CSSimulator {command line options}".<br>

### Endpoints

|Endpoint|Description|
|--------|-----------|
|`http://localhost:<port>`|"intro" description|
|`http://localhost:<port>/docs`|displays this documentation|
|`http://localhost:<port>/api`|API endpoint.  Messages formated as per the City Sourced KB must be **POST**ed to this URL|

### Notes

* The system runs as "localhost".  The port number can be specified on the command line (see [Command Line Options](#CommandLineOptions) below), or in the "config.json" file.  The command line will override the config.json setting.
* The system reads the initial data from the "data.json" file.  Edit this file to create any test data you wish to have initially available at system start.
	* Tip: you can get Lat/Lng easily by: finding the address in Google Maps, right clicking on the map pin (or any location), and selecting "What's here?".
* New reports (created by calling request "CreateThreeOneOne") are kept in memory *until the Simulator is stopped*.  They are **not** saved back to the data.json file.
* Any newly created reports are fully searchable.
* The XML elements are not necessarily returned in the same order as in the CitySourced API... this is typical XML, and there is not an XSD in the CitySourced docs specifying the XML elements be in a particular order.  Your unmarshaler should be ok with this..
* **This currently only supports XML output.**  JSON output can be added if needed.

### Package Contents

* There are zip files for the various runtime environments:
	* CSSimulator_mac.zip - MacOS.
	* CSSimulator_win.zip - Window.
	* CSSimulator_linux.zip - Linux.
* Each zip file contains the following files:
	* CSSimulator program file
	* config.json - configuration settings.
	* data.json - initial data.
	* test/test.paw - [Paw](https://luckymarmot.com/paw) test calls.
	* test/*.sh - [cURL](http://curl.haxx.se/docs/manpage.html) scripts for test calls.

### <a name="CommandLineOptions"></a>Command Line Options
|Option|Description|
|------|-----------|
|-port={portnum}|Run on the specified port number.  The default port number is 5050.  The port can also be set in the config.json file.|
|-debug|Activates debug printing to the console.|

### Config File
Filename: "config.json".
The "api" object contains the mock CitySourced API "key".  It is currently set to "a01234567890z".  This key is present as the first XML element in every API request.
The "server" object contains the port setting.  This can be overridden at the command line using the "-port=xxx" syntax.

### Data File
Filename: "data.json".
The file contains two primary objects: "reports" and "comments".  These should be self explanatory based on the CitySourced API docs.

## Documentation
The Simulator supports the following CitySourced API calls.

----
### Create
#### <a name="CreateThreeOneOne"></a>Create Report
[CitySourced KB: CreateThreeOneOne()](https://citysourced.zendesk.com/entries/30607923-API-Method-CreateThreeOneOne-) - *implemented*  


#### Create Comment
[CitySourced KB: CreateThreeOneOneComent()](https://citysourced.zendesk.com/entries/30542667-API-Method-CreateThreeOneOneComment-)   - *implemented*  
Creates a comment for a specified Report.

#### <a name="CreateThreeOneOneMedia"></a>Save an Image/Video
[CitySourced KB: CreateThreeOneOneMedia()](https://citysourced.zendesk.com/entries/31058586-API-Method-CreateThreeOneOneMedia-)   - *NOT IMPLEMENTED*  
The CitySourced documentation is unclear - is this a multi-part request (XML and Binary image in the same request), or does this request return a URL to which the image is subsequently uploaded?  

#### Vote for a Report
[CitySourced KB: CreateThreeOneOneVote()](https://citysourced.zendesk.com/entries/30608063-API-Method-CreateThreeOneOneVote-)  - *implemented*  
This creates a single upvote for the specified report.

----
### Update
#### Update a Report
[CitySourced KB: UpdateThreeOneOne()](https://citysourced.zendesk.com/entries/30569128-API-Method-UpdateThreeOneOne-) - *implemented*    
Updates the Ticket information attached to the report.  We assume this comes from the agency servicing the report, and would not be called by the mobile device.  
*One field is updated: TicketSla*   
This call is confusing, as the example references only the "TicketId" as any sort of identifier.  We will assume that this is the ReportID.  

----
### Search
#### Get Report
[CitySourced KB: GetReport()](https://citysourced.zendesk.com/entries/30608133-API-Method-GetReport-) - *implemented*  
Get a specific report using the Report ID.

#### Searches Returning Multiple Reports
The following requests can return zero or more reports.  The number of reports can be limited by:

* MaxResults
* DateRangeStart
* DateRangeEnd

#### Get Reports By Address
[CitySourced KB: GetReportsByAddress()](https://citysourced.zendesk.com/entries/30568898-API-Method-GetReportsByAddress-) - *implemented*  
This function first reaches out to the Google Geocode API and converts the specified address to Lat/Lng.  Then [GetReportsByLatLng()](#GetReportsByLatLng) is called to return the reports that are within the specified search radius.

#### <a name="GetReportsByAuthorId"></a>Get Reports By Author ID
[CitySourced KB: GetReportsByAuthorId](https://citysourced.zendesk.com/entries/30542927-API-Method-GetReportsByAuthorId-) - *NOT IMPLEMENTED*  
There does not appear to be an AuthorID in the [CreateThreeOneOne()](#CreateThreeOneOne) call, so it is unclear how to implement this request.

#### Get Reports by Device ID
[CitySourced KB: CitySourced KB: GetReportsByDeviceId](https://citysourced.zendesk.com/entries/30542987-API-Method-GetReportsByDeviceId-) - *implemented*  
Get reports created by the specified Device ID.  **This searches only on the value of `<DeviceId>`** - the `<DeviceType>` and `<DeviceModel>` fields are ignored.

#### <a name="GetReportsByLatLng"></a>Get Reports by Lat/Lng
[CitySourced KB: GetReportsByLatLng()](https://citysourced.zendesk.com/entries/30608083-API-Method-GetAddressByLatLng-) - *implemented*  
Get all reports that within a specified distance (*in meters*) from the specified coordinate.  The routine uses the [Haversine](https://en.wikipedia.org/wiki/Haversine_formula) distance method using [this](https://gist.github.com/cdipaolo/d3f8db3848278b49db68) Go code.  
To limit by MaxResults, the result set is first filtered by distance from the search point ascending, and the limited.  In other words, the reports closest to the search point will be returned.  
<br>
**Limitations**

* The GeoCoderType parameter is ignored.  
* Coordinates must be in the typical Decimal Degrees WGS84 Compliant format, as you would get if you right click on Google Maps and select "What's here?".  For example, San Jose City Hall is: (37.337930, -121.885891).

#### Get Reports by ZipCode
[CitySourced KB: GetReportsByZipCode()](https://citysourced.zendesk.com/entries/30569088-API-Method-GetReportsByZipCode-) - *implemented*  
Gets reports matching the specified Zip Code.  

---
### Not Impemented

* CreateMobileDevice() - purpose and usage of this call is unclear from the City Sourced docs.  
* [GetReportsByAuthorId()](#GetReportsByAuthorId)
* [CreateThreeOneOneMedia()](#CreateThreeOneOneMedia)
