package entity

/*
{
	"pro_nome" : "Jonhson"
}
*/

type Professor struct {
	ID   int    `gorm:"column:pro_id;primaryKey;autoIncrement"`
	Nome string `json:"pro_nome" gorm:"column:pro_nome;size:150;not null"`
}

func (c *Professor) TableName() string {
	return "pro_professores"
}
