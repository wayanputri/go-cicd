package handler

import "be17/main/feature/user"

func RequestToCore(dataInput UserRequest) user.Core{
	return user.Core{
		Name: dataInput.Name,
		Email: dataInput.Email,
		Phone: dataInput.Phone,
		Password: dataInput.Password,
	}
}