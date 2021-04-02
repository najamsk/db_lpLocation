package cockroachdb

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/launchpad7/lp7_micro_location/internal/config"
	"github.com/launchpad7/lp7_micro_location/location"
	"github.com/launchpad7/lp7_micro_location/location/models"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type cockroachRepository struct {
	client *gorm.DB
}

func newCockroachClient() (*gorm.DB, error) {

	fmt.Println("newCockroachClient func invoked")
	configuration, errConfig := config.New()

	if errConfig != nil {
		fmt.Printf("error reading config %v \n", errConfig)
	}

	//reading config for opening connection
	cdbHost := configuration.Items.CockroachDB.Host
	cdbPort := configuration.Items.CockroachDB.Port
	cdbDatabase := configuration.Items.CockroachDB.Database
	// cdbDefaultDatabase := configuration.Items.CockroachDB.DefaultDatabase
	cdbUser := configuration.Items.CockroachDB.User
	cdbSSL := configuration.Items.CockroachDB.SSL
	var cdbSSLMode string = ""
	if cdbSSL == false {
		cdbSSLMode = "?sslmode=disable"
	} else {
		cdbSSLMode = ""
	}

	//this should be formated string
	computedConnectionString := fmt.Sprintf("postgresql://%s@%s:%s/%s%s", cdbUser, cdbHost, cdbPort, cdbDatabase, cdbSSLMode)
	fmt.Printf("computed connection string is = %v \n", computedConnectionString)
	// open default postgres db to create our database
	db, errdb := gorm.Open("postgres", computedConnectionString)
	if errdb != nil {
		fmt.Println("error from db returned")
		fmt.Printf("%v", errdb)
		return nil, errdb
	}
	fmt.Printf("cockroachdb client after error catching \n \n")
	db.LogMode(true)
	//Repo should have dbClose method that close this db
	return db, nil

}

// NewCockroachRepository will be called by other code to get db
func NewCockroachRepository() (location.Repository, error) {
	repo := &cockroachRepository{}

	db, err := newCockroachClient()
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewCockroachRepo")
	}
	repo.client = db
	return repo, nil
}

func (r *cockroachRepository) GetLocationByID(id uuid.UUID) (*models.Location, error) {

	//najam code
	location := &models.Location{}
	err := r.client.Find(location, "ID = ?", id).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return location, nil
}

func (r *cockroachRepository) GetCountries(countryid string) (*[]models.Country, error) {
	//defer r.client.Close()
	//najam code
	cwhere := "1 = 1 "
	countries := &[]models.Country{}

	if len(countryid) > 0 && countryid != "" {

		countryID, err := uuid.FromString(countryid)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND id='%v'", cwhere, countryID)
		}
	}

	err := r.client.Where(cwhere).Find(countries).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.GetCountries", err)
		return nil, err
	}

	return countries, nil
}

func (r *cockroachRepository) GetStates(stateid string, countryid string) (*[]models.State, error) {
	//defer r.client.Close()
	//najam code

	cwhere := "1 = 1 "
	states := &[]models.State{}
	fmt.Println("cockroachRepository: printing input vars")
	fmt.Printf("stateid:%v, countryid: %v", stateid, countryid)

	if len(stateid) > 0 && stateid != "" {

		stateID, err := uuid.FromString(stateid)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND id='%v'", cwhere, stateID)
		}
	}

	if len(countryid) > 0 && countryid != "" {
		cwhere = fmt.Sprintf("%v AND country_id='%v'", cwhere, countryid)
	}

	fmt.Printf("cwhere is = %v \n", cwhere)

	//err := r.client.Where("1 = 1 AND city_id= ? ", stateid).Find(cities).Error
	err := r.client.Where(cwhere).Find(states).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.GetStates", err)
		return nil, err
	}

	return states, nil
}

func (r *cockroachRepository) GetCities(cityid string, stateid string, countryid string) (*[]models.City, error) {
	//defer r.client.Close()
	//najam code
	cwhere := "1 = 1 "
	cities := &[]models.City{}

	fmt.Println("cockroachRepository: printing input vars")
	fmt.Printf("cityid: %v, stateid:%v, countryid: %v", cityid, stateid, countryid)
	if len(cityid) > 0 && cityid != "" {

		cityID, err := uuid.FromString(cityid)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND id='%v'", cwhere, cityID)
		}
	}

	if len(stateid) > 0 && stateid != "" {
		cwhere = fmt.Sprintf("%v AND state_id='%v'", cwhere, stateid)
	}

	if len(countryid) > 0 && countryid != "" {
		cwhere = fmt.Sprintf("%v AND country_id='%v'", cwhere, countryid)
	}

	fmt.Printf("cwhere is = %v \n", cwhere)

	//err := r.client.Where("1 = 1 AND city_id= ? ", cityid).Find(cities).Error
	err := r.client.Where(cwhere).Find(cities).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.GetCities", err)
		return nil, err
	}

	return cities, nil
}

