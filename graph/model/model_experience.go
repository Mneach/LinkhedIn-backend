package model

type Experience struct {
	ID             string `json:"id" gorm:"type:varchar(191)"`
	UserID         string `json:"userId"`
	Title          string `json:"title"`
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
