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

type Education struct {
	ID             string  `json:"id" gorm:"type:varchar(191)"`
	UserID         string  `json:"userId"`
	School         string  `json:"school"`
	Degree         string  `json:"degree"`
	FieldStudy     string  `json:"fieldStudy"`
	Grade          float64 `json:"grade"`
	Activities     string  `json:"activities"`
	Description    string  `json:"description"`
	MonthStartDate string  `json:"monthStartDate"`
	MonthEndDate   string  `json:"monthEndDate"`
	YearStartDate  string  `json:"yearStartDate"`
	YearEndDate    string  `json:"yearEndDate"`
}

type Experience struct {
	ID             string `json:"id" gorm:"type:varchar(191)"`
	UserID         string `json:"userId"`
	EmploymentType string `json:"employmentType"`
	CompanyName    string `json:"companyName"`
	Country        string `json:"country"`
	City           string `json:"city"`
	IsActive       bool   `json:"isActive"`
	Industry       string `json:"industry"`
	MonthStartDate string `json:"monthStartDate"`
	MonthEndDate   string `json:"monthEndDate"`
	YearStartDate  string `json:"yearStartDate"`
	YearEndDate    string `json:"yearEndDate"`
}

type ActivateAccount struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

type ResetPasswordAccount struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}
