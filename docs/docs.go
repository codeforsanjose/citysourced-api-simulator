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

		<p>This program simulates the CitySourced back-end API. It follows the CitySourced API as closely as possible (there are some differences!).<br/>
		The Simulator reads the initial data set from the &#8220;data.json&#8221; file. Once running, data additions and modifications (e.g. new reports created by calls to &#8220;CreateThreeOneOne&#8221;) are appended to the initial set.</p>

		<h3 id="citysourceddocs">CitySourced Docs</h3>

		<p>This is based on the limited CitySource docs available in their <a href="https://citysourced.zendesk.com/home">Knowledge Base</a>. </p>

		<h2 id="operation">Operation</h2>

		<p>To run, simply copy and unzip the appropriate program file and the two JSON files (config.json and data.json) to a suitable directory. You may need to set permissions on Linux/MacOS (&#8220;chmod 777 CSSimulator&#8221;). If you&#8217;re running on MacOS or Linux, start the program in the standard *nix way: &#8220;./CSSimulator {command line options}&#8221;.<br></p>

		<h3 id="endpoints">Endpoints</h3>

		<table>
		<colgroup>
		<col style="text-align:left;"/>
		<col style="text-align:left;"/>
		</colgroup>

		<thead>
		<tr>
			<th style="text-align:left;">Endpoint</th>
			<th style="text-align:left;">Description</th>
		</tr>
		</thead>

		<tbody>
		<tr>
			<td style="text-align:left;"><code>http://localhost:&lt;port&gt;</code></td>
			<td style="text-align:left;">&#8220;intro&#8221; description</td>
		</tr>
		<tr>
			<td style="text-align:left;"><code>http://localhost:&lt;port&gt;/docs</code></td>
			<td style="text-align:left;">displays this documentation</td>
		</tr>
		<tr>
			<td style="text-align:left;"><code>http://localhost:&lt;port&gt;/api</code></td>
			<td style="text-align:left;">API endpoint. Messages formated as per the City Sourced KB must be <strong>POST</strong>ed to this URL</td>
		</tr>
		</tbody>
		</table>

		<h3 id="notes">Notes</h3>

		<ul>
		<li>The system runs as &#8220;localhost&#8221;. The port number can be specified on the command line (see <a href="#CommandLineOptions">Command Line Options</a> below), or in the &#8220;config.json&#8221; file. The command line will override the config.json setting.</li>
		<li>The system reads the initial data from the &#8220;data.json&#8221; file. Edit this file to create any test data you wish to have initially available at system start.

		<ul>
		<li>Tip: you can get Lat/Lng easily by: finding the address in Google Maps, right clicking on the map pin (or any location), and selecting &#8220;What&#8217;s here?&#8221;.</li>
		</ul></li>
		<li>New reports (created by calling request &#8220;CreateThreeOneOne&#8221;) are kept in memory <em>until the Simulator is stopped</em>. They are <strong>not</strong> saved back to the data.json file.</li>
		<li>Any newly created reports are fully searchable.</li>
		<li>The XML elements are not necessarily returned in the same order as in the CitySourced API&#8230; this is typical XML, and there is not an XSD in the CitySourced docs specifying the XML elements be in a particular order. Your unmarshaler should be ok with this..</li>
		<li><strong>This currently only supports XML output.</strong> JSON output can be added if needed.</li>
		</ul>

		<h3 id="packagecontents">Package Contents</h3>

		<ul>
		<li>There are zip files for the various runtime environments:

		<ul>
		<li>CSSimulator_mac.zip - MacOS.</li>
		<li>CSSimulator_win.zip - Window.</li>
		<li>CSSimulator_linux.zip - Linux.</li>
		</ul></li>
		<li>Each zip file contains the following files:

		<ul>
		<li>CSSimulator program file</li>
		<li>config.json - configuration settings.</li>
		<li>data.json - initial data.</li>
		<li>test/test.paw - <a href="https://luckymarmot.com/paw">Paw</a> test calls.</li>
		<li>test/*.sh - <a href="http://curl.haxx.se/docs/manpage.html">cURL</a> scripts for test calls.</li>
		</ul></li>
		</ul>

		<h3 id="anamecommandlineoptionsacommandlineoptions"><a name="CommandLineOptions"></a>Command Line Options</h3>

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
			<td style="text-align:left;">-port={portnum}</td>
			<td style="text-align:left;">Run on the specified port number. The default port number is 5050. The port can also be set in the config.json file.</td>
		</tr>
		<tr>
			<td style="text-align:left;">-debug</td>
			<td style="text-align:left;">Activates debug printing to the console.</td>
		</tr>
		</tbody>
		</table>

		<h3 id="configfile">Config File</h3>

		<p>Filename: &#8220;config.json&#8221;.
		The &#8220;api&#8221; object contains the mock CitySourced API &#8220;key&#8221;. It is currently set to &#8220;a01234567890z&#8221;. This key is present as the first XML element in every API request.
		The &#8220;server&#8221; object contains the port setting. This can be overridden at the command line using the &#8220;-port=xxx&#8221; syntax.</p>

		<h3 id="datafile">Data File</h3>

		<p>Filename: &#8220;data.json&#8221;.
		The file contains two primary objects: &#8220;reports&#8221; and &#8220;comments&#8221;. These should be self explanatory based on the CitySourced API docs.</p>

		<h2 id="documentation">Documentation</h2>

		<p>The Simulator supports the following CitySourced API calls.</p>

		<hr />

		<h3 id="create">Create</h3>

		<h4 id="anamecreatethreeoneoneacreatereport"><a name="CreateThreeOneOne"></a>Create Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30607923-API-Method-CreateThreeOneOne-">CitySourced KB: CreateThreeOneOne()</a> - <em>implemented</em> </p>

		<h4 id="createcomment">Create Comment</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30542667-API-Method-CreateThreeOneOneComment-">CitySourced KB: CreateThreeOneOneComent()</a> - <em>implemented</em><br/>
		Creates a comment for a specified Report.</p>

		<h4 id="anamecreatethreeoneonemediaasaveanimagevideo"><a name="CreateThreeOneOneMedia"></a>Save an Image/Video</h4>

		<p><a href="https://citysourced.zendesk.com/entries/31058586-API-Method-CreateThreeOneOneMedia-">CitySourced KB: CreateThreeOneOneMedia()</a> - <em>NOT IMPLEMENTED</em><br/>
		The CitySourced documentation is unclear - is this a multi-part request (XML and Binary image in the same request), or does this request return a URL to which the image is subsequently uploaded? </p>

		<h4 id="voteforareport">Vote for a Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30608063-API-Method-CreateThreeOneOneVote-">CitySourced KB: CreateThreeOneOneVote()</a> - <em>implemented</em><br/>
		This creates a single upvote for the specified report.</p>

		<hr />

		<h3 id="update">Update</h3>

		<h4 id="updateareport">Update a Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30569128-API-Method-UpdateThreeOneOne-">CitySourced KB: UpdateThreeOneOne()</a> - <em>implemented</em><br/>
		Updates the Ticket information attached to the report. We assume this comes from the agency servicing the report, and would not be called by the mobile device.<br/>
		<em>One field is updated: TicketSla</em><br/>
		This call is confusing, as the example references only the &#8220;TicketId&#8221; as any sort of identifier. We will assume that this is the ReportID. </p>

		<hr />

		<h3 id="search">Search</h3>

		<h4 id="getreport">Get Report</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30608133-API-Method-GetReport-">CitySourced KB: GetReport()</a> - <em>implemented</em><br/>
		Get a specific report using the Report ID.</p>

		<h4 id="searchesreturningmultiplereports">Searches Returning Multiple Reports</h4>

		<p>The following requests can return zero or more reports. The number of reports can be limited by:</p>

		<ul>
		<li>MaxResults</li>
		<li>DateRangeStart</li>
		<li>DateRangeEnd</li>
		</ul>

		<h4 id="getreportsbyaddress">Get Reports By Address</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30568898-API-Method-GetReportsByAddress-">CitySourced KB: GetReportsByAddress()</a> - <em>implemented</em><br/>
		This function first reaches out to the Google Geocode API and converts the specified address to Lat/Lng. Then <a href="#GetReportsByLatLng">GetReportsByLatLng()</a> is called to return the reports that are within the specified search radius.</p>

		<h4 id="anamegetreportsbyauthoridagetreportsbyauthorid"><a name="GetReportsByAuthorId"></a>Get Reports By Author ID</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30542927-API-Method-GetReportsByAuthorId-">CitySourced KB: GetReportsByAuthorId</a> - <em>NOT IMPLEMENTED</em><br/>
		There does not appear to be an AuthorID in the <a href="#CreateThreeOneOne">CreateThreeOneOne()</a> call, so it is unclear how to implement this request.</p>

		<h4 id="getreportsbydeviceid">Get Reports by Device ID</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30542987-API-Method-GetReportsByDeviceId-">CitySourced KB: CitySourced KB: GetReportsByDeviceId</a> - <em>implemented</em><br/>
		Get reports created by the specified Device ID. <strong>This searches only on the value of <code>&lt;DeviceId&gt;</code></strong> - the <code>&lt;DeviceType&gt;</code> and <code>&lt;DeviceModel&gt;</code> fields are ignored.</p>

		<h4 id="anamegetreportsbylatlngagetreportsbylatlng"><a name="GetReportsByLatLng"></a>Get Reports by Lat/Lng</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30608083-API-Method-GetAddressByLatLng-">CitySourced KB: GetReportsByLatLng()</a> - <em>implemented</em><br/>
		Get all reports that within a specified distance (<em>in meters</em>) from the specified coordinate. The routine uses the <a href="https://en.wikipedia.org/wiki/Haversine_formula">Haversine</a> distance method using <a href="https://gist.github.com/cdipaolo/d3f8db3848278b49db68">this</a> Go code.<br/>
		To limit by MaxResults, the result set is first filtered by distance from the search point ascending, and the limited. In other words, the reports closest to the search point will be returned.<br/>
		<br>
		<strong>Limitations</strong></p>

		<ul>
		<li>The GeoCoderType parameter is ignored.</li>
		<li>Coordinates must be in the typical Decimal Degrees WGS84 Compliant format, as you would get if you right click on Google Maps and select &#8220;What&#8217;s here?&#8221;. For example, San Jose City Hall is: (37.337930, &#8211;121.885891).</li>
		</ul>

		<h4 id="getreportsbyzipcode">Get Reports by ZipCode</h4>

		<p><a href="https://citysourced.zendesk.com/entries/30569088-API-Method-GetReportsByZipCode-">CitySourced KB: GetReportsByZipCode()</a> - <em>implemented</em><br/>
		Gets reports matching the specified Zip Code. </p>

		<hr />

		<h3 id="notimpemented">Not Impemented</h3>

		<ul>
		<li>CreateMobileDevice() - purpose and usage of this call is unclear from the City Sourced docs.</li>
		<li><a href="#GetReportsByAuthorId">GetReportsByAuthorId()</a></li>
		<li><a href="#CreateThreeOneOneMedia">CreateThreeOneOneMedia()</a></li>
		</ul>
				`,
	}
)
