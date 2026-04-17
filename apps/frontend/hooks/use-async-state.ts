"use client";

import { useCallback, useState } from "react";

type AsyncState = {
  isLoading: boolean;
  error: string | null;
};

export function useAsyncState(initialLoading = false) {
  const [state, setState] = useState<AsyncState>({
    isLoading: initialLoading,
    error: null
  });

  const start = useCallback(() => setState({ isLoading: true, error: null }), []);
  const fail = useCallback((error: string) => setState({ isLoading: false, error }), []);
  const succeed = useCallback(() => setState({ isLoading: false, error: null }), []);

  return {
    ...state,
    start,
    fail,
    succeed
  };
}
