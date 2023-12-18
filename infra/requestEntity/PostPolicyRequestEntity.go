package requestEntity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
)

const (
	UserIdFieldConst = "o usuário"
	StatusFieldConst = "o status"
)

type PostPolicyRequestEntity struct {
	Status string `json:"status"`
	UserId string `json:"user_id"`
}

func DecodePolicyRequest(request *http.Request) (*PostPolicyRequestEntity, error) {
	var policy *PostPolicyRequestEntity

	if err := json.NewDecoder(request.Body).Decode(&policy); err != nil {
		return nil, err
	}

	return policy, nil
}

func (policy *PostPolicyRequestEntity) Validate() error {
	if policy.Status == "" {
		return fmt.Errorf(helper.PolicyFieldCanNotBeEmptyConst, StatusFieldConst)
	}

	if !slices.Contains(policyDomain.PolicyStatus, policy.Status) {
		return fmt.Errorf(helper.PolicyStatusInvalidConst)
	}

	if err := helper.ValidateUUID(policy.UserId, UserIdFieldConst); err != nil {
		return err
	}

	return nil
}
