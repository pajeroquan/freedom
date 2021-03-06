// Code generated by 'freedom new-crud'
package object

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TestUsers struct {
	changes   map[string]interface{}
	Id        int       `gorm:"primary_key;column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	UserName  string    `gorm:"column:user_name"`
	Password  string    `gorm:"column:password"`
	Age       int       `gorm:"column:age"`
	Status    int       `gorm:"column:status"`
}

func (obj *TestUsers) TableName() string {
	return "test_users"
}

// TakeChanges .
func (obj *TestUsers) TakeChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// SetCreatedAt .
func (obj *TestUsers) SetCreatedAt(createdAt time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.CreatedAt = createdAt
	obj.changes["created_at"] = createdAt
}

// SetUpdatedAt .
func (obj *TestUsers) SetUpdatedAt(updatedAt time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.UpdatedAt = updatedAt
	obj.changes["updated_at"] = updatedAt
}

// SetDeletedAt .
func (obj *TestUsers) SetDeletedAt(deletedAt time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.DeletedAt = deletedAt
	obj.changes["deleted_at"] = deletedAt
}

// SetUserName .
func (obj *TestUsers) SetUserName(userName string) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.UserName = userName
	obj.changes["user_name"] = userName
}

// SetPassword .
func (obj *TestUsers) SetPassword(password string) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Password = password
	obj.changes["password"] = password
}

// SetAge .
func (obj *TestUsers) SetAge(age int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Age = age
	obj.changes["age"] = age
}

// SetStatus .
func (obj *TestUsers) SetStatus(status int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Status = status
	obj.changes["status"] = status
}

// AddAge .
func (obj *TestUsers) AddAge(age int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Age += age
	obj.changes["age"] = gorm.Expr("age + ?", age)
}

// AddStatus .
func (obj *TestUsers) AddStatus(status int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Status += status
	obj.changes["status"] = gorm.Expr("status + ?", status)
}
