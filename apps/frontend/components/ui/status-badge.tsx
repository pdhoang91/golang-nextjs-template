type StatusBadgeProps = {
  status: "ok" | "error";
  label: string;
};

export function StatusBadge({ status, label }: StatusBadgeProps) {
  return (
    <span
      className={`status-badge ${
        status === "ok" ? "status-badge--ok" : "status-badge--error"
      }`}
    >
      <span>{status === "ok" ? "●" : "●"}</span>
      {label}
    </span>
  );
}
