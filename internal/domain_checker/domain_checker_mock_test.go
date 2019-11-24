package domain_checker

// Code generated by http://github.com/gojuno/minimock (3.0.6). DO NOT EDIT.

//go:generate minimock -i github.com/rekby/lets-proxy2/internal/domain_checker.DomainChecker -o ./domain_checker_mock_test.go

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// DomainCheckerMock implements DomainChecker
type DomainCheckerMock struct {
	t minimock.Tester

	funcIsDomainAllowed          func(ctx context.Context, domain string) (b1 bool, err error)
	inspectFuncIsDomainAllowed   func(ctx context.Context, domain string)
	afterIsDomainAllowedCounter  uint64
	beforeIsDomainAllowedCounter uint64
	IsDomainAllowedMock          mDomainCheckerMockIsDomainAllowed
}

// NewDomainCheckerMock returns a mock for DomainChecker
func NewDomainCheckerMock(t minimock.Tester) *DomainCheckerMock {
	m := &DomainCheckerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.IsDomainAllowedMock = mDomainCheckerMockIsDomainAllowed{mock: m}
	m.IsDomainAllowedMock.callArgs = []*DomainCheckerMockIsDomainAllowedParams{}

	return m
}

type mDomainCheckerMockIsDomainAllowed struct {
	mock               *DomainCheckerMock
	defaultExpectation *DomainCheckerMockIsDomainAllowedExpectation
	expectations       []*DomainCheckerMockIsDomainAllowedExpectation

	callArgs []*DomainCheckerMockIsDomainAllowedParams
	mutex    sync.RWMutex
}

// DomainCheckerMockIsDomainAllowedExpectation specifies expectation struct of the DomainChecker.IsDomainAllowed
type DomainCheckerMockIsDomainAllowedExpectation struct {
	mock    *DomainCheckerMock
	params  *DomainCheckerMockIsDomainAllowedParams
	results *DomainCheckerMockIsDomainAllowedResults
	Counter uint64
}

// DomainCheckerMockIsDomainAllowedParams contains parameters of the DomainChecker.IsDomainAllowed
type DomainCheckerMockIsDomainAllowedParams struct {
	ctx    context.Context
	domain string
}

// DomainCheckerMockIsDomainAllowedResults contains results of the DomainChecker.IsDomainAllowed
type DomainCheckerMockIsDomainAllowedResults struct {
	b1  bool
	err error
}

// Expect sets up expected params for DomainChecker.IsDomainAllowed
func (mmIsDomainAllowed *mDomainCheckerMockIsDomainAllowed) Expect(ctx context.Context, domain string) *mDomainCheckerMockIsDomainAllowed {
	if mmIsDomainAllowed.mock.funcIsDomainAllowed != nil {
		mmIsDomainAllowed.mock.t.Fatalf("DomainCheckerMock.IsDomainAllowed mock is already set by Set")
	}

	if mmIsDomainAllowed.defaultExpectation == nil {
		mmIsDomainAllowed.defaultExpectation = &DomainCheckerMockIsDomainAllowedExpectation{}
	}

	mmIsDomainAllowed.defaultExpectation.params = &DomainCheckerMockIsDomainAllowedParams{ctx, domain}
	for _, e := range mmIsDomainAllowed.expectations {
		if minimock.Equal(e.params, mmIsDomainAllowed.defaultExpectation.params) {
			mmIsDomainAllowed.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmIsDomainAllowed.defaultExpectation.params)
		}
	}

	return mmIsDomainAllowed
}

// Inspect accepts an inspector function that has same arguments as the DomainChecker.IsDomainAllowed
func (mmIsDomainAllowed *mDomainCheckerMockIsDomainAllowed) Inspect(f func(ctx context.Context, domain string)) *mDomainCheckerMockIsDomainAllowed {
	if mmIsDomainAllowed.mock.inspectFuncIsDomainAllowed != nil {
		mmIsDomainAllowed.mock.t.Fatalf("Inspect function is already set for DomainCheckerMock.IsDomainAllowed")
	}

	mmIsDomainAllowed.mock.inspectFuncIsDomainAllowed = f

	return mmIsDomainAllowed
}

// Return sets up results that will be returned by DomainChecker.IsDomainAllowed
func (mmIsDomainAllowed *mDomainCheckerMockIsDomainAllowed) Return(b1 bool, err error) *DomainCheckerMock {
	if mmIsDomainAllowed.mock.funcIsDomainAllowed != nil {
		mmIsDomainAllowed.mock.t.Fatalf("DomainCheckerMock.IsDomainAllowed mock is already set by Set")
	}

	if mmIsDomainAllowed.defaultExpectation == nil {
		mmIsDomainAllowed.defaultExpectation = &DomainCheckerMockIsDomainAllowedExpectation{mock: mmIsDomainAllowed.mock}
	}
	mmIsDomainAllowed.defaultExpectation.results = &DomainCheckerMockIsDomainAllowedResults{b1, err}
	return mmIsDomainAllowed.mock
}

//Set uses given function f to mock the DomainChecker.IsDomainAllowed method
func (mmIsDomainAllowed *mDomainCheckerMockIsDomainAllowed) Set(f func(ctx context.Context, domain string) (b1 bool, err error)) *DomainCheckerMock {
	if mmIsDomainAllowed.defaultExpectation != nil {
		mmIsDomainAllowed.mock.t.Fatalf("Default expectation is already set for the DomainChecker.IsDomainAllowed method")
	}

	if len(mmIsDomainAllowed.expectations) > 0 {
		mmIsDomainAllowed.mock.t.Fatalf("Some expectations are already set for the DomainChecker.IsDomainAllowed method")
	}

	mmIsDomainAllowed.mock.funcIsDomainAllowed = f
	return mmIsDomainAllowed.mock
}

