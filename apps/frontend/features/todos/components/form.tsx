"use client";

import { useState } from "react";

import {
  TODO_CREATE_LABEL,
  TODO_CREATING_LABEL,
  TODO_DESCRIPTION_PLACEHOLDER,
  TODO_ERROR_MESSAGE,
  TODO_TITLE_PLACEHOLDER,
  TODO_TITLE_REQUIRED_MESSAGE
} from "@/features/todos/constants";
import type { CreateTodoInput } from "@/features/todos/types";

type TodoCreateFormProps = {
  isSubmitting: boolean;
  onSubmit: (payload: CreateTodoInput) => Promise<void>;
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
      setFormError(TODO_TITLE_REQUIRED_MESSAGE);
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
      setFormError(error instanceof Error ? error.message : TODO_ERROR_MESSAGE);
    }
  }

  return (
    <form className="form" onSubmit={handleSubmit}>
      <input
        className="input"
        type="text"
        placeholder={TODO_TITLE_PLACEHOLDER}
        value={title}
        onChange={(event) => setTitle(event.target.value)}
      />

      <textarea
        className="textarea"
        placeholder={TODO_DESCRIPTION_PLACEHOLDER}
        value={description}
        onChange={(event) => setDescription(event.target.value)}
      />

      {formError ? <div className="error-state">{formError}</div> : null}

      <button className="button" type="submit" disabled={isSubmitting}>
        {isSubmitting ? TODO_CREATING_LABEL : TODO_CREATE_LABEL}
      </button>
    </form>
  );
}
