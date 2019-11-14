package collection

import (
	"time"

	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
)

// TaskCollection タスクコレクションのデータ構造
type TaskCollection struct {
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	Status      string    `firestore:"status"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

// CollectionName コレクション名を返します
func (tc *TaskCollection) CollectionName() string {
	return "tasks"
}

// FromEntity はエンティティからコレクションに内容を入れます
func (tc *TaskCollection) FromEntity(t *entity.Task) string {
	tc.Name = t.Name.String()
	tc.Description = t.Description.String()
	tc.Status = t.Status.String()
	tc.CreatedAt = t.CreatedAt
	tc.UpdatedAt = t.UpdatedAt
	return t.ID.String()
}

// ToEntity はコレクションをエンティティに変換します
func (tc *TaskCollection) ToEntity(id string) (*entity.Task, error) {
	entID, err := task.ToID(id)
	name, err := task.ToName(tc.Name)
	des, err := task.ToDescription(tc.Description)
	status, err := task.ToStatus(tc.Status)
	if err != nil {
		// FIXME
		return nil, err
	}

	return &entity.Task{
		ID:          entID,
		Name:        name,
		Description: des,
		Status:      status,
		CreatedAt:   tc.CreatedAt,
		UpdatedAt:   tc.UpdatedAt,
	}, nil
}
