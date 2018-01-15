package entities

// User is an entity to save user info with username is primary key.
type User struct {
	ID        int    `xorm:"pk autoincr"`
	SessionID string `xorm:"varchar(255) unique"`
	UserName  string `xorm:"varchar(255) notnull unique"`
	Password  string `xorm:"varchar(255) notnull"`
	Email     string `xorm:"varchar(255) notnull"`
	Phone     string `xorm:"varchar(255) notnull"`
}
