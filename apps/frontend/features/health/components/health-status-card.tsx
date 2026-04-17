"use client";

import { Card } from "@/components/ui/card";
import { StatusBadge } from "@/components/ui/status-badge";
import { useHealth } from "@/features/health/hooks/use-health";
import { formatDateTime } from "@/lib/format-date";

export function HealthStatusCard() {
  const { health, isLoading, error } = useHealth();

  return (
    <Card
      title="API health check"
      description="Example frontend-to-backend call to verify the system wiring."
    >
      {isLoading ? <div className="loading-state">Checking backend status...</div> : null}
      {error ? <div className="error-state">{error}</div> : null}

      {health ? (
        <>
          <StatusBadge
            status={health.status === "ok" ? "ok" : "error"}
            label={health.status === "ok" ? "Backend is healthy" : "Backend error"}
          />

          <div className="meta-list">
            <div className="meta-item">
              <span>Service</span>
              <strong>{health.service}</strong>
            </div>
            <div className="meta-item">
              <span>Environment</span>
              <strong>{health.environment}</strong>
            </div>
            <div className="meta-item">
              <span>Timestamp</span>
              <strong>{formatDateTime(health.timestamp)}</strong>
            </div>
          </div>
        </>
      ) : null}
    </Card>
  );
}
