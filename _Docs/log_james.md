# CitySourced Test API

[2015.11.17 - Tue]

* Saved to GIT.
* Cleaned up debug prints and log formatting.
* Added response to CreateThreeOneOne()
* Renamed files in request:
	* request/createthreeoneone.go -> request/create.go
	* request/getreports.go -> get.go
* Started adding request/update.go (for UpdateThreeOneOne).
* Test OK.
* Saved to GIT.
* In data/data.go:
	* Since the data MUST be a singleton, I moved most functions that were methods on the Report struct out and made them exported methods.  This removes the necessity to use "data.D" when calling them... now they are simply "data.FindLL()", for example, rather than "data.D.FindLL()".
* Test OK.
* Saved to GIT.
* Update working.
* Saved to GIT.

[2015.11.16 - Mon]

* Wrote up documentation and added to docs/docs.go.
* Removed copier code.  It is not being used.
* Converted all fmt.PrintXX() debug statements to "log.Debug()".
* Additional testing.  Everything looks good so far.

[2015.11.14 - Sat]

* GetReportsByZipCode() working.
* Test OK.
* Saved to GIT.
* GetReport() working.
	* Does NOT use the "IncludeDetails", "IncludeComments", or "IncludeVotes" fields.  It will always return the above.
* Test OK.
* Saved to GIT.
* 

[2015.11.13 - Fri]

* Comments are loading properly now.
	* Added "comments.go" to "data" package.
	* Test "NewComment()" = OK!
	* Test checks on comments are OK.
* TestOK (data, request)
* Saved to GIT.

[2015.11.12 - Thu]

* Cleaned up the request processing by adding reflection to dynamically create the proper structure for the request, and then use an interface to process the request.
* data/report.go:
	* Needed to have all BaseReport fields exported, so they can be accessed when used in a Request.  So changed the names of the lowercased fields to uppercase, post-pended with a "V":
		* latitude -> LatitudeV
		* longitude -> LongitudeV
		* authorIsAnonymous -> AuthorIsAnonymous
	* Deleted the Lat() and Lng() methods from BaseReport
* In "request" package:
	* Renamed "Request_Type" -> "Request".
	* getreports.go:
		* Deleted GetReportsByAddress()
		* Deleted GetReportsByLatLng()
		* Added "Processor" interface to the GetReportsByAddress struct.
		* Added "Processor" interface to the GetReportsByLatLng struct.
		* Renamed the "GetReportsByAddressType" struct to "GetReportsByAddress".
		* Renamed the "GetReportsByLatLngType" struct to "GetReportsByLatLng".
		* In both request types:
			* Added new style Validate()
			* Moved append code to Run()
	* createthreeoneone.go
		* Deleted CreateThreeOneOne().
		* Added new style Validate()
		* Moved append code to Run()
	* request.go
		* Removed "processors" variable.
		* Added "typeRegistry" variable.
		* Rewrote Process() to use the reflection to dynamically create the appropriate request struct, then Validate() and Run().
		* Standardized all the error and return handling in Process().
* Test OK.
* Saved to GIT.

* Implemented "FindDeviceId" request.
	* Created struct and methods in request/getreports.go.
	* Added test case to request/request_test.go.  Pass OK!
* Added filtering of report output fields.  Previously we were outputing a direct copy of data/BaseReport.  This was outputing too many fields.  So created response/reports.go:
	* response.Reports{} struct is a subset copy of "BaseReport" of the fields that should be output.
	* Created two methods:
		* ConvertReport() maps a data.Report{} to the new Report{}.
		* ConvertReports() converts a slice of pointers to data.Report{} into a slice of pointers to response.Report{}.
* Test OK (data, request)
* Saved to GIT.
* Added Votes to all structs and test data.
	* Added validation for Votes/VotesV.
* Rearranged some of the fields... moved the image URLs down to the bottom of all the structs.

[2015.11.11 - Wed]

* Figured out how to use reflection to create an instance of struct.

[2015.11.10 - Tue]

* Installed new Atom editor... looks like it has a lot of capabilities that VS Code doesn't...  although it does a bit slow.
* Renamed BaseReport_Type to BaseReport.
* Test OK.
* Saved to GIT.
* Refactored the following in the "data" package:
	* BaseReport_Type -> BaseReport
	* Report_Type -> Report
	* Data_type -> Reports
	* Reports.FindId() -> FindID()
	* Reports.FindDeviceId() -> FindDeviceID()
	* Reports.lastId -> lastID
	* Report.Id -> ID
	* BaseReport.RequestTypeId -> RequestTypeID
	* BaseReport.DeviceId -> DeviceID
	* Reports.LastId() -> LastID()
	* Reports.indId -> indID
* Test OK.
* Saved to GIT.
* In "request":
	* Renamed GetReportsByAddress_Type -> GetReportsByAddressType.
* In "data":
	* Added new type ReportListD.  This is a slice of Reports (ReportList type), and a slice of distances.  This will be used to sort and limit the returned values from any calls that search by distance from a center point, like FindAddress() and FindLL().  This type has sorting and limiting capability.
	* In data_test.go, added an additional address, and limited the return report list to 2.  Works OK!
