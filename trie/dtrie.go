package trie

import (
	"errors"
	"fmt"
	"strings"
)

const WildChildIndices = "__WildChildKey"
//dnode 树基础结构
type dnode struct {
	indices  string				//路由名，作为children中的key
	path 	  string			//当前/分隔原始路径
	children  map[string]*dnode	//正常路由节点
	handlers  HandlersChain
	nType     nodeType
}
// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}
type Params []Param

func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

func (ps Params) ByName(name string) (va string) {
	va, _ = ps.Get(name)
	return
}

type nodeValue struct {
	handlers HandlersChain
	params   *Params
	tsr      bool
	fullPath string
}

func NewDNode() *dnode {
	return &dnode{
		children: make(map[string]*dnode),
		nType: root,
	}
}

func (dn *dnode)GetValue(fullPath string, params *Params) nodeValue {
	res := nodeValue{}
	res.params = params
	pathArr := getPathArr(fullPath)
	find := true
	if len(fullPath) == 0 {
		find = false
	}

	findPath := "/"
	for index, path := range pathArr {
		if node, exists := dn.children[WildChildIndices];exists {
			findPath =  findPath + node.path + "/"
			dn = node
			key := string([]byte(node.path)[1:])
			if node.nType == catchAll {
				var value string
				for i := index; i < len(pathArr); i++ {
					value = value + "/" + pathArr[i]
				}
				*res.params = append(*res.params, Param{Key: key,Value: value})
				break
			} else {
				value := path
				*res.params = append(*res.params, Param{Key: key,Value: value})
				continue
			}
		}

		if node, exists := dn.children[path];exists {
			findPath =  findPath + node.path + "/"
			dn = node
		} else {
			//not find
			find = false
			break
		}
	}
	res.fullPath = findPath
	if find {
		res.handlers = dn.handlers
		return res
	}
	if dn.handlers != nil {
		res.fullPath = "/"
		res.handlers = dn.handlers
	}
	return res
}
func (dn *dnode)AddPath(fullPath string, handlers  HandlersChain) error {
	if dn.nType == static {
		dn.children = make(map[string]*dnode)
		dn.nType = root
	}
	if fullPath == "" {
		return errors.New("Illegal path ")
	}

	if fullPath == "/" {
		if dn.handlers != nil {
			return errors.New("path have exists ")
		}
		dn.indices = WildChildIndices
		dn.path = fullPath
		dn.children = make(map[string]*dnode)
		dn.handlers = handlers
		return nil
	}

	pathArr := getPathArr(fullPath)
	if len(pathArr) == 0 {
		return errors.New("Illegal path ")
	}
	pathArrLen := len(pathArr)
	var c byte
	var findPath string
	for index, part := range pathArr {
		findPath = findPath + "/" + part
		countParams := countParams(part)
		if countParams > 2 || part == ""{
			return errors.New(fmt.Sprintf("Illegal path %s", findPath))
		}
		indices := part
		c = []byte(part)[0]
		if c == ':' || c == '*' {
			indices = WildChildIndices
			//* 必须是最后一个
			if c == '*' && index != pathArrLen -1 {
				return errors.New(fmt.Sprintf("catch-all routes are only allowed at the end of the path in path %s",findPath))
			}

			//已有路由冲突检测 '*',":" --> 普通路由
			if c == '*' && len(dn.children) != 0 {
				return errors.New(fmt.Sprintf("conflicts with existing wildcard %s",findPath))
			}
		}

		//已有通配符冲突检测 普通路由 --> '*',':'
		if node, exists := dn.children[WildChildIndices];exists {
			if node.path == part && node.nType == param{
				dn = node
				continue
			}
			return errors.New(fmt.Sprintf("conflicts with existing wildcard %s",part))
		}

		//普通路由
		if node, exists := dn.children[indices]; exists {
			dn = node
		} else {
			child := &dnode{
				indices: indices,
				path: part,
				children: make(map[string]*dnode),
				nType: getNType(part),
			}

			dn.children[indices] = child
			dn = child
		}
	}
	if dn.handlers != nil {
		return errors.New(fmt.Sprintf("conflicts with existing %s",findPath))
	}
	dn.handlers = handlers
	return nil
}

func getNType(part string) nodeType {
	if part != "" && part[0] == ':' { return param }
	if part != "" && part[0] == '*' { return catchAll }
	return normal
}

func getPathArr(fullPath string) []string {
	if fullPath == "" {
		return []string{}
	}
	if []byte(fullPath)[0] == '/' {
		fullPath = string([]byte(fullPath)[1:])
	}
	if fullPath == "" {
		return []string{}
	}
	if []byte(fullPath)[len(fullPath) - 1] == '/' {
		fullPath = string([]byte(fullPath[:len(fullPath) - 1]))
	}

	return  strings.Split(fullPath, "/")
}
