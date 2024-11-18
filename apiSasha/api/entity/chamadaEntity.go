package entity

/*
{
	"cha_dis_disciplina": 1,
	"cha_id_professor" : 1
}
*/

type Chamada struct {
	ID               int        `json:"cha_id" gorm:"column:cha_id;primaryKey;autoIncrement"`
	ChaDisDisciplina int        `json:"cha_dis_disciplina" gorm:"column:cha_dis_disciplina"`
	Disciplina       Disciplina `gorm:"foreignKey:ChaDisDisciplina;references:ID"`
	ChaIDProfessor   string     `json:"cha_id_professor" gorm:"column:cha_id_professor"`
	Professor        User       `gorm:"foreignKey:ChaIDProfessor;references:ID"`
	ChaInAndamento   bool       `json:"cha_inAndamento" gorm:"column:cha_inAndamento"`
}

type ListChamadaResponse struct {
	DetRaAluno  string `json:"det_ra_aluno"`
	Name        string `json:"name"`
	DetPresente bool   `json:"det_presente"`
}

func (Chamada) TableName() string {
	return "cha_chamada"
}

func (Chamada) GetPrimaryKey() string {
	return "cha_id"
}
