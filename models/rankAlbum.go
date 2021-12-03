package models

type RankAlbumParams struct {
	RankIds string `form:"rankIds"`
	Pagination
}

type RankResult struct {
	Count     int    `json:"count"`
	Ids       string `json:"ids"`
	CoverPath string `gorm:"column:coverPath";json:"coverPath"`
	Title     string `json:"title"`
	RankId    int    `gorm:"column:rankId";json:"rankId"`
	Subtitle  string `json:"subtitle"`
	ShareUrl  string `gorm:"column:shareUrl";json:"shareUrl"`
}

type Albums struct {
	ID          int    `json:"id"`
	AlbumTitle  string `gorm:"column:albumTitle";json:"albumTitle"`
	AnchorName  string `gorm:"column:anchorName";json:"anchorName"`
	AlbumUrl    string `gorm:"column:albumUrl";json:"albumUrl"`
	Cover       string `json:"cover"`
	AnchorUrl   string `gorm:"column:anchorUrl";json:"anchorUrl"`
	PlayCount   string `gorm:"column:playCount";json:"playCount"`
	TrackCount  string `gorm:"column:trackCount";json:"trackCount"`
	Description string `json:"description"`
	TagStr      string `gorm:"column:tagStr";json:"tagStr"`
	IsPaid      string `gorm:"column:isPaid";json:"isPaid"`
	Price       string `json:"price"`
	VipType     string `gorm:"column:vipType";json:"vipType"`
}

type AlbumRankPageList struct {
	RankResult RankResult
	Albums     []Albums
}

type RecommendInfoList struct {
	AlbumId           int    `gorm:"column:albumId";json:"albumId"`
	AlbumPlayCount    int64  `gorm:"column:albumPlayCount";json:"albumPlayCount"`
	AlbumTrackCount   int    `gorm:"column:albumTrackCount";json:"albumTrackCount"`
	AlbumCoverPath    string `gorm:"column:albumCoverPath";json:"albumCoverPath"`
	AlbumTitle        string `gorm:"column:albumTitle";json:"albumTitle"`
	AlbumUserNickName string `gorm:"column:albumUserNickName";json:"albumUserNickName"`
	AnchorId          int64  `gorm:"column:anchorId";json:"anchorId"`
	AnchorGrade       int    `gorm:"column:anchorGrade";json:"anchorGrade"`
	IsDeleted         bool   `gorm:"column:isDeleted";json:"isDeleted"`
	IsPaid            bool   `gorm:"column:isPaid";json:"isPaid"`
	IsFinished        int    `gorm:"column:isFinished";json:"isFinished"`
	AnchorUrl         string `gorm:"column:anchorUrl";json:"anchorUrl"`
	AlbumUrl          string `gorm:"column:albumUrl";json:"albumUrl"`
	Intro             string `gorm:"column:intro";json:"intro"`
	VipType           int    `gorm:"column:vipType";json:"vipType"`
	LogoType          int    `gorm:"column:logoType";json:"logoType"`
	AlbumSubscript    int    `gorm:"column:albumSubscript";json:"albumSubscript"`
}

