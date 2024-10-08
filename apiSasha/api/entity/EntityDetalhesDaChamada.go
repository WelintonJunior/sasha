package entity

type DetalheDaChamada struct {
	ID           int     `gorm:"column:det_id;primaryKey;autoIncrement"`
	DetIDChamada int     `json:"det_id_chamada" gorm:"column:det_id_chamada"`
	Chamada      Chamada `gorm:"foreignKey:DetIDChamada;references:ID"`
	DetIDAluno   int     `json:"det_id_aluno" gorm:"column:det_id_aluno"`
	Aluno        Aluno   `gorm:"foreignKey:DetIDAluno;references:ID"`
	Presente     bool    `json:"det_presente" gorm:"column:det_presente"`
}

func (DetalheDaChamada) TableName() string {
	return "det_detalhesdaChamada"
}
