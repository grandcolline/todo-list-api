package task

import (
	"github.com/google/uuid"
	"github.com/grandcolline/todo-list-api/util/errors"
)

// ID タスクID
type ID struct{ value uuid.UUID }

// NewID はタスクIDを生成する
func NewID() ID {
	return ID{value: uuid.New()}
}

// ToID はstringをタスクIDに変換する
func ToID(s string) (ID, error) {
	uuid, err := uuid.Parse(s)
	if err != nil {
		return ID{}, errors.NewFromFmt(errors.Conv, "task.id", err.Error())
	}
	return ID{value: uuid}, nil
}

// String はタスクIDをstringで返す
func (id *ID) String() string {
	return uuid.UUID(id.value).String()
}
