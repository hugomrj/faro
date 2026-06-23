package main

import (
    "embed"
    "log/slog"
    "os"
    "path/filepath"

    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/assetserver"

    "faro/internal/core/database"
    "faro/internal/features/categories"
    "faro/internal/features/products" 

)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
    // Determinar ruta de la base de datos
    // En Linux: ~/.local/share/faro/faro.db
    homeDir, _ := os.UserHomeDir()
    dbPath := filepath.Join(homeDir, ".local", "share", "faro", "faro.db")

    // Inicializar base de datos
    if err := database.Init(dbPath); err != nil {
        slog.Error("Error inicializando base de datos", "error", err)
        os.Exit(1)
    }
    defer database.Close()

    // Crear aplicación Wails
    app := NewApp()

    err := wails.Run(&options.App{
        Title:     "App ERP",
        Width:     1200,
        Height:    800,
        MinWidth:  900,
        MinHeight: 600,
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        BackgroundColour: &options.RGBA{R: 15, G: 15, B: 15, A: 255}, // Dark por defecto
        OnStartup:       app.startup,
        OnShutdown:      app.shutdown,
        Bind: []interface{}{
            app,
            categories.NewController(), 
            products.NewController(), 
        },
    })

    if err != nil {
        slog.Error("Error ejecutando Wails", "error", err)
        os.Exit(1)
    }
}

