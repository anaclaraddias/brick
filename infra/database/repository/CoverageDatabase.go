package repository

import (
	coverageDomain "github.com/anaclaraddias/brick/core/domain/coverage"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	"github.com/anaclaraddias/brick/infra/models"
)

type CoverageDatabase struct {
	connection port.DatabaseConnectionInterface
}

func NewCoverageDatabase(
	connection port.DatabaseConnectionInterface,
) repositoryInterface.CoverageRepositoryInterface {
	connection.Open()

	return &CoverageDatabase{connection: connection}
}

func (coverageDatabase *CoverageDatabase) FindCoverageById(
	coverageId string,
) (*coverageDomain.Coverage, error) {
	var dbCoverage *models.CoverageModel

	query := `SELECT * FROM coverage WHERE id = ?`

	if err := coverageDatabase.connection.Raw(
		query,
		&dbCoverage,
		coverageId,
	); err != nil {
		return nil, err
	}

	if dbCoverage == nil {
		return nil, nil
	}

	coverage := coverageDomain.NewCoverage(
		dbCoverage.Id,
		dbCoverage.Name,
		dbCoverage.Description,
		dbCoverage.RateValue,
	)

	return coverage, nil
}
