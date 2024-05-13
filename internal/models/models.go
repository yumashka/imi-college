package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	RoleRegular string = "regular"
	RoleStaff   string = "staff"
	RoleAdmin   string = "admin"
)

type User struct {
	ID         uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid();not null;" json:"id"`
	UserName   string     `gorm:"uniqueIndex;not null;" json:"username"`
	Email      string     `gorm:"uniqueIndex;not null;" json:"email"`
	IsVerified bool       `gorm:"default:false;not null;" json:"isVerified"`
	Role       string     `gorm:"default:'regular';not null;" json:"role"`
	Tel        *string    `json:"tel"`
	FirstName  string     `gorm:"not null;" json:"firstName"`
	MiddleName string     `gorm:"not null;" json:"middleName"`
	LastName   string     `gorm:"not null;" json:"lastName"`
	Birthday   *time.Time `json:"birthday"`
	CreatedAt  time.Time  `gorm:"default:now();not null;" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"default:now();not null;" json:"updatedAt"`
}

type Password struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();not null;"`
	User   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID uuid.UUID `gorm:"uniqueIndex;not null;"`
	Hash   string    `gorm:"not null;"`
}

type UserToken struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();not null;" json:"id"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null;" json:"-"`
	UserID    uuid.UUID `gorm:"not null;" json:"userId"`
	ExpiresAt time.Time `gorm:"default:now() + interval '2 days';not null;" json:"expiresAt"`
	CreatedAt time.Time `gorm:"default:now();not null;" json:"createdAt"`
	Token     string    `gorm:"uniqueIndex;not null;" json:"token"`
}

type DocumentType uint8

const (
	DocumentPassport              DocumentType = 0
	DocumentInternationalPassport DocumentType = 1
	DocumentForeignPassport       DocumentType = 2
	DocumentBirthCertificate      DocumentType = 3
	DocumentMilitaryCard          DocumentType = 4
	DocumentOther                 DocumentType = 5
)

type UserDocFile struct {
	ID           uuid.UUID    `gorm:"primaryKey;type:uuid;default:gen_random_uuid();not null;" json:"id"`
	User         User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null;" json:"-"`
	UserID       uuid.UUID    `gorm:"not null;index:user_id__type_uindex,unique;" json:"userId"`
	Type         DocumentType `gorm:"not null;index:user_id__type_uindex,unique;" json:"type"`
	Series       string       `gorm:"not null;" json:"series"`
	Number       string       `gorm:"not null;" json:"number"`
	Issuer       string       `gorm:"not null;" json:"issuer"`
	IssuedAt     time.Time    `gorm:"not null;" json:"issuedAt"`
	DivisionCode string       `gorm:"not null;" json:"divisionCode"`
	Country      string       `gorm:"not null;" json:"country"`
}
