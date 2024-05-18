package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Password{},
		&UserToken{},
		&UserIdentity{},
		&UserAddress{},
		&Application{},
		&DictAppState{},
		&DictEduDocType{},
		&DictIdDocType{},
		&DictEduLevel{},
		&DictCountry{},
		&DictRegion{},
		&DictTownType{},
		&DictGender{},
		&EduDoc{},
	)
}

const (
	RoleRegular string = "regular"
	RoleStaff   string = "staff"
	RoleAdmin   string = "admin"
)

type User struct {
	ID         uuid.UUID `gorm:"not null;primaryKey;type:uuid;default:gen_random_uuid();" json:"id"`
	CreatedAt  time.Time `gorm:"not null;default:now();" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"not null;default:now();" json:"updatedAt"`
	UserName   string    `gorm:"not null;uniqueIndex;" json:"username"`
	Email      string    `gorm:"not null;uniqueIndex;" json:"email"`
	IsVerified bool      `gorm:"not null;default:false;" json:"isVerified"`
	Role       string    `gorm:"not null;default:'regular';" json:"role"`
}

type Password struct {
	ID     uuid.UUID `gorm:"not null;primaryKey;type:uuid;default:gen_random_uuid();"`
	User   User      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID uuid.UUID `gorm:"not null;uniqueIndex;"`
	Hash   string    `gorm:"not null;"`
}

type UserToken struct {
	ID        uuid.UUID `gorm:"not null;primaryKey;type:uuid;default:gen_random_uuid();" json:"id"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID    uuid.UUID `gorm:"not null;" json:"userId"`
	CreatedAt time.Time `gorm:"not null;default:now();" json:"createdAt"`
	ExpiresAt time.Time `gorm:"not null;default:now() + interval '2 days';" json:"expiresAt"`
	Token     string    `gorm:"not null;uniqueIndex;" json:"token"`
}

type UserIdentity struct {
	ID         uuid.UUID      `gorm:"not null;primaryKey;type:uuid;default:gen_random_uuid();" json:"id"`
	User       User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID     uuid.UUID      `gorm:"not null;uniqueIndex;" json:"userId"`
	FirstName  string         `gorm:"not null;" json:"firstName"`
	MiddleName string         `gorm:"not null;" json:"middleName"`
	LastName   string         `gorm:"not null;" json:"lastName"`
	Gender     DictGender     `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	GenderID   int            `gorm:"not null;" json:"genderId"`
	Birthday   sql.NullTime   `json:"birthday"`
	Tel        sql.NullString `json:"tel"`
	SNILS      sql.NullString `json:"snils"`
}

type UserAddress struct {
	ID         uuid.UUID    `gorm:"not null;primaryKey;type:uuid;default:gen_random_uuid();" json:"id"`
	Region     DictRegion   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	RegionID   int          `gorm:"not null;" json:"regionId"`
	TownType   DictTownType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	TownTypeID int          `gorm:"not null;" json:"townTypeId"`
	Town       string       `gorm:"not null;" json:"town"`
	PostCode   string       `gorm:"not null;" json:"postCode"`
}

type Application struct {
	ID        uuid.UUID    `gorm:"not null;primaryKey;type:uuid;default:gen_random_uuid();" json:"id"`
	User      User         `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID    uuid.UUID    `json:"userId"`
	State     DictAppState `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	StateID   int          `gorm:"not null;" json:"stateId"`
	CreatedAt time.Time    `gorm:"not null;default:now();" json:"createdAt"`
}

type DictAppState struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
}

type DictEduDocType struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
}

type DictIdDocType struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
}

type DictEduLevel struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
}

type DictCountry struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
	SortPriority int            `gorm:"not null;default:0;" json:"sortPriority"`
}

type DictRegion struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	RegionID     int            `gorm:"not null;uniqueIndex;" json:"regionId"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
	SortPriority int            `gorm:"not null;default:0;" json:"sortPriority"`
}

type DictTownType struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
}

type DictGender struct {
	ID           int            `gorm:"not null;primaryKey;" json:"id"`
	Value        string         `gorm:"not null;" json:"value"`
	DisplayValue sql.NullString `json:"displayValue"`
}

type EduDoc struct {
	ID             uuid.UUID      `gorm:"not null;type:uuid;default:gen_random_uuid();" json:"id"`
	CreatedAt      time.Time      `gorm:"not null;default:now();" json:"createdAt"`
	User           User           `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID         uuid.UUID      `gorm:"not null;type:uuid;" json:"userId"`
	Type           DictEduDocType `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	TypeID         int            `gorm:"not null;" json:"typeId"`
	Series         string         `gorm:"not null;" json:"series"`
	Number         string         `gorm:"not null;" json:"number"`
	Issuer         string         `gorm:"not null;" json:"issuer"`
	IssuedAt       time.Time      `gorm:"not null;" json:"issuedAt"`
	GradYear       uint8          `gorm:"not null;" json:"gradYear"`
	IssuerRegion   DictRegion     `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	IssuerRegionID int            `gorm:"not null;" json:"issuerRegionId"`
}
