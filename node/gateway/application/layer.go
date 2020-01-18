package application

import (
	"context"
	"fmt"
	"time"

	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"github.com/eolinker/goku-api-gateway/node/gateway/application/backend"
	"github.com/eolinker/goku-api-gateway/node/gateway/application/interpreter"
	"github.com/eolinker/goku-api-gateway/node/gateway/response"
)

//LayerApplication layer application
type LayerApplication struct {
	output    response.Encoder
	backsides []*backend.Layer
	static    *staticeResponse

	timeOut time.Duration
}

//Execute execute
func (app *LayerApplication) Execute(ctx *common.Context) {

	orgBody, _ := ctx.ProxyRequest.RawBody()

	bodyObj, _ := ctx.ProxyRequest.BodyInterface()

	variables := interpreter.NewVariables(orgBody, bodyObj, ctx.ProxyRequest.Headers(), ctx.ProxyRequest.Cookies(), ctx.RestfulParam, ctx.ProxyRequest.Querys(), len(app.backsides))

	deadline := context.Background()
	cancelFunc := context.CancelFunc(nil)
	app.timeOut = 0
	if app.timeOut > 0 {
		deadline, cancelFunc = context.WithDeadline(deadline, time.Now().Add(app.timeOut))
	} else {
		deadline, cancelFunc = context.WithCancel(deadline)
	}

	resC := make(chan int, 1)
	errC := make(chan error, 1)
	go app.do(deadline, variables, ctx, resC, errC)

	defer func() {
		close(resC)
		close(errC)
	}()

	select {
	case <-deadline.Done():
		ctx.SetStatus(503, "503")
		ctx.SetBody([]byte("[ERROR]timeout!"))
		// 超时
		return
	case e := <-errC:
		fmt.Println(e)
		cancelFunc()
		ctx.SetStatus(504, "504")
		ctx.SetBody([]byte("[ERROR]Fail to get response after proxy!"))
		//error
		return
	case <-resC:
		//response
		cancelFunc()
		break
	}

	mergeResponse, headers := variables.MergeResponse()

	body, e := app.output.Encode(mergeResponse, nil)
	if e != nil {
		log.Warn("encode response error:", e)
		return
	}
	//if headers.Get("Content-Encoding") == "gzip" {
	//	var b bytes.Buffer
	//	wb := gzip.NewWriter(&b)
	//	wb.Write(body)
	//	wb.Flush()
	//	body, _ = ioutil.ReadAll(&b)
	//}
	ctx.SetProxyResponseHandler(common.NewResponseReader(headers, 200, "200", body))

}
func (app *LayerApplication) do(ctxDeadline context.Context, variables *interpreter.Variables, ctx *common.Context, resC chan<- int, errC chan<- error) {

	l := len(app.backsides)
	for i, b := range app.backsides {

		if deadline, ok := ctxDeadline.Deadline(); ok {
			if time.Now().After(deadline) {
				// 超时
				log.Warn("time out before send step:", i, "/", l)
				return
			}
		}
		r, err := b.Send(ctxDeadline, ctx, variables)

		if deadline, ok := ctxDeadline.Deadline(); ok {
			if time.Now().After(deadline) {
				// 超时
				log.Warn("time out before send step:", i+1, "/", l)
				return
			}
		}
		if err != nil {
			errC <- err
			log.Warn("error by send step:", i+1, "/", l, "\t:", err)
			return
		}
		variables.AppendResponse(r.Header, r.Body)
	}
	if deadline, ok := ctxDeadline.Deadline(); ok {
		if time.Now().After(deadline) {
			// 超时
			log.Warn("time out before send step:", l, "/", l)
			return
		}
	}
	resC <- 1

}

//NewLayerApplication create new layer application
func NewLayerApplication(apiContent *config.APIContent) *LayerApplication {
	app := &LayerApplication{
		output:    response.GetEncoder(apiContent.OutPutEncoder),
		backsides: make([]*backend.Layer, 0, len(apiContent.Steps)),
		static:    nil,
		timeOut:   time.Duration(apiContent.TimeOutTotal) * time.Millisecond,
	}

	for _, step := range apiContent.Steps {
		app.backsides = append(app.backsides, backend.NewLayer(step))
	}

	if apiContent.StaticResponse != "" {
		staticResponseStrategy := config.Parse(apiContent.StaticResponseStrategy)
		app.static = newStaticeResponse(apiContent.StaticResponse, staticResponseStrategy)
	}
	return app
}
