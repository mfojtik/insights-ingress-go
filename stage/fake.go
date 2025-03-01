package stage

import (
	"errors"
	"sync"
)

// Fake is used for tests
type Fake struct {
	Input         *Input
	StageCalledV  bool
	RejectCalledV bool
	GetURLCalledV bool
	RequestID     string
	ShouldError   bool
	URL           string
	lock          sync.Mutex
}

// Stage is used for testing with Fake input
func (f *Fake) Stage(input *Input) (string, error) {
	f.Input = input
	f.lock.Lock()
	f.StageCalledV = true
	f.lock.Unlock()
	if f.ShouldError {
		return "", errors.New("Fake Stager Error")
	}
	return f.URL, nil
}

// Reject is used for testing fake rejections
func (f *Fake) Reject(requestID string) error {
	f.lock.Lock()
	f.RejectCalledV = true
	f.lock.Unlock()
	f.RequestID = requestID
	return nil
}

// GetURL is used to test fake url returns
func (f *Fake) GetURL(requestID string) (string, error) {
	f.lock.Lock()
	f.GetURLCalledV = true
	f.lock.Unlock()
	if f.ShouldError {
		return "", errors.New("Fake Stager Error")
	}
	return f.URL, nil
}

func (f *Fake) GetURLCalled() bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	return f.GetURLCalledV
}

func (f *Fake) StageCalled() bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	return f.StageCalledV
}

func (f *Fake) RejectCalled() bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	return f.RejectCalledV
}
