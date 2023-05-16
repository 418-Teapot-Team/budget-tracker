package budget

const listTable = "lists"

func (List) TableName() string {
	return listTable

}

type List struct {
	Id        int     `json:"id,omitempty" gorm:"column:id"`
	UserId    int     `json:"-" gorm:"column:user_id"`
	Type      string  `json:"type" gorm:"column:type"`
	Category  string  `json:"category" gorm:"column:category"`
	Amount    float64 `json:"amount" gorm:"column:amount"`
	Comment   *string `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt string  `json:"created_at,omitempty" gorm:"column:created_at"`
}