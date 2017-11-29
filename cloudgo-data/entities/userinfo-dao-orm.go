package entities

// import (
// 	"fmt"
// )

func count() int {
	return 	len(UserInfoServiceOrmFindAll())
}

func UserInfoServiceOrmSave(u *UserInfo) error {
	// id, err := myOrm.Insert(u)
	// fmt.Println(id)
	_, err := myOrm.Insert(u)
	u.UID = int(count())
	checkErr(err)
	return err
}

func UserInfoServiceOrmFindAll() []UserInfo {
	everyone := make([]UserInfo, 0)
	err := myOrm.Find(&everyone)
	checkErr(err)
	return everyone
}

func UserInfoServiceOrmFindByID(id int) *UserInfo {
	var user UserInfo
	myOrm.Where("UID=?", id).Get(&user)
	return &user
}