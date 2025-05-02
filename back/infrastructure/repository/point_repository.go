package repository

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/infrastructure/model"
	"gorm.io/gorm"
)

type pointRepository struct {
	orm *gorm.DB
}

func NewPointRepository(orm *gorm.DB) repository.PointRepository {
	return &pointRepository{
		orm: orm,
	}
}

func (r *pointRepository) Create(pe *entity.PointEntity) (*entity.PointEntity, error) {
	const DEFAULT_POINT_AMOUNT = 12500
	pm := &model.PointModel{
		UserID:  pe.UserID,
		GroupID: pe.GroupID,
		Amount:  DEFAULT_POINT_AMOUNT,
	}

	if err := r.orm.Create(pm).Error; err != nil {
		return nil, err
	}

	return &entity.PointEntity{
		UserID:  pm.UserID,
		GroupID: pm.GroupID,
		Amount:  pm.Amount,
	}, nil
}

func (r *pointRepository) FindByUserAndGroup(pe *entity.PointEntity) (*entity.PointEntity, error) {
	var pm model.PointModel
	if err := r.orm.Where("user_id = ? AND group_id = ?", pe.UserID, pe.GroupID).First(&pm).Error; err != nil {
		return nil, err
	}

	return &entity.PointEntity{
		UserID:  pm.UserID,
		GroupID: pm.GroupID,
		Amount:  pm.Amount,
	}, nil
}

func (r *pointRepository) Restore(pe *entity.PointEntity) (*entity.PointEntity, error) {
	pm := &model.PointModel{
		UserID:  pe.UserID,
		GroupID: pe.GroupID,
	}

	if err := r.orm.Where("user_id = ? AND group_id = ?", pe.UserID, pe.GroupID).First(&pm).Error; err != nil {
		return nil, err
	}

	if pm.DeletedAt.Valid {
		if err := r.orm.Model(&pm).Update("deleted_at", nil).Error; err != nil {
			return nil, err
		}
	}

	return &entity.PointEntity{
		UserID:  pm.UserID,
		GroupID: pm.GroupID,
		Amount:  pm.Amount,
	}, nil
}
