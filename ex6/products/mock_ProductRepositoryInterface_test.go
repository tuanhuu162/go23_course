// Code generated by mockery v2.23.4. DO NOT EDIT.

package products

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/tuanhuu162/go23_course/ex6/models"
)

// MockProductRepositoryInterface is an autogenerated mock type for the ProductRepositoryInterface type
type MockProductRepositoryInterface struct {
	mock.Mock
}

type MockProductRepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProductRepositoryInterface) EXPECT() *MockProductRepositoryInterface_Expecter {
	return &MockProductRepositoryInterface_Expecter{mock: &_m.Mock}
}

// AddProduct provides a mock function with given fields: product
func (_m *MockProductRepositoryInterface) AddProduct(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductRepositoryInterface_AddProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddProduct'
type MockProductRepositoryInterface_AddProduct_Call struct {
	*mock.Call
}

// AddProduct is a helper method to define mock.On call
//   - product models.Product
func (_e *MockProductRepositoryInterface_Expecter) AddProduct(product interface{}) *MockProductRepositoryInterface_AddProduct_Call {
	return &MockProductRepositoryInterface_AddProduct_Call{Call: _e.mock.On("AddProduct", product)}
}

func (_c *MockProductRepositoryInterface_AddProduct_Call) Run(run func(product models.Product)) *MockProductRepositoryInterface_AddProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Product))
	})
	return _c
}

func (_c *MockProductRepositoryInterface_AddProduct_Call) Return(_a0 error) *MockProductRepositoryInterface_AddProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductRepositoryInterface_AddProduct_Call) RunAndReturn(run func(models.Product) error) *MockProductRepositoryInterface_AddProduct_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProduct provides a mock function with given fields: product_id
func (_m *MockProductRepositoryInterface) DeleteProduct(product_id uint64) error {
	ret := _m.Called(product_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(product_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductRepositoryInterface_DeleteProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProduct'
type MockProductRepositoryInterface_DeleteProduct_Call struct {
	*mock.Call
}

// DeleteProduct is a helper method to define mock.On call
//   - product_id uint64
func (_e *MockProductRepositoryInterface_Expecter) DeleteProduct(product_id interface{}) *MockProductRepositoryInterface_DeleteProduct_Call {
	return &MockProductRepositoryInterface_DeleteProduct_Call{Call: _e.mock.On("DeleteProduct", product_id)}
}

func (_c *MockProductRepositoryInterface_DeleteProduct_Call) Run(run func(product_id uint64)) *MockProductRepositoryInterface_DeleteProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockProductRepositoryInterface_DeleteProduct_Call) Return(_a0 error) *MockProductRepositoryInterface_DeleteProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductRepositoryInterface_DeleteProduct_Call) RunAndReturn(run func(uint64) error) *MockProductRepositoryInterface_DeleteProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllProducts provides a mock function with given fields:
func (_m *MockProductRepositoryInterface) GetAllProducts() []models.Product {
	ret := _m.Called()

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func() []models.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	return r0
}

// MockProductRepositoryInterface_GetAllProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllProducts'
type MockProductRepositoryInterface_GetAllProducts_Call struct {
	*mock.Call
}

// GetAllProducts is a helper method to define mock.On call
func (_e *MockProductRepositoryInterface_Expecter) GetAllProducts() *MockProductRepositoryInterface_GetAllProducts_Call {
	return &MockProductRepositoryInterface_GetAllProducts_Call{Call: _e.mock.On("GetAllProducts")}
}

func (_c *MockProductRepositoryInterface_GetAllProducts_Call) Run(run func()) *MockProductRepositoryInterface_GetAllProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProductRepositoryInterface_GetAllProducts_Call) Return(_a0 []models.Product) *MockProductRepositoryInterface_GetAllProducts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductRepositoryInterface_GetAllProducts_Call) RunAndReturn(run func() []models.Product) *MockProductRepositoryInterface_GetAllProducts_Call {
	_c.Call.Return(run)
	return _c
}

// GetProduct provides a mock function with given fields: product_id
func (_m *MockProductRepositoryInterface) GetProduct(product_id uint) (models.Product, error) {
	ret := _m.Called(product_id)

	var r0 models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (models.Product, error)); ok {
		return rf(product_id)
	}
	if rf, ok := ret.Get(0).(func(uint) models.Product); ok {
		r0 = rf(product_id)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(product_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductRepositoryInterface_GetProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProduct'
type MockProductRepositoryInterface_GetProduct_Call struct {
	*mock.Call
}

// GetProduct is a helper method to define mock.On call
//   - product_id uint
func (_e *MockProductRepositoryInterface_Expecter) GetProduct(product_id interface{}) *MockProductRepositoryInterface_GetProduct_Call {
	return &MockProductRepositoryInterface_GetProduct_Call{Call: _e.mock.On("GetProduct", product_id)}
}

func (_c *MockProductRepositoryInterface_GetProduct_Call) Run(run func(product_id uint)) *MockProductRepositoryInterface_GetProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockProductRepositoryInterface_GetProduct_Call) Return(_a0 models.Product, _a1 error) *MockProductRepositoryInterface_GetProduct_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductRepositoryInterface_GetProduct_Call) RunAndReturn(run func(uint) (models.Product, error)) *MockProductRepositoryInterface_GetProduct_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateProduct provides a mock function with given fields: product
func (_m *MockProductRepositoryInterface) UpdateProduct(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductRepositoryInterface_UpdateProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProduct'
type MockProductRepositoryInterface_UpdateProduct_Call struct {
	*mock.Call
}

// UpdateProduct is a helper method to define mock.On call
//   - product models.Product
func (_e *MockProductRepositoryInterface_Expecter) UpdateProduct(product interface{}) *MockProductRepositoryInterface_UpdateProduct_Call {
	return &MockProductRepositoryInterface_UpdateProduct_Call{Call: _e.mock.On("UpdateProduct", product)}
}

func (_c *MockProductRepositoryInterface_UpdateProduct_Call) Run(run func(product models.Product)) *MockProductRepositoryInterface_UpdateProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Product))
	})
	return _c
}

func (_c *MockProductRepositoryInterface_UpdateProduct_Call) Return(_a0 error) *MockProductRepositoryInterface_UpdateProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductRepositoryInterface_UpdateProduct_Call) RunAndReturn(run func(models.Product) error) *MockProductRepositoryInterface_UpdateProduct_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProductRepositoryInterface creates a new instance of MockProductRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProductRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProductRepositoryInterface {
	mock := &MockProductRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
