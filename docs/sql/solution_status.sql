SELECT id, is_correct, result, CASE
  WHEN is_correct THEN 2
  WHEN NOT is_correct AND result != '{}' THEN 1
  ELSE 0
END AS status
FROM solutions
