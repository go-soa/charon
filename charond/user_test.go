// +build postgres !unit

package main

import (
	"fmt"
	"testing"

	"reflect"

	"github.com/piotrkowalczuk/nilt"
)

func TestUserRepository_Create(t *testing.T) {
	suite := setupPostgresSuite(t)
	defer suite.teardown(t)

	for res := range generateUserRepositoryData(t, suite) {
		if res.got.CreatedAt.IsZero() {
			t.Errorf("invalid created at field, expected valid time but got %v", res.got.CreatedAt)
		} else {
			t.Logf("user has been properly created at %v", res.got.CreatedAt)
		}

		if res.got.Username != res.given.username {
			t.Errorf("wrong username, expected %s but got %s", res.given.username, res.got.Username)
		}
	}
}

func TestUserRepository_UpdateByID(t *testing.T) {
	suffix := "_modified"
	suite := setupPostgresSuite(t)
	defer suite.teardown(t)

	for res := range generateUserRepositoryData(t, suite) {
		user := res.got
		modified, err := suite.user.UpdateByID(
			user.ID,
			nil,
			nil,
			nilt.Int64{},
			nilt.String{String: user.FirstName + suffix, Valid: true},
			nilt.Bool{Bool: true, Valid: true},
			nilt.Bool{Bool: true, Valid: true},
			nilt.Bool{Bool: true, Valid: true},
			nilt.Bool{Bool: true, Valid: true},
			nil,
			nilt.String{String: user.LastName + suffix, Valid: true},
			nil,
			nil,
			nilt.Int64{},
			nilt.String{String: user.Username + suffix, Valid: true},
		)

		if err != nil {
			t.Errorf("user cannot be modified, unexpected error: %s", err.Error())
			continue
		} else {
			fmt.Println(modified)
			t.Logf("user with id %d has been modified", modified.ID)
		}

		if assertfTime(t, modified.UpdatedAt, "invalid updated at field, expected valid time but got %v", modified.UpdatedAt) {
			t.Logf("user has been properly modified at %v", modified.UpdatedAt)
		}
		assertf(t, modified.Username == user.Username+suffix, "wrong username, expected %s but got %s", user.Username+suffix, modified.Username)
		assertf(t, reflect.DeepEqual(modified.Password, user.Password), "wrong password, expected %s but got %s", user.Password, modified.Password)
		assertf(t, modified.FirstName == user.FirstName+suffix, "wrong first name, expected %s but got %s", user.FirstName+suffix, modified.FirstName)
		assertf(t, modified.LastName == user.LastName+suffix, "wrong last name, expected %s but got %s", user.LastName+suffix, modified.LastName)
		assert(t, modified.IsSuperuser, "user should become a superuser")
		assert(t, modified.IsActive, "user should be active")
		assert(t, modified.IsConfirmed, "user should be confirmed")
		assert(t, modified.IsStaff, "user should become a staff")
	}
}

func TestUserRepository_DeleteOneByID(t *testing.T) {
	suite := setupPostgresSuite(t)
	defer suite.teardown(t)

	for res := range generateUserRepositoryData(t, suite) {
		affected, err := suite.user.DeleteByID(res.got.ID)

		if err != nil {
			t.Errorf("user cannot be deleted, unexpected error: %s", err.Error())
			continue
		}

		assert(t, affected == 1, "user was not deleted, no rows affected")

	}
}

func TestUserRepository_FindOneByID(t *testing.T) {
	suite := setupPostgresSuite(t)
	defer suite.teardown(t)

	for res := range generateUserRepositoryData(t, suite) {
		user := res.got
		found, err := suite.user.FindOneByID(user.ID)

		if err != nil {
			t.Errorf("user cannot be deleted, unexpected error: %s", err.Error())
			continue
		}

		if assert(t, found != nil, "user was not found, nil object returned") {
			assertf(t, reflect.DeepEqual(res.got, *found), "created and retrieved entity should be equal, but its not\ncreated: %#v\nfounded: %#v", res.got, found)
		}
	}
}

func TestUserRepository_UpdateLastLoginAt(t *testing.T) {
	suite := setupPostgresSuite(t)
	defer suite.teardown(t)

	for res := range generateUserRepositoryData(t, suite) {
		affected, err := suite.user.UpdateLastLoginAt(res.got.ID)

		if err != nil {
			t.Errorf("user cannot be updated, unexpected error: %s", err.Error())
			continue
		}

		if assert(t, affected == 1, "user was not updated, no rows affected") {
			entity, err := suite.user.FindOneByID(res.got.ID)
			if err != nil {
				t.Errorf("user cannot be found, unexpected error: %s", err.Error())
				continue
			}

			assertfTime(t, entity.LastLoginAt, "user last login at property was not properly updated, got %v", entity.LastLoginAt)
		}
	}
}

func TestUserRepository_Find(t *testing.T) {
	var (
		err      error
		entities []*userEntity
		all      int64
	)
	suite := setupPostgresSuite(t)
	defer suite.teardown(t)

	for _ = range generateUserRepositoryData(t, suite) {
		all++
	}

	entities, err = suite.user.Find(&userCriteria{limit: all})
	if err != nil {
		t.Errorf("users can not be retrieved, unexpected error: %s", err.Error())
	} else {
		assertf(t, int64(len(entities)) == all, "number of users retrived do not match, expected %d got %d", all, len(entities))
	}

	entities, err = suite.user.Find(&userCriteria{
		offset: all,
		limit:  all,
	})
	if err != nil {
		t.Errorf("users can not be retrieved, unexpected error: %s", err.Error())
	} else {
		assertf(t, len(entities) == 0, "number of users retrived do not match, expected %d got %d", 0, len(entities))
	}
}

type userRepositoryFixture struct {
	got   userEntity
	given struct {
		username, password, firstName, lastName, confirmationToken string
		isSuperuser, isStaff, isActive, isConfirmed                bool
	}
}

func generateUserRepositoryData(t *testing.T, suite *postgresSuite) chan userRepositoryFixture {
	given := []struct {
		username, password, firstName, lastName, confirmationToken string
		isSuperuser, isStaff, isActive, isConfirmed                bool
	}{
		{
			username:          "johnsnow@gmail.com",
			password:          "secret",
			firstName:         "John",
			lastName:          "Snow",
			confirmationToken: "1234567890",
		},
		{
			username:          "1",
			password:          "2",
			firstName:         "3",
			lastName:          "4",
			confirmationToken: "5",
		},
	}
	data := make(chan userRepositoryFixture, 1)

	go func() {
		for _, g := range given {
			entity, err := suite.user.Create(
				g.username,
				[]byte(g.password),
				g.firstName,
				g.lastName,
				[]byte(g.confirmationToken),
				g.isSuperuser,
				g.isStaff,
				g.isActive,
				g.isConfirmed,
			)
			if err != nil {
				t.Errorf("user cannot be created, unexpected error: %s", err.Error())
				continue
			} else {
				t.Logf("user has been created, got id %d", entity.ID)
			}

			data <- userRepositoryFixture{
				got:   *entity,
				given: g,
			}
		}

		close(data)
	}()

	return data
}
