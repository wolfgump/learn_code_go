package filter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestBackendFilter struct {
	Filter
}

func (r RequestBackendFilter) Name() string {
	return "RequestBackend"
}

func (r RequestBackendFilter) Order() int {
	return 300
}

func (r RequestBackendFilter) DoFilter(ctx *FilterContext) {

	if "GET" == ctx.Request.Method {
		client := http.Client{}
		client.Timeout = time.Millisecond * time.Duration(10*1000)
		request, _ := http.NewRequest(http.MethodGet, ctx.finalPath, nil)
		request.Header = ctx.Request.Header
		resp, errDo := client.Do(request)

		if errDo != nil {
			fmt.Printf("request backend error,message:" + errDo.Error())
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		headers := ctx.ResponseWriter.Header()
		for key, _ := range resp.Header {
			headers.Set(key, resp.Header.Get(key))
		}
		ctx.ResponseWriter.Write(bytes)
		//result := string(bytes)
		defer resp.Body.Close()

	}
	if "POST" == ctx.Request.Method {

	}
}

var _ IFilter = (*RequestBackendFilter)(nil)
