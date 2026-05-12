<template>
  <div class="layout">
    <header class="topbar">
      <div>
        <p class="eyebrow">Kimi Workspace</p>
        <h1>Kimi Dashboard</h1>
      </div>
      <nav>
        <RouterLink to="/">Dashboard</RouterLink>
        <RouterLink to="/profile">Profile</RouterLink>
      </nav>
    </header>

    <main class="page-shell">
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<style scoped>
.layout {
  min-height: 100vh;
  background: linear-gradient(180deg, #eef4ff 0%, #f8fafc 100%);
  color: #0f172a;
}

.topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 24px 32px;
}

.eyebrow {
  margin: 0 0 6px;
  color: #2563eb;
  font-size: 14px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

h1 {
  margin: 0;
  font-size: 32px;
}

nav {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

nav a {
  text-decoration: none;
  color: #1e293b;
  font-weight: 600;
  padding: 10px 14px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.7);
}

nav a.router-link-exact-active {
  background: #2563eb;
  color: white;
}

.page-shell {
  padding: 0 32px 32px;
}

@media (max-width: 720px) {
  .topbar {
    flex-direction: column;
    align-items: flex-start;
    padding: 20px;
  }

  .page-shell {
    padding: 0 20px 20px;
  }
}
</style>
