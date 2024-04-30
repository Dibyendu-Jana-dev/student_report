package logic

import (
	"context"
	"mywon/students_reports/constants"
	"mywon/students_reports/graph/model"
	"mywon/students_reports/repository"
	"mywon/students_reports/validation"
	"os"
)

func CreateStudentDetails(ctx context.Context, input model.CreateStudentsInput )(*model.CreateStudentsResponse, error){
	PgSchema := os.Getenv(constants.POSTGRES_SCHEMA)
	response := &model.CreateStudentsResponse{}
	if repository.Pool == nil{
		repository.Pool = repository.GetPool()
	}
	conn := repository.SQLConnDetails{
		PgSchema: PgSchema,
		Pool:     repository.Pool,
	}

	err := validation.UpsertStudentValidation(ctx, input)
	if err != nil{
		return nil, err
	}
	res, err := conn.CreateStudentDetails(ctx, input)
	if err != nil{
		return nil, err
	}
	response = res
	return response, nil
}

func GetStudentDetails(ctx context.Context, input model.GetStudentDetailsInput )(*model.CreateStudentsResponse, error){
	PgSchema := os.Getenv(constants.POSTGRES_SCHEMA)
	response := &model.CreateStudentsResponse{}
	if repository.Pool == nil{
		repository.Pool = repository.GetPool()
	}
	conn := repository.SQLConnDetails{
		PgSchema: PgSchema,
		Pool:     repository.Pool,
	}

	res, err := conn.GetStudentDetails(ctx, input)
	if err != nil{
		return nil, err
	}
	response = res
	return response, nil
}
