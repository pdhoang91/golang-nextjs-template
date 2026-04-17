"use client";

import { useState } from "react";

import type { CreateTodoPayload } from "@/types/todo";

type TodoCreateFormProps = {
  isSubmitting: boolean;
  onSubmit: (payload: CreateTodoPayload) => Promise<void>;
};

export function TodoCreateForm({
  isSubmitting,
  onSubmit
}: TodoCreateFormProps) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [formError, setFormError] = useState<string | null>(null);

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    if (!title.trim()) {
      setFormError("Title is required");
      return;
    }

    setFormError(null);

    try {
      await onSubmit({
        title: title.trim(),
        description: description.trim()
      });

      setTitle("");
      setDescription("");
    } catch (error) {
      setFormError(
        error instanceof Error ? error.message : "Failed to create todo"
      );
    }
  }

  return (
    <form className="form" onSubmit={handleSubmit}>
      <input
        className="input"
        type="text"
        placeholder="Todo title"
        value={title}
        onChange={(event) => setTitle(event.target.value)}
      />

      <textarea
        className="textarea"
        placeholder="Description"
        value={description}
        onChange={(event) => setDescription(event.target.value)}
      />

      {formError ? <div className="error-state">{formError}</div> : null}

      <button className="button" type="submit" disabled={isSubmitting}>
        {isSubmitting ? "Creating..." : "Create todo"}
      </button>
    </form>
  );
}
