package entities

//
// ─── TO BE JSON RESPONSE ───────────────────────────────────────────────────────────
//

// UserInfoResponse -- GetUserByKeyAndID
type UserInfoResponse struct {
	Message  string `json:"message"`
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// SingleMessageResponse -- DeleteUserByKeyAndPassword/ChangeUserPassword/GetUserKey
type SingleMessageResponse struct {
	Message string `json:"message"`
}

// LoginResponse .
type LoginResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// used in UsersInfoResponse
type singleUserInfo struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// UsersInfoResponse -- ListUsersByKeyAndLimit
type UsersInfoResponse struct {
	Message            string           `json:"message"`
	SingleUserInfoList []singleUserInfo `json:"singleuserinfolist"`
}

// UserInfo .
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// ────────────────────────────────────────────────────────────────────────────────

//
// ─── MESSAGE ────────────────────────────────────────────────────────────────────
//

const (
	// EmptyInput . .
	EmptyInput = "empty input"
	// ServerError .
	ServerError = "server error"
	// DuplicateUsername .
	DuplicateUsername = "duplicate username"
	// CreateUserSuceed .
	CreateUserSuceed = "create user successfully"
	//FailCreateUser .
	FailCreateUser = "fail to create user"
	// EmptyUsernameOrPassword .
	EmptyUsernameOrPassword = "empty username or password"
	// EmptyPassword .
	EmptyPassword = "empty password"
	// IncorrectUsernameAndPassword .
	IncorrectUsernameAndPassword = "incorrect username or password"
	// IncorrectPassword .
	IncorrectPassword = "incorrect password"
	// LoginSucceed .
	LoginSucceed = "login successfully"
	// EmptyID .
	EmptyID = "empty id"
	// InvalidID .
	InvalidID = "invalid id"
	// ReLogin .
	ReLogin = "please re-login"
	// GetUserInfoSucceed .
	GetUserInfoSucceed = "get user info successfully"
	// NotExistedID .
	NotExistedID = "the id does not exist"
	// InvalidLimit .
	InvalidLimit = "invalid limit"
	// InvalidOffset .
	InvalidOffset = "invalid offset"
	// LogoutFail .
	LogoutFail = "log out fail"
	// UpdatePasswordSucceed .
	UpdatePasswordSucceed = "update password successfully"
	// EmptyNewPassword .
	EmptyNewPassword = "new password is empty"
	// NotMatchPassword .
	NotMatchPassword = "new password and confirmation do not match"
	// RequestDataError .
	RequestDataError = "request data error"
)

// ────────────────────────────────────────────────────────────────────────────────
