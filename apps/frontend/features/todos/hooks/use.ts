"use client";

import { useCallback, useEffect, useState } from "react";

import { TODO_ERROR_MESSAGE } from "@/features/todos/constants";
import { useAsyncState } from "@/hooks/use-async-state";
import { createTodo, getTodos } from "@/services/todo.service";
import type { CreateTodoInput, Todo } from "@/features/todos/types";

export function useTodos() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [isCreating, setIsCreating] = useState(false);
  const { isLoading, error, start, fail, succeed } = useAsyncState(true);

  const fetchTodos = useCallback(async () => {
    start();
    try {
      const response = await getTodos();
      setTodos(response.data);
      succeed();
    } catch (err) {
      fail(err instanceof Error ? err.message : TODO_ERROR_MESSAGE);
    }
  }, [start, fail, succeed]);

  useEffect(() => {
    void fetchTodos();
  }, [fetchTodos]);

  const addTodo = useCallback(async (payload: CreateTodoInput) => {
    setIsCreating(true);
    try {
      const response = await createTodo(payload);
      setTodos((current) => [response.data, ...current]);
      return response.data;
    } finally {
      setIsCreating(false);
    }
  }, []);

  return {
    todos,
    isLoading,
    error,
    isCreating,
    fetchTodos,
    addTodo
  };
}
