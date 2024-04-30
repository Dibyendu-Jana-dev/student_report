package validation

import (
	"context"
	"errors"
	"mywon/students_reports/constants"
	"mywon/students_reports/graph/model"
	"mywon/students_reports/repository"
	"os"
	"regexp"
	"github.com/agrison/go-commons-lang/stringUtils"
)

func UpsertStudentValidation(ctx context.Context, input model.CreateStudentsInput) error {
	if input.ID == nil {
		if input.StudentName != nil {
			if stringUtils.IsBlank(*input.StudentName) {
				return errors.New("student name can't be empty")
			}
		} else {
			return errors.New("student name can't be nil")
		}
		if input.StudentMobileNumber != nil {
			if stringUtils.IsBlank(*input.StudentMobileNumber) {
				return errors.New("student mobile number can't be empty")
			} else {
				isExist := IsMobileNumberExist(*input.StudentMobileNumber)
				if isExist {
					return errors.New("student mobile number should be unique")
				}
			}
		} else {
			return errors.New("student mobile number can't be nil")
		}
		if input.StudentBloodGroup != nil {
			if stringUtils.IsBlank(*input.StudentBloodGroup) {
				return errors.New("student blood group can't be empty")
			} else {
				err := BloodGroupValidations(*input.StudentBloodGroup)
				if err != nil {
					return errors.New("please provide a valid blood group")
				}
			}
		} else {
			return errors.New("student blood group can't be nil")
		}
		if input.StudentAddress != nil {
			if stringUtils.IsBlank(*input.StudentAddress) {
				return errors.New("student address is required")
			}
		} else {
			return errors.New("student address can't be nil")
		}
		if input.DateOfBirth != nil {
			if stringUtils.IsBlank(*input.DateOfBirth) {
				return errors.New("date of birth is required")
			} else {
				re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
				if !re.MatchString(*input.DateOfBirth) {
					return errors.New("please enter valid DateOfBirth like DD/MM/YYYY")
				}
			}
		} else {
			return errors.New("student DateOfBirth can't be nil")
		}
		if input.StudentClass != nil {
			if *input.StudentClass == 0 {
				return errors.New("student class is required")
			}
		} else {
			return errors.New("student class can't be nil")
		}

		if input.StudentRoll != nil {
			if *input.StudentRoll == 0 {
				return errors.New("student roll is required")
			} else {
				isExist := IsRollNumberExist(*input.StudentRoll)
				if isExist {
					return errors.New("student roll number should be unique")
				}
			}
		} else {
			return errors.New("student roll can't be nil")
		}
	} else {
		
		if input.StudentMobileNumber != nil {
			if stringUtils.IsBlank(*input.StudentMobileNumber) {
				return errors.New("student mobile number can't be empty")
			} else {
				isExist := IsMobileNumberExist(*input.StudentMobileNumber)
				if isExist {
					return errors.New("student mobile number should be unique")
				}
			}
		}

		if input.StudentRoll != nil {
			if *input.StudentRoll == 0 {
				return errors.New("student roll is required")
			} else {
				isExist := IsRollNumberExist(*input.StudentRoll)
				if isExist {
					return errors.New("student Roll number should be unique")
				}
			}
		}

		if input.StudentBloodGroup != nil {
			if stringUtils.IsBlank(*input.StudentBloodGroup) {
				return errors.New("student blood group can't be empty")
			} else {
				err := BloodGroupValidations(*input.StudentBloodGroup)
				if err != nil {
					return errors.New("please provide a valid blood group")
				}
			}
		}

		if input.DateOfBirth != nil {
			if stringUtils.IsBlank(*input.DateOfBirth) {
				return errors.New("date of birth is required")
			} else {
				re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
				if !re.MatchString(*input.DateOfBirth) {
					return errors.New("please enter valid DateOfBirth like DD/MM/YYYY")
				}
			}
		}
	}
	return nil
}

func BloodGroupValidations(bloodGroup string) error {
	blood := make(map[string]bool)
	blood["A+"] = true
	blood["A-"] = true
	blood["B+"] = true
	blood["B-"] = true
	blood["O+"] = true
	blood["O-"] = true
	blood["AB+"] = true
	blood["AB-"] = true
	if _, ok := blood[bloodGroup]; !ok {
		return errors.New("please provide valid blood group")
	}
	return nil
}

func IsMobileNumberExist(mobielNumber string) bool {
	var hasValue int
	var response bool
	var sqlQuery string
	PgSchema := os.Getenv(constants.POSTGRES_SCHEMA)
	if repository.Pool == nil {
		repository.Pool = repository.GetPool()
	}
	conn := repository.SQLConnDetails{
		PgSchema: PgSchema,
		Pool:     repository.Pool,
	}
	sqlQuery = `SELECT 1 FROM ` + conn.PgSchema + `.student s WHERE s.student_mobile_no = $1`
	err := conn.Pool.QueryRowContext(context.Background(), sqlQuery, mobielNumber).Scan(&hasValue)
	if err != nil {
		return false
	}
	if hasValue == 1 {
		response = true
	}
	return response
}

func IsRollNumberExist(RollNumber int) bool {
	var hasValue int
	var response bool
	var sqlQuery string
	PgSchema := os.Getenv(constants.POSTGRES_SCHEMA)
	if repository.Pool == nil {
		repository.Pool = repository.GetPool()
	}
	conn := repository.SQLConnDetails{
		PgSchema: PgSchema,
		Pool:     repository.Pool,
	}
	sqlQuery = `SELECT 1 FROM ` + conn.PgSchema + `.student s WHERE s.student_roll_no = $1`
	err := conn.Pool.QueryRowContext(context.Background(), sqlQuery, RollNumber).Scan(&hasValue)
	if err != nil {
		return false
	}
	if hasValue == 1 {
		response = true
	}
	return response
}
