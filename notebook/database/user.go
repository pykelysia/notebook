package database

type UserModel struct {
	UID    uint   `gorm:"unique;primaryKey" json:"uid"`
	Number uint   `json:"num"`
	Code   []uint `gorm:"type:varint(255)[]" json:"code"`
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
	return database.Save(item).Error
}

func (*UserModel) TableName() string {
	return "user"
}
