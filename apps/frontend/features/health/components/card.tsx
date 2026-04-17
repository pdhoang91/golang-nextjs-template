"use client";

import { Card } from "@/components/ui/card";
import { StatusBadge } from "@/components/ui/status-badge";
import { STATUS_ERROR_LABEL } from "@/constants/status";
import {
  HEALTH_CARD_DESCRIPTION,
  HEALTH_CARD_TITLE,
  HEALTH_ENVIRONMENT_LABEL,
  HEALTH_ERROR_LABEL,
  HEALTH_HEALTHY_LABEL,
  HEALTH_LOADING_MESSAGE,
  HEALTH_SERVICE_LABEL,
  HEALTH_TIMESTAMP_LABEL
} from "@/features/health/constants";
import { HEALTH_STATUS_OK } from "@/features/health/constants";
import { useHealth } from "@/features/health/hooks/use";
import { formatDateTime } from "@/lib/format-date";

export function HealthStatusCard() {
  const { health, isLoading, error } = useHealth();

  return (
    <Card
      title={HEALTH_CARD_TITLE}
      description={HEALTH_CARD_DESCRIPTION}
    >
      {isLoading ? <div className="loading-state">{HEALTH_LOADING_MESSAGE}</div> : null}
      {error ? <div className="error-state">{error}</div> : null}

      {health ? (
        <>
          <StatusBadge
            status={health.status === HEALTH_STATUS_OK ? HEALTH_STATUS_OK : STATUS_ERROR_LABEL}
            label={health.status === HEALTH_STATUS_OK ? HEALTH_HEALTHY_LABEL : HEALTH_ERROR_LABEL}
          />

          <div className="meta-list">
            <div className="meta-item">
              <span>{HEALTH_SERVICE_LABEL}</span>
              <strong>{health.service}</strong>
            </div>
            <div className="meta-item">
              <span>{HEALTH_ENVIRONMENT_LABEL}</span>
              <strong>{health.environment}</strong>
            </div>
            <div className="meta-item">
              <span>{HEALTH_TIMESTAMP_LABEL}</span>
              <strong>{formatDateTime(health.timestamp)}</strong>
            </div>
          </div>
        </>
      ) : null}
    </Card>
  );
}
