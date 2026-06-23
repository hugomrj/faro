// frontend/src/types/index.ts

// En lugar de depender de los tipos generados por Wails,
// definimos los tipos aquí para mayor control

export interface Category {
    id: number
    name: string
    description: string
  }
  
  export interface CategoryCreateRequest {
    name: string
    description: string
  }
  
  export interface CategoryUpdateRequest {
    name: string
    description: string
  }
  
  export interface Product {
    id: number
    name: string
    category_id: number
    price: number
    stock: number
    status: string
    created_at: string
    // Campo opcional para join con categoría
    category_name?: string
  }
  
  export interface ProductCreateRequest {
    name: string
    category_id: number
    price: number
    stock: number
    status: string
  }