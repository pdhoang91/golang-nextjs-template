import { API_ROUTES } from "@/constants/routes";
import { apiClient } from "@/services/api-client";
import type { ApiResponse } from "@/types/api";
import type { CreateTodoPayload, Todo } from "@/types/todo";

export async function getTodos() {
  return apiClient<ApiResponse<Todo[]>>(API_ROUTES.todos, {
    method: "GET"
  });
}

export async function createTodo(payload: CreateTodoPayload) {
  return apiClient<ApiResponse<Todo>>(API_ROUTES.todos, {
    method: "POST",
    json: payload
  });
}
