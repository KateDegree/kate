package entity

type AccessTokenEntity struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
