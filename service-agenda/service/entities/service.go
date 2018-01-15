package entities

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/freakkid/service-agenda/service/tools"
)

// AgendaAtomicService -- a struct to operate service function
type AgendaAtomicService struct{}

// AgendaService -- an instance
var AgendaService = AgendaAtomicService{}

//
// ─── PROVIDE SERICES AND RETURN STATUS CODE AND JSON RESPONSE STRUCT ────────────
//

// CreateUser -- check if input is empty and username is duplicate
func (*AgendaAtomicService) CreateUser(body io.ReadCloser) (int, UserInfoResponse) {
	decoder := json.NewDecoder(body)
	defer body.Close()
	var userInfo UserInfo
	if err := decoder.Decode(&userInfo); err != nil {
		return http.StatusBadRequest, UserInfoResponse{Message: RequestDataError, ID: -1}
	}
	// ---- check input ----
	if userInfo.Username == "" || userInfo.Password == "" || userInfo.Email == "" || userInfo.Phone == "" {
		return http.StatusBadRequest, UserInfoResponse{Message: EmptyInput, ID: -1}
	}
	userInfo.Password = tools.MD5Encryption(userInfo.Password)
	dao := agendaDao{xormEngine}
	// ---- check username ----
	has, err := dao.ifUserExistByConditions(&User{UserName: userInfo.Username})
	if err != nil { // server error
		return http.StatusInternalServerError, UserInfoResponse{Message: ServerError, ID: -1}
	}
	if has { // username exist -- duplicate username
		return http.StatusBadRequest, UserInfoResponse{Message: DuplicateUsername, ID: -1}
	}
	// ---- create user ----
	result, user := dao.createUser(&User{SessionID: tools.GenenrateSessionID(),
		UserName: userInfo.Username, Password: userInfo.Password, Email: userInfo.Email, Phone: userInfo.Phone})
	if result && user != nil { // create user successfully
		return http.StatusCreated, UserInfoResponse{CreateUserSuceed, user.ID, user.UserName, user.Email, user.Phone}
	}
	return http.StatusBadRequest, UserInfoResponse{Message: FailCreateUser, ID: -1}
}

// LoginAndGetSessionID --- check if user exists and generate new sessionID
// if user no exists or occur error, return empty sessionID
// if login success, return sessionID
func (*AgendaAtomicService) LoginAndGetSessionID(
	username string, password string) (string, int, LoginResponse) {
	// ---- check username and password ----
	if username == "" || password == "" { // check if empty username and password
		return "", http.StatusBadRequest, LoginResponse{EmptyUsernameOrPassword, -1}
	}
	password = tools.MD5Encryption(password)
	dao := agendaDao{xormEngine}
	user := User{UserName: username, Password: password}
	has, err := dao.ifUserExistByConditions(&user) // check if exist
	if err != nil {                                // server error
		return "", http.StatusInternalServerError, LoginResponse{ServerError, -1}
	}
	if !has { // user not exist
		return "", http.StatusUnauthorized, LoginResponse{IncorrectUsernameAndPassword, -1}
	}
	// ---- get new sessionID ----
	var sessionID = tools.GenenrateSessionID() // generate new sessionID
	for {                                      // make sure sessionID unique
		if has, err = dao.ifUserExistByConditions(&User{SessionID: sessionID}); err == nil && !has {
			break
		}
		sessionID = tools.GenenrateSessionID()
	}
	affected, _ := dao.updateUser(&User{SessionID: sessionID}, &User{UserName: username, Password: password})
	if affected == 0 { // user not exist
		return "", http.StatusUnauthorized, LoginResponse{IncorrectUsernameAndPassword, -1}
	}
	return sessionID, http.StatusOK, LoginResponse{LoginSucceed, user.ID}
}

// GetUserInfoByID --- check if sessionID is valid and id exsits and belong to the same user
// convert string id to int id, if occur error return empty user and error
// if valid key and exist id, return User struct
func (*AgendaAtomicService) GetUserInfoByID(sessionID string, stringID string) (int, UserInfoResponse) {
	var (
		id   int
		err  error
		has  bool
		user *User
	)
	dao := agendaDao{xormEngine}
	// ---- check sessionID ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: sessionID})
	if err != nil { // server error
		return http.StatusInternalServerError, UserInfoResponse{Message: ServerError, ID: -1}
	}
	if !has { // invalid sessionID
		return http.StatusUnauthorized, UserInfoResponse{Message: ReLogin, ID: -1}
	}
	// ---- check id ----
	if stringID == "" { // empty id
		return http.StatusBadRequest, UserInfoResponse{Message: EmptyID, ID: -1}
	}
	id, err = strconv.Atoi(stringID)
	if err != nil || id <= 0 { // invalid id
		return http.StatusBadRequest, UserInfoResponse{Message: InvalidID, ID: id}
	}
	// ---- find user by id ----
	has, user = dao.findUserByConditions(&User{ID: id})
	if has && user != nil { // user not exist
		return http.StatusOK,
			UserInfoResponse{GetUserInfoSucceed, user.ID, user.UserName, user.Email, user.Phone}
	}
	return http.StatusNotFound, UserInfoResponse{Message: NotExistedID, ID: id}
}

