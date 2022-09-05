package utils

var (
	IDVer                = Rules{"ID": {NotEmpty()}}
	LoginVerify          = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}}
	Register             = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}}
	SendCode             = Rules{"Phone": {NotEmpty()}}
	CountryCode          = Rules{"CountryCode": {NotEmpty()}, "Cname": {NotEmpty()}, "Code": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	BannerVer            = Rules{"Url": {NotEmpty()}, "RedirectUrl": {NotEmpty()}, "Title": {NotEmpty()}}
	RoomCountVer         = Rules{"MinLng": {NotEmpty()}, "MaxLng": {NotEmpty()}, "MinLat": {NotEmpty()}, "MaxLat": {NotEmpty()}}
	RoomCountAddVer      = Rules{"Longitude": {NotEmpty()}, "Latitude": {NotEmpty()}}
	RoomList             = Rules{"MinLng": {NotEmpty()}, "MaxLng": {NotEmpty()}, "MinLat": {NotEmpty()}, "MaxLat": {NotEmpty()}, "Page": {NotEmpty()}}
	RankAlbumVer         = Rules{"IDs": {NotEmpty()}, "Page": {NotEmpty()}, "PageNum": {NotEmpty()}}
	RadioVer             = Rules{"PageNo": {NotEmpty()}, "PageSize": {NotEmpty()}}
	AlbumIdVar           = Rules{"AlbumId": {NotEmpty()}}
	HotWordBillboardVar  = Rules{"CategoryId": {NotEmpty()}}
	BasicInfoVar         = Rules{"Uid": {NotEmpty()}}
	AnchorFollowVar      = Rules{"Uid": {NotEmpty()}, "Follow": {NotEmpty()}}
	APIVar               = Rules{"ApiDesc": {NotEmpty()}, "ApiMethod": {NotEmpty()}, "ApiName": {NotEmpty()}, "ApiPath": {NotEmpty()}}
)
