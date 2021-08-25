package trie

import (
	"context"
	"fmt"
	"testing"
)

func TestDnode_AddPath(t *testing.T) {
	dnode := NewDNode()
	path1 := "/aaa/bbb/ccc"
	path2 := "/aaa/bbb"
	path3 := "/aaa/bbb/ccc/ddd"
	path4 := "/aaa/bbb/ddd"
	path5 := "/aaa/:bbb/ccc"
	path6 := "/aaa/:bbb/"
	path7 := "/aaa/bbb/ccc"
	path8 := "/aaa/:bbb/ccc"
	path9 := "/aaa/*bbb"
	err := dnode.AddPath(path1,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("----------------------------------------")
	err = dnode.AddPath(path2,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("----------------------------------------")
	err = dnode.AddPath(path3,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("----------------------------------------")
	err = dnode.AddPath(path4,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("----------------------------------------")
	err = dnode.AddPath(path5,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("----------------------------------------")
	err = dnode.AddPath(path6,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("-----------------path have exists-----------------------")
	err = dnode.AddPath(path7,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("-----------------path have exists-----------------------")
	err = dnode.AddPath(path8,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("-----------------path have exists-----------------------")
	err = dnode.AddPath(path9,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}
}

//深度遍历
func printDnode(root *dnode, fatherPath string) {
	fmt.Println(fatherPath," --> ", root)
	for _, child := range root.children {
		printDnode(child, root.path)
	}
}