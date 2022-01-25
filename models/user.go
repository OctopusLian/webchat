package models

import (
	"encoding/gob"
	"fmt"
	"time"
	"webchat/utils"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id              int       `orm:"pk;auto" json:"id"`
	Email           string    `orm:"unique;size(32)" json:"email"`
	Code            string    `orm:"size(8)" json:"-"`
	CodeCreatedTime time.Time `orm:"null;type(datetime)" json:"-"`
	CreatedTime     time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"`
	LastLoginedTime time.Time `orm:"null;type(datetime)" json:"last_logined_time"`
}

func (user User) Equal(o User) bool {
	return user.Id == o.Id
}

func Authenticate(email, code string) (*User, error) {
	o := orm.NewOrm()
	user := &User{Email: email}

	// 验证码不为空，验证码生成时间在5分钟之内
	if err := o.Read(user, "Email"); err == nil && user.Code != "" && user.Code == code && time.Now().Sub(user.CodeCreatedTime).Minutes() <= 5 {
		user.Code = ""
		user.LastLoginedTime = time.Now()
		o.Update(user)
		return user, nil
	}
	return nil, fmt.Errorf("email or code error")
}

func CreateLoginCode(email string) error {
	o := orm.NewOrm()
	user := &User{Email: email}
	code := utils.RString(6)
	created, _, err := o.ReadOrCreate(user, "Email")

	if err == nil {
		// 第一次创建或验证码生成时间超过1分钟
		fmt.Println(time.Now(), user.CodeCreatedTime, time.Now().Sub(user.CodeCreatedTime).Minutes())
		if created || time.Now().Sub(user.CodeCreatedTime).Minutes() >= 1 {
			user.Code = code
			user.CodeCreatedTime = time.Now()
			o.Update(user)
			return utils.SendEmail(email, "登陆验证码", code)
		}
		return nil
	}
	return err
}

func AllUser() []User {
	var users []User
	o := orm.NewOrm()
	o.QueryTable("user").All(&users)
	return users
}

func init() {
	gob.Register(User{})
	orm.RegisterModel(new(User))
}
