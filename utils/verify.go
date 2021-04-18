package utils


var(
	IDVer                 = Rules{"ID": {NotEmpty()}}
	LoginVerify           = Rules{"Phone":{NotEmpty()},"Password":{NotEmpty()}}
	Register              = Rules{"Phone": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	SendCode              = Rules{"Phone":{NotEmpty()}}
	CountryCode           = Rules{"CountryCode":{NotEmpty()},"Cname":{NotEmpty()},"Code":{NotEmpty()}}
	ChangePasswordVerify  = Rules{"Phone":{NotEmpty()},"Password":{NotEmpty()},"NewPassword":{NotEmpty()}}
	BannerVer             = Rules{"Url":{NotEmpty()},"RedirectUrl":{NotEmpty()},"Title":{NotEmpty()}}
)