package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
)

type todoModel struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Completed   bool      `gorm:"type:boolean;not null;default:false"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamptz;not null;autoUpdateTime"`
}

func (todoModel) TableName() string {
	return "todos"
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) List(ctx context.Context) ([]todo.Todo, error) {
	var models []todoModel
	if err := r.db.WithContext(ctx).Order("created_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	items := make([]todo.Todo, 0, len(models))
	for _, model := range models {
		items = append(items, mapToDomain(model))
	}

	return items, nil
}

func (r *todoRepository) Create(ctx context.Context, item todo.Todo) (todo.Todo, error) {
	model := mapToModel(item)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return todo.Todo{}, err
	}

	return mapToDomain(model), nil
}

func mapToDomain(model todoModel) todo.Todo {
	return todo.Todo{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Completed:   model.Completed,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func mapToModel(entity todo.Todo) todoModel {
	model := todoModel{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Completed:   entity.Completed,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}

	return model
}
