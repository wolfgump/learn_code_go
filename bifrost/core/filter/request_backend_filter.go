package filter

import (
	log "github.com/go-kratos/kratos/pkg/log"
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

	client := http.Client{}
	client.Timeout = time.Millisecond * time.Duration(10*1000)
	request, _ := http.NewRequest(http.MethodGet, ctx.finalPath, ctx.Request.Body)
	request.Header = ctx.Request.Header
	resp, errDo := client.Do(request)

	if errDo != nil {
		log.Error("request backend error,message:" + errDo.Error())
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

var _ IFilter = (*RequestBackendFilter)(nil)
