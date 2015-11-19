curl -X "POST" "http://localhost:5050/api/" \
	-d $'<?xml version="1.0" encoding="utf-8" ?>
<CsRequest>
   <ApiAuthKey>a01234567890z</ApiAuthKey>
   <ApiRequestType>CreateThreeOneOne</ApiRequestType>
   <ApiRequestVersion>1</ApiRequestVersion>
   <DateCreated>1974-05-20T13:45:30</DateCreated>
   <DeviceType>IPHONE</DeviceType>
   <DeviceModel>3GS</DeviceModel>
   <DeviceId>2222</DeviceId>
   <RequestType>Graffiti Removal</RequestType>
   <RequestTypeId>10</RequestTypeId>
   <Latitude>37.339244</Latitude>
   <Longitude>-121.883638</Longitude>
   <Directionality>25 N NW</Directionality>
   <Description>Graffiti at Punjab cafe near City Hall.</Description>
   <AuthorNameFirst>Fred</AuthorNameFirst>
   <AuthorNameLast>Jones</AuthorNameLast>
   <AuthorEmail>jh@xyz.com</AuthorEmail>
   <AuthorTelephone>2223334444</AuthorTelephone>
   <AuthorIsAnonymous>False</AuthorIsAnonymous>
   <KeyValuePairs>
      <KeyValuePair Key="Key1">Value1</KeyValuePair>
      <KeyValuePair Key="Key2">Value2</KeyValuePair>
   </KeyValuePairs>
</CsRequest>'
