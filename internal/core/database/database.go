package database

import (
    "database/sql"
    "fmt"
    "log/slog"
    "os"
    "path/filepath"

    _ "modernc.org/sqlite"
)

// DB expone la conexión para usarla en los repositorios
var DB *sql.DB

// Schema SQL inicial
const schema = `
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    price REAL NOT NULL DEFAULT 0,
    stock INTEGER NOT NULL DEFAULT 0,
    status TEXT NOT NULL DEFAULT 'Activo',
    created_at TEXT NOT NULL DEFAULT (datetime('now','localtime')),
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT
);

-- Índice para optimizar búsquedas por categoría
CREATE INDEX IF NOT EXISTS idx_products_category ON products(category_id);
`

// Init inicializa la conexión a SQLite y crea las tablas
func Init(dbPath string) error {
    var err error

    // Asegurar que el directorio existe
    dir := filepath.Dir(dbPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return fmt.Errorf("error creando directorio de BD: %w", err)
    }

    // Abrir conexión
    // ?_pragma=... habilita foreign keys y journal mode WAL (mejor rendimiento)
    DB, err = sql.Open("sqlite", dbPath+"?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)")
    if err != nil {
        return fmt.Errorf("error abriendo SQLite: %w", err)
    }

    // Configurar pool de conexiones (SQLite usa 1 conexión real, pero esto
    // ayuda con el manejo de statements concurrentes)
    DB.SetMaxOpenConns(1)
    DB.SetMaxIdleConns(1)

    // Verificar conexión
    if err := DB.Ping(); err != nil {
        return fmt.Errorf("error verificando conexión: %w", err)
    }

    // Ejecutar schema
    if _, err := DB.Exec(schema); err != nil {
        return fmt.Errorf("error creando tablas: %w", err)
    }

    slog.Info("Base de datos inicializada", "path", dbPath)
    return nil
}

// Close cierra la conexión a la base de datos
func Close() error {
    if DB != nil {
        return DB.Close()
    }
    return nil
}