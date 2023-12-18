package requestEntity

import (
	"encoding/json"
	"net/http"

	"github.com/anaclaraddias/brick/core/domain/helper"
)

const (
	PolicyIdFieldConst  = "a apólice"
	VehicleIdFieldConst = "o veículo"
)

type PostPolicyVehicleRequestEntity struct {
	PolicyId  string `json:"policy_id"`
	VehicleId string `json:"vehicle_id"`
}

func DecodePolicyVehicleRequest(request *http.Request) (*PostPolicyVehicleRequestEntity, error) {
	var policyVehicle *PostPolicyVehicleRequestEntity

	if err := json.NewDecoder(request.Body).Decode(&policyVehicle); err != nil {
		return nil, err
	}

	return policyVehicle, nil
}

func (policyVehicle *PostPolicyVehicleRequestEntity) Validate() error {
	if err := helper.ValidateUUID(policyVehicle.PolicyId, PolicyIdFieldConst); err != nil {
		return err
	}

	if err := helper.ValidateUUID(policyVehicle.VehicleId, VehicleIdFieldConst); err != nil {
		return err
	}

	return nil
}
