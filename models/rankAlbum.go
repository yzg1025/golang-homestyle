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
