package policyUsecase

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
)

type CreatePolicyCoverage struct {
	coverageDatabase repositoryInterface.CoverageRepositoryInterface
	policyDatabase   repositoryInterface.PolicyRepositoryInterface
	policyCoverage   *policyDomain.PolicyCoverage
}

func NewCreatePolicyCoverage(
	coverageDatabase repositoryInterface.CoverageRepositoryInterface,
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
	policyCoverage *policyDomain.PolicyCoverage,
) *CreatePolicyCoverage {
	return &CreatePolicyCoverage{
		coverageDatabase: coverageDatabase,
		policyDatabase:   policyDatabase,
		policyCoverage:   policyCoverage,
	}
}

func (createPolicyCoverage *CreatePolicyCoverage) Execute() error {
	if err := createPolicyCoverage.verifyIfPolicyExists(); err != nil {
		return err
	}

	if err := createPolicyCoverage.verifyIfCoverageExists(); err != nil {
		return err
	}

	if err := createPolicyCoverage.policyDatabase.CreatePolicyCoverage(
		createPolicyCoverage.policyCoverage,
	); err != nil {
		return err
	}

	return nil
}

func (createPolicyCoverage *CreatePolicyCoverage) verifyIfPolicyExists() error {
	policy, err := createPolicyCoverage.policyDatabase.FindPolicyById(
		createPolicyCoverage.policyCoverage.PolicyId,
	)

	if err != nil {
		return err
	}

	if policy == nil {
		return fmt.Errorf(helper.PolicyNotFoundConst)
	}

	return nil
}

func (createPolicyCoverage *CreatePolicyCoverage) verifyIfCoverageExists() error {
	coverage, err := createPolicyCoverage.coverageDatabase.FindCoverageById(
		createPolicyCoverage.policyCoverage.CoverageId,
	)

	if err != nil {
		return err
	}

	if coverage == nil {
		return fmt.Errorf(helper.CoverageNotFoundConst)
	}

	return nil
}
