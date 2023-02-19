package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	UserID string `json:"userid"`
	User   User   `json:"user";grom:"foreignkey:UserID"`
	// User User `grom:"foreignkey:UserID"`
}
