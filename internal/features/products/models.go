package products

type Product struct {
    ID           int64   `json:"id"`
    Name         string  `json:"name"`
    CategoryID   int64   `json:"category_id"`
    Price        float64 `json:"price"`
    Stock        int     `json:"stock"`
    Status       string  `json:"status"`
    CreatedAt    string  `json:"created_at"`
    CategoryName string  `json:"category_name,omitempty"` // Para el JOIN
}

type CreateRequest struct {
    Name       string  `json:"name"`
    CategoryID int64   `json:"category_id"`
    Price      float64 `json:"price"`
    Stock      int     `json:"stock"`
    Status     string  `json:"status"`
}