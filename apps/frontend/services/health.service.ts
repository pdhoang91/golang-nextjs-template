import { API_ROUTES } from "@/constants/routes";
import { HTTP_METHOD_GET } from "@/constants/http";
import { apiClient } from "@/services/api-client";
import type { HealthResponse } from "@/features/health/types";

export async function getHealth() {
  return apiClient<HealthResponse>(API_ROUTES.health, {
    method: HTTP_METHOD_GET
  });
}
