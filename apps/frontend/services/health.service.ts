import { API_ROUTES } from "@/constants/routes";
import { apiClient } from "@/services/api-client";
import type { HealthResponse } from "@/types/health";

export async function getHealth() {
  return apiClient<HealthResponse>(API_ROUTES.health, {
    method: "GET"
  });
}
