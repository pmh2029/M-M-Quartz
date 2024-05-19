package entity

var UsersTableName = "tbl_users"

type User struct {
	ID       int    `gorm:"column:id;primaryKey;type:bigint;not null;autoIncrement" mapstructure:"id"`
	Email    string `gorm:"column:email;type:varchar(512);not null;unique" mapstructure:"email"`
	Username string `gorm:"column:username;type:varchar(50);not null;unique" mapstructure:"username"`
	Password string `gorm:"column:password;type:text;not null" mapstructure:"password"`
	BaseEntity
}

func (i *User) TableName() string {
	return UsersTableName
}
