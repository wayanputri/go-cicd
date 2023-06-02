package service

import (
	"be17/main/feature/user"
	"be17/main/mocks"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T){
	userDataLayer :=new(mocks.UserData)


	t.Run("test case s get succes insert data", func(t *testing.T){
		insertData := user.Core{Name:"Andi",Phone:"9865555", Email:"andi@gmail.com",Password:"123"}
		userDataLayer.On("Insert",insertData).Return(nil).Once()
		srv :=New(userDataLayer)
		err:=srv.Create(insertData)
		assert.Nil(t, err)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case s get failed insert data", func(t *testing.T){
		insertData := user.Core{Name:"Andi",Phone:"9865555", Email:"andi@gmail.com",Password:"123"}
		userDataLayer.On("Insert",insertData).Return(errors.New("error insert data ")).Once()
		srv :=New(userDataLayer)
		err:=srv.Create(insertData)
		assert.NotNil(t, err)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case failed insert data, validation error", func(t *testing.T) {
		insertData := user.Core{Name: "Andi", Phone: "0812345"}

		srv := New(userDataLayer)
		err := srv.Create(insertData)
		assert.NotNil(t, err)
		userDataLayer.AssertExpectations(t)
	})
	
}

func TestLogin(t *testing.T){
	userDataLayer :=new(mocks.UserData)
	mockPassword := "mockPasword"
	mockEmail := "mock@gmail.com"
	mockResult := user.Core{Email:"andi@gmail.com",Password:"123"}
	mockToken := "MkBxeN395StiLVyN8W"

	t.Run("test case s get succes login data", func(t *testing.T){
		userDataLayer.On("Login",mockEmail, mockPassword).Return(mockResult,mockToken,nil).Once()
		srv :=New(userDataLayer)
		result, token, err:=srv.Login(mockEmail, mockPassword)
		assert.Nil(t, err)
		assert.Equal(t,mockResult.Email,result.Email)
		assert.Equal(t,mockToken,token)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case s get failed login data", func(t *testing.T){
		userDataLayer.On("Login",mockEmail, mockPassword).Return(user.Core{},"",errors.New("error login ")).Once()
		srv :=New(userDataLayer)
		result, token,err:=srv.Login(mockEmail, mockPassword)
		assert.NotNil(t,err)
		assert.Equal(t, user.Core{}, result)
		assert.Equal(t, "", token)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case email and password is empty", func(t *testing.T){
		mockEmail := ""
		mockPassword := ""
		
		srv :=New(userDataLayer)
		result, token,err:=srv.Login(mockEmail, mockPassword)
		assert.Equal(t, errors.New("error Validation: nama, email, password harus diisi "), err)
		assert.Equal(t, user.Core{}, result)
		assert.Equal(t, "", token)
		userDataLayer.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	userDataLayer := new(mocks.UserData)
	returnData := user.Core{Id:1,Name:"Andi",Phone:"9865555", Email:"andi@gmail.com",Password:"123"}

	t.Run("test case succes get data id", func(t *testing.T){
		id := 1
		userDataLayer.On("SelectById",id).Return(returnData,nil).Once()
		
		srv :=New(userDataLayer)
		response,err:=srv.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t,returnData.Name,response.Name)
		assert.Equal(t,returnData.Email, response.Email)
		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case s get failed get data id", func(t *testing.T){
		id := 2
		userDataLayer.On("SelectById",id).Return(user.Core{Id:0},errors.New("error read data get id")).Once()
		
		srv :=New(userDataLayer)
		response,err:=srv.GetById(id)
		assert.NotNil(t, err)
		assert.Equal(t, user.Core{Id: 0}, response)
		userDataLayer.AssertExpectations(t)
	})
		
}

 func TestUpdate(t *testing.T) {
	userDataLayer := new(mocks.UserData)

	t.Run("test case success Update data id", func(t *testing.T) {
		id := "1"
		newData := user.Core{Name: "Updated Andi", Phone: "987654321"}
		userDataLayer.On("Update", id,newData).Return(nil).Once()

		srv := New(userDataLayer)
		err := srv.Update(id, newData)
		assert.Nil(t, err)

		userDataLayer.AssertExpectations(t)
	})

	t.Run("test case failed Update data id", func(t *testing.T) {
		id := "2"
		NewData := user.Core{Name: "Budi", Phone: "987654"}
		expectedErr := errors.New("update failed")
		userDataLayer.On("Update", id,NewData).Return(expectedErr).Once()

		srv := New(userDataLayer)
		err := srv.Update(id, NewData )
		assert.NotNil(t, err)

		userDataLayer.AssertExpectations(t)
	})
}


func TestDelete(t *testing.T) {
	userDataLayer := new(mocks.UserData)

	t.Run("test case success delete data by ID", func(t *testing.T) {
		id := 1
		userDataLayer.On("Delete", id).Return(nil)

		srv :=New(userDataLayer)
		err := srv.Delete(id)
		assert.Nil(t, err)

		userDataLayer.AssertCalled(t, "Delete", id)
	})

	t.Run("test case failed delete data by ID", func(t *testing.T) {
		id := 2
		expectedError := errors.New("failed to delete data")
		userDataLayer.On("Delete", id).Return(expectedError)
		service := New(userDataLayer)
		err := service.Delete(id)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)

		userDataLayer.AssertCalled(t, "Delete", id)
	})
}





//go test ./feature/user/... -coverprofile=cover.out && go tool cover -html=cover.out