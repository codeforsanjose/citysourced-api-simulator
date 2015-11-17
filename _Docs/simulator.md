## Description
This program simulates the CitySourced back-end API.  It follows the CitySourced API as closely as possible (there are some differences!).  

### CitySourced Docs
The CitySourced docs I used were all I could find in their [Knowledge Base](https://citysourced.zendesk.com/home).  

## Operation
To run, simply copy the program file and the two JSON files (config.json and data.json) to a suitable directory on your computer, and then run it from the command line in the standard *nix way: "./simulator".  
NOTE: The system runs as "localhost".  The port number can be specified (see [Command Line Options](#Command Line Options) below).

### Notes

* The XML elements may not be returned in the same order as in the CitySourced API... this is typical XML, and is not an XSD in the CitySourced docs that would specify the XML elements be in a particular order.

### <a name="Command Line Options"></a>Command Line Options
|Option|Description|
|------|-----------|
|-p {portnum}|Run on the specified port number.  The default port number is 5000.|

### Config File
Filename: "config.json".
The "instrumentation" section controls (somewhat) the output of debug and verbose debug messages.
The "api" section contains the mock CitySourced API "key".  It is currently set to "a01234567890z".  This key is present as the first XML element in every API request.

### Data File
Filename: "data.json".
The file contains two sections: "reports" and "comments".  These should be self explanatory based on the CitySourced API docs.

## Documentation
The Simulator supports the following CitySourced API calls.

----
### Create
#### <a name="CreateThreeOneOne"></a>Create Report
[CitySourced KB: CreateThreeOneOne()](https://citysourced.zendesk.com/entries/30607923-API-Method-CreateThreeOneOne-) - *implemented*  


#### Create Comment
[CitySourced KB: CreateThreeOneOneComent()](https://citysourced.zendesk.com/entries/30542667-API-Method-CreateThreeOneOneComment-)   - *implemented*  
This creates a comment for a specified Report.

#### Save an Image/Video
[CitySourced KB: CreateThreeOneOneMedia()](https://citysourced.zendesk.com/entries/31058586-API-Method-CreateThreeOneOneMedia-)   - *IN PROGRESS*  
This saves an image or video attached to the specified report.  When creating a new report, you would first call CreateThreeOneOne() to save the text of the report, and receive a ReportID.  Then you would call this method with the ReportID to "attach" the image to the report.

#### Vote for a Report
[CitySourced KB: CreateThreeOneOneVote()](https://citysourced.zendesk.com/entries/30608063-API-Method-CreateThreeOneOneVote-)  - *implemented*  
This creates a single upvote for the specified report.


----
### Update
#### Update a Report
[CitySourced KB: UpdateThreeOneOne()](https://citysourced.zendesk.com/entries/30569128-API-Method-UpdateThreeOneOne-) - *TO DO*    
Updates the Ticket information attached to the report.  I assume this comes from the agency servicing the report.  One field is updated:

* TicketSla

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

##### Get Reports By Address
[GetReportsByAddress()](https://citysourced.zendesk.com/entries/30568898-API-Method-GetReportsByAddress-) - *implemented*  
This function first reaches out to the Google Geocode API and converts the specified address to Lat/Lng.  Then [GetReportsByLatLng()](#GetReportsByLatLng) is called to return the reports that re within the specified search radius from the address.

##### Get Reports By Author ID
[GetReportsByAuthorId](https://citysourced.zendesk.com/entries/30542927-API-Method-GetReportsByAuthorId-) - *NOT IMPLEMENTED*  
There does not appear to be an AuthorID in the [CreateThreeOneOne()](#CreateThreeOneOne) call, so it is unclear how to implement this request.

##### Get Reports by Device ID
[GetReportsByDeviceId](https://citysourced.zendesk.com/entries/30542987-API-Method-GetReportsByDeviceId-) - *implemented*  
Get reports created by the specified Device ID.

##### <a name="GetReportsByLatLng"></a>Get Reports by Lat/Lng
[GetReportsByLatLng()](https://citysourced.zendesk.com/entries/30608083-API-Method-GetAddressByLatLng-) - *implemented*  
Get all reports that within a specified distance (*in meters*) from the specified coordinate.  The routine uses the [Haversine](https://en.wikipedia.org/wiki/Haversine_formula) distance method using [this](https://gist.github.com/cdipaolo/d3f8db3848278b49db68) Go code.  
To limit by MaxResults, the result set is first filtered by distance from the search point ascending, and the limited.  In other words, the reports closest to the search point will be returned.  
**Limitations**
* The GeoCoderType parameter is ignored.  
* Coordinates must be in the typical Decimal Degrees WGS84 Compliant format, as you would get if you right click on Google Maps and select "What's here?".  For example, San Jose City Hall is: (37.337930, -121.885891).

##### Get Reports by ZipCode
[GetReportsByZipCode()](https://citysourced.zendesk.com/entries/30569088-API-Method-GetReportsByZipCode-) - *in progress*  
Gets reports matching the specified Zip Code.  

---
### Not Impemented

* CreateMobileDevice()
* GetReportsByAuthorId()
