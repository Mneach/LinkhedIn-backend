package model

import "time"

type Message struct {
	ID             string     `json:"id"`
	Sender         *User      `json:"Sender"`
	Room           *Room      `json:"Room"`
	ShareProfile   *User      `json:"ShareProfile"`
	SharePost      *Post      `json:"SharePost"`
	VideoCall      *VideoCall `json:"VideoCall"`
	Text           string     `json:"text"`
	ImageURL       string     `json:"imageUrl"`
	SenderID       string
	RoomID         string
	ShareProfileID *string
	SharePostID    *string
	VideoCallID    *string
	CreatedAt      time.Time `json:"createdAt"`
}

type Room struct {
	ID        string `json:"id"`
	User1ID   string
	User1     *User `json:"user1"`
	User2ID   string
	User2     *User      `json:"user2"`
	Messages  []*Message `json:"messages"`
	CreatedAt time.Time  `json:"createdAt"`
}

type VideoCall struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Duration string `json:"duration"`
	User1    *User  `json:"user1"`
	User1ID  string
	User2    *User `json:"user2"`
	User2ID  string
	Messages []*Message `json:"messages" gorm:"foreignKey:VideoCallID"`
}
