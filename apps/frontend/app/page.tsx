import { Container } from "@/components/layout/container";
import { HealthStatusCard } from "@/features/health/components/health-status-card";
import { TodoListCard } from "@/features/todos/components/todo-list-card";

export default function HomePage() {
  return (
    <main>
      <Container>
        <section className="hero">
          <p className="eyebrow">Production-ready starter</p>
          <h1>Go + Next.js monorepo template</h1>
          <p className="hero-copy">
            A clean base project with a structured backend, typed frontend,
            Docker Compose, migrations, and an example todo module wired
            end-to-end.
          </p>
        </section>

        <section className="grid">
          <HealthStatusCard />
          <TodoListCard />
        </section>
      </Container>
    </main>
  );
}
