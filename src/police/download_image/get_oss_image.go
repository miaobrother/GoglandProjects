package download_image

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
	"time"
)

var LocalFile = "/Users/zhangyalei/Downloads/"

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func GetImageList([]string) (s []string, err error) {
	client, err := oss.New("http://oss-cn-beijing.aliyuncs.com", "tVvDKCvLkCHCpXQf", "E9zgpaFE52fvxZ9iwtyWUZOILn9V0a")
	if err != nil {
		handleError(err)
	}
	bucket, err := client.Bucket("carlog")
	if err != nil {
		handleError(err)
	}
	bucket.ListObjects()

	/*
		获取线上oss中T-1日照片
		now := time.Now()
		m ,_ := time.ParseDuration("-24h")
		m_1 := now.Add(m)
		t :=m_1.Format("2006/1/2")
		lsRes, _ := bucket.ListObjects( oss.Prefix(t), oss.MaxKeys(1000)) //设置最大对象数
			go func() {
				for _, val := range lsRes.Objects {
					c := strings.Split(val.Key, "/")
					err = bucket.GetObjectToFile(val.Key, LocalFile+"/"+c[3])
				}
			}()
		time.Sleep(time.Second * 20)
		return
	*/
}
