package mvc

import (
	"errors"
	"github.com/Lyo-Shur/gorm"
	"github.com/Lyo-Shur/gorm/info"
	"log"
)

// 定义Service接口
type Service interface {
	GetList(attr interface{}) (gorm.Table, error)
	GetCount(attr interface{}) (int64, error)
	GetModel(attr interface{}) (gorm.Table, error)
	Update(attr interface{}) (int64, error)
	Insert(attr interface{}) (int64, error)
	Delete(attr interface{}) (int64, error)
}

// service层结构体
type serviceImpl struct {
	dao DAO
}

// 获取service层
func GetService(db info.DataBase, tableName string) Service {
	impl := serviceImpl{}
	impl.dao = GetDAO(db, tableName)
	return &impl
}

// 查询列表方法
func (serviceImpl *serviceImpl) GetList(attr interface{}) (gorm.Table, error) {
	table, err := serviceImpl.dao.GetList(attr)
	if err != nil {
		log.Println(err)
		return table, errors.New("system error")
	}
	return table, nil
}

// 查询条数方法
func (serviceImpl *serviceImpl) GetCount(attr interface{}) (int64, error) {
	count, err := serviceImpl.dao.GetCount(attr)
	if err != nil {
		log.Println(err)
		return count, errors.New("system error")
	}
	return count, nil
}

// 查询实体方法
func (serviceImpl *serviceImpl) GetModel(attr interface{}) (gorm.Table, error) {
	table, err := serviceImpl.dao.GetModel(attr)
	if err != nil {
		log.Println(err)
		return table, errors.New("system error")
	}
	return table, nil
}

// 更新记录方法
func (serviceImpl *serviceImpl) Update(attr interface{}) (int64, error) {
	count, err := serviceImpl.dao.Update(attr)
	if err != nil {
		log.Println(err)
		return count, errors.New("system error")
	}
	return count, nil
}

// 添加记录方法
func (serviceImpl *serviceImpl) Insert(attr interface{}) (int64, error) {
	id, err := serviceImpl.dao.Insert(attr)
	if err != nil {
		log.Println(err)
		return id, errors.New("system error")
	}
	return id, nil
}

// 删除记录方法
func (serviceImpl *serviceImpl) Delete(attr interface{}) (int64, error) {
	count, err := serviceImpl.dao.Delete(attr)
	if err != nil {
		log.Println(err)
		return count, errors.New("system error")
	}
	return count, nil
}
