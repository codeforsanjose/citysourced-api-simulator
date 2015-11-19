curl -X "POST" "http://localhost:5050/api/" \
	-d $'<?xml version="1.0" encoding="utf-8" ?>
<CsRequest>
	<ApiAuthKey>a01234567890z</ApiAuthKey>
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
'
