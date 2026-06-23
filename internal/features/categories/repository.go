package categories

import (
    "database/sql"
    "errors"
    "fmt"

    "faro/internal/core/database"
)

// Repository maneja las operaciones de base de datos para categorías
type Repository struct {
    db *sql.DB
}

// NewRepository crea una nueva instancia del repositorio
func NewRepository() *Repository {
    return &Repository{
        db: database.DB,
    }
}

// GetAll retorna todas las categorías ordenadas por nombre
func (r *Repository) GetAll() ([]Category, error) {
    query := `
        SELECT id, name, COALESCE(description, '') 
        FROM categories 
        ORDER BY name ASC
    `

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error consultando categorías: %w", err)
    }
    defer rows.Close()

    var categories []Category

    for rows.Next() {
        var c Category
        if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
            return nil, fmt.Errorf("error leyendo categoría: %w", err)
        }
        categories = append(categories, c)
    }

    // Verificar errores después de iterar
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterando categorías: %w", err)
    }

    // Retornar array vacío en vez de nil (mejor para el frontend JSON)
    if categories == nil {
        categories = []Category{}
    }

    return categories, nil
}

// GetByID retorna una categoría por su ID
func (r *Repository) GetByID(id int64) (*Category, error) {
    query := `
        SELECT id, name, COALESCE(description, '') 
        FROM categories 
        WHERE id = ?
    `

    var c Category
    err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Description)
    
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil // No encontrado no es error, retorna nil
        }
        return nil, fmt.Errorf("error consultando categoría %d: %w", id, err)
    }

    return &c, nil
}

// Create inserta una nueva categoría y retorna su ID
func (r *Repository) Create(req CategoryCreateRequest) (int64, error) {
    query := `
        INSERT INTO categories (name, description) 
        VALUES (?, ?)
    `

    result, err := r.db.Exec(query, req.Name, req.Description)
    if err != nil {
        return 0, fmt.Errorf("error creando categoría: %w", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("error obteniendo ID: %w", err)
    }

    return id, nil
}

// Update actualiza una categoría existente
func (r *Repository) Update(id int64, req CategoryUpdateRequest) error {
    query := `
        UPDATE categories 
        SET name = ?, description = ? 
        WHERE id = ?
    `

    result, err := r.db.Exec(query, req.Name, req.Description, id)
    if err != nil {
        return fmt.Errorf("error actualizando categoría: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

    return nil
}

// Delete elimina una categoría por ID
// NOTA: Fallará si hay productos asociados (FOREIGN KEY RESTRICT)
func (r *Repository) Delete(id int64) error {
    query := `DELETE FROM categories WHERE id = ?`

    result, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error eliminando categoría: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

    return nil
}

// ExistsByName verifica si ya existe una categoría con ese nombre
// Se usa para validación antes de crear/actualizar
func (r *Repository) ExistsByName(name string, excludeID ...int64) (bool, error) {
    query := `SELECT COUNT(1) FROM categories WHERE name = ?`
    var args []interface{}
    args = append(args, name)

    if len(excludeID) > 0 {
        query += ` AND id != ?`
        args = append(args, excludeID[0])
    }

    var count int
    err := r.db.QueryRow(query, args...).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("error verificando nombre: %w", err)
    }

    return count > 0, nil
}