// When sets expectation for the DomainChecker.IsDomainAllowed which will trigger the result defined by the following
// Then helper
func (mmIsDomainAllowed *mDomainCheckerMockIsDomainAllowed) When(ctx context.Context, domain string) *DomainCheckerMockIsDomainAllowedExpectation {
	if mmIsDomainAllowed.mock.funcIsDomainAllowed != nil {
		mmIsDomainAllowed.mock.t.Fatalf("DomainCheckerMock.IsDomainAllowed mock is already set by Set")
	}

	expectation := &DomainCheckerMockIsDomainAllowedExpectation{
		mock:   mmIsDomainAllowed.mock,
		params: &DomainCheckerMockIsDomainAllowedParams{ctx, domain},
	}
	mmIsDomainAllowed.expectations = append(mmIsDomainAllowed.expectations, expectation)
	return expectation
}

// Then sets up DomainChecker.IsDomainAllowed return parameters for the expectation previously defined by the When method
func (e *DomainCheckerMockIsDomainAllowedExpectation) Then(b1 bool, err error) *DomainCheckerMock {
	e.results = &DomainCheckerMockIsDomainAllowedResults{b1, err}
	return e.mock
}

// IsDomainAllowed implements DomainChecker
func (mmIsDomainAllowed *DomainCheckerMock) IsDomainAllowed(ctx context.Context, domain string) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmIsDomainAllowed.beforeIsDomainAllowedCounter, 1)
	defer mm_atomic.AddUint64(&mmIsDomainAllowed.afterIsDomainAllowedCounter, 1)

	if mmIsDomainAllowed.inspectFuncIsDomainAllowed != nil {
		mmIsDomainAllowed.inspectFuncIsDomainAllowed(ctx, domain)
	}

	mm_params := &DomainCheckerMockIsDomainAllowedParams{ctx, domain}

	// Record call args
	mmIsDomainAllowed.IsDomainAllowedMock.mutex.Lock()
	mmIsDomainAllowed.IsDomainAllowedMock.callArgs = append(mmIsDomainAllowed.IsDomainAllowedMock.callArgs, mm_params)
	mmIsDomainAllowed.IsDomainAllowedMock.mutex.Unlock()

	for _, e := range mmIsDomainAllowed.IsDomainAllowedMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmIsDomainAllowed.IsDomainAllowedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIsDomainAllowed.IsDomainAllowedMock.defaultExpectation.Counter, 1)
		mm_want := mmIsDomainAllowed.IsDomainAllowedMock.defaultExpectation.params
		mm_got := DomainCheckerMockIsDomainAllowedParams{ctx, domain}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmIsDomainAllowed.t.Errorf("DomainCheckerMock.IsDomainAllowed got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmIsDomainAllowed.IsDomainAllowedMock.defaultExpectation.results
		if mm_results == nil {
			mmIsDomainAllowed.t.Fatal("No results are set for the DomainCheckerMock.IsDomainAllowed")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmIsDomainAllowed.funcIsDomainAllowed != nil {
		return mmIsDomainAllowed.funcIsDomainAllowed(ctx, domain)
	}
	mmIsDomainAllowed.t.Fatalf("Unexpected call to DomainCheckerMock.IsDomainAllowed. %v %v", ctx, domain)
	return
}

// IsDomainAllowedAfterCounter returns a count of finished DomainCheckerMock.IsDomainAllowed invocations
func (mmIsDomainAllowed *DomainCheckerMock) IsDomainAllowedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsDomainAllowed.afterIsDomainAllowedCounter)
}

// IsDomainAllowedBeforeCounter returns a count of DomainCheckerMock.IsDomainAllowed invocations
func (mmIsDomainAllowed *DomainCheckerMock) IsDomainAllowedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsDomainAllowed.beforeIsDomainAllowedCounter)
}

// Calls returns a list of arguments used in each call to DomainCheckerMock.IsDomainAllowed.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmIsDomainAllowed *mDomainCheckerMockIsDomainAllowed) Calls() []*DomainCheckerMockIsDomainAllowedParams {
	mmIsDomainAllowed.mutex.RLock()

	argCopy := make([]*DomainCheckerMockIsDomainAllowedParams, len(mmIsDomainAllowed.callArgs))
	copy(argCopy, mmIsDomainAllowed.callArgs)

	mmIsDomainAllowed.mutex.RUnlock()

	return argCopy
}

// MinimockIsDomainAllowedDone returns true if the count of the IsDomainAllowed invocations corresponds
// the number of defined expectations
func (m *DomainCheckerMock) MinimockIsDomainAllowedDone() bool {
	for _, e := range m.IsDomainAllowedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsDomainAllowedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsDomainAllowedCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsDomainAllowed != nil && mm_atomic.LoadUint64(&m.afterIsDomainAllowedCounter) < 1 {
		return false
	}
	return true
}

// MinimockIsDomainAllowedInspect logs each unmet expectation
func (m *DomainCheckerMock) MinimockIsDomainAllowedInspect() {
	for _, e := range m.IsDomainAllowedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DomainCheckerMock.IsDomainAllowed with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsDomainAllowedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsDomainAllowedCounter) < 1 {
		if m.IsDomainAllowedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DomainCheckerMock.IsDomainAllowed")
		} else {
			m.t.Errorf("Expected call to DomainCheckerMock.IsDomainAllowed with params: %#v", *m.IsDomainAllowedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsDomainAllowed != nil && mm_atomic.LoadUint64(&m.afterIsDomainAllowedCounter) < 1 {
		m.t.Error("Expected call to DomainCheckerMock.IsDomainAllowed")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DomainCheckerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockIsDomainAllowedInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DomainCheckerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *DomainCheckerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockIsDomainAllowedDone()
}