* Test OK (data & request)
* Saved to GIT.
* Added a new file to the "request" package:
	* validate.go handles all validation for the various request structs.  
	* It validates and converts all necessary fields, and creates a consolidated error message.
* Test OK
* Saved to GIT.

[2015.11.09 - Mon]

* Merged all usages of the "generic" Report struct into data.report.go.  
	* Created two structs: Report_Type (existing), and BaseReport_Type.
	* BaseReport_Type will be used in both the "data" and "request" packages.
	* Added the "copier" package into the project, but am not using it.  It didn't like the CustomTime type in the structs to be copied.
	* Tests on data and request all OK.
	* Saved to GIT.
	* Modified "data" package to return "<Report>" tags within the response.  It was returning "<Request>".
	* Saved to GIT.
	* Created ReportList type - slice: []*Report_Type.
	* Renamed BaseReport.Distance() to CalcDistance() so I could add "Distance" to BaseReport.  Changed my mind... this won't work.
	* Deleted the "_notinuse" folders in geo and request

[2015.11.06 - Fri]

* Moved response.go out of the "request" package, and created it's own package "response".
* Test: request_test.go passed OK.
* Saved to GIT.

[2015.11.05 - Thu]

* Saved to GIT.
* Created "geocode" package.
* Downloaded a geocode API using Google From: https://github.com/nf/geocode/blob/master/geocode.go, by Jonathan Ingram.
* In package "geocode", added mygeocode.go, containing shortcuts:
	* GetLatLongitudeV
	* GetAddress()
* Expanded on Jonathan's test cases.
* Converted Haversin function back to meters.  Everyone uses meters...
* Test OK.
* Saved to GIT.
* Improved CustomTime - will now set a "zero" instance of time.Time if the parse string is empty.
* Added FindAddress() to data/data.go.  Finds all reports within the specified search center point and radius.
	* Added test cases for FindAddress() to data_test.go
	* Test OK.
* Renamed "geocode" package to "geo".
* In request/getreports.go:
	* The GetReportsByAddress struct was failing when parsing the XML, as some of the XML fields are blank.  These were fields previously declared as int64, float64, and boolean - basically anything other than a string.  Strings appear to be the only datatype that will successfully parse if the XML element is blank/empty.  (Columns: Radius, MaxResults, IncludeDetails).  SO...
	* Converted ALL Exportable fields in GetReportsByAddress to string types, and added non-export versions with the original, required datatype.
	* Added code to validate the string data, and parse it into the non-export fields.
* Other miscellaneous fixes from moving code around.
* In request/request_test.go, added several test cases and XML text for testing GetReportsByAddress().

[2015.11.04 - Wed]

* Separated "data" and "config".
* Created separate test files for each.
* Polished the test files.  I think the coverage is pretty good now.
* data test ok.
* config test ok.
* Saved to GIT.
* Moved the CustomTime code from the "common" package to the "data" package.
	* Changed all references to CustomTime from "common" to "data".
* In data:
	* Modified the data.haversin.go routine to return miles rather than meters.
	* Modified the distance test in data_test.go to check for the mileage amount.
* Test ok.

[2015.11.03 - Tue]

* Created "main" file: "cs_api.go".
* Added framework of HTTP handling, with paths:
	* / = home, displays a brief into.
	* /docs/ = returns more detailed info about usage.
	* /api/ = path for using the API.  This only responds to POST's.
* /api POST method reading full post content, and sending to request.ProcessRequest()
* Test OK.
* Saved to GIT.
* In request/request.go:
	* Moved the code creating a new request (from the top of Process()) into the new function newRequest().  
	* Test OK.
	* Saved to GIT.

[2015.11.02 - Mon]

* Added to Report_Type based on City Sourced API.
* Removed all the individual fields from requests.Response_Type, and changed request.response.go to use the Report_Type struct.
* Saved to GIT.
* Separated the data from the config file.  So now there is:
	* config.json
	* data.json
* Test OK.
* Saved to GIT.

[2015.10.30 - Fri]

* config.json
	* Added "api" section.
		* Added "authkey" - will validate the input authkey.
	* Added Auth() function - tests an auth code against the authkey above.
* requests
	* Moved the "common" request header fields into a separate structure "Request_Type".
* Test OK.
* Saved to GIT.
* 


[2015.10.29 - Thu]

* Started on the project.  
* Acomplished:
	* config.json
		* Instrumentation section
		* Data section
			* Sample data for several reports.
	* "requests" package
		* CreateThreeOneOne() struct and CreateThreeOneOne() methods.
		* createthreeoneone_test.go
			* Test loading a single request - added reporter's name, but otherwise a copy of the example on CitySourced API doc.
		* Test OK.
		* Saved to GIT.
	* "data" package
		* load.go
			* Reads config.json
			* Parses OK.
			* Researched and built custom time format.  CitySourced is using a non-standard time format in their examples - like RFC3339 but without the time zone section.
		* haversin.go
			* Functions to compute distance between two geo-locations.
		* data_test.go
			* TestReadConfig()
			* TestFindDeviceId()
			* TestDistance() - compute distance using Haversin function.
			* TestRepeatReadConfig() - fails as it should, with warning that config has already been read.
		* Test OK
		* Saved to GIT.