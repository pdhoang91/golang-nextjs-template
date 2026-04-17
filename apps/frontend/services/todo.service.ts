import { API_ROUTES } from "@/constants/routes";
import { HTTP_METHOD_GET, HTTP_METHOD_POST } from "@/constants/http";
import { apiClient } from "@/services/api-client";
import type { ApiResponse } from "@/types/api";
import type { CreateTodoInput, Todo } from "@/features/todos/types";

export async function getTodos() {
  return apiClient<ApiResponse<Todo[]>>(API_ROUTES.todos, {
    method: HTTP_METHOD_GET
  });
}

export async function createTodo(payload: CreateTodoInput) {
  return apiClient<ApiResponse<Todo>>(API_ROUTES.todos, {
    method: HTTP_METHOD_POST,
    json: payload
  });
}
