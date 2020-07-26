package repository

import (
	"github.com/NeoAssist/NeoAcademy/internal/pkg/database/models"
	"github.com/NeoAssist/NeoAcademy/internal/pkg/database/store"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (as *store.AccountStore) GetByID(id uuid.UUID) (*models.Account, error) {
	var m models.Account
	if err := as.db.Where(&models.Account{ID: id}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// GetByEmail .
func (as *AccountStore) GetByEmail(e string) (*models.Account, error) {
	var m models.Account
	if err := as.db.Where(&models.Account{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// GetByUsername .
func (as *AccountStore) GetByUsername(email string) (*models.Account, error) {
	var m models.Account
	if err := as.db.Where(&models.Account{Email: email}).Preload("Followers").First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// Create .
func (as *AccountStore) Create(u *models.Account) (err error) {
	return as.db.Create(u).Error
}

// Update .
func (as *AccountStore) Update(u *models.Account) error {
	transaction := as.db.Begin()

	if err := transaction.Model(u).Update(u).Error; err != nil {
		return err
	}

	return transaction.Commit().Error
}
