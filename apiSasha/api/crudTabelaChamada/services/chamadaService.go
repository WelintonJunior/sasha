package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"example.com/apiSasha/database"
)

type ChamadaService struct {
	commons.GenericService[entity.Chamada]
}

func NewChamadaService(service commons.GenericService[entity.Chamada]) *ChamadaService {
	return &ChamadaService{service}
}

func (s *ChamadaService) GetAllDetChamadaFromChamadaId(chaID int) ([]entity.ListChamadaResponse, error) {
	var resources []entity.ListChamadaResponse

	err := database.DB.Raw(`
		SELECT d.det_ra_aluno, u.name, d.det_presente 
		FROM det_detalhesDaChamada AS d 
		INNER JOIN cha_chamada AS c ON c.cha_id = d.det_id_chamada 
		INNER JOIN usu_users AS u ON u.usu_ra = d.det_ra_aluno 
		WHERE c.cha_id = ?;
	`, chaID).Scan(&resources).Error

	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (s *ChamadaService) GetTokenFromChamadaId(chaID int) (string, error) {
	var resource string

	err := database.DB.Raw(`
		SELECT cha_token
		FROM cha_chamada
		WHERE cha_id = ?;
	`, chaID).Scan(&resource).Error

	if err != nil {
		return "nil", err
	}

	return resource, nil
}

func (s *ChamadaService) UpdatePresencaAluno(aluRA int) (string, error) {
	err := database.DB.Exec(`
		UPDATE det_detalhesDaChamada SET det_presente = true WHERE det_ra_aluno = ?;
	`, aluRA).Error

	if err != nil {
		return "", err
	}

	return "Success", nil
}

func (s *ChamadaService) UpdateTokenInChamada(token string, chaID int) (string, error) {
	err := database.DB.Exec(`
		UPDATE cha_chamada SET cha_token = ? WHERE cha_id = ?;
	`, token, chaID).Error

	if err != nil {
		return "", err
	}

	return "Success", nil
}
