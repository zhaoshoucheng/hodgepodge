package nexus

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// 上传二进制文件
func UpdateFileToNexus(nexusUrl, file, version, artifactid, groupid string) error {
	if artifactid == "" {
		artifactid = "siteconfig"
	}
	if groupid == "" {
		groupid = "tags"
	}
	targetFile := artifactid + "-" + version + ".tar.gz"
	urlStr := nexusUrl + groupid + "/" + artifactid + "/" + version + "/" + targetFile

	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)

	fh, err := os.Open(file)
	if err != nil {
		return err
	}
	boundary := bodyWriter.Boundary()
	closeBuf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	requestReader := io.MultiReader(bodyBuf, fh, closeBuf)
	fi, err := fh.Stat()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", urlStr, requestReader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.Header.Add("Authorization", "Basic bHNvcDpsc29wXzIwMjI=")
	req.ContentLength = fi.Size() + int64(bodyBuf.Len()) + int64(closeBuf.Len())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("StatusCode code err %d", resp.StatusCode))
	}
	return nil
}
