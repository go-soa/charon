package charond

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/piotrkowalczuk/charon/internal/model"
	"github.com/piotrkowalczuk/charon/internal/password"
	"github.com/piotrkowalczuk/charon/internal/service"
	charonrpc "github.com/piotrkowalczuk/charon/pb/rpc/charond/v1"
	"github.com/piotrkowalczuk/mnemosyne/mnemosynerpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type rpcServer struct {
	opts               DaemonOpts
	logger             *zap.Logger
	session            mnemosynerpc.SessionManagerClient
	passwordHasher     password.Hasher
	permissionRegistry model.PermissionRegistry
	repository         repositories
}

type auth struct {
	*actorHandler
	*loginHandler
	*logoutHandler
	*isGrantedHandler
	*isAuthenticatedHandler
	*belongsToHandler

	charonrpc.UnimplementedAuthServer
}

func newAuth(server *rpcServer) *auth {
	return &auth{
		actorHandler:           &actorHandler{handler: newHandler(server)},
		belongsToHandler:       &belongsToHandler{handler: newHandler(server)},
		isGrantedHandler:       &isGrantedHandler{handler: newHandler(server)},
		isAuthenticatedHandler: &isAuthenticatedHandler{handler: newHandler(server)},
		loginHandler: &loginHandler{
			handler: newHandler(server),
			userFinderFactory: &service.UserFinderFactory{
				Hasher:                 server.passwordHasher,
				UserRepository:         server.repository.user,
				RefreshTokenRepository: server.repository.refreshToken,
			},
		},
		logoutHandler: &logoutHandler{handler: newHandler(server)},
	}
}

func (a auth) Actor(ctx context.Context, req *wrapperspb.StringValue) (*charonrpc.ActorResponse, error) {
	return a.actorHandler.Actor(ctx, req)
}

func (a auth) Login(ctx context.Context, req *charonrpc.LoginRequest) (*wrapperspb.StringValue, error) {
	return a.loginHandler.Login(ctx, req)
}

func (a auth) Logout(ctx context.Context, req *charonrpc.LogoutRequest) (*empty.Empty, error) {
	return a.logoutHandler.Logout(ctx, req)
}

func (a auth) IsGranted(ctx context.Context, req *charonrpc.IsGrantedRequest) (*wrappers.BoolValue, error) {
	return a.isGrantedHandler.IsGranted(ctx, req)
}

func (a auth) IsAuthenticated(ctx context.Context, req *charonrpc.IsAuthenticatedRequest) (*wrappers.BoolValue, error) {
	return a.isAuthenticatedHandler.IsAuthenticated(ctx, req)
}

func (a auth) BelongsTo(ctx context.Context, req *charonrpc.BelongsToRequest) (*wrappers.BoolValue, error) {
	return a.belongsToHandler.BelongsTo(ctx, req)
}

type userManager struct {
	*createUserHandler
	*deleteUserHandler
	*getUserHandler
	*listUserGroupsHandler
	*listUserPermissionsHandler
	*listUsersHandler
	*modifyUserHandler
	*setUserGroupsHandler
	*setUserPermissionsHandler

	charonrpc.UnimplementedUserManagerServer
}

func newUserManager(server *rpcServer) *userManager {
	return &userManager{
		createUserHandler:          &createUserHandler{handler: newHandler(server), hasher: server.passwordHasher},
		deleteUserHandler:          &deleteUserHandler{handler: newHandler(server)},
		getUserHandler:             &getUserHandler{handler: newHandler(server)},
		listUserGroupsHandler:      &listUserGroupsHandler{handler: newHandler(server)},
		listUserPermissionsHandler: &listUserPermissionsHandler{handler: newHandler(server)},
		listUsersHandler:           &listUsersHandler{handler: newHandler(server)},
		modifyUserHandler:          &modifyUserHandler{handler: newHandler(server)},
		setUserGroupsHandler:       &setUserGroupsHandler{handler: newHandler(server)},
		setUserPermissionsHandler:  &setUserPermissionsHandler{handler: newHandler(server)},
	}
}

func (a userManager) Create(ctx context.Context, req *charonrpc.CreateUserRequest) (*charonrpc.CreateUserResponse, error) {
	return a.createUserHandler.Create(ctx, req)
}

func (a userManager) Delete(ctx context.Context, req *charonrpc.DeleteUserRequest) (*wrappers.BoolValue, error) {
	return a.deleteUserHandler.Delete(ctx, req)
}

func (a userManager) Get(ctx context.Context, req *charonrpc.GetUserRequest) (*charonrpc.GetUserResponse, error) {
	return a.getUserHandler.Get(ctx, req)
}

func (a userManager) ListGroups(ctx context.Context, req *charonrpc.ListUserGroupsRequest) (*charonrpc.ListUserGroupsResponse, error) {
	return a.listUserGroupsHandler.ListGroups(ctx, req)
}

func (a userManager) ListPermissions(ctx context.Context, req *charonrpc.ListUserPermissionsRequest) (*charonrpc.ListUserPermissionsResponse, error) {
	return a.listUserPermissionsHandler.ListPermissions(ctx, req)
}

func (a userManager) List(ctx context.Context, req *charonrpc.ListUsersRequest) (*charonrpc.ListUsersResponse, error) {
	return a.listUsersHandler.List(ctx, req)
}

func (a userManager) Modify(ctx context.Context, req *charonrpc.ModifyUserRequest) (*charonrpc.ModifyUserResponse, error) {
	return a.modifyUserHandler.Modify(ctx, req)
}

func (a userManager) SetGroups(ctx context.Context, req *charonrpc.SetUserGroupsRequest) (*charonrpc.SetUserGroupsResponse, error) {
	return a.setUserGroupsHandler.SetGroups(ctx, req)
}

