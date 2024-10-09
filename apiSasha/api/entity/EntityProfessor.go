package entity

/*
{
	"pro_nome" : "Jonhson"
}
*/

type Professor struct {
	ID   int    `json:"pro_id" gorm:"column:pro_id;primaryKey;autoIncrement"`
	Nome string `json:"pro_nome" gorm:"column:pro_nome;size:150;not null"`
}

func (Professor) TableName() string {
	return "pro_professores"
}

func (Professor) GetPrimaryKey() string {
	return "pro_id"
}
