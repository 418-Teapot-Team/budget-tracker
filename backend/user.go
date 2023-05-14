package budget

const (
	usersTable = "users"
)

func (User) TableName() string {
	return usersTable

}

type User struct {
	Id       int    `json:"-" gorm:"column:id"`
	FullName string `json:"fullName" binding:"required" gorm:"column:full_name"`
	Email    string `json:"email" binding:"required" gorm:"column:email"`
	Password string `json:"password" binding:"required" gorm:"column:password_hash"`
}
