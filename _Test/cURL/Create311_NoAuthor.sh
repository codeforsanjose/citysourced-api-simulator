curl -X "POST" "http://localhost:5050/api/" \
	-d $'<?xml version="1.0" encoding="utf-8" ?>
<CsRequest>
   <ApiAuthKey>a01234567890z</ApiAuthKey>
   <ApiRequestType>CreateThreeOneOne</ApiRequestType>
   <ApiRequestVersion>1</ApiRequestVersion>
   <DateCreated>1974-05-20T13:45:30</DateCreated>
   <DeviceType>IPHONE</DeviceType>
   <DeviceModel>3GS</DeviceModel>
   <DeviceId>1111</DeviceId>
   <RequestType>Graffiti Removal</RequestType>
   <RequestTypeId>10</RequestTypeId>
   <Latitude>37.338194</Latitude>
   <Longitude>-121.883925</Longitude>
   <Directionality>25 N NW</Directionality>
   <Description>Graffiti at the Grocery Outleat (near SJ City Hall)</Description>
   <AuthorNameFirst></AuthorNameFirst>
   <AuthorNameLast></AuthorNameLast>
   <AuthorEmail></AuthorEmail>
   <AuthorTelephone></AuthorTelephone>
   <AuthorIsAnonymous>True</AuthorIsAnonymous>
   <KeyValuePairs>
      <KeyValuePair Key="Key1">Value1</KeyValuePair>
      <KeyValuePair Key="Key2">Value2</KeyValuePair>
   </KeyValuePairs>
</CsRequest>'
