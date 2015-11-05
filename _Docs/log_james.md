# CitySourced Test API

[2015.11.05 - Thu]

* Saved to GIT.
* Created "geocode" package.
* Downloaded a geocode API using Google From: https://github.com/nf/geocode/blob/master/geocode.go, by Jonathan Ingram.
* In package "geocode", added mygeocode.go, containing shortcuts:
	* GetLatLng()
	* GetAddress()
* Expanded on Jonathan's test cases.
* Converted Haversin function back to meters.  Everyone uses meters...
* Test OK.
* Saved to GIT.

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