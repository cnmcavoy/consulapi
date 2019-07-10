package consulapi

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// CandidateMock implements Candidate
type CandidateMock struct {
	t minimock.Tester

	funcParticipate          func(l1 LeadershipConfig, a1 AsLeaderFunc) (l2 LeaderSession, err error)
	afterParticipateCounter  uint64
	beforeParticipateCounter uint64
	ParticipateMock          mCandidateMockParticipate
}

// NewCandidateMock returns a mock for Candidate
func NewCandidateMock(t minimock.Tester) *CandidateMock {
	m := &CandidateMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ParticipateMock = mCandidateMockParticipate{mock: m}
	m.ParticipateMock.callArgs = []*CandidateMockParticipateParams{}

	return m
}

type mCandidateMockParticipate struct {
	mock               *CandidateMock
	defaultExpectation *CandidateMockParticipateExpectation
	expectations       []*CandidateMockParticipateExpectation

	callArgs []*CandidateMockParticipateParams
	mutex    sync.RWMutex
}

// CandidateMockParticipateExpectation specifies expectation struct of the Candidate.Participate
type CandidateMockParticipateExpectation struct {
	mock    *CandidateMock
	params  *CandidateMockParticipateParams
	results *CandidateMockParticipateResults
	Counter uint64
}

// CandidateMockParticipateParams contains parameters of the Candidate.Participate
type CandidateMockParticipateParams struct {
	l1 LeadershipConfig
	a1 AsLeaderFunc
}

// CandidateMockParticipateResults contains results of the Candidate.Participate
type CandidateMockParticipateResults struct {
	l2  LeaderSession
	err error
}

// Expect sets up expected params for Candidate.Participate
func (mmParticipate *mCandidateMockParticipate) Expect(l1 LeadershipConfig, a1 AsLeaderFunc) *mCandidateMockParticipate {
	if mmParticipate.mock.funcParticipate != nil {
		mmParticipate.mock.t.Fatalf("CandidateMock.Participate mock is already set by Set")
	}

	if mmParticipate.defaultExpectation == nil {
		mmParticipate.defaultExpectation = &CandidateMockParticipateExpectation{}
	}

	mmParticipate.defaultExpectation.params = &CandidateMockParticipateParams{l1, a1}
	for _, e := range mmParticipate.expectations {
		if minimock.Equal(e.params, mmParticipate.defaultExpectation.params) {
			mmParticipate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmParticipate.defaultExpectation.params)
		}
	}

	return mmParticipate
}

// Return sets up results that will be returned by Candidate.Participate
func (mmParticipate *mCandidateMockParticipate) Return(l2 LeaderSession, err error) *CandidateMock {
	if mmParticipate.mock.funcParticipate != nil {
		mmParticipate.mock.t.Fatalf("CandidateMock.Participate mock is already set by Set")
	}

	if mmParticipate.defaultExpectation == nil {
		mmParticipate.defaultExpectation = &CandidateMockParticipateExpectation{mock: mmParticipate.mock}
	}
	mmParticipate.defaultExpectation.results = &CandidateMockParticipateResults{l2, err}
	return mmParticipate.mock
}

//Set uses given function f to mock the Candidate.Participate method
func (mmParticipate *mCandidateMockParticipate) Set(f func(l1 LeadershipConfig, a1 AsLeaderFunc) (l2 LeaderSession, err error)) *CandidateMock {
	if mmParticipate.defaultExpectation != nil {
		mmParticipate.mock.t.Fatalf("Default expectation is already set for the Candidate.Participate method")
	}

	if len(mmParticipate.expectations) > 0 {
		mmParticipate.mock.t.Fatalf("Some expectations are already set for the Candidate.Participate method")
	}

	mmParticipate.mock.funcParticipate = f
	return mmParticipate.mock
}

