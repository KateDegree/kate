package entity

type UserEntity struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	AccountCode string `json:"accountCode"`
	Password    string `json:"password"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}
