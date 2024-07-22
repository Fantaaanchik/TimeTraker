package repositories

import "time_traker/models"

func (r *UserRepository) GetAllUsers(filter string, page, pageSize int) ([]models.User, error) {
	var users []models.User
	query := r.db.Offset((page - 1) * pageSize).Limit(pageSize)
	if filter != "" {
		query = query.Where("name LIKE ? OR surname LIKE ? OR address LIKE ?", "%"+filter+"%", "%"+filter+"%", "%"+filter+"%")
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
