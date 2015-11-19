curl -X "POST" "http://localhost:5050/api/" \
	-d $'<?xml version="1.0" encoding="utf-8" ?>
<CsRequest>
	<ApiAuthKey>a01234567890z</ApiAuthKey>
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
'
