package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

type TxOptions struct {
	Isolation sql.IsolationLevel
	ReadOnly bool
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// type AddEntryTxParams struct {
// 	FromUserID int64 `json:"from_user_id"`
// 	ToUserID int64 `json:"to_user_id"`
// 	Entry Liftentry `json:"entry"`
// }

// type AddEntryTxResult struct {
// 	AddEntry AddEntry `json:"addentry"`
// 	FromUser UserData `json:"from_user"`
// 	ToUser UserData `json:"to_user"`
// 	FromLiftEntry Liftentry `json:"from_entry"`
// 	ToLiftEntry Liftentry `json:"to_entry"`
// }

// func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

// }