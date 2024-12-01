package user_serivce

import (
	"context"
	"fmt"
	services_errors "server_crm/internal/services/errors"
	"server_crm/internal/services/models"
)

func (s UserService) Authorize(ctx context.Context, aToken string, rToken string) (string, string, models.User, error) {

	const op = "service.user.Authorize"
	
		
	newAToken := ""

	if aToken != "" {
		userIdFromToken, err := s.usC.Validate(ctx, aToken)

		if err != nil {
			return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
		}

		user, err := s.usP.ReadOne(ctx, userIdFromToken)

		if err != nil {
			return "", "", models.User{}, fmt.Errorf("%s: %w", op, services_errors.ErrIncorrectAuthToken)
		}

		newAToken, _, err = s.usC.Generate(ctx, userIdFromToken)

		if err != nil {
			return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return newAToken, rToken, s.fromStorageToDomain(user), nil
	}

	if rToken != "" {
		userIdFromToken, err := s.usC.Validate(ctx, rToken)

		if err != nil {
			return "", "", models.User{}, fmt.Errorf("%s: %w", op, services_errors.ErrNoAuthorizationTokens)
		}

		user, err := s.usP.ReadOne(ctx, userIdFromToken)

		if err != nil {
			return "", "", models.User{}, fmt.Errorf("%s: %w", op, services_errors.ErrIncorrectAuthToken)
		}

		newAToken, _, err = s.usC.Generate(ctx, userIdFromToken)

		if err != nil {
			return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return newAToken, rToken, s.fromStorageToDomain(user), nil
	}
	return "", "", models.User{}, fmt.Errorf("%s: %w", op, services_errors.ErrNoAuthorizationTokens)

}
