package pb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockMovieSearchClient is a mock of MovieSearchClient interface.
type MockMovieSearchClient struct {
	ctrl     *gomock.Controller
	recorder *MockMovieSearchClientMockRecorder
}

// MockMovieSearchClientMockRecorder is the mock recorder for MockMovieSearchClient.
type MockMovieSearchClientMockRecorder struct {
	mock *MockMovieSearchClient
}

// NewMockMovieSearchClient creates a new mock instance.
func NewMockMovieSearchClient(ctrl *gomock.Controller) *MockMovieSearchClient {
	mock := &MockMovieSearchClient{ctrl: ctrl}
	mock.recorder = &MockMovieSearchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieSearchClient) EXPECT() *MockMovieSearchClientMockRecorder {
	return m.recorder
}

// SearchMovie mocks base method.
func (m *MockMovieSearchClient) SearchMovie(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchMovie", varargs...)
	ret0, _ := ret[0].(*SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockMovieSearchClientMockRecorder) SearchMovie(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockMovieSearchClient)(nil).SearchMovie), varargs...)
}

// MockMovieSearchServer is a mock of MovieSearchServer interface.
type MockMovieSearchServer struct {
	ctrl     *gomock.Controller
	recorder *MockMovieSearchServerMockRecorder
}

// MockMovieSearchServerMockRecorder is the mock recorder for MockMovieSearchServer.
type MockMovieSearchServerMockRecorder struct {
	mock *MockMovieSearchServer
}

// NewMockMovieSearchServer creates a new mock instance.
func NewMockMovieSearchServer(ctrl *gomock.Controller) *MockMovieSearchServer {
	mock := &MockMovieSearchServer{ctrl: ctrl}
	mock.recorder = &MockMovieSearchServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieSearchServer) EXPECT() *MockMovieSearchServerMockRecorder {
	return m.recorder
}

// SearchMovie mocks base method.
func (m *MockMovieSearchServer) SearchMovie(arg0 context.Context, arg1 *SearchRequest) (*SearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovie", arg0, arg1)
	ret0, _ := ret[0].(*SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockMovieSearchServerMockRecorder) SearchMovie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockMovieSearchServer)(nil).SearchMovie), arg0, arg1)
}
