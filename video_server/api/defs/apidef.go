package defs

// requests
type UserCredential struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id string `json:"id"`
	VideoId string `json:"video_id"`
	Author string `json:"author"`
	Content string
}

type SimpleSession struct {
	Username string // login_name
	TTL int64 // session过期时间
}