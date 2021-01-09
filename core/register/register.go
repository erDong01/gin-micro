package register

import (
	"github.com/erDong01/micro-kit/cache/impl"
	config2 "github.com/erDong01/micro-kit/config"
	"github.com/erDong01/micro-kit/core"
	"github.com/erDong01/micro-kit/db/mysql"
	"github.com/erDong01/micro-kit/db/redis/drive"
)

// 全局初始化
func GlobalInit() *Register {
	return new(Register).ConfigRegister().
		RedisRegister().
		DbRegister().
		FacadeCacheRegister().
		SetPort(5001)
}

type Register struct {
}

//注册 配置
func (r *Register) ConfigRegister() *Register {
	config2.Init("config")
	return r
}

//注册 orm
func (r *Register) DbRegister() *Register {
	core.New().Db = mysql.New()
	return r
}

//注册 缓存中心
func (r *Register) FacadeCacheRegister() *Register {
	core.New().Cache = impl.New()
	return r
}

//注册 缓存中心
func (r *Register) RedisRegister() *Register {
	core.New().Redis = drive.New()
	return r
}

//设置名称
func (r *Register) SetName(name string) *Register {
	core.New().Info.Name = name
	return r
}

// 设置端口号
func (r *Register) SetPort(port int) *Register {
	core.SetPort(port)
	return r
}