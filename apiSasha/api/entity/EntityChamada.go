package entity

/*
{
	"cha_dis_disciplina": 1,
	"cha_id_professor" : 1
}
*/

type Chamada struct {
	ID               int        `gorm:"column:cha_id;primaryKey;autoIncrement"`
	ChaDisDisciplina int        `json:"cha_dis_disciplina" gorm:"column:cha_dis_disciplina"`
	Disciplina       Disciplina `gorm:"foreignKey:ChaDisDisciplina;references:ID"`
	ChaIDProfessor   int        `json:"cha_id_professor" gorm:"column:cha_id_professor"`
	Professor        Professor  `gorm:"foreignKey:ChaIDProfessor;references:ID"`
}

func (c *Chamada) TableName() string {
	return "cha_chamada"
}
