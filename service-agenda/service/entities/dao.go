package entities

import "github.com/go-xorm/xorm"

type agendaDao struct {
	*xorm.Engine
}

// create user into database
func (dao *agendaDao) createUser(user *User) (bool, *User) {
	affected, _ := dao.Insert(user)
	if affected == 1 {
		return true, user
	}
	return false, nil
}

// uptate userinfo by condition is selectedUser
func (dao *agendaDao) updateUser(user *User, selectedUser *User) (int64, error) {
	return dao.Update(user, selectedUser)
}

// find user
func (dao *agendaDao) ifUserExistByConditions(user *User) (bool, error) {
	return dao.Get(user)
}

// get user
func (dao *agendaDao) findUserByConditions(user *User) (bool, *User) {
	has, err := dao.ifUserExistByConditions(user)
	if has && err == nil {
		return has, user
	}
	return has, nil
}

// get all users info
func (dao *agendaDao) getLimitUsers(limitNumber int, offsetNumber int) ([]User, error) {
	if limitNumber <= 0 {
		limitNumber = 5
	}
	if offsetNumber < 0 {
		offsetNumber = 0
	}
	var userList = make([]User, 0, 0)
	err := dao.Limit(limitNumber, offsetNumber).Find(&userList)
	return userList, err
}

// count all users
func (dao *agendaDao) countAllUsers() (int64, error) {
	return dao.Count(new(User))
}

// delete user by sessionID and password
func (dao *agendaDao) deleteUserBySessionIDAndIDAndPassword(
	sessionID string, id int, password string) (int64, error) {
	return dao.Delete(&User{SessionID: sessionID, ID: id, Password: password})
}
