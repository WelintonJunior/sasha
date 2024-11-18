package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"example.com/apiSasha/database"
	"gorm.io/gorm"
)

type ChamadaService struct {
	commons.GenericService[entity.Chamada]
}

func NewChamadaService(service commons.GenericService[entity.Chamada]) *ChamadaService {
	return &ChamadaService{service}
}

func (s *ChamadaService) GetChamadaByID(chaID int) (entity.Chamada, error) {
	var resources entity.Chamada

	err := database.DB.Raw(`
		SELECT *
		FROM cha_chamada
		WHERE cha_id = ?;
	`, chaID).Scan(&resources).Error

	if err != nil {
		return entity.Chamada{}, err
	}

	return resources, nil
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
	result := database.DB.Exec(`
		UPDATE det_detalhesDaChamada AS det
		INNER JOIN cha_chamada AS cha
		ON det.det_id_chamada = cha.cha_id
		SET det.det_presente = true
		WHERE det.det_ra_aluno = ? AND cha.cha_inAndamento = true;
	`, aluRA)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "Nenhum registro encontrado", result.Error
		}
		return "Erro ao atualizar presença", result.Error
	}

	if result.RowsAffected == 0 {
		return "Nenhum registro atualizado", nil
	}

	return "Presença atualizada com sucesso", nil
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

func (s *ChamadaService) UpdateChamadaTrue(chaID int) (string, error) {
	err := database.DB.Exec(`
		UPDATE cha_chamada SET cha_inAndamento = true WHERE cha_id = ?;`, chaID).Error

	if err != nil {
		return "", err
	}

	return "Success", nil
}

func (s *ChamadaService) UpdateChamadaFalse(chaID int) (string, error) {
	err := database.DB.Exec(`
		UPDATE cha_chamada SET cha_inAndamento = false WHERE cha_id = ?;`, chaID).Error

	if err != nil {
		return "", err
	}

	return "Success", nil
}
