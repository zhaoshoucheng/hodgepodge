package trie

import (
	"errors"
	"strings"
)

const WildChildIndices = "__WildChildKey"
//node 树基础结构
type dnode struct {
	indices  string				//路由名，作为children中的key
	path 	  string			//当前/分隔原始路径
	children  map[string]*dnode	//正常路由节点
	handlers  HandlersChain
	nType     nodeType
}

func NewDNode() *dnode {
	return &dnode{
		children: make(map[string]*dnode),
		nType: root,
	}
}

func (dn *dnode)AddPath(path string, handlers  HandlersChain) error {
	if dn.nType == static {
		return errors.New("dnode need init")
	}

	if path == "/" {
		if dn.handlers != nil {
			return errors.New("path have exists ")
		}
		child := &dnode{
			indices: WildChildIndices,
			path: path,
			children: make(map[string]*dnode),
			handlers: handlers,
			nType:catchAll,
		}
		dn.children[WildChildIndices] = child
		return nil
	}

	if []byte(path)[0] == '/' {
		path = string([]byte(path)[1:])
	}
	if []byte(path)[len(path) - 1] == '/' {
		path = string([]byte(path[:len(path) - 1]))
	}

	pathArr := strings.Split(path, "/")
	if len(pathArr) == 0 {
		return errors.New("Illegal path ")
	}
	pathArrLen := len(pathArr)
	var c byte
	for index, part := range pathArr {
		countParams := countParams(part)
		if countParams > 2 || part == ""{
			return errors.New("Illegal path ")
		}
		indices := part
		c = []byte(part)[0]
		//处理通配符
		if c ==  ':' || c == '*' {
			indices = WildChildIndices
		}

		if node, exists := dn.children[indices]; exists {
			if index == pathArrLen - 1 && node.handlers != nil {
				return errors.New("path have exists ")
			}
			if index == pathArrLen - 1 {
				node.handlers = handlers
				return nil
			}
			dn = node
		} else {
			child := &dnode{
				indices: indices,
				path: part,
				children: make(map[string]*dnode),
				nType: getNType(part),
			}

			dn.children[indices] = child
			//尾部推出
			if index == pathArrLen - 1 {
				child.handlers = handlers
				return nil
			}
			dn = child
		}
	}

	//dn.handlers = handlers
	//dn.nType = normal
	//if c == ':' {
	//	dn.nType = param
	//}
	//if c == '*' {
	//	dn.nType = catchAll
	//}
	return nil
}

func getNType(part string) nodeType {
	if part != "" && part[0] == ':' { return param }
	if part != "" && part[0] == '*' { return catchAll }
	return normal
}

func (dn *dnode)getHandlers(path string) (handlers  HandlersChain) {
	return nil
}