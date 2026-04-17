"use client";

import { Card } from "@/components/ui/card";
import { StatusBadge } from "@/components/ui/status-badge";
import { STATUS_ERROR_LABEL, STATUS_OK_LABEL } from "@/constants/status";
import {
  TODO_COMPLETED_LABEL,
  TODO_NO_DESCRIPTION_MESSAGE,
  TODOS_CARD_DESCRIPTION,
  TODOS_CARD_TITLE,
  TODOS_EMPTY_MESSAGE,
  TODOS_LOADING_MESSAGE,
  TODO_OPEN_LABEL
} from "@/features/todos/constants";
import { TodoCreateForm } from "@/features/todos/components/form";
import { useTodos } from "@/features/todos/hooks/use";
import { formatDateTime } from "@/lib/format-date";
import type { CreateTodoInput } from "@/features/todos/types";

export function TodoListCard() {
  const { todos, isLoading, error, isCreating, addTodo } = useTodos();

  async function handleCreate(payload: CreateTodoInput) {
    await addTodo(payload);
  }

  return (
    <Card
      title={TODOS_CARD_TITLE}
      description={TODOS_CARD_DESCRIPTION}
    >
      <TodoCreateForm isSubmitting={isCreating} onSubmit={handleCreate} />

      {isLoading ? <div className="loading-state">{TODOS_LOADING_MESSAGE}</div> : null}
      {error ? <div className="error-state">{error}</div> : null}

      {!isLoading && !error && todos.length === 0 ? (
        <div className="empty-state">{TODOS_EMPTY_MESSAGE}</div>
      ) : null}

      <div className="todo-list">
        {todos.map((todo) => (
          <article className="todo-item" key={todo.id}>
            <h3>{todo.title}</h3>
            <p>{todo.description || TODO_NO_DESCRIPTION_MESSAGE}</p>

            <div className="todo-meta">
              <StatusBadge
                status={todo.completed ? STATUS_OK_LABEL : STATUS_ERROR_LABEL}
                label={todo.completed ? TODO_COMPLETED_LABEL : TODO_OPEN_LABEL}
              />
              <span>{formatDateTime(todo.created_at)}</span>
            </div>
          </article>
        ))}
      </div>
    </Card>
  );
}
