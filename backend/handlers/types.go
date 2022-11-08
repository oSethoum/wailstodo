package handlers

type ManyQuery struct {
	OneQuery
	Like   *string `json:"like"`
	Limit  *int    `json:"limit"`
	Offset *uint   `json:"offset"`
}

type OneQuery struct {
	Deleted *bool   `json:"deleted"`
	Preload *string `json:"preload"`
	Select  *string `json:"select"`
}
