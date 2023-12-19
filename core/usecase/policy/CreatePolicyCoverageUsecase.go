package policyUsecase

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	sharedMethodInterface "github.com/anaclaraddias/brick/core/port/sharedMethod"
)

type CreatePolicyCoverage struct {
	coverageDatabase   repositoryInterface.CoverageRepositoryInterface
	policyDatabase     repositoryInterface.PolicyRepositoryInterface
	policySharedMethod sharedMethodInterface.PolicySharedMethodInterface
	policyCoverage     *policyDomain.PolicyCoverage
}

func NewCreatePolicyCoverage(
	coverageDatabase repositoryInterface.CoverageRepositoryInterface,
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
	policySharedMethod sharedMethodInterface.PolicySharedMethodInterface,
	policyCoverage *policyDomain.PolicyCoverage,
) *CreatePolicyCoverage {
	return &CreatePolicyCoverage{
		coverageDatabase:   coverageDatabase,
		policyDatabase:     policyDatabase,
		policySharedMethod: policySharedMethod,
		policyCoverage:     policyCoverage,
	}
}

func (createPolicyCoverage *CreatePolicyCoverage) Execute() error {
	if err := createPolicyCoverage.policySharedMethod.VerifyIfPolicyExists(
		createPolicyCoverage.policyCoverage.PolicyId,
	); err != nil {
		return err
	}

	if err := createPolicyCoverage.verifyIfCoverageExists(); err != nil {
		return err
	}

	if err := createPolicyCoverage.verifyIfCoverageIsInPolicy(); err != nil {
		return err
	}

	if err := createPolicyCoverage.policyDatabase.CreatePolicyCoverage(
		createPolicyCoverage.policyCoverage,
	); err != nil {
		return err
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

func (createPolicyCoverage *CreatePolicyCoverage) verifyIfCoverageIsInPolicy() error {
	policyCoverage, err := createPolicyCoverage.policyDatabase.FindPolicyCoverageByPolicyIdAndCoverageId(
		createPolicyCoverage.policyCoverage,
	)

	if err != nil {
		return err
	}

	if policyCoverage != nil {
		return fmt.Errorf(helper.CoverageAlreadyInPolicyConst)
	}

	return nil
}
