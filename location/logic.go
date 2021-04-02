package location

import (
	"errors"

	"fmt"

	"github.com/launchpad7/lp7_micro_location/location/models"
	errs "github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrRecordNotFound  = errors.New("Record Not Found")
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

//TODO: figure out []*entity and *[]entity
type locationService struct {
	locationRepo Repository
}

// NewLocationService hold repo to work
func NewLocationService(locationRepo Repository) Service {
	return &locationService{
		locationRepo,
	}
}

func (r *locationService) GetLocationByID(id uuid.UUID) (*models.Location, error) {
	loc, err := r.locationRepo.GetLocationByID(id)

	if err != nil {
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetLocationByID")
	}

	return loc, nil
}

//func (r *locationService) GetCountryList() (*[]models.Country, error) {
//	countries, err := r.locationRepo.GetCountryList()
//
//	if err != nil {
//		fmt.Printf("err loading countries from locationRepo.GetCountryList: %+v \n", err)
//		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetCountryList")
//	}
//
//	return countries, nil
//}

func (r *locationService) GetCountries(countryid string) (*[]models.Country, error) {
	countries, err := r.locationRepo.GetCountries(countryid)

	if err != nil {
		fmt.Printf("err loading countries from locationRepo.GetCountryList: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetCountryList")
	}

	return countries, nil
}
func (r *locationService) GetStates(
	stateid string,
	countryid string,
) (*[]models.State, error) {
	countries, err := r.locationRepo.GetStates(stateid, countryid)

	if err != nil {
		fmt.Printf("err loading countries from locationRepo.GetCountryList: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetCountryList")
	}

	return countries, nil
}

func (r *locationService) GetCities(
	cityid string,
	stateid string,
	countryid string,
) (*[]models.City, error) {
	cities, err := r.locationRepo.GetCities(cityid, stateid, countryid)

	if err != nil {
		fmt.Printf("err loading cities from locationRepo.GetCountryList: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetCountryList")
	}

	return cities, nil
}

func (r *locationService) GetLocations(
	locationid string,
	cityid string,
	stateid string,
	countryid string,
) (*[]models.Location, error) {

	locations, err := r.locationRepo.GetLocations(locationid, cityid, stateid, countryid)

	if err != nil {
		fmt.Printf("err loading locations from locationRepo.GetCountryList: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetCountryList")
	}

	return locations, nil
}

func (r *locationService) CreateLocation(
	location *models.Location,
) (*models.Location, error) {

	rlocation, err := r.locationRepo.CreateLocation(location)

	if err != nil {
		fmt.Printf("err loading locations from locationRepo.CreateLocation: %+v \n", err)
		return nil, errs.Wrap(err, "service.Location.CreateLocation")
	}

	return rlocation, nil
}

//UpdateLocationFeatures(locationid string, featureids []string) error
func (r *locationService) UpdateLocationFeatures(
	locationid string, featureids []string, userid string) error {

	err := r.locationRepo.UpdateLocationFeatures(locationid, featureids, userid)

	if err != nil {
		fmt.Printf("err loading locations from locationRepo.CreateLocation: %+v \n", err)
		return errs.Wrap(err, "service.Location.UpdateLocationFeatures")
	}

	return nil
}

func (r *locationService) GetFeatures(id string) (*[]models.Feature, error) {
	features, err := r.locationRepo.GetFeatures(id)

	if err != nil {
		fmt.Printf("err loading features from locationRepo.GetFeatures: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetFeatures")
	}

	return features, nil
}

func (r *locationService) GetLocationFeatures(locationid string) (*[]models.Feature, error) {
	locationFeatures, err := r.locationRepo.GetLocationFeatures(locationid)

	if err != nil {
		fmt.Printf("err loading countries from locationRepo.GetLocationFeatures: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetLocationFeatures")
	}

	return locationFeatures, nil
}

func (r *locationService) SaveFeature(
	feature *models.Feature,
) (*models.Feature, error) {

	rfeature, err := r.locationRepo.SaveFeature(feature)

	if err != nil {
		fmt.Printf("err loading features from locationRepo.SaveFeature: %+v \n", err)
		return nil, errs.Wrap(err, "service.Location.SaveFeature")
	}

	return rfeature, nil
}

func (r *locationService) GetCategoryDetail(id string) (*[]models.CategoryDetail, error) {
	categoryDetail, err := r.locationRepo.GetCategoryDetail(id)

	if err != nil {
		fmt.Printf("err loading category details from locationRepo.GetCategoryDetail: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetCategoryDetail")
	}

	return categoryDetail, nil
}

func (r *locationService) SaveCategoryDetail(
	catDetail *models.CategoryDetail,
) (*models.CategoryDetail, error) {

	rCatDetail, err := r.locationRepo.SaveCategoryDetail(catDetail)

	if err != nil {
		fmt.Printf("err loading features from locationRepo.SaveCategoryDetail: %+v \n", err)
		return nil, errs.Wrap(err, "service.Location.SaveCategoryDetail")
	}

	return rCatDetail, nil
}

func (r *locationService) GetCategory(id string) (*[]models.Category, error) {
	category, err := r.locationRepo.GetCategory(id)

	if err != nil {
		fmt.Printf("err loading category details from locationRepo.GetCategory: %+v \n", err)
		return nil, errs.Wrap(ErrRecordNotFound, "service.Location.GetCategory")
	}

	return category, nil
}

func (r *locationService) SaveCategory(
	category *models.Category,
) (*models.Category, error) {

	rCategory, err := r.locationRepo.SaveCategory(category)

	if err != nil {
		fmt.Printf("err loading features from locationRepo.SaveCategory: %+v \n", err)
		return nil, errs.Wrap(err, "service.Location.SaveCategory")
	}

	return rCategory, nil
}

//func (r *locationService) Store(redirect *Redirect) error {
//	if err := validate.Validate(redirect); err != nil {
//		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
//	}
//	redirect.Code = shortid.MustGenerate()
//	redirect.CreatedAt = time.Now().UTC().Unix()
//	return r.locationRepo.Store(redirect)
//}
