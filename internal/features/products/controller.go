package products

import "log/slog"

type Controller struct {
    repo *Repository
}

func NewController() *Controller {
    return &Controller{
        repo: NewRepository(),
    }
}

func (c *Controller) GetAll() ([]Product, error) {
    return c.repo.GetAll()
}

func (c *Controller) GetByID(id int64) (*Product, error) {
    return c.repo.GetByID(id)
}

func (c *Controller) Create(req CreateRequest) (int64, error) {
    id, err := c.repo.Create(req)
    if err != nil {
        slog.Error("Error creando producto", "error", err)
        return 0, err
    }
    return id, nil
}

func (c *Controller) Update(id int64, req CreateRequest) error {
    return c.repo.Update(id, req)
}

func (c *Controller) Delete(id int64) error {
    err := c.repo.Delete(id)
    if err != nil {
        slog.Error("Error eliminando producto", "id", id, "error", err)
        return err
    }
    return nil
}