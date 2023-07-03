package utils

import (
	"github.com/duke-git/lancet/v2/slice"
	"golang.org/x/crypto/bcrypt"
	"pixiu-panel/pkg/constant"
	"unicode/utf8"
)

// Password
// @description: 密码工具包
type Password interface {
	CheckWeakPassword(password string) bool   // 检查是否弱口令
	CheckPasswordRule(password string) bool   // 检查是否合规
	HashPassword(password *string)            // 加密密码
	ComparePassword(dbPass, pass string) bool // 校验密码
}

type password struct{}

// PasswordUtils
// @description: 密码工具包
// @return Password
func PasswordUtils() Password {
	return &password{}
}

// CheckWeakPassword
// @description: 检查是否弱口令
// @receiver password
// @param password
// @return bool
func (password) CheckWeakPassword(password string) bool {
	return slice.Contain(constant.WeakPasswords, password)
}

// CheckPasswordRule
// @description: TODO 检查密码规则是否健全
// @receiver password
// @param password
// @return bool
func (password) CheckPasswordRule(password string) bool {
	// 正则判断密码必须包含数字、字母大小写
	//reg := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,16}$`)
	////return regexp.MustCompile(`^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*]).{6,32}$`).MatchString(password)
	//matched, err := regexp.MatchString(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z]{8,32}$`, password)
	//if err != nil {
	//	log.Errorf("正则匹配错误: %v", err.Error())
	//	return false
	//}
	//return matched
	return utf8.RuneCountInString(password) >= 8 && utf8.RuneCountInString(password) <= 32
}

// HashPassword
// @description: 密码加密
// @receiver password
// @param pass
func (password) HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

// ComparePassword
// @description: 密码比较
// @receiver password
// @param dbPass
// @param pass
// @return bool
func (password) ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
