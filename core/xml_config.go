/*
xml_config是解析sql配置xml文件而生成的对象
可以理解为xml文件对应的实体
解析xml得到的CPS（映射数组）会交给sql_mapper处理
解析xml得到的KVS（SQL语句数组）会交给sql_holder处理
xml_config只是为了解析xml存在，自身不参与功能
*/

package core

import (
	"encoding/xml"
	"log"
)

// 定义xml根元素
// CPS Column和Parameter的映射关系 数组
// KVS SQL的键和SQL的值对应关系 数组
type XmlConfig struct {
	// 映射数组
	CPS []cp `xml:"mapper"`
	// 键值数组
	KVS []kv `xml:"sql"`
}

// cp Column和Parameter的映射关系
type cp struct {
	// 数据库列名
	Column string `xml:"column,attr"`
	// 结构体参数名
	Parameter string `xml:"parameter,attr"`
}

// kv SQL的键和SQL的值对应关系
type kv struct {
	// SQL的键，相当于sql的别名
	Key string `xml:"key"`
	// SQL的值，具体的SQL
	Value string `xml:"value"`
}

// 初始化配置
func GetXmlConfig(str string) XmlConfig {
	// 保存内容结构体
	x := XmlConfig{}
	// 读取配置文件并缓存到结构体
	err := xml.Unmarshal([]byte(str), &x)
	if err != nil {
		log.Fatal(err)
	}
	return x
}
