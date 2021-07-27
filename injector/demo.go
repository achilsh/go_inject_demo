//+build wireinject
//The build tag makes sure the stub is not built in the final build.
//必须和包申明有一空行隔开，否则编译标签会被当做包声明的注释而不是编译标签; 条件编译是: +build xxx, 其中在前面加!表示否定的意思
//一个源文件可以有多个编译标签，多个编译标签之间是逻辑“与”的关系，一个编译标签可以包括由空格分割的多个标签，这些标签是逻辑“或”的关系
//参考：https://blog.csdn.net/phantom_111/article/details/79451635

package injector

import (
	"demo/ProviderDemo"
	"github.com/google/wire"
)

func InitializeIjk() (ProviderDemo.IJK, error) {
	wire.Build(ProviderDemo.SetDemoOne)
	return ProviderDemo.IJK{}, nil
}
