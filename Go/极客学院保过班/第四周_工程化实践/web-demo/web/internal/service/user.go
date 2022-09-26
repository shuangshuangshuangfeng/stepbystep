package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web-demo/web/internal/dao"
	. "web-demo/web/internal/handler"
	"web-demo/web/internal/pkg/errno"
)

func AddUser(c *gin.Context)  {
	var r dao.User
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := dao.User{
		UserName: r.UserName,
		Password: r.Password,
	}
	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Insert the user to the database.
	if _,err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	// Show the user information.
	SendResponse(c, nil, u)
}

// SelectUser 查询用户
func SelectUser(c *gin.Context)  {
	name := c.Query("user_name")
	if name == ""{
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	var  user dao.User
	if err := user.SelectUserByName(name);nil != err {
		fmt.Println(err)
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// Validate the data.
	if err := user.Validate(); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}

func SayHiUser(c *gin.Context){
	println("hello!")
}
