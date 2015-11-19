curl -X "POST" "http://localhost:5050/api/" \
	-d $'<?xml version="1.0" encoding="utf-8" ?>
<CsRequest>
	<ApiAuthKey>a01234567890z</ApiAuthKey>
	<ApiRequestType>GetReportsByDeviceId</ApiRequestType>
	<ApiRequestVersion>1</ApiRequestVersion>
	<DeviceId>3333</DeviceId>
    <DeviceType>IPHONE</DeviceType>
	<MaxResults></MaxResults>
	<IncludeDetails>False</IncludeDetails>
	<DateRangeStart>2015-05-20T13:45:30</DateRangeStart>
	<DateRangeEnd>2015-05-20T13:45:30</DateRangeEnd>
	<CurrentStatus></CurrentStatus>
</CsRequest>
'
