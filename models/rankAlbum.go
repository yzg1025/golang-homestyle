package models

type RankAlbumParams struct {
	RankIds string `form:"rankIds"`
	Pagination
}

type RankResult struct {
	Count     int    `json:"count"`
	Ids       string `json:"ids"`
	CoverPath string `gorm:"column:coverPath" json:"coverPath"`
	Title     string `json:"title"`
	RankId    int    `gorm:"column:rankId" json:"rankId"`
	Subtitle  string `json:"subtitle"`
	ShareUrl  string `gorm:"column:shareUrl" json:"shareUrl"`
}

type Albums struct {
	ID          int    `json:"id"`
	AlbumTitle  string `gorm:"column:albumTitle" json:"albumTitle"`
	AnchorName  string `gorm:"column:anchorName" json:"anchorName"`
	AlbumUrl    string `gorm:"column:albumUrl" json:"albumUrl"`
	Cover       string `json:"cover"`
	AnchorUrl   string `gorm:"column:anchorUrl" json:"anchorUrl"`
	PlayCount   string `gorm:"column:playCount" json:"playCount"`
	TrackCount  string `gorm:"column:trackCount" json:"trackCount"`
	Description string `json:"description"`
	TagStr      string `gorm:"column:tagStr" json:"tagStr"`
	IsPaid      string `gorm:"column:isPaid" json:"isPaid"`
	Price       string `json:"price"`
	VipType     string `gorm:"column:vipType" json:"vipType"`
}

type AlbumRankPageList struct {
	RankResult RankResult `gorm:"column:rankResult" json:"rankResult"`
	Albums     []Albums   `gorm:"column:albums" json:"albums"`
}

type RecommendInfoList struct {
	AlbumId           int    `gorm:"column:albumId" json:"albumId"`
	AlbumPlayCount    int64  `gorm:"column:albumPlayCount" json:"albumPlayCount"`
	AlbumTrackCount   int    `gorm:"column:albumTrackCount" json:"albumTrackCount"`
	AlbumCoverPath    string `gorm:"column:albumCoverPath" json:"albumCoverPath"`
	AlbumTitle        string `gorm:"column:albumTitle" json:"albumTitle"`
	AlbumUserNickName string `gorm:"column:albumUserNickName" json:"albumUserNickName"`
	AnchorId          int64  `gorm:"column:anchorId" json:"anchorId"`
	AnchorGrade       int    `gorm:"column:anchorGrade" json:"anchorGrade"`
	IsDeleted         bool   `gorm:"column:isDeleted" json:"isDeleted"`
	IsPaid            bool   `gorm:"column:isPaid" json:"isPaid"`
	IsFinished        int    `gorm:"column:isFinished" json:"isFinished"`
	AnchorUrl         string `gorm:"column:anchorUrl" json:"anchorUrl"`
	AlbumUrl          string `gorm:"column:albumUrl" json:"albumUrl"`
	Intro             string `gorm:"column:intro" json:"intro"`
	VipType           int    `gorm:"column:vipType" json:"vipType"`
	LogoType          int    `gorm:"column:logoType" json:"logoType"`
	AlbumSubscript    int    `gorm:"column:albumSubscript" json:"albumSubscript"`
}

