package articleslistcollector

// Generated by https://quicktype.io

type ArticlesList struct {
	Result   int64     `json:"result"`
	Cursor   string    `json:"cursor"`
	HostName string    `json:"host-name"`
	Articles []Article `json:"data"`
}

type Article struct {
	CommentCount int64  `json:"commentCount"`
	ArticleID    int64  `json:"articleId"`
	RealmID      int64  `json:"realmId"`
	Title        string `json:"title"`
	UserID       int64  `json:"userId"`
	CreateTime   int64  `json:"createTime"`
	Description  string `json:"description"`
	UserName     string `json:"userName"`
	IsOriginal   bool   `json:"isOriginal"`
}
