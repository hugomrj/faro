package categories

// Controller maneja la lógica de negocio y expone los métodos al frontend
type Controller struct {
    repo *Repository
}

// NewController crea un nuevo controlador inicializando su repositorio
func NewController() *Controller {
    return &Controller{
        repo: NewRepository(),
    }
}

// GetAll retorna todas las categorías
func (c *Controller) GetAll() ([]Category, error) {
    return c.repo.GetAll()
}

// GetByID retorna una categoría específica
func (c *Controller) GetByID(id int64) (*Category, error) {
    return c.repo.GetByID(id)
}

// Create crea una nueva categoría
func (c *Controller) Create(req CategoryCreateRequest) (int64, error) {
    return c.repo.Create(req)
}

// Update actualiza una categoría existente
func (c *Controller) Update(id int64, req CategoryUpdateRequest) error {
    return c.repo.Update(id, req)
}

// Delete elimina una categoría
func (c *Controller) Delete(id int64) error {
    return c.repo.Delete(id)
}