type RecommendAnchorList struct {
	Uid            int    `json:"uid"`
	CoverPath      string `gorm:"column:coverPath" json:"coverPath"`
	AnchorNickName string `gorm:"column:anchorNickName" json:"anchorNickName"`
	Background     string `json:"background"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Grade          int64  `json:"grade"`
	GradeType      int    `gorm:"column:gradeType" json:"gradeType"`
	TrackCount     int64  `gorm:"column:trackCount" json:"trackCount"`
	AlbumCount     int    `gorm:"column:albumCount" json:"albumCount"`
	FollowerCount  int    `gorm:"column:followerCount" json:"followerCount"`
	FollowingCount string `gorm:"column:followingCount" json:"followingCount"`
	IsFollow       bool   `gorm:"column:isFollow" json:"isFollow"`
	beFollow       bool   `gorm:"column:beFollow" json:"beFollow"`
	isBlack        bool   `gorm:"column:isBlack" json:"isBlack"`
	LogoType       int    `gorm:"column:logoType" json:"logoType"`
	Ptitle         string `gorm:"column:ptitle" json:"ptitle"`
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

type Albumsimple struct {
	AlbumId             string `gorm:"column:albumId" json:"albumId"`
	AnchorUid           int    `gorm:"column:anchorUid" json:"anchorUid"`
	AlbumStatus         int    `gorm:"column:albumStatus" json:"albumStatus"`
	ShowApplyFinishBtn  bool   `gorm:"column:showApplyFinishBtn" json:"showApplyFinishBtn"`
	ShowEditBtn         bool   `gorm:"column:showEditBtn" json:"showEditBtn"`
	ShowTrackManagerBtn bool   `gorm:"column:showTrackManagerBtn" json:"showTrackManagerBtn"`
	ShowInformBtn       bool   `gorm:"column:showInformBtn" json:"showInformBtn"`
	Cover               string `json:"cover"`
	AlbumTitle          string `gorm:"column:albumTitle" json:"albumTitle"`
	UpdateDate          string `gorm:"column:updateDate" json:"updateDate"`
	CreateDate          string `gorm:"column:createDate" json:"createDate"`
	PlayCount           int64  `gorm:"column:playCount" json:"playCount"`
	IsPaid              bool   `gorm:"column:isPaid" json:"isPaid"`
	IsFinished          string `gorm:"column:isFinished" json:"isFinished"`
	IsSubscribe         bool   `gorm:"column:isSubscribe" json:"isSubscribe"`
	IsPublic            bool   `gorm:"column:isPublic" json:"isPublic"`
	HasBuy              bool   `gorm:"column:hasBuy" json:"hasBuy"`
	VipType             int    `gorm:"column:vipType" json:"vipType"`
	CustomTitle         string `gorm:"column:customTitle" json:"customTitle"`
	RecommendReason     string `gorm:"column:recommendReason" json:"recommendReason"`
	AlbumSubscript      int    `gorm:"column:albumSubscript" json:"albumSubscript"`
	Tags                string `json:"tags"`
	XimaVipFreeType     int    `gorm:"column:ximaVipFreeType" json:"ximaVipFreeType"`
	SubscribeCount      int    `gorm:"column:subscribeCount" json:"subscribeCount"`
	ShortIntro          string `gorm:"type:text;column:shortIntro" json:"shortIntro"`
	RichIntro           string `gorm:"type:text;column:richIntro" json:"richIntro"`
	DetailRichIntro     string `gorm:"type:text;column:detailRichIntro" json:"detailRichIntro"`
}

// https://www.ximalaya.com/revision/seo/hotWordAlbums?id=4615999&queryType=1
type HotWordAlbums struct {
	CoverPath  string `gorm:"column:coverPath" json:"coverPath"`
	ID         int    `json:"id"`
	Intro      string `json:"intro"`
	IsPaid     bool   `gorm:"column:isPaid" json:"isPaid"`
	Link       string `json:"link"`
	Nickname   string `json:"nickname"`
	PlayCount  int    `gorm:"column:playCount" json:"playCount"`
	Title      string `json:"title"`
	TrackCount int    `gorm:"column:trackCount" json:"trackCount"`
	VipType    int    `gorm:"column:vipType" json:"vipType"`
}

// https://www.ximalaya.com/revision/album/v1/getTracksList?albumId=4615999&pageNum=40
type TracksAlbum struct {
	ID               int    `json:"id"`
	Index            int    `json:"index"`
	AlbumCoverPath   string `gorm:"column:albumCoverPath" json:"albumCoverPath"`
	AlbumId          int    `gorm:"column:albumId" json:"albumId"`
	AlbumTitle       string `gorm:"column:albumTitle" json:"albumTitle"`
	AnchorId         int    `gorm:"column:anchorId" json:"anchorId"`
	AnchorName       string `gorm:"column:anchorName" json:"anchorName"`
	BreakSecond      int    `gorm:"column:breakSecond" json:"breakSecond"`
	CreateDateFormat string `gorm:"column:createDateFormat" json:"createDateFormat"`
	Duration         int    `gorm:"column:duration" json:"duration"`
	IsLike           bool   `gorm:"column:isLike" json:"isLike"`
	IsPaid           bool   `gorm:"column:isPaid" json:"isPaid"`
	IsVideo          bool   `gorm:"column:isVideo" json:"isVideo"`
	IsVipFirst       bool   `gorm:"column:isVipFirst" json:"isVipFirst"`
	Length           int    `json:"length"`
	PlayCount        int    `gorm:"column:playCount" json:"playCount"`
	Tag              int    `json:"tag"`
	Title            string `json:"title"`
	TrackId          int    `gorm:"column:trackId" json:"trackId"`
	Url              string `json:"url"`
}

type TabList struct {
	GroupId     int    `gorm:"column:groupId" json:"groupId"`
	GroupName   string `gorm:"column:groupName" json:"groupName"`
	Position    int    `gorm:"column:position" json:"position"`
	ChannelId   int    `gorm:"column:channelId" json:"channelId"`
	ChannelName string `gorm:"column:channelName" json:"channelName"`
	ChannelLink string `gorm:"column:channelLink" json:"channelLink"`
}

type TabViewList struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ChannelCount int    `gorm:"column:channelCount" json:"channelCount"`
	Position     int    `json:"position"`
}

// hotWordBillboardCategory

type HotWordBillboardCategory struct {
	CategoryId   int    `gorm:"column:categoryId" json:"categoryId"`
	CategoryName string `gorm:"column:categoryName" json:"categoryName"`
}

type HotWordBillboardParams struct {
	CategoryId string `form:"categoryId"`
}

// SearchCateList
type HotWordBillboard struct {
	Word                 string `json:"word"`
	Shift                int    `json:"shift"`
	DisplayType          int    `gorm:"column:display_type" json:"display_type"`
	IsThrough            bool   `gorm:"column:is_through" json:"is_through"`
	ThroughType          int    `gorm:"column:through_type" json:"through_type"`
	TgtId                int    `gorm:"column:tgt_id" json:"tgt_id"`
	Count                int    `json:"count"`
	Intro                string `json:"intro"`
	Url                  string `json:"url"`
	DisplayWord          string `gorm:"column:display_word" json:"display_word"`
	OutsideHotSearchType int    `gorm:"column:outside_hot_search_type" json:"outside_hot_search_type"`
	Score                int    `json:"score"`
	CategoryInt          int    `gorm:"column:categoryInt" json:"categoryInt"`
}

type AnchorBasicInfoParams struct {
	Uid string `form:"uid"`
}

type AnchorFollowParams struct {
	Uid    string `form:"uid"`
	Follow string `form:"follow"`
}

type AnchorBasicInfo struct {
	Uid                 int    `json:"uid"`
	NickName            string `gorm:"column:nickName" json:"nickName"`
	Cover               string `json:"cover"`
	Background          string `json:"background"`
	IsVip               bool   `gorm:"column:isVip" json:"isVip"`
	ConstellationType   int    `gorm:"column:constellationType" json:"constellationType"`
	PersonalSignature   string `gorm:"column:personalSignature" json:"personalSignature"`
	PersonalDescription string `gorm:"column:personalDescription" json:"personalDescription"`
	FansCount           int    `gorm:"column:fansCount" json:"fansCount"`
	Gender              int    `json:"gender"`
	BirthMonth          int    `gorm:"column:birthMonth" json:"birthMonth"`
	BirthDay            int    `gorm:"column:birthDay" json:"birthDay"`
	Province            string `json:"province"`
	City                string `json:"city"`
	AnchorGrade         int    `gorm:"column:anchorGrade" json:"anchorGrade"`
	AnchorGradeType     int    `gorm:"column:anchorGradeType" json:"anchorGradeType"`
	IsMusician          bool   `gorm:"column:isMusician" json:"isMusician"`
	AnchorUrl           string `gorm:"column:anchorUrl" json:"anchorUrl"`
	LogoType            int    `gorm:"column:logoType" json:"logoType"`
	FollowingCount      int    `gorm:"column:followingCount" json:"followingCount"`
	TracksCount         int    `gorm:"column:tracksCount" json:"tracksCount"`
	AlbumsCount         int    `gorm:"column:albumsCount" json:"albumsCount"`
	AlbumCountReal      int    `gorm:"column:albumCountReal" json:"albumCountReal"`
	UserCompany         string `gorm:"column:userCompany" json:"userCompany"`
	IsFollow            bool   `gorm:"column:isFollow" json:"isFollow"`
	BeFollow            bool   `gorm:"column:beFollow" json:"beFollow"`
	IsBlack             bool   `gorm:"column:isBlack" json:"isBlack"`
}

type AlbumComments struct {
	AlbumId       int    `gorm:"column:albumId" json:"albumId"`
	AlbumPaid     bool   `gorm:"column:albumPaid" json:"albumPaid"`
	AlbumUid      int    `gorm:"column:albumUid" json:"albumUid"`
	AuditStatus   int    `gorm:"column:auditStatus" json:"auditStatus"`
	CommentId     int    `gorm:"column:commentId" json:"commentId"`
	Content       string `gorm:"type:text" json:"content"`
	CreatedAt     int    `gorm:"column:createdAt" json:"createdAt"`
	IsHighQuality bool   `gorm:"column:isHighQuality" json:"isHighQuality"`
	Liked         bool   `json:"liked"`
	Likes         int    `json:"likes"`
	NewAlbumScore int    `gorm:"column:newAlbumScore" json:"newAlbumScore"`
	Nickname      string `json:"nickname"`
	ReplyCount    int    `gorm:"column:replyCount" json:"replyCount"`
	SmallHeader   string `gorm:"column:smallHeader" json:"smallHeader"`
	Uid           int    `json:"uid"`
	UpdatedAt     int    `gorm:"column:updatedAt" json:"updatedAt"`
	VLogoType     int    `gorm:"column:vLogoType" json:"vLogoType"`
}

type ScoreDiagram struct {
	AlbumId   int     `gorm:"column:albumId" json:"albumId"`
	FiveStar  float32 `gorm:"column:fiveStar" json:"fiveStar"`
	FourStar  float32 `gorm:"column:fourStar" json:"fourStar"`
	IsExited  bool    `gorm:"column:isExited" json:"isExited"`
	OneStar   float32 `gorm:"column:oneStar" json:"oneStar"`
	ThreeStar float32 `gorm:"column:threeStar" json:"threeStar"`
	TwoStar   float32 `gorm:"column:twoStar" json:"twoStar"`
}

type Anchoralbums struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	SubTitle       string `gorm:"column:subTitle" json:"subTitle"`
	CoverPath      string `gorm:"column:coverPath" json:"coverPath"`
	IsFinished     bool   `gorm:"column:isFinished" json:"isFinished"`
	IsPaid         bool   `gorm:"column:isPaid" json:"isPaid"`
	AnchorUrl      string `gorm:"column:anchorUrl" json:"anchorUrl"`
	AnchorUid      int32  `gorm:"column:anchorUid" json:"anchorUid"`
	PlayCount      int    `gorm:"column:playCount" json:"playCount"`
	TrackCount     int    `gorm:"column:trackCount" json:"trackCount"`
	AlbumUrl       string `gorm:"column:albumUrl" json:"albumUrl"`
	Description    string `gorm:"column:description" json:"description"`
	VipType        int    `gorm:"column:vipType" json:"vipType"`
	AlbumSubscript int    `gorm:"column:albumSubscript" json:"albumSubscript"`
}
