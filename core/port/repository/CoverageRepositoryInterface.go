package repositoryInterface

import coverageDomain "github.com/anaclaraddias/brick/core/domain/coverage"

type CoverageRepositoryInterface interface {
	FindCoverageById(coverageId string) (*coverageDomain.Coverage, error)
}
