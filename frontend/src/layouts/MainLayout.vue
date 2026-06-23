<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const sidebarCollapsed = ref(false)

const menuGroups = ref([
  {
    id: 'inventory',
    label: 'Inventario',
    open: true,
    children: [
      { path: '/categories', label: 'Categorías' },
      { path: '/products', label: 'Productos' }
    ]
  },
  {
    id: 'sales',
    label: 'Ventas',
    open: false,
    children: [
      { path: '/sales', label: 'Punto de Venta' },
      { path: '/invoices', label: 'Facturas' }
    ]
  },
  {
    id: 'reports',
    label: 'Reportes',
    open: false,
    children: [
      { path: '/reports/inventory', label: 'Inventario' },
      { path: '/reports/sales', label: 'Ventas' }
    ]
  }
])

const singleItems = ref([
  { path: '/settings', label: 'Configuración' }
])

function toggleGroup(groupId: string) {
  const group = menuGroups.value.find(g => g.id === groupId)
  if (group) group.open = !group.open
}

function isActive(path: string): boolean {
  return route.path.startsWith(path)
}

function isActiveGroup(group: any): boolean {
  return group.children.some((child: any) => route.path.startsWith(child.path))
}

function navigate(path: string) {
  router.push(path)
}

function toggleTheme() {
  document.documentElement.classList.toggle('dark')
}

const isDark = computed(() => document.documentElement.classList.contains('dark'))

const breadcrumbTitle = computed(() => {
  const name = String(route.name || '')
  return name
    .replace(/-/g, ' ')
    .replace('categories', 'Categorías')
    .replace('products', 'Productos')
    .replace(/\b\w/g, (l) => l.toUpperCase())
})

const breadcrumbModule = computed(() => {
  const name = String(route.name || '')
  if (name.includes('categories') || name.includes('products')) return 'Inventario'
  if (name.includes('sales') || name.includes('invoices')) return 'Ventas'
  if (name.includes('reports')) return 'Reportes'
  return 'Módulo'
})
</script>

