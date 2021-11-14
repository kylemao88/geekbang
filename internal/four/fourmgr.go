package fourmgr

import (
	"fmt"
	"git.code.oa.com/gongyi/yqz/pkg/activitymgr/config" // 引用了现负责项目的配置文件包。。呃呃。。。
)

// InitComponents init all fourmgr client
func InitComponents(path string) bool {
	// get config from file
	config.InitConfig(path)
	conf := config.GetConfig()

	if conf.PluginConf != nil {
		err := conf.PluginConf.Setup()
		if err != nil {
			return false
		}
	}

	////  学习wire用法   // 安装好wire工具后，在对应目录下执行
	u := InitializeUser()
	sentence := u.shout()
	fmt.Println("u shout :", sentence)

	// init mysql
	// ....

	// init redis
	// ....
	return true
}

/////////////////////////////

type Head struct {
	name  string
	score int
}

func (h *Head) say() string {
	return h.name
}

type Nick struct {
	name string
	len  int
}

func (n *Nick) speak() string {
	return n.name
}

type user struct {
	h *Head
	n *Nick
}

func (u *user) shout() string {
	return u.h.say() + u.n.speak()
}

func NewUser(h *Head, n *Nick) user {
	s := user{
		h: h,
		n: n,
	}
	return s
}

func NewNick() *Nick {
	b := &Nick{
		name: "/cccc",
		len:  4,
	}
	return b
}

func NewHead() *Head {
	a := &Head{
		name:  "wang xiao ya",
		score: 100,
	}
	return a
}
