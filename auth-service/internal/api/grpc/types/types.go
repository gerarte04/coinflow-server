package types

import (
	"coinflow/coinflow-server/auth-service/internal/models"
	pb "coinflow/coinflow-server/gen/auth_service/golang"
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/pkg/vars"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

// Requests -------------------------------------------

type RegisterRequestObject struct {
	User *models.User
}

func CreateRegisterRequestObject(r *pb.RegisterRequest) (*RegisterRequestObject, error) {
	const op = "CreateRegisterRequestObject"

	var usr models.User

	if err := copier.Copy(&usr, r.Usr); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &RegisterRequestObject{User: &usr}, nil
}

type GetUserDataRequestObject struct {
	UsrId uuid.UUID
}

func CreateGetUserDataRequestObject(r *pb.GetUserDataRequest) (*GetUserDataRequestObject, error) {
	const op = "CreateUserDataRequestObject"

	usrId, err := utils.ParseStringToUuid(r.UserId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetUserDataRequestObject{UsrId: usrId}, nil
}

// Responses -------------------------------------------

func CreateGetUserDataResponse(usr *models.User) (*pb.GetUserDataResponse, error) {
	const op = "CreateGetTransactionResponse"

	var pbUsr pb.User

	if err := copier.Copy(&pbUsr, usr); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pbUsr.RegistrationTimestamp = usr.RegistrationTimestamp.Format(vars.TimeLayout)

	return &pb.GetUserDataResponse{Usr: &pbUsr}, nil
}
