package handler

import (
	"be17/main/feature/user"
	"time"
)

type UserResponse struct {
	Id        uint      
	Name      string    
	Email     string    
	CreatedAt time.Time 
}

func CoreToResponse(input user.Core) UserResponse{
	return UserResponse{
		Id: input.Id,
		Name: input.Name,
		Email: input.Email,
		CreatedAt: input.CreatedAt,
	}
}