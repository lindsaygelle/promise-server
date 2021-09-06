package account

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/database"
)

type Setting struct {
	SettingBiography

	Account  uint      `json:"account"`
	Country  uint      `json:"country"`
	Edited   time.Time `json:"edited"`
	Language uint      `json:"language"`
}

type SettingBiography struct {
	Biography string `json:"biography"`
}

type SettingMake struct {
	SettingBiography

	Account  uint `json:"account"`
	Country  uint `json:"country"`
	Language uint `json:"language"`
}

type SettingMakeValidator func(SettingMake) error

var (
	ErrNoCountry  = errors.New("country is empty")
	ErrNoLanguage = errors.New("language is empty")
)

var (
	settingMakeValidators = [...]SettingMakeValidator{
		validateSettingCountry,
		validateSettingLanguage}
)

func EditSettingsBiography(database *sql.DB, ID string, request *http.Request) (err error) {
	return
}

// NewSetting returns a new Setting from the database.
func NewSetting(scanner database.Scanner) (setting Setting, err error) {
	err = scanner.Scan(&setting.Account, &setting.Biography, &setting.Country, &setting.Edited, &setting.Language)
	return
}

// NewSettingBiography returns a new SettingBiography from an io.ReadCloser.
func NewSettingBiography(reader io.ReadCloser) (settingBiography SettingBiography, err error) {
	err = json.NewDecoder(reader).Decode(&settingBiography)
	return
}

// NewSettingMake returns a new SettingMake from an io.ReadCloser.
func NewSettingMake(reader io.ReadCloser) (settingMake SettingMake, err error) {
	err = json.NewDecoder(reader).Decode(&settingMake)
	return
}

// ReadSetting reads an setting from the database by its ID.
func ReadSetting(database *sql.DB, ID string) (setting Setting, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(`SELECT account, biography, country, edited, language FROM account.setting WHERE account=$1`, ID)
	setting, err = NewSetting(row)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// ReadSettings reads all settings from the database.
func ReadSettings(database *sql.DB) (settings []Setting, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	rows, err := transaction.Query(`SELECT account, biography, country, edited, language FROM account.setting`)
	if err != nil {
		transaction.Rollback()
		return
	}
	err = processSettings(&settings, rows)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// WriteSetting writes a new Setting to the database.
func WriteSetting(context *gin.Context, database *sql.DB) (setting Setting, err error) {
	defer context.Request.Body.Close()
	settingMake, err := NewSettingMake(context.Request.Body)
	if err != nil {
		return
	}
	err = validateSettingMake(settingMake)
	if err != nil {
		return
	}
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(`INSERT INTO account.setting (biography, country, language) VALUES ($1, $2, $3) RETURNING account, biography, country, edited, language`, settingMake.Biography, setting.Country, setting.Language)
	setting, err = NewSetting(row)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// addSetting scans a account.setting record from the database rows and adds it to the collection.
//
// addSetting returns an error on the condition the setting cannot be scanned.
func addSetting(settings *[]Setting, rows *sql.Rows) (err error) {
	setting, err := NewSetting(rows)
	if err != nil {
		return err
	}
	*settings = append(*settings, setting)
	return err
}

// processSettings processes all rows from the database and adds them to the collection.
//
// processSettings returns an error on the condition the row cannot be added.
func processSettings(settings *[]Setting, rows *sql.Rows) (err error) {
	for rows.Next() {
		err = addSetting(settings, rows)
		if err != nil {
			return
		}
	}
	return
}

// validateSettingMake validates an SettingMake.
func validateSettingMake(settingMake SettingMake) (err error) {
	for _, fn := range settingMakeValidators {
		err = fn(settingMake)
		if err != nil {
			break
		}
	}
	return err
}

// validateSettingCountry validates SettingMake.Country.
func validateSettingCountry(settingMake SettingMake) (err error) {
	if settingMake.Country == 0 {
		err = ErrNoCountry
	}
	return
}

// validateSettingLanguage validates SettingMake.Language.
func validateSettingLanguage(settingMake SettingMake) (err error) {
	if settingMake.Language == 0 {
		err = ErrNoLanguage
	}
	return
}
