package entity

/*
{
	"det_id_chamada": 9,
	"det_id_aluno" : 1,
	"det_presente": true
}
*/

type DetalheDaChamada struct {
	ID           int     `json:"det_id" gorm:"column:det_id;primaryKey;autoIncrement"`
	DetIDChamada int     `json:"det_id_chamada" gorm:"column:det_id_chamada"`
	Chamada      Chamada `gorm:"foreignKey:DetIDChamada;references:ID"`
	DetRAAluno   int     `json:"det_ra_aluno" gorm:"column:det_ra_aluno"`
	Aluno        User    `gorm:"foreignKey:DetIDAluno;references:ID"`
	Presente     bool    `json:"det_presente" gorm:"column:det_presente"`
}

func (DetalheDaChamada) TableName() string {
	return "det_detalhesdaChamada"
}

func (DetalheDaChamada) GetPrimaryKey() string {
	return "det_id"
}
