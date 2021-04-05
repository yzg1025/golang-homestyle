package utils


var(
	LoginVerify           = Rules{"Phone":{NotEmpty()},"Password":{NotEmpty()}}
	SendCode              = Rules{"Phone":{NotEmpty()}}
)
