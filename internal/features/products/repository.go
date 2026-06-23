package products

import (
    "database/sql"
    "errors"
    "fmt"

    "faro/internal/core/database"
)

// Repository maneja las operaciones de base de datos para productos
type Repository struct {
    db *sql.DB
}

// NewRepository crea una nueva instancia del repositorio
func NewRepository() *Repository {
    return &Repository{
        db: database.DB,
    }
}

// GetAll retorna todos los productos con el nombre de su categoría (JOIN)
func (r *Repository) GetAll() ([]Product, error) {
    query := `
        SELECT p.id, p.name, p.category_id, p.price, p.stock, p.status, p.created_at, COALESCE(c.name, '') 
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.id
        ORDER BY p.id DESC
    `

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error consultando productos: %w", err)
    }
    defer rows.Close()

    var products []Product

    for rows.Next() {
        var p Product
        if err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock, &p.Status, &p.CreatedAt, &p.CategoryName); err != nil {
            return nil, fmt.Errorf("error leyendo producto: %w", err)
        }
        products = append(products, p)
    }

    // Verificar errores después de iterar
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterando productos: %w", err)
    }

    // Retornar array vacío en vez de nil (evita null en el JSON del frontend)
    if products == nil {
        products = []Product{}
    }

    return products, nil
}

// GetByID retorna un producto por su ID junto con el nombre de categoría
func (r *Repository) GetByID(id int64) (*Product, error) {
    query := `
        SELECT p.id, p.name, p.category_id, p.price, p.stock, p.status, p.created_at, COALESCE(c.name, '') 
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.id
        WHERE p.id = ?
    `

    var p Product
    err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock, &p.Status, &p.CreatedAt, &p.CategoryName)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil // No encontrado no es error fatal
        }
        return nil, fmt.Errorf("error consultando producto %d: %w", id, err)
    }

    return &p, nil
}

// Create inserta un nuevo producto y retorna su ID
func (r *Repository) Create(req CreateRequest) (int64, error) {
    query := `
        INSERT INTO products (name, category_id, price, stock, status) 
        VALUES (?, ?, ?, ?, ?)
    `

    result, err := r.db.Exec(query, req.Name, req.CategoryID, req.Price, req.Stock, req.Status)
    if err != nil {
        return 0, fmt.Errorf("error creando producto: %w", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("error obteniendo ID del producto: %w", err)
    }

    return id, nil
}

// Update actualiza un producto existente
func (r *Repository) Update(id int64, req CreateRequest) error {
    query := `
        UPDATE products 
        SET name = ?, category_id = ?, price = ?, stock = ?, status = ? 
        WHERE id = ?
    `

    result, err := r.db.Exec(query, req.Name, req.CategoryID, req.Price, req.Stock, req.Status, id)
    if err != nil {
        return fmt.Errorf("error actualizando producto: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

    return nil
}


// Delete elimina un producto por ID
func (r *Repository) Delete(id int64) error {
    query := `DELETE FROM products WHERE id = ?`

    result, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error eliminando producto: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

    return nil
}