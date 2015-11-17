package docs

type Page struct {
	Title string
	Body  string
}

var (
	Home = &Page{
		Title: "CitySourced Test API",
		Body: `
			<p>This is the documentation for the City Sourced test API System.
			Enter URL &#8220;docs&#8221; for more information.</p>
		`,
	}

	Detail = &Page{
		Title: "CitySourced API Simulator",
		Body: `
		<h2 id="description">Description</h2>

		<p>This program simulates the CitySourced back-end API. It follows the CitySourced API as closely as possible (there are some differences!). </p>

		<h3 id="citysourceddocs">CitySourced Docs</h3>

		<p>The CitySourced docs I used were all I could find in their <a href="https://citysourced.zendesk.com/home">Knowledge Base</a>. </p>

		<h2 id="operation">Operation</h2>

		<p>To run, simply copy the program file and the two JSON files (config.json and data.json) to a suitable directory on your computer, and then run it from the command line in the standard *nix way: &#8220;./simulator&#8221;.<br/>
		NOTE: The system runs as &#8220;localhost&#8221;. The port number can be specified (see [Command Line Options](#Command Line Options) below).</p>

		<h3 id="notes">Notes</h3>

		<ul>
		<li>The XML elements may not be returned in the same order as in the CitySourced API&#8230; this is typical XML, and I can&#8217;t find an XSD anywhere to tell me that they expect the XML in a particular order.</li>
		</ul>

		<h3 id="anamecommandlineoptionsacommandlineoptions"><a name="Command Line Options"></a>Command Line Options</h3>

		<table>
		<colgroup>
		<col style="text-align:left;"/>
		<col style="text-align:left;"/>
		</colgroup>

		<thead>
		<tr>
			<th style="text-align:left;">Option</th>
			<th style="text-align:left;">Description</th>
		</tr>
		</thead>

		<tbody>
		<tr>
			<td style="text-align:left;">-p {portnum}</td>
			<td style="text-align:left;">Run on the specified port number. The default port number is 5000.</td>
		</tr>
		</tbody>
		</table>

		<h3 id="configfile">Config File</h3>

		<p>Filename: &#8220;config.json&#8221;.
		The &#8220;instrumentation&#8221; section controls (somewhat) the output of debug and verbose debug messages.
		The &#8220;api&#8221; section contains the mock CitySourced API &#8220;key&#8221;. It is currently set to &#8220;a01234567890z&#8221;. This key is present as the first XML element in every API request.</p>

		<h3 id="datafile">Data File</h3>

		<p>Filename: &#8220;data.json&#8221;.
		The file contains two sections: &#8220;reports&#8221; and &#8220;comments&#8221;. These should be self explanatory based on the CitySourced API docs.</p>

		<h2 id="documentation">Documentation</h2>

		<p>The Simulator supports the following CitySourced API calls.</p>

		<hr />

		<h3 id="create">Create</h3>

		<h4 id="anamecreatethreeoneoneacreatereport"><a name="CreateThreeOneOne"></a>Create Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30607923-API-Method-CreateThreeOneOne-">CitySourced KB: CreateThreeOneOne()</a> - <em>implemented</em> </p>

		<h4 id="createcomment">Create Comment</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30542667-API-Method-CreateThreeOneOneComment-">CitySourced KB: CreateThreeOneOneComent()</a> - <em>implemented</em><br/>
		This creates a comment for a specified Report.</p>

		<h4 id="saveanimagevideo">Save an Image/Video</h4>

		<p><em>IN PROGRESS</em><br/>
		<a href="https://citysourced.zendesk.com/entries/31058586-API-Method-CreateThreeOneOneMedia-">CitySourced KB: CreateThreeOneOneMedia()</a> - <em>IN PROGRESS</em><br/>
		This saves an image or video attached to the specified report. When creating a new report, you would first call CreateThreeOneOne() to save the text of the report, and receive a ReportID. Then you would call this method with the ReportID to &#8220;attach&#8221; the image to the report.</p>

		<h4 id="voteforareport">Vote for a Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30608063-API-Method-CreateThreeOneOneVote-">CitySourced KB: CreateThreeOneOneVote()</a> - <em>implemented</em><br/>
		This creates a single upvote for the specified report.</p>

		<hr />

		<h3 id="update">Update</h3>

		<h4 id="updateareport">Update a Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30569128-API-Method-UpdateThreeOneOne-">CitySourced KB: UpdateThreeOneOne()</a> - <em>TO DO</em><br/>
		Updates the Ticket information attached to the report. I assume this comes from the agency servicing the report. Two fields are updated:</p>

		<ul>
		<li>TicketId</li>
		<li>TicketSla</li>
		</ul>

		<hr />

		<h3 id="search">Search</h3>

		<h4 id="getreport">Get Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30608133-API-Method-GetReport-">CitySourced KB: GetReport()</a> - <em>implemented</em><br/>
		Get a specific report using the Report ID.</p>

		<h4 id="searchesreturningmultiplereports">Searches Returning Multiple Reports</h4>

		<p>The following requests can return zero or more reports. The number of reports can be limited by:<br/>
		* MaxResults
		* DateRangeStart
		* DateRangeEnd</p>

		<h5 id="getreportsbydeviceid">Get Reports by Device ID</h5>

		<p><a href="https://citysourced.zendesk.com/entries/30542987-API-Method-GetReportsByDeviceId-">GetReportsByDeviceId</a> - <em>implemented</em><br/>
		Get reports created by the specified Device ID.</p>

		<h5 id="anamegetreportsbylatlngagetreportsbylatlng"><a name="GetReportsByLatLng"></a>Get Reports by Lat/Lng</h5>

		<p><a href="https://citysourced.zendesk.com/entries/30608083-API-Method-GetAddressByLatLng-">GetReportsByLatLng()</a> - <em>implemented</em><br/>
		Get all reports that within a specified distance (<em>in meters</em>) from the specified coordinate. The routine uses the <a href="https://en.wikipedia.org/wiki/Haversine_formula">Haversine</a> distance method using <a href="https://gist.github.com/cdipaolo/d3f8db3848278b49db68">this</a> Go code.<br/>
		To limit by MaxResults, the result set is first filtered by distance from the search point ascending, and the limited. In other words, the reports closest to the search point will be returned.<br/>
		<strong>Limitations</strong>
		* The GeoCoderType parameter is ignored.<br/>
		* Coordinates must be in the typical Decimal Degrees WGS84 Compliant format, as you would get if you right click on Google Maps and select &#8220;What&#8217;s here?&#8221;. For example, San Jose City Hall is: (37.337930, &#8211;121.885891).</p>

		<h5 id="getreportsbyzipcode">Get Reports by ZipCode</h5>

		<p><a href="https://citysourced.zendesk.com/entries/30569088-API-Method-GetReportsByZipCode-">GetReportsByZipCode()</a> - <em>in progress</em><br/>
		Gets reports matching the specified Zip Code. </p>

		<h5 id="getreportsbyaddress">Get Reports By Address</h5>

		<p><a href="https://citysourced.zendesk.com/entries/30568898-API-Method-GetReportsByAddress-">GetReportsByAddress()</a> - <em>implemented</em><br/>
		This function first reaches out to the Google Geocode API and converts the specified address to Lat/Lng. Then <a href="#GetReportsByLatLng">GetReportsByLatLng()</a> is called to return the reports that re within the specified search radius from the address.</p>

		<h5 id="getreportsbyauthorid">Get Reports By Author ID</h5>

		<p><a href="https://citysourced.zendesk.com/entries/30542927-API-Method-GetReportsByAuthorId-">GetReportsByAuthorId</a> - <em>NOT IMPLEMENTED</em><br/>
		There does not appear to be an AuthorID in the <a href="#CreateThreeOneOne">CreateThreeOneOne()</a> call, so it is unclear how to implement this request.</p>

		<hr />

		<h3 id="notimpemented">Not Impemented</h3>

		<ul>
		<li>CreateMobileDevice()</li>
		<li>GetReportsByAuthorId()</li>
		</ul>
				`,
	}
)
