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
	Experiences          []*Experience        `json:"Experiences" gorm:"foreignKey:UserID"`
	Educations           []*Education         `json:"Educations" gorm:"foreignKey:UserID"`
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
