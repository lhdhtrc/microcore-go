package download

import (
	"fmt"
	"github.com/lhdhtrc/microservice-go/model/base"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
)

func (s EntranceEntity) RemoteCert(dir string, config *base.TLSEntity) {
	logPrefix := "DownloadRemoteCert"
	dirPath := fmt.Sprintf("dep/cert/%s", dir)

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		if e := os.MkdirAll(dirPath, 755); e != nil {
			s.Logger.Error(fmt.Sprintf("%s mkdir path %s", logPrefix, e.Error()))
		}
	}

	var wg sync.WaitGroup

	ref := reflect.ValueOf(config).Elem()
	for i := 0; i < ref.NumField(); i++ {
		t := ref.Field(i)
		if t.IsValid() && !t.IsZero() {
			wg.Add(1)
			go func(url string) {
				hRes, hErr := http.Get(url)
				if hErr != nil {
					s.Logger.Error(fmt.Sprintf("%s http get %s", logPrefix, hErr.Error()))
					return
				}
				defer func(Body io.ReadCloser) {
					if ce := Body.Close(); ce != nil {
						s.Logger.Error(fmt.Sprintf("%s close http body %s", logPrefix, ce.Error()))
					}
				}(hRes.Body)

				strTemp := strings.Split(url, "/")
				file := fmt.Sprintf("%s/%s", dirPath, strings.Join(strTemp[len(strTemp)-1:], ""))

				bytes, _ := io.ReadAll(hRes.Body)
				if re := os.WriteFile(file, bytes, 0666); re != nil {
					s.Logger.Error(fmt.Sprintf("%s write error %s", logPrefix, re.Error()))
					return
				}

				t.SetString(file)
				wg.Done()
			}(t.String())
		}
	}

	wg.Wait()
}
