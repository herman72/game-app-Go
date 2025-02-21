package mysql

import (
	"database/sql"
	"fmt"
	"game-app-go/entity"
	"game-app-go/pkg/errmsg"
	"game-app-go/pkg/richerror"
)



func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		// fmt.Errorf("can't scan query result: %w", err)
		return false, richerror.New(op).WithError(err).
		WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return false, nil
}

func(d *MySQLDB)Register(u entity.User)(entity.User, error){
	res, err := d.db.Exec("insert into users(name, phone_number, password) values(?, ?, ?)", u.Name, u.PhoneNumber, u.Password)

	if err != nil {
		return entity.User{}, fmt.Errorf("cant excute command %w", err)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)


	return u, nil
}


func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user,err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}
		// fmt.Errorf("can't scan query result: %w", err)
		return entity.User{}, false, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		
	}
	return user, true, nil
}

func (d *MySQLDB) GetUserByID(userID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"
	row := d.db.QueryRow(`select * from users where id = ?`, userID)
	user, err := scanUser(row)
	
	if err != nil {
		// fmt.Errorf("record not found")
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
			
			
		}
		return entity.User{}, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}	
		
	
	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password)
	return user, err
}