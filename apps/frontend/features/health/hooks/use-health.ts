"use client";

import { useEffect, useState } from "react";

import { useAsyncState } from "@/hooks/use-async-state";
import { getHealth } from "@/services/health.service";
import type { HealthResponse } from "@/types/health";

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
        fail(err instanceof Error ? err.message : "Failed to fetch health status");
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