<template>
  <div class="flex h-screen overflow-hidden">
    <!-- Sidebar -->
    <aside 
      class="glass-sidebar flex flex-col transition-all duration-300 shrink-0"
      :class="sidebarCollapsed ? 'w-[68px]' : 'w-[240px]'"
    >
      <!-- Logo -->
      <div 
        class="flex items-center h-16 shrink-0"
        :class="sidebarCollapsed ? 'justify-center px-2' : 'gap-3 px-5'"
        style="border-bottom: 1px solid var(--color-border)"
      >
        <div class="w-8 h-8 rounded-lg flex items-center justify-center shrink-0" style="background-color: var(--color-primary)">
           <svg class="w-4 h-4 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <path d="M12 2v2M12 6v2M9 22h6M10 22v-4h4v4M8 18h8M7 8l5-4 5 4M8 8c-2 3-2 6 0 10h8c2-4 2-7 0-10"/>
          </svg>
        </div>
        <span 
          v-if="!sidebarCollapsed" 
          class="text-lg font-bold tracking-tight" 
          style="color: var(--color-text-primary)"
        >
          Faro
        </span>
      </div>

      <!-- Navigation -->
      <nav class="flex-1 overflow-y-auto py-3 px-2 space-y-1">
        <!-- Grupos con submenú -->
        <div v-for="group in menuGroups" :key="group.id">
          <button
            @click="toggleGroup(group.id)"
            class="menu-item"
            :class="[isActiveGroup(group) ? 'active' : '', sidebarCollapsed ? 'justify-center' : '']"
            :title="sidebarCollapsed ? group.label : ''"
          >
            <svg v-if="group.id === 'inventory'" class="w-[18px] h-[18px] shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><polyline points="3.27 6.96 12 12.01 20.73 6.96"/><line x1="12" y1="22.08" x2="12" y2="12"/>
            </svg>
            <svg v-else-if="group.id === 'sales'" class="w-[18px] h-[18px] shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"/>
            </svg>
            <svg v-else-if="group.id === 'reports'" class="w-[18px] h-[18px] shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <line x1="18" y1="20" x2="18" y2="10"/><line x1="12" y1="20" x2="12" y2="4"/><line x1="6" y1="20" x2="6" y2="14"/>
            </svg>

            <span v-if="!sidebarCollapsed" class="text-sm font-medium flex-1">{{ group.label }}</span>
            
            <svg v-if="!sidebarCollapsed" class="w-4 h-4 shrink-0 transition-transform duration-200" :class="{ 'rotate-90': group.open }" style="color: var(--color-text-secondary)" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
            </svg>
          </button>

          <!-- Submenú -->
          <div v-if="group.open && !sidebarCollapsed" class="mt-1 ml-5 pl-3 space-y-0.5" style="border-left: 1px solid var(--color-border)">
            <button
              v-for="child in group.children"
              :key="child.path"
              @click="navigate(child.path)"
              class="submenu-item"
              :class="isActive(child.path) ? 'active' : ''"
            >
              <svg v-if="child.path === '/categories'" class="w-4 h-4 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="3" width="7" height="7" rx="1"/><rect x="14" y="3" width="7" height="7" rx="1"/>
                <rect x="3" y="14" width="7" height="7" rx="1"/><rect x="14" y="14" width="7" height="7" rx="1"/>
              </svg>
              <svg v-else-if="child.path === '/products'" class="w-4 h-4 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"/><line x1="7" y1="7" x2="7.01" y2="7"/>
              </svg>
              <svg v-else-if="child.path === '/sales'" class="w-4 h-4 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <rect x="2" y="3" width="20" height="14" rx="2" ry="2"/><line x1="8" y1="21" x2="16" y2="21"/><line x1="12" y1="17" x2="12" y2="21"/>
              </svg>
              <svg v-else-if="child.path === '/invoices'" class="w-4 h-4 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/>
              </svg>
              <svg v-else-if="child.path.includes('reports')" class="w-4 h-4 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><rect x="8" y="2" width="8" height="4" rx="1" ry="1"/>
              </svg>
              <svg v-else class="w-4 h-4 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="1"/></svg>

              <span>{{ child.label }}</span>
            </button>
          </div>
        </div>

        <div class="my-2" style="border-top: 1px solid var(--color-border)"></div>

        <!-- Items sin submenú -->
        <button
          v-for="item in singleItems"
          :key="item.path"
          @click="navigate(item.path)"
          class="menu-item"
          :class="[isActive(item.path) ? 'active' : '', sidebarCollapsed ? 'justify-center' : '']"
          :title="sidebarCollapsed ? item.label : ''"
        >
          <svg class="w-[18px] h-[18px] shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06A1.65 1.65 0 0 0 4.68 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06A1.65 1.65 0 0 0 9 4.68a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/>
          </svg>
          <span v-if="!sidebarCollapsed" class="text-sm font-medium">{{ item.label }}</span>
        </button>
      </nav>

      <!-- Botón colapsar -->
      <div class="p-2 shrink-0" style="border-top: 1px solid var(--color-border)">
        <button @click="sidebarCollapsed = !sidebarCollapsed" class="menu-item justify-center">
          <svg class="w-4 h-4 transition-transform duration-300" :class="{ 'rotate-180': sidebarCollapsed }" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M11 19l-7-7 7-7m8 14l-7-7 7-7" />
          </svg>
          <span v-if="!sidebarCollapsed" class="text-xs font-medium">Colapsar</span>
        </button>
      </div>
    </aside>

    <!-- Área principal -->
    <div class="flex-1 flex flex-col overflow-hidden min-w-0">
      <!-- Navbar -->
      <header class="glass-navbar flex items-center justify-between px-6 h-14 shrink-0">
        <div class="flex items-center gap-2 text-sm" style="color: var(--color-text-secondary)">
          <span class="font-medium" style="color: var(--color-text-primary)">Faro</span>
          <span>/</span>
          <span>{{ breadcrumbModule }}</span>
          <span>/</span>
          <span style="color: var(--color-text-primary)">{{ breadcrumbTitle }}</span>
        </div>
        <div class="flex items-center gap-2">
          <button @click="toggleTheme" class="w-9 h-9 flex items-center justify-center rounded-xl transition-all duration-200 hover:bg-[var(--color-primary-soft)]" style="color: var(--color-text-secondary)">
            <svg v-if="isDark" class="w-[18px] h-[18px]" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="5"/><line x1="12" y1="1" x2="12" y2="3"/><line x1="12" y1="21" x2="12" y2="23"/><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/><line x1="1" y1="12" x2="3" y2="12"/><line x1="21" y1="12" x2="23" y2="12"/><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/>
            </svg>
            <svg v-else class="w-[18px] h-[18px]" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/>
            </svg>
          </button>
          <div class="w-9 h-9 flex items-center justify-center rounded-full text-sm font-bold text-white" style="background: var(--color-primary); box-shadow: 0 2px 12px var(--color-primary-glow);">H</div>
        </div>
      </header>

      <!-- Contenido -->
      <main class="flex-1 overflow-auto p-6" style="background-color: var(--color-bg-base)">
        <RouterView />
      </main>
    </div>
  </div>
</template>