// When sets expectation for the Candidate.Participate which will trigger the result defined by the following
// Then helper
func (mmParticipate *mCandidateMockParticipate) When(l1 LeadershipConfig, a1 AsLeaderFunc) *CandidateMockParticipateExpectation {
	if mmParticipate.mock.funcParticipate != nil {
		mmParticipate.mock.t.Fatalf("CandidateMock.Participate mock is already set by Set")
	}

	expectation := &CandidateMockParticipateExpectation{
		mock:   mmParticipate.mock,
		params: &CandidateMockParticipateParams{l1, a1},
	}
	mmParticipate.expectations = append(mmParticipate.expectations, expectation)
	return expectation
}

// Then sets up Candidate.Participate return parameters for the expectation previously defined by the When method
func (e *CandidateMockParticipateExpectation) Then(l2 LeaderSession, err error) *CandidateMock {
	e.results = &CandidateMockParticipateResults{l2, err}
	return e.mock
}

// Participate implements Candidate
func (mmParticipate *CandidateMock) Participate(l1 LeadershipConfig, a1 AsLeaderFunc) (l2 LeaderSession, err error) {
	mm_atomic.AddUint64(&mmParticipate.beforeParticipateCounter, 1)
	defer mm_atomic.AddUint64(&mmParticipate.afterParticipateCounter, 1)

	params := &CandidateMockParticipateParams{l1, a1}

	// Record call args
	mmParticipate.ParticipateMock.mutex.Lock()
	mmParticipate.ParticipateMock.callArgs = append(mmParticipate.ParticipateMock.callArgs, params)
	mmParticipate.ParticipateMock.mutex.Unlock()

	for _, e := range mmParticipate.ParticipateMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.l2, e.results.err
		}
	}

	if mmParticipate.ParticipateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmParticipate.ParticipateMock.defaultExpectation.Counter, 1)
		want := mmParticipate.ParticipateMock.defaultExpectation.params
		got := CandidateMockParticipateParams{l1, a1}
		if want != nil && !minimock.Equal(*want, got) {
			mmParticipate.t.Errorf("CandidateMock.Participate got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmParticipate.ParticipateMock.defaultExpectation.results
		if results == nil {
			mmParticipate.t.Fatal("No results are set for the CandidateMock.Participate")
		}
		return (*results).l2, (*results).err
	}
	if mmParticipate.funcParticipate != nil {
		return mmParticipate.funcParticipate(l1, a1)
	}
	mmParticipate.t.Fatalf("Unexpected call to CandidateMock.Participate. %v %v", l1, a1)
	return
}

// ParticipateAfterCounter returns a count of finished CandidateMock.Participate invocations
func (mmParticipate *CandidateMock) ParticipateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmParticipate.afterParticipateCounter)
}

// ParticipateBeforeCounter returns a count of CandidateMock.Participate invocations
func (mmParticipate *CandidateMock) ParticipateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmParticipate.beforeParticipateCounter)
}

// Calls returns a list of arguments used in each call to CandidateMock.Participate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmParticipate *mCandidateMockParticipate) Calls() []*CandidateMockParticipateParams {
	mmParticipate.mutex.RLock()

	argCopy := make([]*CandidateMockParticipateParams, len(mmParticipate.callArgs))
	copy(argCopy, mmParticipate.callArgs)

	mmParticipate.mutex.RUnlock()

	return argCopy
}

// MinimockParticipateDone returns true if the count of the Participate invocations corresponds
// the number of defined expectations
func (m *CandidateMock) MinimockParticipateDone() bool {
	for _, e := range m.ParticipateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ParticipateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterParticipateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcParticipate != nil && mm_atomic.LoadUint64(&m.afterParticipateCounter) < 1 {
		return false
	}
	return true
}

// MinimockParticipateInspect logs each unmet expectation
func (m *CandidateMock) MinimockParticipateInspect() {
	for _, e := range m.ParticipateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CandidateMock.Participate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ParticipateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterParticipateCounter) < 1 {
		if m.ParticipateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CandidateMock.Participate")
		} else {
			m.t.Errorf("Expected call to CandidateMock.Participate with params: %#v", *m.ParticipateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcParticipate != nil && mm_atomic.LoadUint64(&m.afterParticipateCounter) < 1 {
		m.t.Error("Expected call to CandidateMock.Participate")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CandidateMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockParticipateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CandidateMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CandidateMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockParticipateDone()
}