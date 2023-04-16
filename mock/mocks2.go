package mock_models

import "gorm.io/gorm"

func (m *MockSkuMock) First(arg0 interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	return m.ctrl.Call(m, "First", arg0).(*gorm.DB)
}

func (m *MockSkuMock) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	return nil // return your mock gorm.DB instance here
}
