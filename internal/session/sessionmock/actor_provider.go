// Code generated by mockery v1.0.0. DO NOT EDIT.

package sessionmock

import context "context"
import mock "github.com/stretchr/testify/mock"
import session "github.com/piotrkowalczuk/charon/internal/session"

// ActorProvider is an autogenerated mock type for the ActorProvider type
type ActorProvider struct {
	mock.Mock
}

// Actor provides a mock function with given fields: _a0
func (_m *ActorProvider) Actor(_a0 context.Context) (*session.Actor, error) {
	ret := _m.Called(_a0)

	var r0 *session.Actor
	if rf, ok := ret.Get(0).(func(context.Context) *session.Actor); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.Actor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}