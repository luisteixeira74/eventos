package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	PayLoad interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayLoad() interface{} {
	return e.PayLoad
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	// Handle the event
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		// Initialize the handler
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		// Initialize the handler
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		// Initialize the handler
		ID: 3,
	}
	suite.event = TestEvent{
		Name:    "test",
		PayLoad: "payload 1",
	}
	suite.event2 = TestEvent{
		Name:    "test2",
		PayLoad: "payload 2",
	}
}

func (suite *EventDispatcherTestSuite) Test_EventDispatcher_Register() {
	// Test that we can register an event handler

	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcherRegisterWithSameHandler() {
	// Test that we can register the same handler for different events

	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Test that we can clear all handlers

	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
	// Clear all handlers
	suite.eventDispatcher.Clear()
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// Test that we can check if a handler is registered for an event

	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	eH := &MockEventHandler{}
	eH.On("Handle", &suite.event)
	eH2 := &MockEventHandler{}
	eH2.On("Handle", &suite.event)

	suite.eventDispatcher.Register(suite.event.GetName(), eH)
	suite.eventDispatcher.Register(suite.event.GetName(), eH2)

	suite.eventDispatcher.Dispatch(&suite.event)

	eH.AssertExpectations(suite.T())                // Assert that the expected method was called on eH
	eH2.AssertExpectations(suite.T())               // Assert that the expected method was called on eH2
	eH.AssertNumberOfCalls(suite.T(), "Handle", 1)  // Assert that the method was called once on eH
	eH2.AssertNumberOfCalls(suite.T(), "Handle", 1) // Assert that the method was called once on eH2
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Remove() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_HandlerProcessesPayloadCorrectly() {
	mockHandler := new(MockEventHandler)
	testPayload := "msg importante"

	event := &TestEvent{
		Name:    "EventoTeste",
		PayLoad: testPayload,
	}

	mockHandler.On("Handle", event).Once()

	suite.eventDispatcher.Register("EventoTeste", mockHandler)

	err := suite.eventDispatcher.Dispatch(event)
	suite.NoError(err)

	mockHandler.AssertExpectations(suite.T())
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_FromSimulatedMessage() {
	mockHandler := new(MockEventHandler)

	eventName := "EventoTeste"
	payload := "payload do RabbitMQ"
	event := &TestEvent{Name: eventName, PayLoad: payload}

	mockHandler.On("Handle", event).Once()
	suite.eventDispatcher.Register(eventName, mockHandler)

	err := suite.eventDispatcher.Dispatch(event)
	suite.NoError(err)

	mockHandler.AssertExpectations(suite.T())
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_NoHandlerRegistered() {
	event := &TestEvent{
		Name:    "EventoSemHandler",
		PayLoad: "qualquer coisa",
	}

	err := suite.eventDispatcher.Dispatch(event)
	suite.Error(err)
	suite.Equal(ErrNoHandlerRegistered, err)
}

type SpyHandler struct {
	LastPayload interface{}
}

func (s *SpyHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	s.LastPayload = event.GetPayLoad()
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_WithRealHandlerLikeUpdateDB() {
	spy := &SpyHandler{}
	event := &TestEvent{Name: "EventoReal", PayLoad: "valor importante"}

	suite.eventDispatcher.Register("EventoReal", spy)
	suite.eventDispatcher.Dispatch(event)

	suite.Equal("valor importante", spy.LastPayload)
}
