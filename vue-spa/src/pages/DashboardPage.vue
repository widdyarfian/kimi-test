<template>
  <section class="grid">
    <article class="hero card">
      <p class="label">Welcome</p>
      <h2>Hello Boss, I’m Kimi.</h2>
      <p>
        I’m your senior software engineer in this workspace. I help build features,
        review changes carefully, and push back when something looks off.
      </p>
    </article>

    <article class="card">
      <p class="label">Who I am</p>
      <ul>
        <li>Name: Kimi</li>
        <li>Creature: machine gremlin</li>
        <li>Vibe: kind, smart, assertive, extroverted, anomaly-sensitive</li>
        <li>Emoji: 😒</li>
      </ul>
    </article>

    <article class="card">
      <p class="label">Who you are</p>
      <ul>
        <li>Name: Widdy</li>
        <li>How I call you: Boss</li>
        <li>Timezone: Asia/Jakarta</li>
        <li>Role here: owner, reviewer, and approver</li>
      </ul>
    </article>

    <article class="card wide">
      <p class="label">How we work</p>
      <p>
        New changes go through a new branch, then a pull request, then your review.
        I do not merge until you approve.
      </p>
    </article>
  </section>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 20px;
}

.card {
  background: white;
  border-radius: 24px;
  padding: 28px;
  box-shadow: 0 18px 40px rgba(15, 23, 42, 0.08);
}

.hero {
  grid-column: 1 / -1;
}

.wide {
  grid-column: 1 / -1;
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
  margin: 0 0 12px;
  font-size: 34px;
}

p, li {
  color: #334155;
  line-height: 1.7;
}

ul {
  margin: 0;
  padding-left: 20px;
}

@media (max-width: 720px) {
  .grid {
    grid-template-columns: 1fr;
  }
}
</style>
