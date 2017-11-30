package entities

import "util"

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	tx := mydb.Begin()
	util.PanicIf(tx.Error)

	if tx.Save(u).Error != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	var rows []UserInfo
	mydb.Find(&rows)
	return rows

	// TODEL: @@binly:
	// dao := userInfoDao{mydb}
	// return dao.FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var uInfo UserInfo
	mydb.First(&uInfo, UserInfo{UID: id})
	return &uInfo

	// TODEL: @@binly:
	// dao := userInfoDao{mydb}
	// return dao.FindByID(id)
}
