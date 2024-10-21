package service

import (
	"arlog/core/gorm"
	log "arlog/core/logs"
	"arlog/model"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用 Bcrypt 算法生成密码哈希值
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords 比较输入的密码与哈希值是否匹配
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// AddUser 增加用户
func AddUser(user model.User) (model.User, int) {

	//密码加密重新赋值
	user.Password, _ = HashPassword(user.Password)
	//创建用户
	gorm.Db.Create(&user)

	log.Info("创建用户成功!", user)
	return user, 200

}

// DelUser 删除用户
func DelUser(user model.User) int {
	//查询用户是否存在
	gorm.Db.Select("id").Where("username = ?", user.Username).Find(&user)
	if user.Id > 0 {
		//删除用户
		gorm.Db.Delete(&user)
		log.Info("删除用户成功!", user)
		return 200
	}
	return 204
}

// QueUser 查询用户
func QueUser(user model.User) int {
	//查询用户是否存在
	gorm.Db.Select("id").Where("username = ?", user.Username).Find(&user)
	if user.Id > 0 {
		log.Info("查询到当前用户", user)
		return 202
	}

	return 200
}

// AllUser 查询所有用户
func AllUser(pageSize int, page int) ([]model.User, int, int64) {
	var users []model.User
	//查询返回所有用户,只返回默认字段

	var total int64

	//分页查询分类

	gorm.Db.Omit("password").Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&users)

	// 统计有多少条数据
	gorm.Db.Model(&users).Count(&total)

	log.Info(users)
	return users, 200, total
}

// UpUser 更新用户数据
func UpUser(user model.User) int {
	//密码加密
	ps, _ := HashPassword(user.Password)
	//更新用户数据
	db := gorm.Db.Model(&user).Updates(model.User{
		Id:       user.Id,
		Username: user.Username,
		Password: ps,
		AImg:     user.AImg,
		Phone:    user.Phone,
		Mail:     user.Mail,
		Gitee:    user.Gitee,
		Qq:       user.Qq,
	})
	//判断更新结果
	if db.RowsAffected == 0 {
		log.Info("更新用户失败")
		return 201
	}
	return 200
}
