package modle

import (
	"api/logger"
	"go.uber.org/zap"
)

// User represents a registered user.
type UserModel struct {
	Uid       int64  `json:"uid" db:"uid"`
	Email     string `json:"email" db:"email"`
	Passwd    string `json:"passwd" db:"passwd"`
}


// 示例：通过邮箱和密码查询id
func GetUser(email string,passwd string) (*UserModel, error) {
	logger.Log.Info("get info:",zap.String("email:",email),zap.String("passwd:",passwd))
	sqlStr := `select uid from user where email = ? and passwd = ?`
	u := &UserModel{}
	err := DB.Get(&u.Uid, sqlStr, email,passwd)
	if err != nil {
		logger.Log.Error("get failed, err:",zap.Error(err))
		return u,err
	}
	return u, err
}

