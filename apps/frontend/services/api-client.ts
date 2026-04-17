import { env } from "@/config/env";
import type { ApiErrorPayload } from "@/types/api";

type RequestInitWithJson = RequestInit & {
  json?: unknown;
};

export async function apiClient<T>(
  path: string,
  options: RequestInitWithJson = {}
): Promise<T> {
  const headers = new Headers(options.headers ?? {});
  headers.set("Content-Type", "application/json");

  const response = await fetch(`${env.apiBaseUrl}${path}`, {
    ...options,
    headers,
    body: options.json ? JSON.stringify(options.json) : options.body
  });

  if (!response.ok) {
    let message = "Request failed";

    try {
      const payload = (await response.json()) as ApiErrorPayload;
      message = payload.error || message;
    } catch {
      // ignore JSON parse failure and keep fallback message
    }

    throw new Error(message);
  }

  return (await response.json()) as T;
}
