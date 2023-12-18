package policyUsecase

import (
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
)

type CreatePolicy struct {
	policyDatabase repositoryInterface.PolicyRepositoryInterface
	policy         *policyDomain.Policy
}

func NewCreatePolicy(
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
	policy *policyDomain.Policy,
) *CreatePolicy {
	return &CreatePolicy{
		policyDatabase: policyDatabase,
		policy:         policy,
	}
}

func (createPolicy *CreatePolicy) Execute() error {
	err := createPolicy.policyDatabase.CreatePolicy(createPolicy.policy)

	if err != nil {
		return err
	}

	return nil
}
