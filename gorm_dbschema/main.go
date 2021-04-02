package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/launchpad7/lp7_micro_location/internal/config"

	// Import GORM-related packages.

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//global varriable required in different functions for referencing entities in insert into statements.

var catDrink = Category{Name: "Drinks", IsActive: true}
var catMeals = Category{Name: "Meals", IsActive: true}
var catParking = Category{Name: "Parking", IsActive: true}
var catIndoorGames = Category{Name: "Indoor Games", IsActive: true}
var catVideoGames = Category{Name: "Video Games", IsActive: true}
var catRoom = Category{Name: "Room", IsActive: true}
var catShiftTiming = Category{Name: "Shift", IsActive: true}

var catMealsLunch = CategoryDetail{Name: "Lunch", IsActive: true}
var catMealsDinner = CategoryDetail{Name: "Dinner", IsActive: true}
var catParkingBike = CategoryDetail{Name: "Bike", IsActive: true}
var catParkingCar = CategoryDetail{Name: "Car", IsActive: true}
var catInGamesTable = CategoryDetail{Name: "Table Games", IsActive: true}
var catInGamesBoard = CategoryDetail{Name: "Board Games", IsActive: true}
var catVidGamesXbox = CategoryDetail{Name: "Xbox Games", IsActive: true}
var catRoomMeeting = CategoryDetail{Name: "Meeting Room", IsActive: true}
var catRoomConference = CategoryDetail{Name: "Conference Room", IsActive: true}
var catRoomPrivate = CategoryDetail{Name: "Private Room", IsActive: true}
var catRoomShared = CategoryDetail{Name: "Shared Room", IsActive: true}
var catShiftFirst = CategoryDetail{Name: "First Shift", IsActive: true}
var catShiftSecond = CategoryDetail{Name: "Second Shift", IsActive: true}
var catShiftExtended = CategoryDetail{Name: "Extended Shift", IsActive: true}
var catShiftFull = CategoryDetail{Name: "24 hours Shift", IsActive: true}

var featureBikeParking = Feature{Name: "Bike Parking", IsActive: true}
var featureCarParking = Feature{Name: "Car Parking", IsActive: true}
var featureLunch = Feature{Name: "Lunch", IsActive: true}
var featureDinner = Feature{Name: "Dinner", IsActive: true}
var featureTableTennis = Feature{Name: "Table Tennis", IsActive: true}
var featureXbox = Feature{Name: "XBox Gaming", IsActive: true}
var featureMeeting = Feature{Name: "Meeting Room", IsActive: true}
var featureConference = Feature{Name: "Conference Room", IsActive: true}
var featurePrivate = Feature{Name: "Private Room", IsActive: true}
var featureShared = Feature{Name: "Shared Space", IsActive: true}
var featureShiftFull = Feature{Name: "24 hours open", IsActive: true}
var featureShiftFirst = Feature{Name: "9am to 7pm open", IsActive: true}

var countryPakistan = Country{
	Name:        "Pakistan",
	DisplayName: "Islami Jamhooriya Pakistan",
	IsActive:    true,
	//States:      []*State{&statePunjab, &stateIslamabad},
}

var countryUk = Country{
	Name:        "UK",
	DisplayName: "United Kingdom",
	IsActive:    false,
	//States:      []*State{&statePunjab, &stateIslamabad},
}
var countryUSA = Country{
	Name:        "USA",
	DisplayName: "United States Of America",
	IsActive:    false,
	//States:      []*State{&statePunjab, &stateIslamabad},
}
var statePunjab = State{
	Name:        "Punjab",
	DisplayName: "Punjab",
	IsActive:    true,
	CountryID:   countryPakistan.ID,
	//Cities:      []*City{&cityRawalpinidi}
}

var stateIslamabad = State{
	Name:        "Islamabad",
	DisplayName: "Islamabad",
	IsActive:    true,
	CountryID:   countryPakistan.ID,
	//Cities:      []*City{&cityIslamabad}
}

