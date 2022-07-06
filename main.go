package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/zhaoshoucheng/hodgepodge/zlog"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}
func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo1 val = %d\n", x)
	}
}
func main() {
	ipstr := "82.157.71.193"
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return
	}
	fmt.Println(ip)
	ip4 := ip.To4()
	fmt.Println(ip4)
	fmt.Println(binary.BigEndian.Uint32(ip4))
	return

	path1 := "/Users/zhaoshoucheng/Desktop/ufe_n/custom"
	files1 := make(map[string]string)
	path2 := "/Users/zhaoshoucheng/Desktop/ufe_o/custom"
	files2 := make(map[string]string)
	err := filepath.Walk(path1, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files1[info.Name()] = path
		return nil
	})
	if err != nil {
		fmt.Println("err")
		return
	}
	err = filepath.Walk(path2, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files2[info.Name()] = path
		return nil
	})
	for key, path2 := range files2 {
		if path1, exists := files1[key]; exists {
			if "luckymarketgrowthtooladmin" == key {
				_ = key
			}
			flag, err := DiffFile(path1, path2)
			if err != nil {
				fmt.Println(key, "------false ", err)
				continue
			}
			if !flag {
				fmt.Println(key, "------false ")
			}
		}
	}
	return

	files, err := os.ReadDir("/Users/zhaoshoucheng/data/ssrc/mysrc/hodgepodge/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(files))
	return

	file, err := os.OpenFile("/Users/zhaoshoucheng/data/src/mysrc/hodgepodge/text.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write([]byte(fmt.Sprintf("# Version: %s\n", "1.0.0")))
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write([]byte("xx111xxssss"))

	targetFIleName := "/Users/zhaoshoucheng/data/src/mysrc/hodgepodge/text2.log"
	targetFile, err := os.Create(targetFIleName)
	if err != nil {
		fmt.Println("open file fail (file: %s, err: %v )", targetFIleName, err)
		return
	}

	_, err = io.Copy(file, targetFile)
	if err != nil {
		fmt.Println("copy file fail (file: %s, to %s, err: %v )", err)
		return
	}

	fmt.Println(runtime.GOOS)
	fmt.Println(err)
	return
	timestamp := int64(1641888018)
	t := time.Unix(timestamp, 0)
	fmt.Println(t)
	return
	writer := zlog.MultiLevelWriter(zlog.NewStdoutWriter())
	logger := zlog.New(writer)
	logger.Write([]byte("hello world!"))
	time.Sleep(time.Second * 5)
	return
}

func DiffFile(file1, file2 string) (bool, error) {

	fi1, err := os.Open(file1)
	if err != nil {
		return false, err
	}
	fi2, err := os.Open(file2)
	if err != nil {
		return false, err
	}

	// 创建 Reader
	r1 := bufio.NewReader(fi1)
	r2 := bufio.NewReader(fi2)

	var err1 error
	var err2 error
	for {
		var line1 string
		for {
			var lineBytes1 []byte
			lineBytes1, err1 = r1.ReadBytes('\n')
			line1 = strings.TrimSpace(string(lineBytes1))
			if err1 != nil {
				if err1 == io.EOF {
					break
				}
				return false, err
			}
			//去掉空格
			line1 = strings.ReplaceAll(line1, " ", "")
			//去掉注释
			if len([]byte(line1)) != 0 && []byte(line1)[0] == '#' {
				continue
			}
			if line1 != "" {
				break
			}
		}
		var line2 string
		for {
			var lineBytes2 []byte
			lineBytes2, err2 = r2.ReadBytes('\n')
			line2 = strings.TrimSpace(string(lineBytes2))
			if err2 != nil {
				if err2 == io.EOF {
					break
				}
				return false, err
			}
			//去掉空格
			line2 = strings.ReplaceAll(line2, " ", "")
			//去掉注释
			if len([]byte(line2)) != 0 && []byte(line2)[0] == '#' {
				continue
			}
			if line2 != "" {
				break
			}
		}
		if err1 == io.EOF || err2 == io.EOF {
			break
		}
		if line1 == line2 {
			continue
		}
		//fmt.Println(line1, line2)
		return false, nil
	}
	if err1 == io.EOF && err2 == io.EOF {
		return true, nil
	}
	if err1 == io.EOF {
		for {
			lineBytes1, err := r2.ReadBytes('\n')
			line := strings.TrimSpace(string(lineBytes1))
			if err != nil {
				if err == io.EOF {
					break
				}
				return false, err
			}
			//去掉空格
			line = strings.ReplaceAll(line, " ", "")
			if line == "" {
				continue
			}
			//	fmt.Println("r2 ", line)
			return false, nil
		}
	}
	if err2 == io.EOF {
		for {
			lineBytes1, err := r1.ReadBytes('\n')
			line := strings.TrimSpace(string(lineBytes1))
			if err != nil {
				if err == io.EOF {
					break
				}
				return false, err
			}
			//去掉空格
			line = strings.ReplaceAll(line, " ", "")
			if line == "" {
				continue
			}
			//			fmt.Println("r1 ", line)
			return false, nil
		}
	}
	return true, nil
}
