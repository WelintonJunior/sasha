package entity

/*
{
	"dis_nome": "GGTI"
}
*/

type Disciplina struct {
	ID   int    `json:"dis_id" gorm:"column:dis_id;primaryKey;autoIncrement"`
	Nome string `json:"dis_nome" gorm:"column:dis_nome;size:50;not null"`
}

func (Disciplina) TableName() string {
	return "dis_disciplina"
}

func (Disciplina) GetPrimaryKey() string {
	return "dis_id"
}
