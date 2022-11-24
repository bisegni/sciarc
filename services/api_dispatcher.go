package services

import "sync"

var submission_channel chan *RequestSubmission = make(chan *RequestSubmission)
var result_channel chan *ApiResult = make(chan *ApiResult)
var quit chan bool

// RequestSubmission ...
type RequestSubmission struct {
	conn_id     int32
	result_type ResultType
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

	wg sync.WaitGroup
}

// NewApiDispatcher ...
func NewApiDispatcher(thread_number int) (*ApiDispatcher, error) {
	d := &ApiDispatcher{
		submission_channel:   make(chan *RequestSubmission),
		result_channel:       make(chan *ApiResult),
		quit:                 make(chan bool),
		stringAnswerCallback: nil,
		binaryAnswerCallback: nil,
	}

	d.start(thread_number)
	return d, nil
}

// Start ...
func (ad *ApiDispatcher) start(thread_number int) {
	for i := 1; i <= thread_number; i++ {
		ad.wg.Add(1)
		go ad.requesthandler(ad.submission_channel, ad.result_channel)
	}
}

// Stop ...
func (ad *ApiDispatcher) Stop() {
	close(quit)
	ad.wg.Wait()
}

func (ad *ApiDispatcher) SetStringAnswerHandler(callback func(id int, resutl *string)) {
	ad.stringAnswerCallback = callback
}

func (ad *ApiDispatcher) SetBinaryAnswerHandler(callback func(id int, resutl *[]byte)) {
	ad.binaryAnswerCallback = callback
}

// ProcessStringRequest ...
func (ad *ApiDispatcher) SubmitStringRequest(stringRequest string) error {

}

// ProcessBinaryRequest ...
func (ad *ApiDispatcher) SubmitBinaryRequest(binaryRequest []byte) error {

}

// requesthandler ...
func (ad *ApiDispatcher) requesthandler(submission_channel chan *RequestSubmission, result_channel chan *ApiResult) {
	var work bool = true
	for work {
		select {
		case <-quit:
			work = false
			return
		case request := <-submission_channel:
			break
		case answer := <-result_channel:
			break
		}
	}
	// signal exit
	ad.wg.Done()
}