var cityRawalpinidi = City{
	Name:        "Rawalpindi",
	DisplayName: "Rawalpindi",
	IsActive:    true,
	//Locations:   []*Location{&locationMoftakpk}
}

var cityIslamabad = City{
	Name:        "Islamabad",
	DisplayName: "Islamabad",
	IsActive:    true,
	//Locations:   []*Location{&locationMoftakblue}
}

var userzohaib = Staff{
	IsActive:      true,
	FirstName:     "zohaib",
	LastName:      "Hassan",
	Email:         "zz@jsp.com",
	Mobile:        "0345-990901",
	Title:         "Mr",
	Designation:   "Biz guru",
	IsSecondShift: false,
}

var usershakeela = Staff{
	IsActive:      true,
	FirstName:     "Shakeela",
	LastName:      "Saqaat",
	Email:         "shakeela@jsp.com",
	Mobile:        "0345-98982111",
	Title:         "Miss",
	Designation:   "Admin Officer",
	IsSecondShift: false,
}

var useribti = Staff{
	IsActive:      true,
	FirstName:     "Ibtihaaj",
	LastName:      "Shaikh",
	Email:         "ibti@jsp.com",
	Mobile:        "0300-9898123",
	Title:         "Mr",
	Designation:   "Admin Night Shift",
	IsSecondShift: true,
}

