package pg

import (
	"database/sql"
	dto2 "employee-hierarchy-api/internal/dto"
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnector interface {
	Connect() error
	Close() error
	GetDB() *gorm.DB
}

type PostgreSQLConnector struct {
	db    *gorm.DB
	err   error
	sqlDB *sql.DB
}

// Connect establishes a new database connection
func (p *PostgreSQLConnector) Connect() error {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	port := 5432
	sslMode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host, username, password, database, port, sslMode)

	p.db, p.err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if p.err != nil {
		return errors.New(fmt.Sprintf("failed to connect database %v", p.err))
	}

	// Migrate
	p.err = p.db.AutoMigrate(&dto2.User{}, &dto2.Employee{})
	if p.err != nil {
		return errors.New(fmt.Sprintf("failed to migrate database %v", p.err))
	}
	if p.err = p.seedData(); p.err != nil {
		return p.err
	}

	return nil
}

func (p *PostgreSQLConnector) Close() error {
	if p.db != nil {
		p.sqlDB, p.err = p.db.DB()
		if p.err != nil {
			return p.err
		}
		p.err = p.sqlDB.Close()
		if p.err != nil {
			return p.err
		}
	}
	return nil
}

func (p *PostgreSQLConnector) GetDB() *gorm.DB {
	return p.db
}

func (p *PostgreSQLConnector) seedData() error {
	employees := []dto2.Employee{
		{Name: "Jonas"},
		{Name: "Sophie"},
		{Name: "Nick"},
		{Name: "Pete"},
		{Name: "Barbara"},
	}

	for _, employee := range employees {
		if p.err = p.db.Where("name = ?", employee.Name).First(&dto2.Employee{}).Error; p.err != nil {
			if errors.Is(p.err, gorm.ErrRecordNotFound) {
				if p.err = p.db.Create(&employee).Error; p.err != nil {
					log.Fatal("Error seeding data: ", p.err)
				}
			} else {
				log.Fatal("Error checking existing data: ", p.err)
			}
		}
	}

	// Set supervisors
	if p.err = p.setSupervisor("Sophie", "Nick"); p.err != nil {
		log.Fatal("Error setting supervisor: ", p.err)
	}
	if p.err = p.setSupervisor("Nick", "Jonas"); p.err != nil {
		log.Fatal("Error setting supervisor: ", p.err)
	}
	if p.err = p.setSupervisor("Pete", "Nick"); p.err != nil {
		log.Fatal("Error setting supervisor: ", p.err)
	}
	if p.err = p.setSupervisor("Barbara", "Nick"); p.err != nil {
		log.Fatal("Error setting supervisor: ", p.err)
	}
	return nil
}

func (p *PostgreSQLConnector) setSupervisor(employeeName, supervisorName string) error {
	var employee, supervisor dto2.Employee

	if p.err = p.db.Where("name = ?", employeeName).First(&employee).Error; p.err != nil {
		return p.err
	}
	if p.err = p.db.Where("name = ?", supervisorName).First(&supervisor).Error; p.err != nil {
		return p.err
	}

	employee.SupervisorID = &supervisor.ID
	return p.db.Save(&employee).Error
}
