// Code generated by mockery v1.0.0
package charontest

import charonrpc "github.com/piotrkowalczuk/charon/charonrpc"
import context "context"
import mock "github.com/stretchr/testify/mock"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// UserManagerServer is an autogenerated mock type for the UserManagerServer type
type UserManagerServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Create(_a0 context.Context, _a1 *charonrpc.CreateUserRequest) (*charonrpc.CreateUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.CreateUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.CreateUserRequest) *charonrpc.CreateUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.CreateUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.CreateUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Delete(_a0 context.Context, _a1 *charonrpc.DeleteUserRequest) (*wrappers.BoolValue, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *wrappers.BoolValue
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.DeleteUserRequest) *wrappers.BoolValue); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrappers.BoolValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.DeleteUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Get(_a0 context.Context, _a1 *charonrpc.GetUserRequest) (*charonrpc.GetUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.GetUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.GetUserRequest) *charonrpc.GetUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.GetUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.GetUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) List(_a0 context.Context, _a1 *charonrpc.ListUsersRequest) (*charonrpc.ListUsersResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.ListUsersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.ListUsersRequest) *charonrpc.ListUsersResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.ListUsersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.ListUsersRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) ListGroups(_a0 context.Context, _a1 *charonrpc.ListUserGroupsRequest) (*charonrpc.ListUserGroupsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.ListUserGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.ListUserGroupsRequest) *charonrpc.ListUserGroupsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.ListUserGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.ListUserGroupsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPermissions provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) ListPermissions(_a0 context.Context, _a1 *charonrpc.ListUserPermissionsRequest) (*charonrpc.ListUserPermissionsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.ListUserPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.ListUserPermissionsRequest) *charonrpc.ListUserPermissionsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.ListUserPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.ListUserPermissionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Modify(_a0 context.Context, _a1 *charonrpc.ModifyUserRequest) (*charonrpc.ModifyUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.ModifyUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.ModifyUserRequest) *charonrpc.ModifyUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.ModifyUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.ModifyUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetGroups provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) SetGroups(_a0 context.Context, _a1 *charonrpc.SetUserGroupsRequest) (*charonrpc.SetUserGroupsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.SetUserGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.SetUserGroupsRequest) *charonrpc.SetUserGroupsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.SetUserGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.SetUserGroupsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPermissions provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) SetPermissions(_a0 context.Context, _a1 *charonrpc.SetUserPermissionsRequest) (*charonrpc.SetUserPermissionsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charonrpc.SetUserPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.SetUserPermissionsRequest) *charonrpc.SetUserPermissionsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.SetUserPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.SetUserPermissionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}