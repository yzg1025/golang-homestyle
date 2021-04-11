package utils


var(
	LoginVerify           = Rules{"Phone":{NotEmpty()},"Password":{NotEmpty()}}
	SendCode              = Rules{"Phone":{NotEmpty()}}
	CountryCode           = Rules{"CountryCode":{NotEmpty()},"Cname":{NotEmpty()},"Code":{NotEmpty()}}
	ChangePasswordVerify  = Rules{"Phone":{NotEmpty()},"Password":{NotEmpty()},"NewPassword":{NotEmpty()}}
)
