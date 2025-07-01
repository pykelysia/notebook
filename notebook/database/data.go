package database

type DataModel struct {
	UID  uint   `gorm:"primaryKey;unique" json:"id"`
	Data string `json:"data"`
	Done bool   `json:"done"`
}

func NewDataModel() *DataModel {
	return &DataModel{}
}

func (*DataModel) Get(id uint) (res DataModel, err error) {
	err = database.Model(&DataModel{}).First(&res, id).Error
	return
}

func (*DataModel) Create(item *DataModel) error {
	return database.Model(&DataModel{}).Create(item).Error
}

func (*DataModel) Delete(id uint) error {
	return database.Delete(&DataModel{}, id).Error
}

func (*DataModel) Updata(item *DataModel) error {
	return database.Model(&DataModel{}).Where("uid = ?", item.UID).Updates(item).Error
}

func (*DataModel) TableName() string {
	return "data"
}
