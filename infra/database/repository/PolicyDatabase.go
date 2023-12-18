package repository

import (
	"database/sql"
	"time"

	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	"github.com/anaclaraddias/brick/infra/models"
)

type PolicyDatabase struct {
	connection port.DatabaseConnectionInterface
}

func NewPolicyDatabase(
	connection port.DatabaseConnectionInterface,
) repositoryInterface.PolicyRepositoryInterface {
	connection.Open()

	return &PolicyDatabase{connection: connection}
}

func (policyDatabase *PolicyDatabase) CreatePolicy(
	policy *policyDomain.Policy,
) error {
	var dbPolicy *models.PolicyModel

	createdAt := time.Now()
	updatedAt := sql.NullTime{Valid: false}

	query := `INSERT INTO policy (
			id,
			status,
			user_id,
			update_date,
			creation_date
		) values (?,?,?,?,?);`

	if err := policyDatabase.connection.Raw(
		query,
		&dbPolicy,
		policy.Id,
		policy.Status,
		policy.UserId,
		updatedAt,
		createdAt,
	); err != nil {
		return err
	}

	return nil
}

func (policyDatabase *PolicyDatabase) FindPolicyById(
	policyId string,
) (*policyDomain.Policy, error) {
	var dbPolicy *models.PolicyModel

	query := `SELECT * FROM policy WHERE id = ?`

	if err := policyDatabase.connection.Raw(
		query,
		&dbPolicy,
		policyId,
	); err != nil {
		return nil, err
	}

	if dbPolicy == nil {
		return nil, nil
	}

	policy := policyDomain.NewPolicy(
		dbPolicy.Id,
		dbPolicy.Status,
		&dbPolicy.StartDate,
		&dbPolicy.EndDate,
		&dbPolicy.CoverageLimit,
		&dbPolicy.Value,
		dbPolicy.UserId,
	)

	return policy, nil
}

func (policyDatabase *PolicyDatabase) CreatePolicyVehicle(
	policyVehicle *policyDomain.PolicyVehicle,
) error {
	var dbPolicyVehicle *models.LinkedPolicyVehicleModel

	createdAt := time.Now()

	query := `INSERT INTO linked_policy_vehicle (
			id,
			vehicle_id,
			policy_id,
			creation_date
		) values (?,?,?,?);`

	if err := policyDatabase.connection.Raw(
		query,
		&dbPolicyVehicle,
		policyVehicle.Id,
		policyVehicle.VehicleId,
		policyVehicle.PolicyId,
		createdAt,
	); err != nil {
		return err
	}

	return nil
}

func (policyDatabase *PolicyDatabase) CreatePolicyCoverage(
	policyCoverage *policyDomain.PolicyCoverage,
) error {
	var dbPolicyCoverage *models.LinkedPolicyCoverageModel

	createdAt := time.Now()

	query := `INSERT INTO linked_policy_coverage (
			id,
			coverage_id,
			policy_id,
			creation_date
		) values (?,?,?,?);`

	if err := policyDatabase.connection.Raw(
		query,
		&dbPolicyCoverage,
		policyCoverage.Id,
		policyCoverage.CoverageId,
		policyCoverage.PolicyId,
		createdAt,
	); err != nil {
		return err
	}

	return nil
}
