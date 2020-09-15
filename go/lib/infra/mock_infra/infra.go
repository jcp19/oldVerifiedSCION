// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/lib/infra (interfaces: Messenger,ResponseWriter,Verifier,Handler)

// Package mock_infra is a generated GoMock package.
package mock_infra

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	addr "github.com/scionproto/scion/go/lib/addr"
	ack "github.com/scionproto/scion/go/lib/ctrl/ack"
	ifid "github.com/scionproto/scion/go/lib/ctrl/ifid"
	path_mgmt "github.com/scionproto/scion/go/lib/ctrl/path_mgmt"
	infra "github.com/scionproto/scion/go/lib/infra"
	signed "github.com/scionproto/scion/go/lib/scrypto/signed"
	crypto "github.com/scionproto/scion/go/pkg/proto/crypto"
	net "net"
	reflect "reflect"
)

// MockMessenger is a mock of Messenger interface
type MockMessenger struct {
	ctrl     *gomock.Controller
	recorder *MockMessengerMockRecorder
}

// MockMessengerMockRecorder is the mock recorder for MockMessenger
type MockMessengerMockRecorder struct {
	mock *MockMessenger
}

// NewMockMessenger creates a new mock instance
func NewMockMessenger(ctrl *gomock.Controller) *MockMessenger {
	mock := &MockMessenger{ctrl: ctrl}
	mock.recorder = &MockMessengerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessenger) EXPECT() *MockMessengerMockRecorder {
	return m.recorder
}

// AddHandler mocks base method
func (m *MockMessenger) AddHandler(arg0 infra.MessageType, arg1 infra.Handler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHandler", arg0, arg1)
}

// AddHandler indicates an expected call of AddHandler
func (mr *MockMessengerMockRecorder) AddHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockMessenger)(nil).AddHandler), arg0, arg1)
}

// CloseServer mocks base method
func (m *MockMessenger) CloseServer() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseServer")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseServer indicates an expected call of CloseServer
func (mr *MockMessengerMockRecorder) CloseServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseServer", reflect.TypeOf((*MockMessenger)(nil).CloseServer))
}

// GetHPCfgs mocks base method
func (m *MockMessenger) GetHPCfgs(arg0 context.Context, arg1 *path_mgmt.HPCfgReq, arg2 net.Addr, arg3 uint64) (*path_mgmt.HPCfgReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHPCfgs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*path_mgmt.HPCfgReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHPCfgs indicates an expected call of GetHPCfgs
func (mr *MockMessengerMockRecorder) GetHPCfgs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHPCfgs", reflect.TypeOf((*MockMessenger)(nil).GetHPCfgs), arg0, arg1, arg2, arg3)
}

// GetHPSegs mocks base method
func (m *MockMessenger) GetHPSegs(arg0 context.Context, arg1 *path_mgmt.HPSegReq, arg2 net.Addr, arg3 uint64) (*path_mgmt.HPSegReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHPSegs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*path_mgmt.HPSegReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHPSegs indicates an expected call of GetHPSegs
func (mr *MockMessengerMockRecorder) GetHPSegs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHPSegs", reflect.TypeOf((*MockMessenger)(nil).GetHPSegs), arg0, arg1, arg2, arg3)
}

// ListenAndServe mocks base method
func (m *MockMessenger) ListenAndServe() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ListenAndServe")
}

// ListenAndServe indicates an expected call of ListenAndServe
func (mr *MockMessengerMockRecorder) ListenAndServe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockMessenger)(nil).ListenAndServe))
}

