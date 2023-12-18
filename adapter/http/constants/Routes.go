package routesConsts

import "net/http"

const (
	PostVehicleConst = "/vehicle"

	PostUserConst = "/user"

	PostPolicyConst         = "/policy"
	PostPolicyVehicleConst  = "/policy/vehicle"
	PostPolicyCoverageConst = "/policy/coverage"
)

const (
	DefaultPortConst = ":8083"
)

const (
	OriginLocalhostConst = "http://localhost"
)

const (
	BadRequestConst          = http.StatusBadRequest
	ForbiddenRequestConst    = http.StatusForbidden
	Unauthorized             = http.StatusUnauthorized
	StatusOk                 = http.StatusOK
	InternarServerErrorConst = http.StatusInternalServerError
	CreatedConst             = http.StatusCreated
	TimeoutConst             = http.StatusRequestTimeout
)

const (
	DataKeyConst    = "data"
	MessageKeyConst = "message"
)
