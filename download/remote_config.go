package download

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func (s EntranceEntity) ReadRemoteConfig(Source []string, Config []*interface{}) {
	logPrefix := "ReadRemoteConfig"
	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	var wg sync.WaitGroup
	for i, it := range Source {
		wg.Add(1)
		go func(url string, index int) {
			hRes, hErr := http.Get(url)
			if hErr != nil {
				s.logger.Error(fmt.Sprintf("%s http get %s", logPrefix, hErr.Error()))
				return
			}
			defer func(Body io.ReadCloser) {
				if ce := Body.Close(); ce != nil {
					s.logger.Error(fmt.Sprintf("%s close http body %s", logPrefix, ce.Error()))
				}
			}(hRes.Body)

			bytes, _ := io.ReadAll(hRes.Body)
			if re := json.Unmarshal(bytes, Config[index]); re != nil {
				s.logger.Error(fmt.Sprintf("%s json unmarshal %s", logPrefix, re.Error()))
				return
			}

			wg.Done()
		}(it, i)
	}
	wg.Wait()

	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))
}