// SendAck mocks base method
func (m *MockMessenger) SendAck(arg0 context.Context, arg1 *ack.Ack, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAck", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAck indicates an expected call of SendAck
func (mr *MockMessengerMockRecorder) SendAck(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAck", reflect.TypeOf((*MockMessenger)(nil).SendAck), arg0, arg1, arg2, arg3)
}

// SendHPCfgReply mocks base method
func (m *MockMessenger) SendHPCfgReply(arg0 context.Context, arg1 *path_mgmt.HPCfgReply, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHPCfgReply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHPCfgReply indicates an expected call of SendHPCfgReply
func (mr *MockMessengerMockRecorder) SendHPCfgReply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHPCfgReply", reflect.TypeOf((*MockMessenger)(nil).SendHPCfgReply), arg0, arg1, arg2, arg3)
}

// SendHPSegReg mocks base method
func (m *MockMessenger) SendHPSegReg(arg0 context.Context, arg1 *path_mgmt.HPSegReg, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHPSegReg", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHPSegReg indicates an expected call of SendHPSegReg
func (mr *MockMessengerMockRecorder) SendHPSegReg(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHPSegReg", reflect.TypeOf((*MockMessenger)(nil).SendHPSegReg), arg0, arg1, arg2, arg3)
}

// SendHPSegReply mocks base method
func (m *MockMessenger) SendHPSegReply(arg0 context.Context, arg1 *path_mgmt.HPSegReply, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHPSegReply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHPSegReply indicates an expected call of SendHPSegReply
func (mr *MockMessengerMockRecorder) SendHPSegReply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHPSegReply", reflect.TypeOf((*MockMessenger)(nil).SendHPSegReply), arg0, arg1, arg2, arg3)
}

// SendIfId mocks base method
func (m *MockMessenger) SendIfId(arg0 context.Context, arg1 *ifid.IFID, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendIfId", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendIfId indicates an expected call of SendIfId
func (mr *MockMessengerMockRecorder) SendIfId(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendIfId", reflect.TypeOf((*MockMessenger)(nil).SendIfId), arg0, arg1, arg2, arg3)
}

// MockResponseWriter is a mock of ResponseWriter interface
type MockResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockResponseWriterMockRecorder
}

// MockResponseWriterMockRecorder is the mock recorder for MockResponseWriter
type MockResponseWriterMockRecorder struct {
	mock *MockResponseWriter
}

// NewMockResponseWriter creates a new mock instance
func NewMockResponseWriter(ctrl *gomock.Controller) *MockResponseWriter {
	mock := &MockResponseWriter{ctrl: ctrl}
	mock.recorder = &MockResponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponseWriter) EXPECT() *MockResponseWriterMockRecorder {
	return m.recorder
}

// SendAckReply mocks base method
func (m *MockResponseWriter) SendAckReply(arg0 context.Context, arg1 *ack.Ack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAckReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAckReply indicates an expected call of SendAckReply
func (mr *MockResponseWriterMockRecorder) SendAckReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAckReply", reflect.TypeOf((*MockResponseWriter)(nil).SendAckReply), arg0, arg1)
}

// SendHPCfgReply mocks base method
func (m *MockResponseWriter) SendHPCfgReply(arg0 context.Context, arg1 *path_mgmt.HPCfgReply) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHPCfgReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHPCfgReply indicates an expected call of SendHPCfgReply
func (mr *MockResponseWriterMockRecorder) SendHPCfgReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHPCfgReply", reflect.TypeOf((*MockResponseWriter)(nil).SendHPCfgReply), arg0, arg1)
}

// SendHPSegReply mocks base method
func (m *MockResponseWriter) SendHPSegReply(arg0 context.Context, arg1 *path_mgmt.HPSegReply) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHPSegReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHPSegReply indicates an expected call of SendHPSegReply
func (mr *MockResponseWriterMockRecorder) SendHPSegReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHPSegReply", reflect.TypeOf((*MockResponseWriter)(nil).SendHPSegReply), arg0, arg1)
}

// MockVerifier is a mock of Verifier interface
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method
func (m *MockVerifier) Verify(arg0 context.Context, arg1 *crypto.SignedMessage, arg2 ...[]byte) (*signed.Message, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Verify", varargs...)
	ret0, _ := ret[0].(*signed.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockVerifierMockRecorder) Verify(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifier)(nil).Verify), varargs...)
}

// WithIA mocks base method
func (m *MockVerifier) WithIA(arg0 addr.IA) infra.Verifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithIA", arg0)
	ret0, _ := ret[0].(infra.Verifier)
	return ret0
}

// WithIA indicates an expected call of WithIA
func (mr *MockVerifierMockRecorder) WithIA(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithIA", reflect.TypeOf((*MockVerifier)(nil).WithIA), arg0)
}

// WithServer mocks base method
func (m *MockVerifier) WithServer(arg0 net.Addr) infra.Verifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithServer", arg0)
	ret0, _ := ret[0].(infra.Verifier)
	return ret0
}

// WithServer indicates an expected call of WithServer
func (mr *MockVerifierMockRecorder) WithServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithServer", reflect.TypeOf((*MockVerifier)(nil).WithServer), arg0)
}

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockHandler) Handle(arg0 *infra.Request) *infra.HandlerResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0)
	ret0, _ := ret[0].(*infra.HandlerResult)
	return ret0
}

// Handle indicates an expected call of Handle
func (mr *MockHandlerMockRecorder) Handle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), arg0)
}
