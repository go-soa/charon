package charond

import (
	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/ntypes"
	"github.com/piotrkowalczuk/pqt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type modifyUserHandler struct {
	*handler
}

func (muh *modifyUserHandler) handle(ctx context.Context, req *charon.ModifyUserRequest) (*charon.ModifyUserResponse, error) {
	if req.Id <= 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "charond: user cannot be modified, invalid id: %d", req.Id)
	}

	muh.loggerWith("user_id", req.Id)

	actor, err := muh.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}

	ent, err := muh.repository.user.FindOneByID(req.Id)
	if err != nil {
		return nil, err
	}

	if hint, ok := muh.firewall(req, ent, actor); !ok {
		return nil, grpc.Errorf(codes.PermissionDenied, "charond: "+hint)
	}

	ent, err = muh.repository.user.UpdateByID(
		req.Id,
		nil,
		nil,
		&ntypes.Int64{Int64: actor.user.ID, Valid: actor.user.ID != 0},
		req.FirstName,
		req.IsActive,
		req.IsConfirmed,
		req.IsStaff,
		req.IsSuperuser,
		nil,
		req.LastName,
		req.SecurePassword,
		nil,
		nil,
		req.Username,
	)
	if err != nil {
		switch pqt.ErrorConstraint(err) {
		case tableUserConstraintUsernameUnique:
			return nil, grpc.Errorf(codes.AlreadyExists, charon.ErrDescUserWithUsernameExists)
		default:
			return nil, err
		}
	}

	return muh.response(ent)
}

func (muh *modifyUserHandler) firewall(req *charon.ModifyUserRequest, entity *userEntity, actor *actor) (string, bool) {
	isOwner := actor.user.ID == entity.ID

	if !actor.user.IsSuperuser {
		switch {
		case entity.IsSuperuser:
			return "only superuser can modify a superuser account", false
		case entity.IsStaff && !isOwner && actor.permissions.Contains(charon.UserCanModifyStaffAsStranger):
			return "missing permission to modify an account as a stranger", false
		case entity.IsStaff && isOwner && actor.permissions.Contains(charon.UserCanModifyStaffAsOwner):
			return "missing permission to modify an account as an owner", false
		case req.IsSuperuser != nil && req.IsSuperuser.Valid:
			return "only superuser can change existing account to superuser", false
		case req.IsStaff != nil && req.IsStaff.Valid && !actor.permissions.Contains(charon.UserCanCreateStaff):
			return "user is not allowed to create user with is_staff property that has custom value", false
		}
	}

	return "", true
}

func (muh *modifyUserHandler) response(u *userEntity) (*charon.ModifyUserResponse, error) {
	msg, err := u.message()
	if err != nil {
		return nil, err
	}
	return &charon.ModifyUserResponse{
		User: msg,
	}, nil
}