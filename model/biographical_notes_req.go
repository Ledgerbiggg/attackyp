package model

type BiographicalNotesReq struct {
	Current      int    `json:"current"`
	PageSize     int    `json:"pageSize"`
	SortField    string `json:"sortField"`
	SortOrder    string `json:"sortOrder"`
	SearchText   string `json:"searchText"`
	ReviewStatus int    `json:"reviewStatus"`
}

func NewBiographicalNotesReq(current int, pageSize int, sortField string, sortOrder string, searchText string, reviewStatus int) *BiographicalNotesReq {
	return &BiographicalNotesReq{Current: current, PageSize: pageSize, SortField: sortField, SortOrder: sortOrder, SearchText: searchText, ReviewStatus: reviewStatus}
}
