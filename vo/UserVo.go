package vo

type UserVo struct {
	Id        int    `json:"id"`
	UserName  string `json:"username"`
	Nickname  string `json:"nickname"`
	IsRoot    bool   `json:"isRoot"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
