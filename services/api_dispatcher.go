package services

import (
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/shettyh/threadpool"
)

type DataType int32

const (
	Binary DataType = iota
	String
)

var command_regex, _ = regexp.Compile("(\\w*) ([a-zA-Z0-9 _\\-\\:]*)")

// RequestSubmission ...
type RequestSubmission struct {
	conn_id      int
	request_type DataType
	data         interface{}
}

// ApiDefinition ...
type ApiFactory interface {
	Name() string
	BuildApiByString(result_channel chan<- *ApiResult, conn_id int, conf *string) interface{}
	BuildApiByBinary(result_channel chan<- *ApiResult, conn_id int, conf *[]byte) interface{}
}

// Api result
type ApiResult struct {
	conn_id     int
	result_type DataType
	result      interface{}
	err         error
}

// ApiDispatcher ...
type ApiDispatcher struct {
	submission_channel chan *RequestSubmission
	result_channel     chan *ApiResult
	quit               chan bool

	stringAnswerCallback func(id int, resutl *string)
	binaryAnswerCallback func(id int, resutl *[]byte)

	api_catalog map[string]ApiFactory
	pool        *threadpool.ThreadPool
	wg          sync.WaitGroup
}

// NewApiDispatcher ...
func NewApiDispatcher(thread_number int) (*ApiDispatcher, error) {
	d := &ApiDispatcher{
		submission_channel:   make(chan *RequestSubmission),
		result_channel:       make(chan *ApiResult),
		quit:                 make(chan bool),
		stringAnswerCallback: nil,
		binaryAnswerCallback: nil,
		api_catalog:          make(map[string]ApiFactory),
		pool:                 threadpool.NewThreadPool(1, 1000000),
	}

	d.start(thread_number)
	return d, nil
}

// RegisterApi ...
func (ad *ApiDispatcher) RegisterApi(def ApiFactory) error {
	if _, ok := ad.api_catalog[def.Name()]; ok {
		log.Printf("Api name %s has been already registered\n", def.Name())
	}
	ad.api_catalog[def.Name()] = def
	return nil
}

// Start ...
func (ad *ApiDispatcher) start(thread_number int) {
	for i := 1; i <= thread_number; i++ {
		ad.wg.Add(1)
		go ad.requesthandler()
	}
}

// Stop ...
func (ad *ApiDispatcher) Stop() {
	close(ad.quit)
	ad.wg.Wait()
}

func (ad *ApiDispatcher) SetStringAnswerHandler(callback func(id int, resutl *string)) {
	ad.stringAnswerCallback = callback
}

func (ad *ApiDispatcher) SetBinaryAnswerHandler(callback func(id int, resutl *[]byte)) {
	ad.binaryAnswerCallback = callback
}

func (ad *ApiDispatcher) getApiFromStringRequest(req *RequestSubmission) threadpool.Runnable {
	command := req.data.(*string)
	var command_param strings.Builder
	var api_factory ApiFactory
	var ok bool
	match := command_regex.FindStringSubmatch(*command)

	if len(match) == 1 {
		log.Printf("the command '%s' is not well formed\n", *command)
	}

	if api_factory, ok = ad.api_catalog[match[1]]; !ok {
		log.Printf("Api factory %s has not registered\n", match[1])
		return nil
	}

	if len(match) > 2 {
		for i := 2; i < len(match); i++ {
			command_param.WriteString(match[i])
		}
	}
	command_param_string := command_param.String()
	return api_factory.BuildApiByString(
		ad.result_channel,
		req.conn_id,
		&command_param_string).(threadpool.Runnable)
}

func (ad *ApiDispatcher) getApiFromBinaryRequest(req *RequestSubmission) threadpool.Runnable {
	return nil
}

// ProcessStringRequest ...
func (ad *ApiDispatcher) processRequest(req *RequestSubmission) {
	var api threadpool.Runnable
	switch req.request_type {
	case String:
		{
			api = ad.getApiFromStringRequest(req)
			break
		}
	case Binary:
		{
			api = ad.getApiFromBinaryRequest(req)
			break
		}
	}
	if api != nil {
		ad.pool.Execute(api)
	}
}

// ProcessBinaryRequest ...
func (ad *ApiDispatcher) SubmitRequest(req *RequestSubmission) error {
	ad.submission_channel <- req
	return nil
}

// requesthandler ...
func (ad *ApiDispatcher) requesthandler() {
	var work bool = true
	for work {
		select {
		case <-ad.quit:
			work = false
			return
		case request := <-ad.submission_channel:
			{
				ad.processRequest(request)
				break
			}
		case answer := <-ad.result_channel:
			switch answer.result_type {
			case Binary:
				{
					if ad.binaryAnswerCallback != nil {
						ad.binaryAnswerCallback(int(answer.conn_id), answer.result.(*[]byte))
					}
					break
				}
			case String:
				{
					if ad.binaryAnswerCallback != nil {
						ad.stringAnswerCallback(int(answer.conn_id), answer.result.(*string))
					}
					break
				}
			}
		}
	}
	// signal exit
	ad.wg.Done()
}
