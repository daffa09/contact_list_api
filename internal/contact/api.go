package contact

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/<id>", res.get)
	r.Get("", res.query)

	// r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("", res.create)
	r.Patch("/<id>", res.update)
	r.Delete("/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	contact, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(contact)
}

func (r resource) query(c *routing.Context) error {

	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}

	pages := pagination.NewFromRequest(c.Request, count)
	contact, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}

	pages.Items = contact
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateContactRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	contact, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(contact, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateContactRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	contact, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(contact)
}

func (r resource) delete(c *routing.Context) error {
	contact, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(contact)
}
