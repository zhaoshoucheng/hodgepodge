package nexus

import (
	"fmt"
	"testing"
)

func TestUpdateFileToNexus(t *testing.T) {
	target := "/Users/zhaoshoucheng/Downloads/globalconfig-v37.tar.gz"
	nexusRepo := "http://10.218.21.130:8081/nexus/content/repositories/test3-packages/"
	err := UpdateFileToNexus(nexusRepo, target, "v37", "globalconfig", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("done")
}
