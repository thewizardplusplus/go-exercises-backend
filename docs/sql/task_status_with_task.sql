SELECT tasks.*, COALESCE(statuses.status, 0) AS status
FROM tasks
LEFT JOIN (
  SELECT tasks.id AS tasks_id, MAX(CASE
    WHEN is_correct THEN 2
    WHEN NOT is_correct AND result != '{}' THEN 1
    ELSE 0
  END) AS status
  FROM tasks
  LEFT JOIN solutions ON solutions.task_id = tasks.id
  GROUP BY tasks.id
) statuses ON statuses.tasks_id = tasks.id
