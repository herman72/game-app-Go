package uservalidator

import (
	"fmt"
	"game-app-go/dto"
	"game-app-go/pkg/errmsg"
	"game-app-go/pkg/richerror"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type Repository interface{
	IsPhoneNumberUnique(phonenumber string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository)Validator{
	return Validator{repo: repo}
}

func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) (error, map[string]string) {
	const op = "uservalidator.ValidateRegisterRequest"
	if err := validation.ValidateStruct(&req,
		
		validation.Field(&req.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`))),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^09[0-9]{9}$")), validation.By(v.checkPhoneNumberUniquenes)),
	); err != nil {
		fieldErrors:= make(map[string]string)
		errV, ok := err.(validation.Errors)

		if ok{
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
		WithKind(richerror.KindInvalid).WithMeta(map[string]interface{}{"req": req}).WithError(err),
		fieldErrors
	}
	
	return nil, nil
}


func (v Validator)checkPhoneNumberUniquenes(value interface{}) error {
	PhoneNumber := value.(string)
	if isUnique, err := v.repo.IsPhoneNumberUnique(PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return err 
		}

		if !isUnique {
			return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotUnique) 
		}
		
	}
	return nil
}
