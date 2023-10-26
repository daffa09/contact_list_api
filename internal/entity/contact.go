package entity

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}
