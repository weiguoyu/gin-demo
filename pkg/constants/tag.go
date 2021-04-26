package constants

const (
	ALL           = "all"     //所有字段
	FUZZY         = "fuzzy"   //模糊查询的tag fuzzy:"true"表示支持模糊查询
	SELECTEXCLUDE = "exclude" //在查询时，Select(iam.User{}...) exclude:"true" 表示该字段不会被select
	GTE           = "gte"     //时间字段上加上后，在查询时>=  gte:"true"
	LTE           = "lte"     //时间字段上加上后，在查询时<=  lte:"true"
)

type TagInfo struct {
	TagName  string
	TagValue string
	TagType  string
}

var TAG_SET = []*TagInfo{
	{FUZZY, "true", "fuzzy"},
	{SELECTEXCLUDE, "true", "select"},
	{GTE, "true", "condition"},
	{LTE, "true", "condition"},
}
