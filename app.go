package main

import (
    "context"
    "log/slog"
)

type App struct {
    ctx context.Context
}

func NewApp() *App {
    return &App{}
}

func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    slog.Info("Aplicación Faro iniciada")
}

func (a *App) shutdown(ctx context.Context) {
    slog.Info("Aplicación Faro cerrándose")
}
