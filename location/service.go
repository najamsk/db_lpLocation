package location

import (
	// "github.com/launchpad7/lp7_micro_location/location/models"
	"github.com/launchpad7/lp7_micro_location/location/models"
	uuid "github.com/satori/go.uuid"
)

// Service interface
type Service interface {
	GetLocationByID(id uuid.UUID) (*models.Location, error)

	GetCountries(countryid string) (*[]models.Country, error)

	GetStates(
		stateid string,
		countryid string,
	) (*[]models.State, error)

	GetCities(
		cityid string,
		stateid string,
		countryid string,
	) (*[]models.City, error)
	GetLocations(
		locationid string,
		cityid string,
		stateid string,
		countryid string) (*[]models.Location, error)

	GetFeatures(id string) (*[]models.Feature, error)
	GetLocationFeatures(locationid string) (*[]models.Feature, error)

	CreateLocation(location *models.Location) (*models.Location, error)
	UpdateLocationFeatures(locationid string, featureids []string, userid string) error
	SaveFeature(feature *models.Feature) (*models.Feature, error)
	GetCategoryDetail(id string) (*[]models.CategoryDetail, error)
	SaveCategoryDetail(categoryDetail *models.CategoryDetail) (*models.CategoryDetail, error)
	SaveCategory(category *models.Category) (*models.Category, error)
	GetCategory(id string) (*[]models.Category, error)
}