// DeleteUserByPassword --- check if sessionID valid and password correct
func (*AgendaAtomicService) DeleteUserByPassword(
	sessionID string, stringID string, password string) (int, SingleMessageResponse) {
	var (
		err      error
		id       int
		has      bool
		affected int64
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: sessionID})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{Message: ServerError}
	}
	if !has { // invalid sessionID
		return http.StatusUnauthorized, SingleMessageResponse{Message: ReLogin}
	}
	// ---- check id ----
	id, err = strconv.Atoi(stringID)
	if err != nil || id <= 0 {
		return http.StatusBadRequest, SingleMessageResponse{Message: InvalidID}
	}
	// ---- check password ----
	if password == "" { // empty input
		return http.StatusBadRequest, SingleMessageResponse{EmptyPassword}
	}
	affected, err = dao.deleteUserBySessionIDAndIDAndPassword(sessionID, id, tools.MD5Encryption(password))
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{Message: ServerError}
	}
	if affected == 0 { // delete user fail
		return http.StatusUnauthorized, SingleMessageResponse{Message: IncorrectPassword}
	}
	// delete successfully
	return http.StatusNoContent, SingleMessageResponse{}
}

// ListUsersByLimit --- check sessionID is valid or not
// if limit is invalid, default set to 5
// if offset is invalid, default set to 0
func (*AgendaAtomicService) ListUsersByLimit(
	sessionID string, stringLimit string, stringOffset string) (int, UsersInfoResponse) {
	var (
		limit  int
		offset int
		has    bool
		err    error
		users  []User
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: sessionID})
	if err != nil { // server error
		return http.StatusInternalServerError, UsersInfoResponse{ServerError, []singleUserInfo{}}
	}
	if !has { // if sessionID not exist -- invalid sessionID
		return http.StatusUnauthorized, UsersInfoResponse{ReLogin, []singleUserInfo{}}
	}
	// ---- check limit ----
	if stringLimit == "" {
		limit = 5
	} else {
		limit, err = strconv.Atoi(stringLimit)
		if err != nil || limit <= 0 { // invalid limit
			return http.StatusBadRequest, UsersInfoResponse{InvalidLimit, []singleUserInfo{}}
		}
	}
	// ---- check offset ----
	if stringOffset == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(stringOffset)
		if err != nil || offset < 0 { // invalid limit
			return http.StatusBadRequest, UsersInfoResponse{InvalidOffset, []singleUserInfo{}}
		}
	}
	// ---- get limit, offset users ----
	users, err = dao.getLimitUsers(limit, offset)
	if err != nil { // server error
		return http.StatusInternalServerError, UsersInfoResponse{ServerError, []singleUserInfo{}}
	}
	singleUserInfoList := make([]singleUserInfo, 0, 0)
	for _, userInfo := range users {
		singleUserInfoList = append(singleUserInfoList,
			singleUserInfo{userInfo.ID, userInfo.UserName, userInfo.Email, userInfo.Phone})
	}
	return http.StatusOK, UsersInfoResponse{GetUserInfoSucceed, singleUserInfoList}
}

// ChangeUserPassword -- check if sessionID valid and password correct
// check if new password valid and match confirmation
func (*AgendaAtomicService) ChangeUserPassword(sessionID string, stringID string, password string, newPassword string, confirmation string) (int, SingleMessageResponse) {
	var (
		has      bool
		err      error
		id       int
		affected int64
	)
	dao := agendaDao{xormEngine}
	// ---- check sessionID ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: sessionID})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{ServerError}
	}
	if !has { // if sessionID not exist -- invalid sessionID
		return http.StatusUnauthorized, SingleMessageResponse{ReLogin}
	}
	// ---- check id ----
	id, err = strconv.Atoi(stringID)
	if err != nil || id <= 0 {
		return http.StatusBadRequest, SingleMessageResponse{Message: InvalidID}
	}
	// ---- check old password ----
	password = tools.MD5Encryption(password)
	has, err = dao.ifUserExistByConditions(&User{SessionID: sessionID, ID: id, Password: password})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{ServerError}
	}
	if !has { // password incorrect
		return http.StatusUnauthorized, SingleMessageResponse{IncorrectPassword}
	}
	// ---- check new password ----
	if newPassword == "" {
		return http.StatusBadRequest, SingleMessageResponse{EmptyNewPassword}
	}
	if newPassword != confirmation {
		return http.StatusBadRequest, SingleMessageResponse{NotMatchPassword}
	}
	// ---- update new password ----
	newPassword = tools.MD5Encryption(newPassword)
	affected, _ = dao.updateUser(&User{Password: newPassword},
		&User{SessionID: sessionID, ID: id, Password: password})
	if affected == 0 { // user not exist
		return http.StatusUnauthorized, SingleMessageResponse{IncorrectPassword}
	}
	return http.StatusOK, SingleMessageResponse{UpdatePasswordSucceed}
}

// LogoutAndDeleteSessionID -- logout and update sessionid
func (*AgendaAtomicService) LogoutAndDeleteSessionID(
	sessionID string, stringID string) (int, SingleMessageResponse) {
	// ---- check sessionID ----
	dao := agendaDao{xormEngine}
	has, err := dao.ifUserExistByConditions(&User{SessionID: sessionID})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{ServerError}
	}
	if !has { // user not exist
		return http.StatusUnauthorized, SingleMessageResponse{LogoutFail}
	}
	// ---- check id ----
	var id int
	id, err = strconv.Atoi(stringID)
	if err != nil || id <= 0 {
		return http.StatusBadRequest, SingleMessageResponse{Message: InvalidID}
	}
	// ---- get new sessionID ----
	var newSessionID = tools.GenenrateSessionID() // generate new sessionID
	for {                                         // make sure new sessionID unique
		if has, err = dao.ifUserExistByConditions(&User{SessionID: newSessionID}); err == nil && !has {
			break
		}
		newSessionID = tools.GenenrateSessionID()
	}
	affected, _ := dao.updateUser(&User{SessionID: newSessionID}, &User{SessionID: sessionID, ID: id}) // replace old sessionID
	if affected == 0 {                                                                                 // user not exist
		return http.StatusUnauthorized, SingleMessageResponse{LogoutFail}
	}
	return http.StatusNoContent, SingleMessageResponse{}
}
