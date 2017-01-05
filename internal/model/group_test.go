package model

import (
	"context"
	"testing"
)

var (
	groupTestFixtures = []*GroupEntity{}
)

func TestGroupRepository_IsGranted(t *testing.T) {
	suite := &postgresSuite{}
	suite.setup(t)
	defer suite.teardown(t)

	for ur := range loadGroupFixtures(t, suite.repository.group, groupPermissionsTestFixtures) {
		for pr := range loadPermissionFixtures(t, suite.repository.permission, ur.given.Permissions) {
			add := []*GroupPermissionsEntity{{
				GroupID:             ur.got.ID,
				PermissionSubsystem: pr.got.Subsystem,
				PermissionModule:    pr.got.Module,
				PermissionAction:    pr.got.Action,
			}}
			for range loadGroupPermissionsFixtures(t, suite.repository.groupPermissions, add) {
				exists, err := suite.repository.group.IsGranted(context.TODO(), ur.given.ID, pr.given.Permission())

				if err != nil {
					t.Errorf("group permission cannot be found, unexpected error: %s", err.Error())
					continue
				}

				if !exists {
					t.Errorf("group permission not found for group %d and permission %d", ur.given.ID, pr.given.ID)
				} else {
					t.Logf("group permission relationship exists for group %d and permission %d", ur.given.ID, pr.given.ID)
				}
			}
		}
	}
}

type groupFixtures struct {
	got, given GroupEntity
}

func loadGroupFixtures(t *testing.T, r GroupProvider, f []*GroupEntity) chan groupFixtures {
	data := make(chan groupFixtures, 1)

	go func() {
		for _, given := range f {
			entity, err := r.Insert(context.TODO(), given)
			if err != nil {
				t.Errorf("group cannot be created, unexpected error: %s", err.Error())
				continue
			} else {
				t.Logf("group has been created, got id %d", entity.ID)
			}

			data <- groupFixtures{
				got:   *entity,
				given: *given,
			}
		}

		close(data)
	}()

	return data
}
