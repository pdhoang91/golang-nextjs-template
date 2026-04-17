"use client";

import { Card } from "@/components/ui/card";
import { StatusBadge } from "@/components/ui/status-badge";
import { TodoCreateForm } from "@/features/todos/components/todo-create-form";
import { useTodos } from "@/features/todos/hooks/use-todos";
import { formatDateTime } from "@/lib/format-date";
import type { CreateTodoPayload } from "@/types/todo";

export function TodoListCard() {
  const { todos, isLoading, error, isCreating, addTodo } = useTodos();

  async function handleCreate(payload: CreateTodoPayload) {
    await addTodo(payload);
  }

  return (
    <Card
      title="Todo module example"
      description="This feature demonstrates a full end-to-end module from database migration to UI."
    >
      <TodoCreateForm isSubmitting={isCreating} onSubmit={handleCreate} />

      {isLoading ? <div className="loading-state">Loading todos...</div> : null}
      {error ? <div className="error-state">{error}</div> : null}

      {!isLoading && !error && todos.length === 0 ? (
        <div className="empty-state">No todos yet. Create your first item.</div>
      ) : null}

      <div className="todo-list">
        {todos.map((todo) => (
          <article className="todo-item" key={todo.id}>
            <h3>{todo.title}</h3>
            <p>{todo.description || "No description provided."}</p>

            <div className="todo-meta">
              <StatusBadge
                status={todo.completed ? "ok" : "error"}
                label={todo.completed ? "Completed" : "Open"}
              />
              <span>{formatDateTime(todo.created_at)}</span>
            </div>
          </article>
        ))}
      </div>
    </Card>
  );
}
