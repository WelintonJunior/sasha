package entity

/*
{
	"alu_aluno": "aluno"
}
*/

type User struct {
	ID       string `json:"usu_ra" gorm:"column:usu_ra;primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"column:name;size:100;not null"`
	Password string `json:"password" gorm:"column:password;size:100;not null"`
	Type     int    `json:"type" gorm:"column:type;size:1;not null"`
}

func (User) TableName() string {
	return "usu_users"
}

func (User) GetPrimaryKey() string {
	return "usu_ra"
}
