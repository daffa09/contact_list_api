package contact

import (
	"strconv"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"

	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for contact.
type Service interface {
	Get(ctx context.Context, id string) (Contact, error)
	Query(ctx context.Context, offset, limit int) ([]Contact, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateContactRequest) (Contact, error)
	Update(ctx context.Context, id string, input UpdateContactRequest) (Contact, error)
	Delete(ctx context.Context, id string) (Contact, error)
}

// contact represents the data about an contact.
type Contact struct {
	entity.Contact
}

// CreateContactRequest represents an contact creation request.
type CreateContactRequest struct {
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}

// Validate validates the CreateContactRequest fields.
func (m CreateContactRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateContactRequest represents an contact update request.
type UpdateContactRequest struct {
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}

// Validate validates the CreateContactRequest fields.
func (m UpdateContactRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new contact service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the contact with the specified the contact ID.
func (s service) Get(ctx context.Context, id string) (Contact, error) {
	contact, err := s.repo.Get(ctx, id)
	if err != nil {
		return Contact{}, err
	}
	return Contact{contact}, nil
}

func (s service) Create(ctx context.Context, req CreateContactRequest) (Contact, error) {
	if err := req.Validate(); err != nil {
		return Contact{}, err
	}
	id := entity.GenerateIDInt()
	err := s.repo.Create(ctx, entity.Contact{
		ID:    id,
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
		Phone: req.Phone,
	})
	if err != nil {
		return Contact{}, err
	}

	idNew := strconv.Itoa(id)

	return s.Get(ctx, idNew)
}

// Update updates the contact with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateContactRequest) (Contact, error) {
	if err := req.Validate(); err != nil {
		return Contact{}, err
	}

	contact, err := s.Get(ctx, id)
	if err != nil {
		return contact, err
	}
	contact.Name = req.Name
	contact.Age = req.Age
	contact.Email = req.Email
	contact.Phone = req.Phone

	if err := s.repo.Update(ctx, contact.Contact); err != nil {
		return contact, err
	}
	return contact, nil
}

// Delete deletes the contact with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Contact, error) {
	contact, err := s.Get(ctx, id)
	if err != nil {
		return Contact{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Contact{}, err
	}
	return contact, nil
}

// Count returns the number of contact.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the contact with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Contact, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Contact{}
	for _, item := range items {
		result = append(result, Contact{item})
	}
	return result, nil
}
