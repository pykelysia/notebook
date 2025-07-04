package database

type DataModel struct {
	ID   uint   `gorm:"primaryKey;unique" json:"id"`
	Data string `json:"data"`
	Done bool   `json:"done"`
}

var DataNum uint = 1

func NewDataModel() *DataModel {
	return &DataModel{}
}

func (*DataModel) Get(id uint) (res DataModel, err error) {
	err = database.Model(&DataModel{}).First(&res, id).Error
	return
}

func (*DataModel) Create(item *DataModel) error {
	DataNum++
	return database.Model(&DataModel{}).Create(item).Error
}

func (*DataModel) Delete(id uint) error {
	return database.Delete(&DataModel{}, id).Error
}

func (*DataModel) Updata(item *DataModel) error {
	return database.Model(&DataModel{}).Where("id = ?", item.ID).Updates(item).Error
}

func (*DataModel) TableName() string {
	return "data"
}
