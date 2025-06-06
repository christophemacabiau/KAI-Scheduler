// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/binder/plugins/k8s-plugins/common/interface.go
//
// Generated by this command:
//
//	mockgen -source=pkg/binder/plugins/k8s-plugins/common/interface.go -destination=pkg/binder/plugins/k8s-plugins/common/mock/mock_common_plugins.go -package=mock_common_plugins
//

// Package mock_common_plugins is a generated GoMock package.
package mock_common_plugins

import (
	context "context"
	reflect "reflect"

	v1alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v1alpha2"
	common "github.com/NVIDIA/KAI-scheduler/pkg/binder/plugins/k8s-plugins/common"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockK8sPlugin is a mock of K8sPlugin interface.
type MockK8sPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockK8sPluginMockRecorder
	isgomock struct{}
}

// MockK8sPluginMockRecorder is the mock recorder for MockK8sPlugin.
type MockK8sPluginMockRecorder struct {
	mock *MockK8sPlugin
}

// NewMockK8sPlugin creates a new mock instance.
func NewMockK8sPlugin(ctrl *gomock.Controller) *MockK8sPlugin {
	mock := &MockK8sPlugin{ctrl: ctrl}
	mock.recorder = &MockK8sPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8sPlugin) EXPECT() *MockK8sPluginMockRecorder {
	return m.recorder
}

// Allocate mocks base method.
func (m *MockK8sPlugin) Allocate(ctx context.Context, pod *v1.Pod, hostname string, state common.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allocate", ctx, pod, hostname, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// Allocate indicates an expected call of Allocate.
func (mr *MockK8sPluginMockRecorder) Allocate(ctx, pod, hostname, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allocate", reflect.TypeOf((*MockK8sPlugin)(nil).Allocate), ctx, pod, hostname, state)
}

// Bind mocks base method.
func (m *MockK8sPlugin) Bind(ctx context.Context, pod *v1.Pod, request *v1alpha2.BindRequest, state common.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", ctx, pod, request, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockK8sPluginMockRecorder) Bind(ctx, pod, request, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockK8sPlugin)(nil).Bind), ctx, pod, request, state)
}

// Filter mocks base method.
func (m *MockK8sPlugin) Filter(ctx context.Context, pod *v1.Pod, node *v1.Node, state common.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", ctx, pod, node, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockK8sPluginMockRecorder) Filter(ctx, pod, node, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockK8sPlugin)(nil).Filter), ctx, pod, node, state)
}

// IsRelevant mocks base method.
func (m *MockK8sPlugin) IsRelevant(pod *v1.Pod) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRelevant", pod)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRelevant indicates an expected call of IsRelevant.
func (mr *MockK8sPluginMockRecorder) IsRelevant(pod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRelevant", reflect.TypeOf((*MockK8sPlugin)(nil).IsRelevant), pod)
}

// Name mocks base method.
func (m *MockK8sPlugin) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockK8sPluginMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockK8sPlugin)(nil).Name))
}

// PostBind mocks base method.
func (m *MockK8sPlugin) PostBind(ctx context.Context, pod *v1.Pod, hostname string, state common.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostBind", ctx, pod, hostname, state)
}

// PostBind indicates an expected call of PostBind.
func (mr *MockK8sPluginMockRecorder) PostBind(ctx, pod, hostname, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBind", reflect.TypeOf((*MockK8sPlugin)(nil).PostBind), ctx, pod, hostname, state)
}

// PreFilter mocks base method.
func (m *MockK8sPlugin) PreFilter(ctx context.Context, pod *v1.Pod, state common.State) (error, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreFilter", ctx, pod, state)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// PreFilter indicates an expected call of PreFilter.
func (mr *MockK8sPluginMockRecorder) PreFilter(ctx, pod, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreFilter", reflect.TypeOf((*MockK8sPlugin)(nil).PreFilter), ctx, pod, state)
}

// UnAllocate mocks base method.
func (m *MockK8sPlugin) UnAllocate(ctx context.Context, pod *v1.Pod, hostname string, state common.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnAllocate", ctx, pod, hostname, state)
}

// UnAllocate indicates an expected call of UnAllocate.
func (mr *MockK8sPluginMockRecorder) UnAllocate(ctx, pod, hostname, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnAllocate", reflect.TypeOf((*MockK8sPlugin)(nil).UnAllocate), ctx, pod, hostname, state)
}
