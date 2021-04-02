package location

import (
	"github.com/launchpad7/lp7_micro_location/location/models"
	uuid "github.com/satori/go.uuid"
	// "github.com/jinzhu/gorm"
	// "github.com/launchpad7/lp7_micro_location/internal/config"
	// "golang.org/x/crypto/bcrypt"
)

// Repository is here
type Repository interface {
	// GetAll(isActive bool, skip int, limit int) (models.Location, error)

	GetLocationByID(uuid.UUID) (*models.Location, error)
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
	ClientClose()
}
