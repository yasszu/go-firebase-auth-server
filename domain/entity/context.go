package entity

import "context"

type contextKey int

const (
	ContextKeyUser contextKey = iota
	ContextKeyAccountID
)

func GetCurrentUser(ctx context.Context) (*User, error) {
	accountID, ok := ctx.Value(ContextKeyUser).(*User)
	if !ok {
		return nil, &UnauthorizedError{Massage: "Unauthorized"}
	}

	return accountID, nil
}
