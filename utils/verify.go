package utils


var(
	LoginVerify           = Rules{"Phone":{NotEmpty()},"Password":{NotEmpty()}}
	Register              = Rules{"Phone": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	SendCode              = Rules{"Phone":{NotEmpty()}}
	CountryCode           = Rules{"CountryCode":{NotEmpty()},"Cname":{NotEmpty()},"Code":{NotEmpty()}}
	ChangePasswordVerify  = Rules{"Phone":{NotEmpty()},"Password":{NotEmpty()},"NewPassword":{NotEmpty()}}
)
