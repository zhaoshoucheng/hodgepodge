package trie

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestNode_InsertChild(t *testing.T) {
	node := &node{}
	path := "/aaa/:bbb/ccc"
	t.Log(countParams(path))
	node.InsertChild(countParams(path),path, path, HandlersChain{func(*context.Context){}})
	print(node,"root")
}

func TestNode_AddRoute(t *testing.T) {
	node := &node{}
	//path1 := "/aaa/eee/ttt"
	path2 := "/aaa/:bbb"
	path3 := "/aaa/:bbb/ccc/:ddd"
	path4 := "/"

	node.AddRoute(path4,HandlersChain{func(*context.Context){}})
	print(node, "root")

	node.AddRoute(path2,HandlersChain{func(*context.Context){}})
	print(node, "root")
	fmt.Println("----------------------------------------")
	node.AddRoute(path3,HandlersChain{func(*context.Context){}})
	print(node, "root")
	fmt.Println("----------------------------------------")
	//node.AddRoute(path1,HandlersChain{func(*context.Context){}})
	//print(node, "root")
	fmt.Println("----------------------------------------")
}

func TestNewDNode(t *testing.T) {
	str := "/aaa/bbb/cccc/"
	arr := strings.Split(str, "/")
	t.Log(len(arr))
	t.Log(arr)
}

//深度遍历
func print(root *node, fatherPath string) {
	fmt.Println(fatherPath," --> ", root)
	for _, child := range root.children {
		print(child, root.path)
	}
}