func (r *cockroachRepository) GetLocations(locationid string, cityid string, stateid string, countryid string) (*[]models.Location, error) {
	//defer r.client.Close()
	//najam code
	cwhere := "1 = 1 "
	locations := &[]models.Location{}

	fmt.Println("cockroachRepository: printing input vars")
	fmt.Printf("locationid: %v, cityid: %v, stateid:%v, countryid: %v", locationid, cityid, stateid, countryid)
	if len(locationid) > 0 && locationid != "" {

		locationID, err := uuid.FromString(locationid)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND locations.id='%v'", cwhere, locationID)
		}
	}

	if len(cityid) > 0 && cityid != "" {
		cwhere = fmt.Sprintf("%v AND locations.city_id='%v'", cwhere, cityid)
	}

	if len(stateid) > 0 && stateid != "" {
		cwhere = fmt.Sprintf("%v AND locations.state_id='%v'", cwhere, stateid)
	}

	if len(countryid) > 0 && countryid != "" {
		cwhere = fmt.Sprintf("%v AND locations.country_id='%v'", cwhere, countryid)
	}

	fmt.Printf("cwhere is = %v \n", cwhere)

	//err := r.client.Where("1 = 1 AND city_id= ? ", cityid).Find(locations).Error
	// err := r.client.Where(cwhere).Find(locations).Error

	err := r.client.Table("locations").Select("locations.*,cities.name as city_name,states.name as state_name, countries.name as country_name").Joins(" inner join countries on locations.country_id = countries.id").Joins(" inner join cities  on locations.city_id = cities.id").Joins("inner join states on locations.state_id = states.id").Where(cwhere).Scan(&locations).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.GetLocations", err)
		return nil, err
	}

	return locations, nil
}

func (r *cockroachRepository) GetFeatures(id string) (*[]models.Feature, error) {

	cwhere := "1 = 1 "
	features := &[]models.Feature{}

	if len(id) > 0 && id != "" {

		featureID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND features.id='%v'", cwhere, featureID)
		}
	}

	//err := r.client.Where(cwhere).Find(features).Error
	err := r.client.Table("features").Select("features.*,category_details.name as category_name").Joins(" inner join category_details on features.category_detail_id = category_details.id").Where(cwhere).Scan(&features).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.GetFeatures", err)
		return nil, err
	}

	return features, nil
}

func (r *cockroachRepository) GetLocationFeatures(locationid string) (*[]models.Feature, error) {

	cwhere := "1 = 1 "
	features := &[]models.Feature{}

	if len(locationid) > 0 && locationid != "" {

		locationID, err := uuid.FromString(locationid)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND location_features.location_id='%v'", cwhere, locationID)
		}
	}

	//err := r.client.Where(cwhere).Find(features).Error
	err := r.client.Table("features").Select("features.*").Joins("inner join location_features on features.id = location_features.feature_id").Where(cwhere).Scan(&features).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.GetLocationFeatures", err)
		return nil, err
	}

	return features, nil
}

func (r *cockroachRepository) CreateLocation(location *models.Location) (*models.Location, error) {

	//najam code

	fmt.Println("cockroachRepository: printing input vars")
	fmt.Printf("location: %v", location)

	cityid := location.CityID.String()
	stateid := location.StateID.String()
	countryid := location.CountryID.String()

	fmt.Printf("locationrepo: cityid= %v \n", cityid)
	fmt.Printf("locationrepo: len - cityid= %v \n", len(cityid))
	fmt.Printf("locationrepo: len - cityid= %v \n", len(cityid))

	if len(cityid) == 0 || cityid == "" {
		return nil, errors.New("city ID is required")
	}

	if len(stateid) == 0 || stateid == "" {
		return nil, errors.New("state ID is required")
	}

	if len(countryid) == 0 || countryid == "" {
		return nil, errors.New("country ID is required")
	}

	if location.ID == uuid.Nil {
		fmt.Printf("location.id is nil = %v \n", location.ID)
		location.ID, _ = uuid.NewV4()
		rID := r.client.Create(location)
		fmt.Printf("rID = %+v \n", rID)
	} else {
		rID := r.client.Save(location)
		fmt.Printf("rID = %+v \n", rID)
	}

	fmt.Printf("after create statement \n")
	fmt.Printf("location = %+v \n", location)
	fmt.Printf("location.ID = %+v \n", location.ID)

	return location, nil
}

