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
	Experiences          []*Experience        `json:"Experiences" gorm:"foreignKey:UserID"`
	Educations           []*Education         `json:"Educations" gorm:"foreignKey:UserID"`
}

type ActivateAccount struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

type ResetPasswordAccount struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}
