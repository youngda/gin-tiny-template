package user

import (
	"api/modle"
	"api/pkg/errno"
	. "api/handler"
	"github.com/gin-gonic/gin"
)

// Get user info
func GetUserInfo(c *gin.Context) {

	var r GetUserInfoRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u,err := modle.GetUser(r.Email,r.Password)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}


	rsp := GetUserInfoResponse{
		Uid: u.Uid,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
