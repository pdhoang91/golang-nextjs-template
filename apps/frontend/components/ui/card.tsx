type CardProps = {
  title: string;
  description?: string;
  children: React.ReactNode;
};

export function Card({ title, description, children }: CardProps) {
  return (
    <section className="card">
      <header className="card-header">
        <h2 className="card-title">{title}</h2>
        {description ? <p className="card-description">{description}</p> : null}
      </header>
      {children}
    </section>
  );
}
