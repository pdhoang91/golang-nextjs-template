export type ApiResponse<T> = {
  data: T;
};

export type ApiErrorPayload = {
  error: string;
};
