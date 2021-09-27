package trie

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestDnode_AddPath(t *testing.T) {
	dnode := NewDNode()
	path1 := "/aaa/bbb/ccc"			//s
	path2 := "/aaa/bbb"				//s
	path3 := "/aaa/bbb/ccc/ddd" 	//s
	path4 := "/aaa/bbb/ddd"			//s
	dnode2 := NewDNode()
	path5 := "/aaa/:bbb/ccc"		//s
	path6 := "/aaa/:bbb/"			//s
	path7 := "/aaa/bbb/ccc"			//f
	dnode3 := NewDNode()
	path8 := "/aaa/:bbb/ccc"		//s
	path9 := "/aaa/*bbb"			//f
	dnode4 := NewDNode()
	path10 := "/aaa/*bbb"			//f
	fmt.Println("-----------------path1-----------------------")
	err := dnode.AddPath(path1,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("-----------------path2-----------------------")
	err = dnode.AddPath(path2,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("------------------path3----------------------")
	err = dnode.AddPath(path3,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("------------------path4----------------------")
	err = dnode.AddPath(path4,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode, "root")
	}

	fmt.Println("-----------------path5-----------------------")
	err = dnode2.AddPath(path5,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode2, "root")
	}

	fmt.Println("-----------------path6-----------------------")
	err = dnode2.AddPath(path6,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode2, "root")
	}

	fmt.Println("-----------------path7-----------------------")
	err = dnode2.AddPath(path7,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode2, "root")
	}

	fmt.Println("-----------------path8-----------------------")
	err = dnode3.AddPath(path8,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode3, "root")
	}

	fmt.Println("-----------------path9-----------------------")
	err = dnode3.AddPath(path9,HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode3, "root")
	}

	fmt.Println("-----------------path10-----------------------")
	err = dnode4.AddPath(path10, HandlersChain{func(*context.Context){}})
	if err != nil {
		t.Log(err)
	} else {
		printDnode(dnode4, "root")
	}
}

//深度遍历
func printDnode(root *dnode, fatherPath string) {
	fmt.Println(fatherPath," --> ", root)
	for _, child := range root.children {
		printDnode(child, root.path)
	}
}

var fakeHandlerValue string

func fakeHandler(val string) HandlersChain {
	return HandlersChain{func(c *context.Context) {
		fakeHandlerValue = val
	}}
}

func TestDnode_AddPath2(t *testing.T) {
	tree := &dnode{}
	routes := [...]string{
		"/",
		"/cmd/:tool/",
		"/cmd/:tool/:sub",
		"/cmd/whoami",
		"/cmd/whoami/root",
		"/cmd/whoami/root/",
		"/src/*filepath",
		"/search/",
		"/search/:query",
		"/search/gin-gonic",
		"/search/google",
		"/user_:name",
		"/user_:name/about",
		"/files/:dir/*filepath",
		"/doc/",
		"/doc/go_faq.html",
		"/doc/go1.html",
		"/info/:user/public",
		"/info/:user/project/:project",
		"/info/:user/project/golang",
	}
	for _, route := range routes {
		err := tree.AddPath(route, fakeHandler(route))
		if err != nil {
			t.Log(route, err)
		}
	}
	t.Log("done")
}

type testRequests []struct {
	path       string
	nilHandler bool
	route      string
	ps         Params
}
func getParams() *Params {
	ps := make(Params, 0, 20)
	return &ps
}


func checkRequests(t *testing.T, tree *dnode, requests testRequests, unescapes ...bool) {
	for _, request := range requests {
		value := tree.GetValue(request.path, getParams())
		if value.handlers == nil {
			if !request.nilHandler {
				t.Errorf("handle mismatch for route '%s': Expected non-nil handle", request.path)
				continue
			}
		} else if request.nilHandler {
			t.Errorf("handle mismatch for route '%s': Expected nil handle", request.path)
			continue
		} else {
			value.handlers[0](nil)
			if fakeHandlerValue != request.route {
				t.Errorf("handle mismatch for route '%s': Wrong handle (%s != %s)", request.path, fakeHandlerValue, request.route)
				continue
			}
		}

		if value.params != nil {
			if !reflect.DeepEqual(*value.params, request.ps) {
				t.Errorf("Params mismatch for route '%s'", request.path)
				continue
			}
		}
		t.Log("succeed ",request.path)
	}
}

