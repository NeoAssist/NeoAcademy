package store

import (
	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/model"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// AccountStore .
type AccountStore struct {
	db *gorm.DB
}

// NewAccountStore .
func NewAccountStore(db *gorm.DB) *AccountStore {
	return &AccountStore{
		db: db,
	}
}

// GetByID .
func (as *AccountStore) GetByID(id uuid.UUID) (*model.Account, error) {
	var m model.Account
	if err := as.db.Where(&model.Account{ID: id}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// GetByEmail .
func (as *AccountStore) GetByEmail(e string) (*model.Account, error) {
	var m model.Account
	if err := as.db.Where(&model.Account{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// GetByUsername .
func (as *AccountStore) GetByUsername(email string) (*model.Account, error) {
	var m model.Account
	if err := as.db.Where(&model.Account{Email: email}).Preload("Followers").First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// Create .
func (as *AccountStore) Create(u *model.Account) (err error) {
	return as.db.Create(u).Error
}

// Update .
func (as *AccountStore) Update(u *model.Account) error {
	transaction := as.db.Begin()

	if err := transaction.Model(u).Update(u).Error; err != nil {
		return err
	}

	return transaction.Commit().Error
}