func (a userManager) SetPermissions(ctx context.Context, req *charonrpc.SetUserPermissionsRequest) (*charonrpc.SetUserPermissionsResponse, error) {
	return a.setUserPermissionsHandler.SetPermissions(ctx, req)
}

type permissionManager struct {
	*registerPermissionsHandler
	*getPermissionHandler
	*listPermissionsHandler

	charonrpc.UnimplementedPermissionManagerServer
}

func newPermissionManager(server *rpcServer) *permissionManager {
	return &permissionManager{
		registerPermissionsHandler: &registerPermissionsHandler{handler: newHandler(server), registry: server.permissionRegistry},
		listPermissionsHandler:     &listPermissionsHandler{handler: newHandler(server)},
		getPermissionHandler:       &getPermissionHandler{handler: newHandler(server)},
	}
}

func (a permissionManager) Register(ctx context.Context, req *charonrpc.RegisterPermissionsRequest) (*charonrpc.RegisterPermissionsResponse, error) {
	return a.registerPermissionsHandler.Register(ctx, req)
}

func (a permissionManager) Get(ctx context.Context, req *charonrpc.GetPermissionRequest) (*charonrpc.GetPermissionResponse, error) {
	return a.getPermissionHandler.Get(ctx, req)
}

func (a permissionManager) List(ctx context.Context, req *charonrpc.ListPermissionsRequest) (*charonrpc.ListPermissionsResponse, error) {
	return a.listPermissionsHandler.List(ctx, req)
}

type groupManager struct {
	*getGroupHandler
	*deleteGroupHandler
	*modifyGroupHandler
	*listGroupsHandler
	*setGroupPermissionsHandler
	*createGroupHandler
	*listGroupPermissionsHandler

	charonrpc.UnimplementedGroupManagerServer
}

func newGroupManager(server *rpcServer) *groupManager {
	return &groupManager{
		getGroupHandler:             &getGroupHandler{handler: newHandler(server)},
		deleteGroupHandler:          &deleteGroupHandler{handler: newHandler(server)},
		modifyGroupHandler:          &modifyGroupHandler{handler: newHandler(server)},
		listGroupsHandler:           &listGroupsHandler{handler: newHandler(server)},
		setGroupPermissionsHandler:  &setGroupPermissionsHandler{handler: newHandler(server)},
		createGroupHandler:          &createGroupHandler{handler: newHandler(server)},
		listGroupPermissionsHandler: &listGroupPermissionsHandler{handler: newHandler(server)},
	}
}

func (a groupManager) Get(ctx context.Context, req *charonrpc.GetGroupRequest) (*charonrpc.GetGroupResponse, error) {
	return a.getGroupHandler.Get(ctx, req)
}

func (a groupManager) Delete(ctx context.Context, req *charonrpc.DeleteGroupRequest) (*wrappers.BoolValue, error) {
	return a.deleteGroupHandler.Delete(ctx, req)
}

func (a groupManager) Modify(ctx context.Context, req *charonrpc.ModifyGroupRequest) (*charonrpc.ModifyGroupResponse, error) {
	return a.modifyGroupHandler.Modify(ctx, req)
}

func (a groupManager) List(ctx context.Context, req *charonrpc.ListGroupsRequest) (*charonrpc.ListGroupsResponse, error) {
	return a.listGroupsHandler.List(ctx, req)
}

func (a groupManager) SetPermissions(ctx context.Context, req *charonrpc.SetGroupPermissionsRequest) (*charonrpc.SetGroupPermissionsResponse, error) {
	return a.setGroupPermissionsHandler.SetPermissions(ctx, req)
}

func (a groupManager) Create(ctx context.Context, req *charonrpc.CreateGroupRequest) (*charonrpc.CreateGroupResponse, error) {
	return a.createGroupHandler.Create(ctx, req)
}

func (a groupManager) ListPermissions(ctx context.Context, req *charonrpc.ListGroupPermissionsRequest) (*charonrpc.ListGroupPermissionsResponse, error) {
	return a.listGroupPermissionsHandler.ListPermissions(ctx, req)
}

type refreshTokenManager struct {
	*createRefreshTokenHandler
	*listRefreshTokensHandler
	*revokeRefreshTokenHandler

	charonrpc.UnimplementedRefreshTokenManagerServer
}

func newRefreshTokenManager(server *rpcServer) *refreshTokenManager {
	return &refreshTokenManager{
		createRefreshTokenHandler: &createRefreshTokenHandler{handler: newHandler(server)},
		listRefreshTokensHandler:  &listRefreshTokensHandler{handler: newHandler(server)},
		revokeRefreshTokenHandler: &revokeRefreshTokenHandler{handler: newHandler(server)},
	}
}

func (a refreshTokenManager) Create(ctx context.Context, req *charonrpc.CreateRefreshTokenRequest) (*charonrpc.CreateRefreshTokenResponse, error) {
	return a.createRefreshTokenHandler.Create(ctx, req)
}

func (a refreshTokenManager) List(ctx context.Context, req *charonrpc.ListRefreshTokensRequest) (*charonrpc.ListRefreshTokensResponse, error) {
	return a.listRefreshTokensHandler.List(ctx, req)
}

func (a refreshTokenManager) Revoke(ctx context.Context, req *charonrpc.RevokeRefreshTokenRequest) (*charonrpc.RevokeRefreshTokenResponse, error) {
	return a.revokeRefreshTokenHandler.Revoke(ctx, req)
}

func unaryServerInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		wrap := func(current grpc.UnaryServerInterceptor, next grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return current(currentCtx, currentReq, info, next)
			}
		}
		chain := handler
		for _, i := range interceptors {
			chain = wrap(i, chain)
		}
		return chain(ctx, req)
	}
}
