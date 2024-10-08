package entity

/*
{
	"alu_aluno": "aluno"
}
*/

type Aluno struct {
	ID   int    `gorm:"column:alu_id;primaryKey;autoIncrement"`
	Nome string `json:"alu_aluno" gorm:"column:alu_nome;size:150;not null"`
	//Imagem []byte `gorm:"column:alu_imagem;type:longblob"`
}

func (Aluno) TableName() string {
	return "alu_alunos"
}