func (r *cockroachRepository) UpdateLocationFeatures(
	locationid string, featureids []string, userid string) error {

	var deleteQuery strings.Builder
	var insertQuery strings.Builder
	ids := strings.Join(featureids, "','")

	deleteQuery.WriteString(" delete from location_features ")
	deleteQuery.WriteString(" where  location_id = ? ")
	if len(ids) > 0 {
		deleteQuery.WriteString(fmt.Sprintf(` and feature_id not in ('%s')`, ids))
	}

	insertQuery.WriteString(" insert into  location_features (id, location_id, feature_id, created_by, created_at, is_active, is_add_on) ")
	insertQuery.WriteString(" select  gen_random_uuid(), ?,  id, ?, now(), true, false  from features  ")
	insertQuery.WriteString(fmt.Sprintf(` where id in ('%s')`, ids))
	insertQuery.WriteString(" and id not in(select feature_id from location_features where location_id = ?)")

	fmt.Println("deleteQuery:", deleteQuery.String())
	fmt.Println("insertQuery:", insertQuery.String())

	// Note the use of tx as the database handle once you are within a transaction
	thandler := r.client.Begin()
	defer func() {
		if r := recover(); r != nil {
			thandler.Rollback()
		}
	}()

	if err := thandler.Error; err != nil {
		return err
	}

	if err := thandler.Exec(deleteQuery.String(), locationid).Error; err != nil {
		thandler.Rollback()
		return err
	}

	if len(ids) > 0 {
		if err := thandler.Exec(insertQuery.String(), locationid, userid, locationid).Error; err != nil {
			thandler.Rollback()
			return err
		}

	}

	return thandler.Commit().Error
}

func (r *cockroachRepository) SaveFeature(feature *models.Feature) (*models.Feature, error) {

	//najam code

	fmt.Println("cockroachRepository: printing input vars")
	fmt.Printf("feature: %v", feature)

	categorydetailid := feature.CategoryDetailID.String()

	fmt.Printf("repo: categorydetailid= %v \n", categorydetailid)

	if len(categorydetailid) == 0 || categorydetailid == "" {
		return nil, errors.New("CategoryDetailID is required")
	}

	if feature.ID == uuid.Nil {
		fmt.Printf("feature.id is nil = %v \n", feature.ID)
		feature.ID, _ = uuid.NewV4()
		rID := r.client.Create(feature)
		fmt.Printf("rID = %+v \n", rID)
	} else {
		rID := r.client.Save(feature)
		fmt.Printf("rID = %+v \n", rID)
	}

	fmt.Printf("after create statement \n")
	fmt.Printf("feature = %+v \n", feature)
	fmt.Printf("feature.ID = %+v \n", feature.ID)

	return feature, nil
}

func (r *cockroachRepository) GetCategoryDetail(id string) (*[]models.CategoryDetail, error) {

	cwhere := "1 = 1 "
	categoryDetails := &[]models.CategoryDetail{}

	if len(id) > 0 && id != "" {

		catDetailID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND id='%v'", cwhere, catDetailID)
		}
	}

	err := r.client.Where(cwhere).Find(categoryDetails).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.categoryDetails", err)
		return nil, err
	}

	return categoryDetails, nil
}

func (r *cockroachRepository) SaveCategoryDetail(catDetail *models.CategoryDetail) (*models.CategoryDetail, error) {

	fmt.Println("cockroachRepository: printing input vars")
	fmt.Printf("catDetail: %v", catDetail)

	categoryid := catDetail.CategoryID.String()

	fmt.Printf("repo: categoryid= %v \n", categoryid)

	if len(categoryid) == 0 || categoryid == "" {
		return nil, errors.New("categoryid is required")
	}

	if catDetail.ID == uuid.Nil {
		fmt.Printf("feature.id is nil = %v \n", catDetail.ID)
		catDetail.ID, _ = uuid.NewV4()
		rID := r.client.Create(catDetail)
		fmt.Printf("rID = %+v \n", rID)
	} else {
		rID := r.client.Save(catDetail)
		fmt.Printf("rID = %+v \n", rID)
	}

	fmt.Printf("after create statement \n")
	fmt.Printf("catDetail = %+v \n", catDetail)
	fmt.Printf("catDetail.ID = %+v \n", catDetail.ID)

	return catDetail, nil
}

func (r *cockroachRepository) SaveCategory(category *models.Category) (*models.Category, error) {

	fmt.Println("cockroachRepository: printing input vars")
	fmt.Printf("category: %v", category)

	if category.ID == uuid.Nil {
		fmt.Printf("feature.id is nil = %v \n", category.ID)
		category.ID, _ = uuid.NewV4()
		rID := r.client.Create(category)
		fmt.Printf("rID = %+v \n", rID)
	} else {
		rID := r.client.Save(category)
		fmt.Printf("rID = %+v \n", rID)
	}

	fmt.Printf("after create statement \n")
	fmt.Printf("feature = %+v \n", category)
	fmt.Printf("feature.ID = %+v \n", category.ID)

	return category, nil
}

func (r *cockroachRepository) ClientClose() {
	fmt.Println("server shutting down: closing cockroach client conneciton")
	defer r.client.Close()
}

func (r *cockroachRepository) GetCategory(id string) (*[]models.Category, error) {

	cwhere := "1 = 1 "
	categories := &[]models.Category{}

	if len(id) > 0 && id != "" {

		catID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("Something went wrong: %s \n", err)
		} else {
			cwhere = fmt.Sprintf("%v AND id='%v'", cwhere, catID)
		}
	}

	err := r.client.Where(cwhere).Find(categories).Error

	if err != nil {
		fmt.Printf("error= %+v \n from cockroachRepository.GetCategory", err)
		return nil, err
	}

	return categories, nil
}
