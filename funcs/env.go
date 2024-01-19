package funcs

import (
	"fmt"
	"os"
)

// 环境变量 - 添加PATH, PATH不本应用运行目录下的子目录
func EnvAddPathFromRoot(path string) {
	origPath := os.Getenv("path")
	os.Setenv("path", fmt.Sprintf("%s\\%s;%s", Getwd(), path, origPath))
}

func Getwd() string {
	var _pwd_, _ = os.Getwd()
	pwd := _pwd_
	return pwd
}
