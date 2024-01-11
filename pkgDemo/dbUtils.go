package pkgDemo

// 演示go的模块导入
import "fmt"

func GetDbConn() { // 首字母大写是公共方法
	fmt.Println("执行GetDbConn函数连接数据库...")
}
