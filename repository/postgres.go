package repository

import (
	"fmt"

	"github.com/AntonyIS/go-joint/app"
	"github.com/AntonyIS/go-joint/config"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func postgresClient() (*gorm.DB, error) {
	var (
		host     = config.GetEnvVariable("HOST")
		port     = config.GetEnvVariable("PORT")
		user     = "postgres" //config.GetEnvVariable("USER")
		dbname   = config.GetEnvVariable("DBNAME")
		password = config.GetEnvVariable("PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type attendeeRepo struct {
	client *gorm.DB
}

func NewAttendeeRepository() (app.AttendeeRepository, error) {
	repo := attendeeRepo{}
	client, err := postgresClient()
	if err != nil {
		return nil, err
	}
	repo.client = client
	repo.client.AutoMigrate(&app.Attendee{})
	return repo, nil
}

func (repo attendeeRepo) Create(attendee *app.Attendee) (*app.Attendee, error) {
	res := repo.client.Create(&attendee)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not created")
	}
	return attendee, nil
}

func (repo attendeeRepo) Read(id string) (*app.Attendee, error) {
	var attendee app.Attendee
	res := repo.client.First(&attendee, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not found")
	}
	return &attendee, nil
}

func (repo attendeeRepo) ReadAll() ([]*app.Attendee, error) {
	var attendees []*app.Attendee
	res := repo.client.Find(&attendees)

	if res.Error != nil {
		return nil, errors.New("attendeees not found")
	}
	return attendees, nil
}

func (repo attendeeRepo) Update(attendee *app.Attendee) (*app.Attendee, error) {
	var updateAttendee app.Attendee
	res := repo.client.First(&updateAttendee, "id = ?", attendee.ID)
	if res.RowsAffected == 0 {
		return nil, errors.Wrap(app.ErrInternalServer, res.Error.Error())
	}
	res = repo.client.Model(updateAttendee).Where("id = ?", attendee.ID).Updates(attendee)

	if res.RowsAffected == 0 {
		return nil, errors.Wrap(app.ErrInternalServer, res.Error.Error())
	}
	return attendee, nil
}

func (repo attendeeRepo) Delete(id string) error {
	var deleteAttendee app.Attendee
	res := repo.client.Where("id = ?", id).Delete(&deleteAttendee)
	if res.RowsAffected == 0 {
		return errors.New("attendee not deleted")
	}
	return nil
}
