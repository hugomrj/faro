// frontend/src/router/index.ts
import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      component: () => import('../layouts/MainLayout.vue'),
      children: [
        {
          path: '',
          redirect: '/categories'
        },
        // --- Inventario ---
        {
          path: 'categories',
          name: 'categories-list',
          component: () => import('../views/categories/ListView.vue')
        },
        {
          path: 'categories/create',
          name: 'categories-create',
          component: () => import('../views/categories/FormView.vue')
        },
        {
          path: 'categories/edit/:id',
          name: 'categories-edit',
          component: () => import('../views/categories/FormView.vue')
        },
        {
          path: 'products',
          name: 'products-list',
          component: () => import('../views/products/ListView.vue')
        },
        // --- AÑADE ESTAS DOS RUTAS ---
        {
          path: 'products/create',
          name: 'products-create',
          component: () => import('../views/products/FormView.vue')
        },
        {
          path: 'products/edit/:id',
          name: 'products-edit',
          component: () => import('../views/products/FormView.vue')
        },
        // --------------------------------

        // --- Ventas (placeholders) ---
        {
          path: 'sales',
          name: 'sales-pos',
          component: () => import('../views/PlaceholderView.vue')
        },
        {
          path: 'invoices',
          name: 'invoices-list',
          component: () => import('../views/PlaceholderView.vue')
        },
        // --- Reportes (placeholders) ---
        {
          path: 'reports/inventory',
          name: 'reports-inventory',
          component: () => import('../views/PlaceholderView.vue')
        },
        {
          path: 'reports/sales',
          name: 'reports-sales',
          component: () => import('../views/PlaceholderView.vue')
        },
        // --- Configuración ---
        {
          path: 'settings',
          name: 'settings',
          component: () => import('../views/PlaceholderView.vue')
        }
      ]
    }
  ]
})

export default router