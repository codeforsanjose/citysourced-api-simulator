# CitySourced Test API

[2015.11.02 - Mon]

* Added to Report_Type based on City Sourced API.
* Removed all the individual fields from requests.Response_Type, and changed request.response.go to use the Report_Type struct.
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