func TestDnode_GetValue(t *testing.T) {
	tree := &dnode{}
	routes := [...]string{
		"/",
		"/cmd/:tool/",
		"/cmd/:tool/:sub",
		"/cmd/whoami",
		"/cmd/whoami/root",
		"/cmd/whoami/root/",
		"/src/*filepath",
		"/search/",
		"/search/:query",
		"/search/gin-gonic",
		"/search/google",
		"/user_:name",
		"/user_:name/about",
		"/files/:dir/*filepath",
		"/doc/",
		"/doc/go_faq.html",
		"/doc/go1.html",
		"/info/:user/public",
		"/info/:user/project/:project",
		"/info/:user/project/golang",
	}
	for _, route := range routes {
		err := tree.AddPath(route, fakeHandler(route))
		if err != nil {
			t.Log(route, err)
		}
	}
	checkRequests(t, tree, testRequests{
		{"/", false, "/", *getParams()},
		{"/cmd/test", true, "/cmd/:tool/", Params{Param{"tool", "test"}}},
		{"/cmd/test/", false, "/cmd/:tool/", Params{Param{"tool", "test"}}},
		{"/cmd/test/3", false, "/cmd/:tool/:sub", Params{Param{Key: "tool", Value: "test"}, Param{Key: "sub", Value: "3"}}},
		{"/cmd/who", true, "/cmd/:tool/", Params{Param{"tool", "who"}}},
		{"/cmd/who/", false, "/cmd/:tool/", Params{Param{"tool", "who"}}},
		{"/cmd/whoami", false, "/cmd/whoami", nil},
		{"/cmd/whoami/", true, "/cmd/whoami", nil},
		{"/cmd/whoami/r", false, "/cmd/:tool/:sub", Params{Param{Key: "tool", Value: "whoami"}, Param{Key: "sub", Value: "r"}}},
		{"/cmd/whoami/r/", true, "/cmd/:tool/:sub", Params{Param{Key: "tool", Value: "whoami"}, Param{Key: "sub", Value: "r"}}},
		{"/cmd/whoami/root", false, "/cmd/whoami/root", nil},
		{"/cmd/whoami/root/", false, "/cmd/whoami/root/", nil},
		{"/src/", false, "/src/*filepath", Params{Param{Key: "filepath", Value: "/"}}},
		{"/src/some/file.png", false, "/src/*filepath", Params{Param{Key: "filepath", Value: "/some/file.png"}}},
		{"/search/", false, "/search/", *getParams()},
		{"/search/someth!ng+in+ünìcodé", false, "/search/:query", Params{Param{Key: "query", Value: "someth!ng+in+ünìcodé"}}},
		{"/search/someth!ng+in+ünìcodé/", true, "", Params{Param{Key: "query", Value: "someth!ng+in+ünìcodé"}}},
		{"/search/gin", false, "/search/:query", Params{Param{"query", "gin"}}},
		{"/search/gin-gonic", false, "/search/gin-gonic", nil},
		{"/search/google", false, "/search/google", nil},
		{"/user_gopher", false, "/user_:name", Params{Param{Key: "name", Value: "gopher"}}},
		{"/user_gopher/about", false, "/user_:name/about", Params{Param{Key: "name", Value: "gopher"}}},
		{"/files/js/inc/framework.js", false, "/files/:dir/*filepath", Params{Param{Key: "dir", Value: "js"}, Param{Key: "filepath", Value: "/inc/framework.js"}}},
		{"/info/gordon/public", false, "/info/:user/public", Params{Param{Key: "user", Value: "gordon"}}},
		{"/info/gordon/project/go", false, "/info/:user/project/:project", Params{Param{Key: "user", Value: "gordon"}, Param{Key: "project", Value: "go"}}},
		{"/info/gordon/project/golang", false, "/info/:user/project/golang", Params{Param{Key: "user", Value: "gordon"}}},
	})
	t.Log("done")
}