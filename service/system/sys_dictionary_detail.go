package system

import (
	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/model/system"
	"github.com/wangrui19970405/wu-shi-admin/server/model/system/request"
)

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: CreateSysDictionaryDetail
//@description: 创建字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

type DictionaryDetailService struct{}

func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.WUSHI_DB.Create(&sysDictionaryDetail).Error
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: DeleteSysDictionaryDetail
//@description: 删除字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.WUSHI_DB.Delete(&sysDictionaryDetail).Error
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: UpdateSysDictionaryDetail
//@description: 更新字典详情数据
//@param: sysDictionaryDetail *model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *system.SysDictionaryDetail) (err error) {
	err = global.WUSHI_DB.Save(sysDictionaryDetail).Error
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetSysDictionaryDetail
//@description: 根据id获取字典详情单条数据
//@param: id uint
//@return: sysDictionaryDetail system.SysDictionaryDetail, err error

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.WUSHI_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetSysDictionaryDetailInfoList
//@description: 分页获取字典详情列表
//@param: info request.SysDictionaryDetailSearch
//@return: list interface{}, total int64, err error

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.WUSHI_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}