type RecommendAnchorList struct {
	Uid            int    `json:"uid"`
	CoverPath      string `gorm:"column:coverPath";json:"coverPath"`
	AnchorNickName string `gorm:"column:anchorNickName";json:"anchorNickName"`
	Background     string `json:"background"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Grade          int64  `json:"grade"`
	GradeType      int    `gorm:"column:gradeType";json:"gradeType"`
	TrackCount     int64  `gorm:"column:trackCount";json:"trackCount"`
	AlbumCount     int    `gorm:"column:albumCount";json:"albumCount"`
	FollowerCount  int    `gorm:"column:followerCount";json:"followerCount"`
	FollowingCount string `gorm:"column:followingCount";json:"followingCount"`
	IsFollow       bool   `gorm:"column:isFollow";json:"isFollow"`
	beFollow       bool   `gorm:"column:beFollow";json:"beFollow"`
	isBlack        bool   `gorm:"column:isBlack";json:"isBlack"`
	LogoType       int    `gorm:"column:logoType";json:"logoType"`
	Ptitle         string `gorm:"column:ptitle";json:"ptitle"`
}

type CategoryAll struct {
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Position      int          `json:"position"`
	GroupType     int          `gorm:"column:groupType" json:"groupType"`
	DisplayStatus int          `gorm:"column:displayStatus" json:"displayStatus"`
	Categories    []Categories `json:"categories" gorm:"foreignKey:ID"`
}

type Categories struct {
	ID            int             `json:"id"`
	CategoryType  int             `gorm:"column:categoryType" json:"categoryType"`
	DisplayStatus int             `gorm:"column:displayStatus" json:"displayStatus"`
	DisplayName   string          `gorm:"column:displayName" json:"displayName"`
	Link          string          `json:"link"`
	Name          string          `json:"name"`
	PicPath       string          `gorm:"column:picPath" json:"picPath"`
	Pinyin        string          `json:"pinyin"`
	Position      int             `json:"position"`
	Subcategories []Subcategories `json:"subcategories" gorm:"foreignKey:CategoryId"`
}

type Subcategories struct {
	ID            int    `json:"id"`
	CategoryId    int    `gorm:"column:categoryId" json:"categoryId"`
	Position      int    `json:"position"`
	MetadataId    int64  `gorm:"column:metadataId" json:"metadataId"`
	MetadataValue string `gorm:"column:metadataValue" json:"metadataValue"`
	Code          string `json:"code"`
	DisplayStatus int    `gorm:"column:displayStatus" json:"displayStatus"`
	Link          string `json:"link"`
	DisplayValue  int    `gorm:"column:displayValue" json:"displayValue"`
	IsKeyword     bool   `gorm:"column:isKeyword" json:"isKeyword"`
}

type Radio struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	CoverSmall        string `gorm:"column:coverSmall" json:"coverSmall"`
	CoverLarge        string `gorm:"column:coverLarge" json:"coverLarge"`
	ProgramScheduleId int    `gorm:"column:programScheduleId" json:"programScheduleId"`
	ProgramId         int    `gorm:"column:programId" json:"programId"`
	ProgramName       string `gorm:"column:programName" json:"programName"`
	Start             string `json:"start"`
	End               string `json:"end"`
	NextStart         string `gorm:"column:nextStart" json:"nextStart"`
	NextEnd           string `gorm:"column:nextEnd" json:"nextEnd"`
	NextProgramName   string `gorm:"column:nextProgramName" json:"nextProgramName"`
	PlayCount         int    `gorm:"column:playCount" json:"playCount"`
	FmUid             int    `gorm:"column:fmUid" json:"fmUid"`
	Type              string `json:"type"`
	Sort              int    `json:"sort"`
	IsFavorite        bool   `gorm:"column:isFavorite" json:"isFavorite"`
	CategoryId        int    `gorm:"column:categoryId" json:"categoryId"`
	CategoryName      string `gorm:"column:categoryName" json:"categoryName"`
	CategoryRank      int    `gorm:"column:categoryRank" json:"categoryRank"`
}

// https://www.ximalaya.com/revision/album/v1/simple?albumId=4615999
type Albumsimple struct{
	AnchorUid   int `gorm:"column:anchorUid" json:"anchorUid"`,
	AlbumStatus   int `gorm:"column:albumStatus" json:"albumStatus"`
	ShowApplyFinishBtn   bool `gorm:"column:showApplyFinishBtn" json:"showApplyFinishBtn"`
	ShowEditBtn   bool `gorm:"column:showEditBtn" json:"showEditBtn"`
	ShowTrackManagerBtn   bool `gorm:"column:showTrackManagerBtn" json:"showTrackManagerBtn"`
	ShowInformBtn   bool `gorm:"column:showInformBtn" json:"showInformBtn"`
	Cover   string `json:"cover"`
	AlbumTitle   string `gorm:"column:albumTitle" json:"albumTitle"`
	UpdateDate   string `gorm:"column:updateDate" json:"updateDate"`
	CreateDate   string `gorm:"column:createDate" json:"createDate"`
	PlayCount   int64 `gorm:"column:playCount" json:"playCount"`
	isPaid   bool `gorm:"column:isPaid" json:"isPaid"`
	IsFinished   string `gorm:"column:isFinished" json:"isFinished"`
	IsSubscribe   bool `gorm:"column:isSubscribe" json:"isSubscribe"`
	shortIntro   string `gorm:"column:shortIntro" json:"shortIntro"`
	IsPublic   bool `gorm:"column:isPublic" json:"isPublic"`
	HasBuy   bool `gorm:"column:hasBuy" json:"hasBuy"`
	VipType   int `gorm:"column:vipType" json:"vipType"`
	CustomTitle   string `gorm:"column:customTitle" json:"customTitle"`
	RecommendReason   string `gorm:"column:recommendReason" json:"recommendReason"`
	AlbumSubscript   int `gorm:"column:albumSubscript" json:"albumSubscript"`
	Tags   string `json:"tags"`
	XimiVipFreeType   int `gorm:"column:ximiVipFreeType" json:"ximiVipFreeType"`
	PriceOp  PriceOp  `json:'priceOp'`
}
type PriceOp struct{
	gin.Model
	AlbumDiscountedPrice string `gorm:"column:albumDiscountedPrice" json:"albumDiscountedPrice"`
	AlbumPrice string `gorm:"column:albumPrice" json:"albumPrice"`
	PriceType int `gorm:"column:priceType" json:"priceType"`
	RemainAlbumTotalPrice string `gorm:"column:remainAlbumTotalPrice" json:"remainAlbumTotalPrice"`
}

// https://www.ximalaya.com/revision/seo/hotWordAlbums?id=4615999&queryType=1
type HotWordAlbums struct{
	CoverPath   string `gorm:"column:coverPath" json:"coverPath"`
	ID   int `json:"id"`
	Intro   string `json:"intro"`
	IsPaid   bool `gorm:"column:isPaid" json:"isPaid"`
	Link   string `json:"link"`
	Nickname   string `json:"nickname"`
	PlayCount   int `gorm:"column:playCount" json:"playCount"`
	Title   string `json:"title"`
	TrackCount   int `gorm:"column:trackCount" json:"trackCount"`
	VipType   int `gorm:"column:vipType" json:"vipType"`
}

// https://www.ximalaya.com/revision/album/v1/getTracksList?albumId=4615999&pageNum=40
type TracksAlbum struct{
	AlbumCoverPath   string `gorm:"column:albumCoverPath" json:"albumCoverPath"`
	AlbumId   int `gorm:"column:albumId" json:"albumId"`
	AlbumTitle   string `gorm:"column:albumTitle" json:"albumTitle"`
	AnchorId   int `gorm:"column:anchorId" json:"anchorId"`
	AnchorName   string `gorm:"column:anchorName" json:"anchorName"`
    BreakSecond   int `gorm:"column:breakSecond" json:"breakSecond"`
	CreateDateFormat   string `gorm:"column:createDateFormat" json:"createDateFormat"`
	Duration   int `gorm:"column:duration" json:"duration"`
	IsLike   bool `gorm:"column:isLike" json:"isLike"`
	IsPaid   bool `gorm:"column:isPaid" json:"isPaid"`
	IsVideo   bool `gorm:"column:isVideo" json:"isVideo"`
	IsVipFirst   bool `gorm:"column:isVipFirst" json:"isVipFirst"`
	Length   int `json:"length"`
	PlayCount   int `gorm:"column:playCount" json:"playCount"`
	Tag   int `json:"tag"`
	Title   string `titlejson:"title"`
	TrackId   int `gorm:"column:trackId" json:"trackId"`
	Url   string `json:"url"`
}
