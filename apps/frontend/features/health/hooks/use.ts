"use client";

import { useEffect, useState } from "react";

import { HEALTH_ERROR_MESSAGE } from "@/features/health/constants";
import { useAsyncState } from "@/hooks/use-async-state";
import { getHealth } from "@/services/health.service";
import type { HealthResponse } from "@/features/health/types";

export function useHealth() {
  const [health, setHealth] = useState<HealthResponse | null>(null);
  const { isLoading, error, start, fail, succeed } = useAsyncState(true);

  useEffect(() => {
    let isMounted = true;

    async function run() {
      start();
      try {
        const response = await getHealth();
        if (!isMounted) return;
        setHealth(response);
        succeed();
      } catch (err) {
        if (!isMounted) return;
        fail(err instanceof Error ? err.message : HEALTH_ERROR_MESSAGE);
      }
    }

    void run();

    return () => {
      isMounted = false;
    };
  }, [start, fail, succeed]);

  return {
    health,
    isLoading,
    error
  };
}
