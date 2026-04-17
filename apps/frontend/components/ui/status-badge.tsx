import { STATUS_BULLET, STATUS_ERROR_LABEL, STATUS_OK_LABEL } from "@/constants/status";

type StatusBadgeProps = {
  status: typeof STATUS_OK_LABEL | typeof STATUS_ERROR_LABEL;
  label: string;
};

export function StatusBadge({ status, label }: StatusBadgeProps) {
  return (
    <span
      className={`status-badge ${
        status === STATUS_OK_LABEL ? "status-badge--ok" : "status-badge--error"
      }`}
    >
      <span>{STATUS_BULLET}</span>
      {label}
    </span>
  );
}
