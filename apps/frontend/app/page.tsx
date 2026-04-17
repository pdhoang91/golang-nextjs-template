import { Container } from "@/components/layout/container";
import { APP_HERO_COPY, APP_HERO_EYEBROW, APP_HERO_TITLE } from "@/constants/copy";
import { HealthStatusCard } from "@/features/health/components/card";
import { TodoListCard } from "@/features/todos/components/list";

export default function HomePage() {
  return (
    <main>
      <Container>
        <section className="hero">
          <p className="eyebrow">{APP_HERO_EYEBROW}</p>
          <h1>{APP_HERO_TITLE}</h1>
          <p className="hero-copy">{APP_HERO_COPY}</p>
        </section>

        <section className="grid">
          <HealthStatusCard />
          <TodoListCard />
        </section>
      </Container>
    </main>
  );
}
