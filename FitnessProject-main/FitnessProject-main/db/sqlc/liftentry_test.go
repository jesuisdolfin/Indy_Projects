package db

import (
	"context"
	"database/sql"
	"fmt"
	"FitnessProject/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomLiftentry(t *testing.T) Liftentry {

	user1 := createRandomUser(t)

	arg := CreateLiftentryParams{
		UserID: fmt.Sprint(user1.ID),
		WeightLifted: util.RandomWeightLifted(),
		Reps: fmt.Sprint(util.RandomReps()),
	}

	liftentry, err := testQueries.CreateLiftEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, liftentry)

	require.Equal(t, arg.WeightLifted, liftentry.WeightLifted)
	require.Equal(t, arg.Reps, liftentry.Reps)

	require.NotZero(t, liftentry.ID)
	require.NotZero(t, liftentry.CreateAt)

	return liftentry
}

func TestCreateLiftentry(t *testing.T) {
	createRandomLiftentry(t)
}

func TestGetLiftentry(t *testing.T) {
	liftentry1 := createRandomLiftentry(t)
	liftentry2, err := testQueries.GetLiftEntry(context.Background(), liftentry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, liftentry2)

	require.Equal(t, liftentry1.ID, liftentry2.ID)
	require.Equal(t, liftentry1.WeightLifted, liftentry2.WeightLifted)
	require.Equal(t, liftentry1.Reps, liftentry2.Reps)

	require.WithinDuration(t, liftentry1.CreateAt, liftentry2.CreateAt, time.Second)
}

func TestUpdateLiftentry(t *testing.T) {
	liftentry1 := createRandomLiftentry(t)

	arg := UpdateLiftentryParams{
		ID: liftentry1.ID,
		Reps: fmt.Sprint(util.RandomReps()),
	}

	liftentry2, err := testQueries.UpdateLiftEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, liftentry2)

	require.Equal(t, liftentry1.ID, liftentry2.ID)
	require.Equal(t, liftentry1.WeightLifted, liftentry2.WeightLifted)
	require.Equal(t, arg.Reps, liftentry2.Reps)

	require.WithinDuration(t, liftentry1.CreateAt, liftentry2.CreateAt, time.Second)
}

func TestDeleteLiftentry(t *testing.T) {
	liftentry1 := createRandomLiftentry(t)
	err := testQueries.DeleteLiftEntry(context.Background(), liftentry1.ID)
	require.NoError(t, err)

	liftentry2, err := testQueries.GetLiftEntry(context.Background(), liftentry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, liftentry2)
}

func TestListLiftentrys(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomLiftentry(t)
	}

	arg := ListLiftEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	users, err := testQueries.ListLiftEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}