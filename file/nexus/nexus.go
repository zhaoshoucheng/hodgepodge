package nexus

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 上传二进制文件
func UpdateFileToNexus(nexusUrl, file, version, artifactid, groupid string) error {
	if artifactid == "" {
		artifactid = "siteconfig"
	}
	if groupid == "" {
		groupid = "lucky/tags/lfeserver"
	}
	targetFile := artifactid + "-" + version + ".tar.gz"
	urlStr := nexusUrl + groupid + "/" + artifactid + "/" + version + "/" + targetFile

	err := uploadPkgFile(file, urlStr)
	if err != nil {
		return err
	}

	err = uploadMd5File(file, urlStr)
	if err != nil {
		return err
	}

	err = uploadSha1File(file, urlStr)
	if err != nil {
		return err
	}
	return nil
}
func uploadPkgFile(file string, urlStr string) error {
	fh, err := os.Open(file)
	if err != nil {
		return err
	}
	//bodyBuf := bytes.NewBufferString("")
	//bodyWriter := multipart.NewWriter(bodyBuf)

	//boundary := bodyWriter.Boundary()
	//closeBuf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	//requestReader := io.MultiReader(bodyBuf, fh, closeBuf)
	//_ = closeBuf
	//requestReader := io.MultiReader(fh)
	//requestReader, err := gzip.NewReader(fh)
	//if err != nil {
	//	return err
	//}
	return PostFile(urlStr, fh, "multipart/form-data;")
}
func uploadMd5File(file string, urlStr string) error {
	fh, err := os.Open(file)
	if err != nil {
		return err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, fh)
	md5Value := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(md5Value)
	payload := strings.NewReader(md5Value)
	urlStr += ".md5"
	return PostFile(urlStr, payload, "multipart/form-data;")
}
func uploadSha1File(file string, urlStr string) error {
	fh, err := os.Open(file)
	if err != nil {
		return err
	}
	shaObj := sha1.New()
	_, _ = io.Copy(shaObj, fh)
	shaValue := hex.EncodeToString(shaObj.Sum(nil))
	payload := strings.NewReader(shaValue)
	urlStr += ".sha1"
	return PostFile(urlStr, payload, "multipart/form-data;")
}
func PostFile(urlStr string, payload io.Reader, ty string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, payload)
	if err != nil {
		return err
	}
	//req.Header.Add("Authorization", "Basic bHNvcDpsc29wXzIwMjI=")
	req.Header.Add("Authorization", "Basic ***********")
	req.Header.Add("Content-Type", ty)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("StatusCode code err %d", resp.StatusCode))
	}
	return nil
}
