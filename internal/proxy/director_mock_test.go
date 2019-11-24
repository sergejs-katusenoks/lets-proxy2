package proxy

// Code generated by http://github.com/gojuno/minimock (3.0.6). DO NOT EDIT.

//go:generate minimock -i github.com/rekby/lets-proxy2/internal/proxy.Director -o ./director_mock_test.go

import (
	"net/http"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// DirectorMock implements Director
type DirectorMock struct {
	t minimock.Tester

	funcDirector          func(request *http.Request)
	inspectFuncDirector   func(request *http.Request)
	afterDirectorCounter  uint64
	beforeDirectorCounter uint64
	DirectorMock          mDirectorMockDirector
}

// NewDirectorMock returns a mock for Director
func NewDirectorMock(t minimock.Tester) *DirectorMock {
	m := &DirectorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DirectorMock = mDirectorMockDirector{mock: m}
	m.DirectorMock.callArgs = []*DirectorMockDirectorParams{}

	return m
}

type mDirectorMockDirector struct {
	mock               *DirectorMock
	defaultExpectation *DirectorMockDirectorExpectation
	expectations       []*DirectorMockDirectorExpectation

	callArgs []*DirectorMockDirectorParams
	mutex    sync.RWMutex
}

// DirectorMockDirectorExpectation specifies expectation struct of the Director.Director
type DirectorMockDirectorExpectation struct {
	mock   *DirectorMock
	params *DirectorMockDirectorParams

	Counter uint64
}

// DirectorMockDirectorParams contains parameters of the Director.Director
type DirectorMockDirectorParams struct {
	request *http.Request
}

// Expect sets up expected params for Director.Director
func (mmDirector *mDirectorMockDirector) Expect(request *http.Request) *mDirectorMockDirector {
	if mmDirector.mock.funcDirector != nil {
		mmDirector.mock.t.Fatalf("DirectorMock.Director mock is already set by Set")
	}

	if mmDirector.defaultExpectation == nil {
		mmDirector.defaultExpectation = &DirectorMockDirectorExpectation{}
	}

	mmDirector.defaultExpectation.params = &DirectorMockDirectorParams{request}
	for _, e := range mmDirector.expectations {
		if minimock.Equal(e.params, mmDirector.defaultExpectation.params) {
			mmDirector.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDirector.defaultExpectation.params)
		}
	}

	return mmDirector
}

// Inspect accepts an inspector function that has same arguments as the Director.Director
func (mmDirector *mDirectorMockDirector) Inspect(f func(request *http.Request)) *mDirectorMockDirector {
	if mmDirector.mock.inspectFuncDirector != nil {
		mmDirector.mock.t.Fatalf("Inspect function is already set for DirectorMock.Director")
	}

	mmDirector.mock.inspectFuncDirector = f

	return mmDirector
}

// Return sets up results that will be returned by Director.Director
func (mmDirector *mDirectorMockDirector) Return() *DirectorMock {
	if mmDirector.mock.funcDirector != nil {
		mmDirector.mock.t.Fatalf("DirectorMock.Director mock is already set by Set")
	}

	if mmDirector.defaultExpectation == nil {
		mmDirector.defaultExpectation = &DirectorMockDirectorExpectation{mock: mmDirector.mock}
	}

	return mmDirector.mock
}

//Set uses given function f to mock the Director.Director method
func (mmDirector *mDirectorMockDirector) Set(f func(request *http.Request)) *DirectorMock {
	if mmDirector.defaultExpectation != nil {
		mmDirector.mock.t.Fatalf("Default expectation is already set for the Director.Director method")
	}

	if len(mmDirector.expectations) > 0 {
		mmDirector.mock.t.Fatalf("Some expectations are already set for the Director.Director method")
	}

	mmDirector.mock.funcDirector = f
	return mmDirector.mock
}

// Director implements Director
func (mmDirector *DirectorMock) Director(request *http.Request) {
	mm_atomic.AddUint64(&mmDirector.beforeDirectorCounter, 1)
	defer mm_atomic.AddUint64(&mmDirector.afterDirectorCounter, 1)

	if mmDirector.inspectFuncDirector != nil {
		mmDirector.inspectFuncDirector(request)
	}

	mm_params := &DirectorMockDirectorParams{request}

	// Record call args
	mmDirector.DirectorMock.mutex.Lock()
	mmDirector.DirectorMock.callArgs = append(mmDirector.DirectorMock.callArgs, mm_params)
	mmDirector.DirectorMock.mutex.Unlock()

	for _, e := range mmDirector.DirectorMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmDirector.DirectorMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDirector.DirectorMock.defaultExpectation.Counter, 1)
		mm_want := mmDirector.DirectorMock.defaultExpectation.params
		mm_got := DirectorMockDirectorParams{request}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDirector.t.Errorf("DirectorMock.Director got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmDirector.funcDirector != nil {
		mmDirector.funcDirector(request)
		return
	}
	mmDirector.t.Fatalf("Unexpected call to DirectorMock.Director. %v", request)

}

// DirectorAfterCounter returns a count of finished DirectorMock.Director invocations
func (mmDirector *DirectorMock) DirectorAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDirector.afterDirectorCounter)
}

// DirectorBeforeCounter returns a count of DirectorMock.Director invocations
func (mmDirector *DirectorMock) DirectorBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDirector.beforeDirectorCounter)
}

// Calls returns a list of arguments used in each call to DirectorMock.Director.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDirector *mDirectorMockDirector) Calls() []*DirectorMockDirectorParams {
	mmDirector.mutex.RLock()

	argCopy := make([]*DirectorMockDirectorParams, len(mmDirector.callArgs))
	copy(argCopy, mmDirector.callArgs)

	mmDirector.mutex.RUnlock()

	return argCopy
}

// MinimockDirectorDone returns true if the count of the Director invocations corresponds
// the number of defined expectations
func (m *DirectorMock) MinimockDirectorDone() bool {
	for _, e := range m.DirectorMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DirectorMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDirectorCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDirector != nil && mm_atomic.LoadUint64(&m.afterDirectorCounter) < 1 {
		return false
	}
	return true
}

// MinimockDirectorInspect logs each unmet expectation
func (m *DirectorMock) MinimockDirectorInspect() {
	for _, e := range m.DirectorMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DirectorMock.Director with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DirectorMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDirectorCounter) < 1 {
		if m.DirectorMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DirectorMock.Director")
		} else {
			m.t.Errorf("Expected call to DirectorMock.Director with params: %#v", *m.DirectorMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDirector != nil && mm_atomic.LoadUint64(&m.afterDirectorCounter) < 1 {
		m.t.Error("Expected call to DirectorMock.Director")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DirectorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDirectorInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DirectorMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *DirectorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDirectorDone()
}
