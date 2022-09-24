package model

type User struct {
	ID                   string               `json:"id"`
	Email                string               `json:"email"`
	Password             string               `json:"password"`
	IsActive             bool                 `json:"isActive"`
	FirstName            string               `json:"firstName"`
	LastName             string               `json:"lastName"`
	AdditionalName       string               `json:"additionalName"`
	ProfileImageURL      string               `json:"profileImageUrl"`
	BackgroundImageURL   string               `json:"backgroundImageUrl"`
	Pronouns             string               `json:"pronouns"`
	Headline             string               `json:"headline"`
	About                string               `json:"about"`
	Country              string               `json:"country"`
	City                 string               `json:"city"`
	ProfileLink          string               `json:"profileLink"`
	ActivateAccount      ActivateAccount      `json:"activeAcount" gorm:"foreignKey:UserID"`
	ResetPasswordAccount ResetPasswordAccount `json:"ResetPasswordAccount" gorm:"foreignKey:UserID"`
	Visits               []*User              `json:"Visit" gorm:"many2many:user_visits"`
	Follows              []*User              `json:"Follow" gorm:"many2many:user_follows"`
	Comment              []*Comment           `json:"Comment" gorm:"foreignKey:CommenterID;"`
	LikeComment          []*LikeComment       `json:"LikeComment" gorm:"foreignKey:UserID"`
	Connection           []*Connection        `json:"Connection" gorm:"foreignKey:User1ID;foreignKey:User2ID"`
	ConnectRequest       []*ConnectRequest    `json:"ConnectRequest" gorm:"foreignKey:FromUserID;foreignKey:ToUserID"`
	Block                []*User              `json:"Block" gorm:"many2many:user_blocks"`
	Experiences          []*Experience        `json:"Experiences" gorm:"foreignKey:UserID"`
	Educations           []*Education         `json:"Educations" gorm:"foreignKey:UserID"`
	Notification         []*Notification      `json:"Notification" gorm:"foreignKey:FromUserID;foreignKey:ToUserID"`
	Messages             []*Message           `json:"messages" gorm:"foreignKey:ShareProfileID;foreignKey:SenderID"`
	Rooms                []*Room              `json:"rooms" gorm:"foreignKey:User1ID;foreignKey:User2ID"`
	VideoCall            []*VideoCall         `json:"VideoCall" gorm:"foreignKey:User1ID;foreignKey:User2ID"`
}

type Follow struct {
	UserID   string `json:"userId"`
	FollowID string `json:"followId"`
}

type Visit struct {
	UserID  string `json:"userId"`
	VisitID string `json:"visitId"`
}

type ActivateAccount struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

type ResetPasswordAccount struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

type Connection struct {
	ID      string `json:"id"`
	User1   *User  `json:"user1"`
	User1ID string `json:"user1Id"`
	User2   *User  `json:"user2"`
	User2ID string `json:"user2Id"`
}

type ConnectRequest struct {
	ID         string `json:"id"`
	FromUserID string `json:"fromUserId"`
	FromUser   *User  `json:"fromUser"`
	ToUser     *User  `json:"toUser"`
	ToUserID   string `json:"toUserId"`
	Message    string `json:"message"`
}
