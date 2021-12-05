package fileType

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFileType(t *testing.T) {
	//f, err := os.Open("C:\\Users\\Administrator\\Desktop\\api.html")
	f, err := os.Open("d:\\Danny\\Desktop\\2021-07-10桌面\\合同2.pdf")
	if err != nil {
		fmt.Println("open error: ", err)
	}

	fSrc, err := ioutil.ReadAll(f)
	fmt.Println(GetFileType(fSrc[:10]))
	fmt.Println("GetFileType(fSrc[:10])")
}
