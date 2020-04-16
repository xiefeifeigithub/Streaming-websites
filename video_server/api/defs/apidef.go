package defs

// requests
type UserCredential struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

// response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type SignedIn struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type UserInfo struct {
	Id int `json:"id"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

// Data model
type User struct {
	Id int
	LoginName string
	Pwd string
}

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