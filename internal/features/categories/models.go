package categories

// Category representa una categoría del sistema
type Category struct {
    ID          int64  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

// CategoryCreateRequest se usa para crear nuevas categorías
// Separamos del modelo principal para no exponer ID ni campos auto-generados
type CategoryCreateRequest struct {
    Name        string `json:"name" binding:"required,min=2,max=100"`
    Description string `json:"description" binding:"max=500"`
}

// CategoryUpdateRequest se usa para actualizar categorías existentes
type CategoryUpdateRequest struct {
    Name        string `json:"name" binding:"required,min=2,max=100"`
    Description string `json:"description" binding:"max=500"`
}