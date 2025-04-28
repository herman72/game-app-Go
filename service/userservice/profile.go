package userservice

import (
	"game-app-go/dto"
	"game-app-go/pkg/richerror"
)

// All req inputs for intractor/service should be sanitized

func (s Service)Profile(req dto.ProfileRequest)(dto.ProfileResponse, error){
	const op =  "userservice.Profile"
	user, err := s.repo.GetUserByID(req.UserID)
	// I assume data is already sanitized
	if err != nil{
		// TODO: we can use rich error for better error handeling
		return dto.ProfileResponse{}, richerror.New(op).
											WithError(err).
											WithMeta(map[string]interface{}{"req":req})
		
		// return ProfileResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return dto.ProfileResponse{Name: user.Name}, nil
}
