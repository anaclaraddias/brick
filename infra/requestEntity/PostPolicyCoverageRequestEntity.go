package requestEntity

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anaclaraddias/brick/core/domain/helper"
)

const (
	CoverageIdFieldConst = "a cobertura"
)

type PostPolicyCoverageRequestEntity struct {
	CoverageId string `json:"coverage_id"`
	PolicyId   string `json:"policy_id"`
}

func DecodePolicyCoverageRequest(request *http.Request) (*PostPolicyCoverageRequestEntity, error) {
	var policyCoverage *PostPolicyCoverageRequestEntity

	if err := json.NewDecoder(request.Body).Decode(&policyCoverage); err != nil {
		return nil, err
	}

	return policyCoverage, nil
}

func (policyCoverage *PostPolicyCoverageRequestEntity) Validate() error {
	if err := helper.ValidateUUID(policyCoverage.PolicyId, PolicyIdFieldConst); err != nil {
		return err
	}

	if policyCoverage.CoverageId == "" {
		return fmt.Errorf(helper.InformFieldConst, CoverageIdFieldConst)
	}

	return nil
}
