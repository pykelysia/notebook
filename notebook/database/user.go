package database

import (
	"database/sql/driver"
	"encoding/json"
)

type Code []string

type UserModel struct {
	UID    uint `gorm:"unique;primaryKey" json:"uid"`
	Number uint `json:"password"`
	Code   Code `json:"code"`
}

func (t *Code) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Code) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (*UserModel) Get(id uint) (res UserModel, err error) {
	err = database.Model(&UserModel{}).First(&res, id).Error
	return
}

func (*UserModel) Create(item *UserModel) error {
	return database.Model(&UserModel{}).Create(item).Error
}

func (*UserModel) Delete(id uint) error {
	return database.Delete(&UserModel{}, id).Error
}

func (*UserModel) Updata(item *UserModel) error {
	return database.Model(&UserModel{}).Where("uid = ?", item.UID).Updates(item).Error
}

func (*UserModel) TableName() string {
	return "user"
}
