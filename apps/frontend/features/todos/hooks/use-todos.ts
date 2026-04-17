"use client";

import { useCallback, useEffect, useState } from "react";

import { useAsyncState } from "@/hooks/use-async-state";
import { createTodo, getTodos } from "@/services/todo.service";
import type { CreateTodoPayload, Todo } from "@/types/todo";

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
      fail(err instanceof Error ? err.message : "Failed to fetch todos");
    }
  }, [start, fail, succeed]);

  useEffect(() => {
    void fetchTodos();
  }, [fetchTodos]);

  const addTodo = useCallback(async (payload: CreateTodoPayload) => {
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
