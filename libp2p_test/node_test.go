package libp2p_test

import (
	"fmt"
	"github.com/libp2p/go-libp2p"
	"testing"
)

func TestNodeEmptyArgsNew(t *testing.T) {

	// 建立一个p2p节点，默认配置
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	// 打印节点的所有地址
	fmt.Println("Listen addresses:", node.Addrs())

	// 关闭节点，然后退出
	if err := node.Close(); err != nil {
		panic(err)
	}
}
