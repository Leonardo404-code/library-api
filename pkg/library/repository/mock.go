// Code generated by mockery v2.38.0. DO NOT EDIT.

package repository

import (
	library "library-api/pkg/library"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// Mock is an autogenerated mock type for the Repository type
type Mock struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: bookInfo, bookFile, uploadToBucket
func (_m *Mock) CreateBook(bookInfo *library.Book, bookFile multipart.File, uploadToBucket func() error) error {
	ret := _m.Called(bookInfo, bookFile, uploadToBucket)

	if len(ret) == 0 {
		panic("no return value specified for CreateBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*library.Book, multipart.File, func() error) error); ok {
		r0 = rf(bookInfo, bookFile, uploadToBucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: filter, deleteFromBucket
func (_m *Mock) DeleteBook(filter *library.Filter, deleteFromBucket func() error) error {
	ret := _m.Called(filter, deleteFromBucket)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*library.Filter, func() error) error); ok {
		r0 = rf(filter, deleteFromBucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: filter
func (_m *Mock) Search(filter *library.Filter) ([]*library.Book, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 []*library.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(*library.Filter) ([]*library.Book, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(*library.Filter) []*library.Book); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*library.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(*library.Filter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMock creates a new instance of Mock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMock(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Mock {
	mock := &Mock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
