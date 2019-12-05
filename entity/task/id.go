package task

import "github.com/google/uuid"

// ID タスクID
type ID struct{ value uuid.UUID }

// NewID はタスクIDを生成する
func NewID() ID {
	return ID{value: uuid.New()}
}

// ToID はstringをタスクIDに変換する
func ToID(s string) (ID, error) {
	uuid, err := uuid.Parse(s)
	// FIXME: errのラップはしなくても大丈夫か検討
	return ID{value: uuid}, err
}

// String はタスクIDをstringで返す
func (id ID) String() string {
	return uuid.UUID(id.value).String()
}
