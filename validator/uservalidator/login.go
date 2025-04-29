package uservalidator

import (
	"fmt"
	"game-app-go/param"
	"game-app-go/pkg/errmsg"
	"game-app-go/pkg/richerror"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateLoginRequest(req dto.LoginRequest) (error, map[string]string) {
	const op = "uservalidator.ValidateLoginRequest"
	if err := validation.ValidateStruct(&req,
		
		// validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`))),
		validation.Field(&req.PhoneNumber, validation.Required, 
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid), 
			validation.By(v.doesPhoneNumberExist)),
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


func (v Validator)doesPhoneNumberExist(value interface{}) error {
	PhoneNumber := value.(string)
	_, err := v.repo.GetUserByPhoneNumber(PhoneNumber)

	if err != nil {
		return fmt.Errorf(errmsg.ErrorMsgNotFound)
	}

	// if isUnique, err := v.repo.IsPhoneNumberUnique(PhoneNumber); err != nil || !isUnique {
	// 	if err != nil {
	// 		return err 
	// 	}

	// 	if !isUnique {
	// 		return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotUnique) 
	// 	}
		
	// }
	return nil
}
