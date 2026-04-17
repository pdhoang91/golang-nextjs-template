import { env } from "@/config/env";
import { API_ERROR_MESSAGE } from "@/constants/copy";
import { CONTENT_TYPE_JSON, HTTP_HEADER_CONTENT_TYPE } from "@/constants/http";
import type { ApiErrorPayload } from "@/types/api";

type RequestInitWithJson = RequestInit & {
  json?: unknown;
};

export async function apiClient<T>(
  path: string,
  options: RequestInitWithJson = {}
): Promise<T> {
  const headers = new Headers(options.headers ?? {});
  headers.set(HTTP_HEADER_CONTENT_TYPE, CONTENT_TYPE_JSON);

  const response = await fetch(`${env.apiBaseUrl}${path}`, {
    ...options,
    headers,
    body: options.json ? JSON.stringify(options.json) : options.body
  });

  if (!response.ok) {
    let message = API_ERROR_MESSAGE;

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
