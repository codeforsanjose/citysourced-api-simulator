curl -X "POST" "http://localhost:5050/api/" \
	-d $'<?xml version="1.0" encoding="utf-8" ?>
<CsRequest>
	<ApiAuthKey>a01234567890z</ApiAuthKey>
	<ApiRequestType>UpdateThreeOneOne</ApiRequestType>
	<ApiRequestVersion>1</ApiRequestVersion>
	<StatusType>Received</StatusType>
	<TicketId>102</TicketId>
	<TicketSla>Every attempt will be made to remove this graffiti within 2 weeks.</TicketSla>
</CsRequest>
'