var locationMoftakblue = Location{
	IsActive:     true,
	Name:         "moftak Blue Area",
	DisplayName:  "Moftak blue area branch",
	Mobile:       "0345-3077412",
	PhoneNumber:  "051-585858",
	Fax:          "0092-555458",
	AddressLine1: "blue ki pheli gali",
	AddressLine2: "meri nani kali apnay kaam par chali",
}
var locationMoftakpk = Location{
	IsActive:     true,
	Name:         "moftak chandani chowk",
	DisplayName:  "Moftak chandani chowk branch",
	Mobile:       "0346-3077000",
	PhoneNumber:  "051-585000",
	Fax:          "0092-0099909",
	AddressLine1: "murree road",
	AddressLine2: "yeah chandni chowk yeh agay ja kar sajay",
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

func getPassword(key string) string {
	var hashpass string
	hash, err := bcrypt.GenerateFromPassword([]byte(strings.ToLower(key)), bcrypt.DefaultCost)
	if err == nil {
		hashpass = string(hash)
	}

	return hashpass
}

func seedDataCountries(db *gorm.DB) {
	db.DropTableIfExists(&Country{})
	db.AutoMigrate(&Country{})
	db.Create(&countryPakistan)
	db.Create(&countryUk)
	db.Create(&countryUSA)
	fmt.Println("seedDataCountry Done")

}

func seedDataStates(db *gorm.DB) {
	db.DropTableIfExists(&State{})
	db.AutoMigrate(&State{})

	stateIslamabad.CountryID = countryPakistan.ID
	statePunjab.CountryID = countryPakistan.ID
	db.Create(&stateIslamabad)
	db.Create(&statePunjab)

	fmt.Println("seedDataStates Done")

}

func seedDataCity(db *gorm.DB) {
	db.DropTableIfExists(&City{})
	db.AutoMigrate(&City{})

	cityIslamabad.StateID = stateIslamabad.ID
	cityIslamabad.CountryID = countryPakistan.ID

	cityRawalpinidi.CountryID = countryPakistan.ID
	cityRawalpinidi.StateID = statePunjab.ID

	db.Create(&cityRawalpinidi)
	db.Create(&cityIslamabad)

	fmt.Println("seedDataCity Done")

}

func seedStaff(db *gorm.DB) {
	db.DropTableIfExists(&Staff{})
	db.AutoMigrate(&Staff{})

	cityIslamabad.StateID = stateIslamabad.ID
	cityIslamabad.CountryID = countryPakistan.ID

	cityRawalpinidi.CountryID = countryPakistan.ID
	cityRawalpinidi.StateID = statePunjab.ID

	db.Create(&userzohaib)
	db.Create(&usershakeela)
	db.Create(&useribti)

	fmt.Println("seedStaff Done")

}

func seedDataLocation(db *gorm.DB) {
	db.DropTableIfExists(&Location{})
	db.AutoMigrate(&Location{})

	locationMoftakblue.CityID = cityIslamabad.ID
	locationMoftakblue.StateID = stateIslamabad.ID
	locationMoftakblue.CountryID = countryPakistan.ID

	locationMoftakpk.CityID = cityRawalpinidi.ID
	locationMoftakpk.StateID = statePunjab.ID
	locationMoftakpk.CountryID = countryPakistan.ID

	db.Create(&locationMoftakpk)
	db.Create(&locationMoftakblue)

	fmt.Println("seedDataLocation Done")

}

func seedLocationStaff(db *gorm.DB) {
	db.DropTableIfExists(&LocationStaff{})
	db.AutoMigrate(&LocationStaff{})

	db.Create(&LocationStaff{UserID: usershakeela.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationStaff{UserID: usershakeela.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationStaff{UserID: userzohaib.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationStaff{UserID: useribti.ID, LocationID: locationMoftakblue.ID})

	fmt.Println("seedLocationStaff Done")
}

func seedDataCategory(db *gorm.DB) {
	db.DropTableIfExists(&Category{})
	db.AutoMigrate(&Category{})

	db.Create(&catDrink)
	db.Create(&catMeals)
	db.Create(&catParking)
	db.Create(&catIndoorGames)
	db.Create(&catVideoGames)
	db.Create(&catRoom)
	db.Create(&catShiftTiming)

	fmt.Println("seedDataCategory Done")

}

func seedDataCategoryDetail(db *gorm.DB) {
	db.DropTableIfExists(&CategoryDetail{})
	db.AutoMigrate(&CategoryDetail{})

	catMealsLunch.CategoryID = catMeals.ID
	catMealsDinner.CategoryID = catMeals.ID
	catParkingBike.CategoryID = catParking.ID
	catParkingCar.CategoryID = catParking.ID
	catInGamesBoard.CategoryID = catIndoorGames.ID
	catInGamesTable.CategoryID = catIndoorGames.ID
	catVidGamesXbox.CategoryID = catVideoGames.ID
	catRoomConference.CategoryID = catRoom.ID
	catRoomMeeting.CategoryID = catRoom.ID
	catRoomPrivate.CategoryID = catRoom.ID
	catRoomShared.CategoryID = catRoom.ID
	catShiftExtended.CategoryID = catShiftTiming.ID
	catShiftFirst.CategoryID = catShiftTiming.ID
	catShiftSecond.CategoryID = catShiftTiming.ID
	catShiftFull.CategoryID = catShiftTiming.ID

	db.Create(&catMealsLunch)
	db.Create(&catMealsDinner)
	db.Create(&catParkingBike)
	db.Create(&catParkingCar)
	db.Create(&catInGamesTable)
	db.Create(&catInGamesBoard)
	db.Create(&catVidGamesXbox)
	db.Create(&catRoomShared)
	db.Create(&catRoomPrivate)
	db.Create(&catRoomConference)
	db.Create(&catRoomMeeting)
	db.Create(&catShiftFull)
	db.Create(&catShiftSecond)
	db.Create(&catShiftExtended)
	db.Create(&catShiftFirst)

	fmt.Println("seedDataCategoryDetail Done")

}

func seedDataFeatures(db *gorm.DB) {
	db.DropTableIfExists(&Feature{})
	db.AutoMigrate(&Feature{})

	featureXbox.CategoryDetailID = catVidGamesXbox.ID
	featureTableTennis.CategoryDetailID = catInGamesTable.ID
	featureBikeParking.CategoryDetailID = catParkingBike.ID
	featureCarParking.CategoryDetailID = catParkingCar.ID
	featureLunch.CategoryDetailID = catMealsLunch.ID
	featureDinner.CategoryDetailID = catMealsDinner.ID
	featureMeeting.CategoryDetailID = catRoomMeeting.ID
	featureShared.CategoryDetailID = catRoomShared.ID
	featureConference.CategoryDetailID = catRoomConference.ID
	featurePrivate.CategoryDetailID = catRoomPrivate.ID
	featureShiftFirst.CategoryDetailID = catShiftFirst.ID
	featureShiftFull.CategoryDetailID = catShiftFull.ID

	db.Create(&featureXbox)
	db.Create(&featureTableTennis)
	db.Create(&featureBikeParking)
	db.Create(&featureCarParking)
	db.Create(&featureLunch)
	db.Create(&featureDinner)
	db.Create(&featureMeeting)
	db.Create(&featureShared)
	db.Create(&featureConference)
	db.Create(&featurePrivate)
	db.Create(&featureShiftFull)
	db.Create(&featureShiftFirst)

	fmt.Println("seedDataFeatures Done")

}

func seedDataLocationFeatures(db *gorm.DB) {
	db.DropTableIfExists(&LocationFeature{})
	db.AutoMigrate(&LocationFeature{})

	db.Create(&LocationFeature{FeatureID: featureXbox.ID, LocationID: locationMoftakblue.ID})
	db.Create(&LocationFeature{FeatureID: featureTableTennis.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationFeature{FeatureID: featureBikeParking.ID, LocationID: locationMoftakblue.ID})
	db.Create(&LocationFeature{FeatureID: featureCarParking.ID, LocationID: locationMoftakblue.ID})
	db.Create(&LocationFeature{FeatureID: featureBikeParking.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationFeature{FeatureID: featureLunch.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationFeature{FeatureID: featureDinner.ID, LocationID: locationMoftakpk.ID})

	db.Create(&LocationFeature{FeatureID: featurePrivate.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationFeature{FeatureID: featureShared.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationFeature{FeatureID: featureConference.ID, LocationID: locationMoftakpk.ID})
	db.Create(&LocationFeature{FeatureID: featureMeeting.ID, LocationID: locationMoftakpk.ID})

	db.Create(&LocationFeature{FeatureID: featurePrivate.ID, LocationID: locationMoftakblue.ID})
	db.Create(&LocationFeature{FeatureID: featureShared.ID, LocationID: locationMoftakblue.ID})
	db.Create(&LocationFeature{FeatureID: featureConference.ID, LocationID: locationMoftakblue.ID})
	db.Create(&LocationFeature{FeatureID: featureMeeting.ID, LocationID: locationMoftakblue.ID})

	db.Create(&LocationFeature{FeatureID: featureShiftFirst.ID, LocationID: locationMoftakblue.ID})
	db.Create(&LocationFeature{FeatureID: featureShiftFull.ID, LocationID: locationMoftakpk.ID})

	fmt.Println("seedDataLocationFeatures Done")

}

func seedDataLocationProductType(db *gorm.DB) {
	db.DropTableIfExists(&LocationProductType{})
	db.AutoMigrate(&LocationProductType{})
	fmt.Println("seedDataLocationProductType Done")
}

func seedDataLocationGallery(db *gorm.DB) {
	db.DropTableIfExists(&LocationGallery{})
	db.AutoMigrate(&LocationGallery{})
	fmt.Println("seedDataLocationGallery done")
}
func seedDataInventory(db *gorm.DB) {
	db.DropTableIfExists(&Inventory{})
	db.AutoMigrate(&Inventory{})
	fmt.Println("seedDataInventory Done")
}

func seedDataStaff(db *gorm.DB) {
	db.DropTableIfExists(&Staff{})
	db.AutoMigrate(&Staff{})
	fmt.Println("staff Done")
}

func seedDataLocationStaff(db *gorm.DB) {
	db.DropTableIfExists(&LocationStaff{})
	db.AutoMigrate(&LocationStaff{})
	fmt.Println("LocationStaff Done")
}
func seedDataLocationUser(db *gorm.DB) {
	db.DropTableIfExists(&LocationUser{})
	db.AutoMigrate(&LocationUser{})
	fmt.Println("LocationUser Done")
}
func seedData(db *gorm.DB) {

	//seedDataCountries(db)
	seedDataCountries(db)
	seedDataStates(db)
	seedDataCity(db)
	seedDataLocation(db)
	seedStaff(db)
	seedLocationStaff(db)
	seedDataCategory(db)
	seedDataCategoryDetail(db)
	seedDataFeatures(db)
	seedDataLocationFeatures(db)
	//seedDataLocationFeatures(db)
	//seedDataFeatures(db)
	//seedDataLocationProductType(db)
	//seedDataLocationGallery(db)
	//seedDataInventory(db)
	//seedDataLocationContactUser(db)
	//seedDataLocationUser(db)
}

func setupDBConnection(configuration *config.Config) (string, string) {

	// Connect to the "bank" database as the "maxroach" user.
	cdbHost := configuration.Items.CockroachDB.Host
	cdbPort := configuration.Items.CockroachDB.Port
	cdbDatabase := configuration.Items.CockroachDB.Database
	cdbDefaultDatabase := configuration.Items.CockroachDB.DefaultDatabase
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
	computedDefaultConnectionString := fmt.Sprintf("postgresql://%s@%s:%s/%s%s", cdbUser, cdbHost, cdbPort, cdbDefaultDatabase, cdbSSLMode)

	fmt.Printf("computed connection string is = %v \n", computedConnectionString)
	fmt.Printf("computed connection string is = %v \n", computedDefaultConnectionString)
	return computedConnectionString, computedDefaultConnectionString
}

func main() {

	configuration, errConfig := config.New()

	if errConfig != nil {
		fmt.Printf("error reading config %v \n", errConfig)
	}

	//setupDBConnection(configuration)
	const addr = "postgresql://root@192.168.0.11:26257/lp7_location?sslmode=disable"
	//const addr = "postgres://root@162.252.80.136:26257/eventvisor_test?sslmode=verify-full&sslrootcert=certs-162.252.80.136/ca.crt&sslcert=certs-162.252.80.136/client.root.crt&sslkey=certs-162.252.80.136/client.root.key"
	fmt.Println("openning db")
	fmt.Println("configuration jwtpass = " + configuration.Items.JWT.JWTPASS)

	dbCon, dbConDefualt := setupDBConnection(configuration)

	// open default postgres db to create our database
	dbDefault, errdbDefault := gorm.Open("postgres", dbConDefualt)
	if errdbDefault != nil {
		fmt.Println("error from db returned")
		fmt.Printf("%v", errdbDefault)
		log.Fatal(errdbDefault)
	}
	dbDefault.LogMode(true)
	defer dbDefault.Close()
	// creating our target database and close default connection
	targetDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v;", configuration.Items.CockroachDB.Database)
	fmt.Printf("targetDB command is = %v \n", targetDBQuery)

	res := dbDefault.Exec(targetDBQuery)
	fmt.Printf("create database res :%+v \n", res)

	// creating connection to our target database and pass to seeding func
	db, err := gorm.Open("postgres", dbCon)
	if err != nil {
		fmt.Println("error from db returned")
		fmt.Printf("%v", err)
		log.Fatal(err)
	}
	db.LogMode(true)
	defer db.Close()

	//CREATE DATABASE eventvisor;

	// if you want to seed data then uncomment following line.

	//blocking this since database is generated. @najam[]
	fmt.Println("configuration.Items.Environment:", configuration.Items)
	switch envMode := configuration.Items.Environment; envMode {
	case "dev":
		seedData(db)
	case "local":
		seedData(db)
	default:
		fmt.Println("not seeding data becuse of Environment settings")
	}

}
