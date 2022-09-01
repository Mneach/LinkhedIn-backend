package model

type Post struct {
	ID       string  `json:"ID"`
	Text     string  `json:"text"`
	URL      string  `json:"url"`
	PhotoUrl string  `json:"photoUrl"`
	VideoUrl string  `json:"videoUrl"`
	SenderId string  `json:"senderId"`
	Sender   *User   `json:"Sender" gorm:"reference:User"`
	Likes    []*User `json:"Likes" gorm:"many2many:like_posts"`
}

type LikePosts struct {
	PostId string `json:"PostId"`
	UserId string `json:"UserId"`
}
