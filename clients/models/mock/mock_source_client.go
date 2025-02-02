// Code generated by MockGen. DO NOT EDIT.
// Source: source_client.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	os "os"
	reflect "reflect"

	models "github.com/fastenhealth/fasten-sources/clients/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
	gomock "github.com/golang/mock/gomock"
)

// MockSourceClient is a mock of SourceClient interface.
type MockSourceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSourceClientMockRecorder
}

// MockSourceClientMockRecorder is the mock recorder for MockSourceClient.
type MockSourceClientMockRecorder struct {
	mock *MockSourceClient
}

// NewMockSourceClient creates a new mock instance.
func NewMockSourceClient(ctrl *gomock.Controller) *MockSourceClient {
	mock := &MockSourceClient{ctrl: ctrl}
	mock.recorder = &MockSourceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceClient) EXPECT() *MockSourceClientMockRecorder {
	return m.recorder
}

// ExtractPatientId mocks base method.
func (m *MockSourceClient) ExtractPatientId(bundleFile *os.File) (string, pkg.FhirVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractPatientId", bundleFile)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(pkg.FhirVersion)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExtractPatientId indicates an expected call of ExtractPatientId.
func (mr *MockSourceClientMockRecorder) ExtractPatientId(bundleFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractPatientId", reflect.TypeOf((*MockSourceClient)(nil).ExtractPatientId), bundleFile)
}

// GetRequest mocks base method.
func (m *MockSourceClient) GetRequest(resourceSubpath string, decodeModelPtr interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", resourceSubpath, decodeModelPtr)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockSourceClientMockRecorder) GetRequest(resourceSubpath, decodeModelPtr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockSourceClient)(nil).GetRequest), resourceSubpath, decodeModelPtr)
}

// GetResourceBundle mocks base method.
func (m *MockSourceClient) GetResourceBundle(relativeResourcePath string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceBundle", relativeResourcePath)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceBundle indicates an expected call of GetResourceBundle.
func (mr *MockSourceClientMockRecorder) GetResourceBundle(relativeResourcePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceBundle", reflect.TypeOf((*MockSourceClient)(nil).GetResourceBundle), relativeResourcePath)
}

// GetSourceCredential mocks base method.
func (m *MockSourceClient) GetSourceCredential() models.SourceCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceCredential")
	ret0, _ := ret[0].(models.SourceCredential)
	return ret0
}

// GetSourceCredential indicates an expected call of GetSourceCredential.
func (mr *MockSourceClientMockRecorder) GetSourceCredential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceCredential", reflect.TypeOf((*MockSourceClient)(nil).GetSourceCredential))
}

// GetUsCoreResources mocks base method.
func (m *MockSourceClient) GetUsCoreResources() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsCoreResources")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetUsCoreResources indicates an expected call of GetUsCoreResources.
func (mr *MockSourceClientMockRecorder) GetUsCoreResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsCoreResources", reflect.TypeOf((*MockSourceClient)(nil).GetUsCoreResources))
}

// SyncAll mocks base method.
func (m *MockSourceClient) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncAll", db)
	ret0, _ := ret[0].(models.UpsertSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncAll indicates an expected call of SyncAll.
func (mr *MockSourceClientMockRecorder) SyncAll(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncAll", reflect.TypeOf((*MockSourceClient)(nil).SyncAll), db)
}

// SyncAllBundle mocks base method.
func (m *MockSourceClient) SyncAllBundle(db models.DatabaseRepository, bundleFile *os.File, bundleFhirVersion pkg.FhirVersion) (models.UpsertSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncAllBundle", db, bundleFile, bundleFhirVersion)
	ret0, _ := ret[0].(models.UpsertSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncAllBundle indicates an expected call of SyncAllBundle.
func (mr *MockSourceClientMockRecorder) SyncAllBundle(db, bundleFile, bundleFhirVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncAllBundle", reflect.TypeOf((*MockSourceClient)(nil).SyncAllBundle), db, bundleFile, bundleFhirVersion)
}

// SyncAllByPatientEverythingBundle mocks base method.
func (m *MockSourceClient) SyncAllByPatientEverythingBundle(db models.DatabaseRepository, bundleModel interface{}) (models.UpsertSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncAllByPatientEverythingBundle", db, bundleModel)
	ret0, _ := ret[0].(models.UpsertSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncAllByPatientEverythingBundle indicates an expected call of SyncAllByPatientEverythingBundle.
func (mr *MockSourceClientMockRecorder) SyncAllByPatientEverythingBundle(db, bundleModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncAllByPatientEverythingBundle", reflect.TypeOf((*MockSourceClient)(nil).SyncAllByPatientEverythingBundle), db, bundleModel)
}

// SyncAllByResourceName mocks base method.
func (m *MockSourceClient) SyncAllByResourceName(db models.DatabaseRepository, resourceNames []string) (models.UpsertSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncAllByResourceName", db, resourceNames)
	ret0, _ := ret[0].(models.UpsertSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncAllByResourceName indicates an expected call of SyncAllByResourceName.
func (mr *MockSourceClientMockRecorder) SyncAllByResourceName(db, resourceNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncAllByResourceName", reflect.TypeOf((*MockSourceClient)(nil).SyncAllByResourceName), db, resourceNames)
}
