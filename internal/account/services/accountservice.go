package services

import (
	"FinanceApp/internal/account"
	"FinanceApp/internal/db/sqlc"
	"context"
)

type AccountService interface {
	Create(context.Context, account.CreateAccountRequest) (*account.AccountResponse, error)
	Get(context.Context, account.AccountUriRequest) (*account.AccountResponse, error)
	Update(context.Context, account.AccountUriRequest, account.UpdateAccountRequest) (*account.AccountResponse, error)
	Delete(context.Context, account.AccountUriRequest) error
}

type AccountServiceImp struct {
	store db.Store
}

func (a *AccountServiceImp) Create(ctx context.Context, accRequest account.CreateAccountRequest) (*account.AccountResponse, error) {

	acc, err := a.store.CreateAccount(ctx, db.CreateAccountParams{
		Number:   accRequest.PhoneNumber,
		Owner:    accRequest.Owner,
		Currency: accRequest.Currency,
		Balance:  0,
	})

	if err != nil {
		return nil, err
	}
	return (&account.AccountResponse{
		ID:          acc.ID,
		Currency:    acc.Currency,
		Balance:     acc.Balance,
		PhoneNumber: acc.Number,
		Owner:       acc.Owner,
	}), nil
}

func (a *AccountServiceImp) Update(ctx context.Context, uri account.AccountUriRequest, reqDto account.UpdateAccountRequest) (*account.AccountResponse, error) {
	// TODO will fix
	accountEntity, err := a.store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:      uri.ID,
		Balance: reqDto.Balance,
	})

	if err != nil {
		return nil, err
	}
	_ = accountEntity
	return nil, nil
}

func (a *AccountServiceImp) Get(ctx context.Context, uri account.AccountUriRequest) (*account.AccountResponse, error) {

	accountEntity, err := a.store.GetAccountForUpdate(ctx, uri.ID)

	_ = accountEntity

	if err != nil {
		return nil, err
	}

	return &account.AccountResponse{
		ID:          accountEntity.ID,
		Email:       accountEntity.Email,
		PhoneNumber: accountEntity.Number,
		Owner:       accountEntity.Owner,
		Balance:     accountEntity.Balance,
		Currency:    accountEntity.Currency,
	}, nil
}

func (a *AccountServiceImp) Delete(ctx context.Context, uri account.AccountUriRequest) error {
	err := a.store.DeleteAccount(ctx, uri.ID)
	if err != nil {
		return err
	}
	return nil
}
