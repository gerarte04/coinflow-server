package types

import (
	"coinflow/coinflow-server/api-gateway/internal/models"
	pb "coinflow/coinflow-server/gen/auth_service/golang"
	"coinflow/coinflow-server/pkg/vars"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

// Requests --------------------------------------------

func CreateRegisterRequest(usr *models.User) (*pb.RegisterRequest, error) {
	const op = "CreateRegisterRequest"

	var pbUsr pb.User

	if err := copier.Copy(&pbUsr, usr); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &pb.RegisterRequest{Usr: &pbUsr}, nil
}

// Responses -------------------------------------------

func CreateGetUserDataResponse(resp *pb.GetUserDataResponse) (*models.User, error) {
	const op = "CreateGetUserDataResponse"

	var usr models.User

	if err := copier.Copy(&usr, resp.Usr); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tm, err := time.Parse(vars.TimeLayout, resp.Usr.RegistrationTimestamp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	usr.RegistrationTimestamp = tm

	return &usr, nil
}
