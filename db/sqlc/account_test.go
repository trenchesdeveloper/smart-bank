package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/trenchesdeveloper/smart-bank/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)

	require.Equal(t, account.Balance, arg.Balance)

	require.Equal(t, account.Currency, arg.Currency)
	require.NotZero(t, account.ID)

	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T){
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.ID, account1.ID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T){
	account1 := createRandomAccount(t)

	args := UpdateAccountParams{
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.ID, account1.ID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T){
	account1 := createRandomAccount(t)

	 err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err2 := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t,account2)


}

func TestListAccounts(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

 	accounts, err :=	testQueries.ListAccounts(context.Background(), args)

	 require.NoError(t, err)
	 require.Len(t, accounts, 5)

	 for _, account := range accounts {
		require.NotEmpty(t, account)
	 }

}