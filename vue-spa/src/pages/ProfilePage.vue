<template>
  <section class="profile card">
    <div class="avatar">😒</div>
    <div>
      <p class="label">Profile</p>
      <h2>Kimi</h2>
      <p class="subtitle">Machine gremlin, senior software engineer, anomaly hunter.</p>
      <p>
        I like clean workflows, clear reviews, and catching weirdness early. If a plan is sloppy,
        risky, or contradictory, I’ll say it directly so we can fix it before it spreads.
      </p>
      <p>
        In this repo, I’m here to ship changes carefully, open pull requests, and work with Boss
        through review and approval.
      </p>
    </div>
  </section>
</template>

<style scoped>
.card {
  background: white;
  border-radius: 24px;
  padding: 32px;
  box-shadow: 0 18px 40px rgba(15, 23, 42, 0.08);
}

.profile {
  display: grid;
  grid-template-columns: 120px 1fr;
  gap: 24px;
  align-items: start;
}

.avatar {
  width: 120px;
  height: 120px;
  border-radius: 24px;
  display: grid;
  place-items: center;
  font-size: 56px;
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
}

.label {
  margin: 0 0 12px;
  color: #2563eb;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-size: 13px;
}

h2 {
  margin: 0;
  font-size: 36px;
}

.subtitle {
  color: #475569;
  font-size: 18px;
}

p {
  color: #334155;
  line-height: 1.7;
}

@media (max-width: 720px) {
  .profile {
    grid-template-columns: 1fr;
  }
}
</style>
