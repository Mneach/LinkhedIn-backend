package model

type Education struct {
	ID             string `json:"id" gorm:"type:varchar(191)"`
	UserID         string `json:"userId"`
	School         string `json:"school"`
	Degree         string `json:"degree"`
	FieldStudy     string `json:"fieldStudy"`
	Grade          string `json:"grade"`
	Activities     string `json:"activities"`
	Description    string `json:"description"`
	MonthStartDate string `json:"monthStartDate"`
	MonthEndDate   string `json:"monthEndDate"`
	YearStartDate  string `json:"yearStartDate"`
	YearEndDate    string `json:"yearEndDate"`
}
