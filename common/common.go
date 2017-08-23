package common

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"time"
)

//返回提示信息
func GetMsg(codeNum int, msg string, data ...interface{}) map[string]interface{} {
	if data == nil {
		return map[string]interface{}{"code": codeNum, "message": msg}
	} else {
		return map[string]interface{}{"code": codeNum, "message": msg, "data": data[0]}
	}
}

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		nums = append(nums, num)
	}

	return nums
}

//md5加密
func Common_md5(md5_string string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(md5_string))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//获取表名
func Common_getTableName(table_name string) string {
	return beego.AppConfig.String("mysqlprefix") + table_name
}
