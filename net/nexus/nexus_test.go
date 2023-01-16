package nexus

import (
	"fmt"
	"testing"
)

func TestUpdateFileToNexus(t *testing.T) {
	target := "/Users/zhaoshoucheng/Desktop/cks_old.tar.gz"
	nexusRepo := "https://127.0.0.1/nexus/content/repositories/test1-packages/tags/siteconfig/v000000/"
	err := UpdateFileToNexus(nexusRepo, target, "v000000", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("done")
}
