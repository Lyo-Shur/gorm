package mvc

import (
	"github.com/Lyo-Shur/gorm"
	"github.com/Lyo-Shur/gorm/info"
)

// 定义DAO接口
type DAO interface {
	GetList(attr interface{}) (gorm.Table, error)
	GetCount(attr interface{}) (int64, error)
	GetModel(attr interface{}) (gorm.Table, error)
	Update(attr interface{}) (int64, error)
	Insert(attr interface{}) (int64, error)
	Delete(attr interface{}) (int64, error)
}

// dao层结构体
type daoImpl struct {
	tableManager *gorm.TableManager
}

// 获取dao层
func GetDAO(db info.DataBase, tableName string) DAO {
	daoImpl := daoImpl{}
	xml := GetMapperXML(db, tableName)
	daoImpl.tableManager = gorm.GetTableManager(gorm.XmlConfig(xml))
	daoImpl.tableManager.Engine.Use(db.Alias)
	return &daoImpl
}

// 查询列表方法
func (daoImpl *daoImpl) GetList(attr interface{}) (gorm.Table, error) {
	return daoImpl.tableManager.Query("GetList", attr)
}

// 查询条数方法
func (daoImpl *daoImpl) GetCount(attr interface{}) (int64, error) {
	return daoImpl.tableManager.Count("GetCount", attr)
}

// 查询实体方法
func (daoImpl *daoImpl) GetModel(attr interface{}) (gorm.Table, error) {
	return daoImpl.tableManager.Query("GetModel", attr)
}

// 更新记录方法
func (daoImpl *daoImpl) Update(attr interface{}) (int64, error) {
	_, num, err := daoImpl.tableManager.Exec("Update", attr)
	return num, err
}

// 添加记录方法
func (daoImpl *daoImpl) Insert(attr interface{}) (int64, error) {
	id, _, err := daoImpl.tableManager.Exec("Insert", attr)
	return id, err
}

// 删除记录方法
func (daoImpl *daoImpl) Delete(attr interface{}) (int64, error) {
	_, num, err := daoImpl.tableManager.Exec("Delete", attr)
	return num, err
}
