package entity

type Disciplina struct {
	ID   int    `gorm:"column:dis_id;primaryKey;autoIncrement"`
	Nome string `json:"dis_nome" gorm:"column:dis_nome;size:50;not null"`
}

func (c *Disciplina) TableName() string {
	return "dis_disciplina"
}
