package contact

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id string) (entity.Contact, error)
	// Count returns the number of albums.
	Count(ctx context.Context) (int, error)
	// Query returns the list of albums with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Contact, error)
	// Create saves a new album in the storage.
	Create(ctx context.Context, contact entity.Contact) error
	// Update updates the album with given ID in the storage.
	Update(ctx context.Context, contact entity.Contact) error
	// Delete removes the album with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists albums in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.Contact, error) {
	var contact entity.Contact
	err := r.db.With(ctx).Select().Model(id, &contact)
	return contact, err
}

func (r repository) Create(ctx context.Context, contact entity.Contact) error {
	return r.db.With(ctx).Model(&contact).Insert()
}

func (r repository) Update(ctx context.Context, contact entity.Contact) error {
	return r.db.With(ctx).Model(&contact).Update()
}

func (r repository) Delete(ctx context.Context, id string) error {
	contact, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&contact).Delete()
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("contact").Row(&count)
	return count, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Contact, error) {
	var contact []entity.Contact
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&contact)
	return contact, err
}
