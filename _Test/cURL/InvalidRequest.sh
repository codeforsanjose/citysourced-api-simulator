curl -X "POST" "http://localhost:5050/api/" \
	-d $'<?xml version="1.0" encoding="utf-8" ?>
<CsRequest>
   <ApiAuthKey>a01234567890z</ApiAuthKey>
   <ApiRequestType>CreateSomething</ApiRequestType>
   <ApiRequestVersion>1</ApiRequestVersion>
</CsRequest>'
