package pg

import (
	"employee-hierarchy-api/external/dto"
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Connect establishes a new database connection
func Connect() error {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	port := 5432
	sslMode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host, username, password, database, port, sslMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect database %v", err))
	}

	// Migrate
	err = db.AutoMigrate(&dto.User{}, &dto.Employee{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to migrate database %v", err))
	}

	if err = seedData(db); err != nil {
		return err
	}

	return nil
}

func ConnectDBTest() error {
	host := "localhost"
	username := "fram"
	password := "fram"
	database := "employeeHierarchy"
	port := 5432
	sslMode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host, username, password, database, port, sslMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect database %v", err))
	}
	return nil
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}

func seedData(db *gorm.DB) error {
	employees := []dto.Employee{
		{Name: "Jonas"},
		{Name: "Sophie"},
		{Name: "Nick"},
		{Name: "Pete"},
		{Name: "Barbara"},
	}

	for _, employee := range employees {
		if err = db.Where("name = ?", employee.Name).First(&dto.Employee{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&employee).Error; err != nil {
					log.Fatal("Error seeding data: ", err)
				}
			} else {
				log.Fatal("Error checking existing data: ", err)
			}
		}
	}

	// Set supervisors
	if err = setSupervisor(db, "Sophie", "Nick"); err != nil {
		log.Fatal("Error setting supervisor: ", err)
	}
	if err = setSupervisor(db, "Nick", "Jonas"); err != nil {
		log.Fatal("Error setting supervisor: ", err)
	}
	if err = setSupervisor(db, "Pete", "Nick"); err != nil {
		log.Fatal("Error setting supervisor: ", err)
	}
	if err = setSupervisor(db, "Barbara", "Nick"); err != nil {
		log.Fatal("Error setting supervisor: ", err)
	}
	return nil
}

func setSupervisor(db *gorm.DB, employeeName, supervisorName string) error {
	var employee, supervisor dto.Employee

	if err = db.Where("name = ?", employeeName).First(&employee).Error; err != nil {
		return err
	}
	if err = db.Where("name = ?", supervisorName).First(&supervisor).Error; err != nil {
		return err
	}

	employee.SupervisorID = &supervisor.ID
	return db.Save(&employee).Error
}
