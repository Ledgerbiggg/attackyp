package model

// InterviewReq 结构体用于表示整个 JSON 数据
type InterviewReq struct {
	Category     string   `json:"category"`
	Current      int      `json:"current"`
	PageSize     int      `json:"pageSize"`
	ReviewStatus int      `json:"reviewStatus"`
	SortField    string   `json:"sortField"`
	SortOrder    string   `json:"sortOrder"`
	Tags         []string `json:"tags"`
}

func NewInterviewReq(category string, current int, pageSize int, reviewStatus int, sortField string, sortOrder string, tags []string) *InterviewReq {
	return &InterviewReq{Category: category, Current: current, PageSize: pageSize, ReviewStatus: reviewStatus, SortField: sortField, SortOrder: sortOrder, Tags: tags}
}
