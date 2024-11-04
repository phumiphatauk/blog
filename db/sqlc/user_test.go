package db

import (
	"context"
	"testing"
	"time"

	"blog_api/util"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Code:           util.RandomOwner(),
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FirstName:      util.RandomOwner(),
		LastName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		Phone:          util.RandomPhone(),
		Description: pgtype.Text{
			String: util.RandomString(6),
			Valid:  true,
		},
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Phone, user2.Phone)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserOnlyFirstName(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomOwner()
	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		UserID: oldUser.ID,
		FirstName: pgtype.Text{
			String: newFirstName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := util.RandomEmail()
	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		UserID: oldUser.ID,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	err = testStore.UpdateUserPassword(context.Background(), UpdateUserPasswordParams{
		ID:             oldUser.ID,
		HashedPassword: newHashedPassword,
	})
	require.NoError(t, err)

	updatedUser, err := testStore.GetUserByUsername(context.Background(), oldUser.Username)
	require.NoError(t, err)

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomOwner()
	newLastName := util.RandomString(6)
	newEmail := util.RandomEmail()
	newPhone := util.RandomPhone()
	newDescription := util.RandomString(6)

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		FirstName: pgtype.Text{
			String: newFirstName,
			Valid:  true,
		},
		LastName: pgtype.Text{
			String: newLastName,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
		Phone: pgtype.Text{
			String: newPhone,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: newDescription,
			Valid:  true,
		},
		UserID: oldUser.ID,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.NotEqual(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, newLastName, updatedUser.LastName)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.NotEqual(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, newPhone, updatedUser.Phone)
	require.NotEqual(t, oldUser.Description, updatedUser.Description.String)
	require.Equal(t, newDescription, updatedUser.Description.String)

}
