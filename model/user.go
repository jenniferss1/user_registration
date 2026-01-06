package model

type Users struct {
	ID          int     `json:"user_id"`
	User_name   string  `json:"name"`
	User_age    int     `json:"age"`
	User_weight float64 `json:"weight"`
	User_height float64 `json:"height"`
}
