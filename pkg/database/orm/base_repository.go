package orm

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository[E any] struct {
	*gorm.DB
}

func (r *BaseRepository[E]) Create(ctx context.Context, entity E) (*E, error) {
	if err := r.DB.WithContext(ctx).Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// FindAll Retrieve all data of repository
func (r *BaseRepository[E]) FindAll(ctx context.Context) ([]*E, error) {
	entities := make([]*E, 0)
	if err := r.DB.WithContext(ctx).Find(&entities).Error; err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return entities, nil
}

// FindFirst Find first data by field and value
func (r *BaseRepository[E]) FindFirst(ctx context.Context, field string, value any) (*E, error) {
	var entity E
	if err := r.DB.WithContext(ctx).First(&entity, fmt.Sprintf("%s = ?", field), value).Error; err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &entity, nil
}

// FindByField Find data by field and value
func (r *BaseRepository[E]) FindByField(ctx context.Context, field string, value any) ([]*E, error) {
	entities := make([]*E, 0)
	if err := r.DB.WithContext(ctx).Find(&entities, fmt.Sprintf("%s = ?", field), value).Error; err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return entities, nil
}

// FindWhereIn Find data by multiple values in one field
func (r *BaseRepository[E]) FindWhereIn(ctx context.Context, field string, values any) ([]*E, error) {
	entities := make([]*E, 0)
	query := fmt.Sprintf("%s IN ?", field)
	if err := r.DB.WithContext(ctx).Where(query, values).Find(&entities).Error; err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return entities, nil
}

func (r *BaseRepository[E]) Updates(ctx context.Context, id any, values map[string]any) (*E, error) {
	var entity E
	if err := r.DB.WithContext(ctx).Model(&entity).Clauses(clause.Returning{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete a entity in repository by id
func (r *BaseRepository[E]) Delete(ctx context.Context, id any) error {
	var entity E
	err := r.DB.WithContext(ctx).Delete(&entity, id).Error
	if err != nil {
		slog.Error(err.Error())
	}

	return err
}
