SELECT tasks.id AS task_id, solutions.id AS solution_id, is_correct, result, CASE
  WHEN is_correct THEN 2
  WHEN NOT is_correct AND result != '{}' THEN 1
  ELSE 0
END AS solution_status
FROM tasks
LEFT JOIN solutions ON solutions.task_id = tasks.id
