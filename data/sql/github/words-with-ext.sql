SELECT
  path,
  COUNT(*) count
FROM
  `bigquery-public-data.github_repos.files`
WHERE
  REGEXP_CONTAINS (path, r'({{extensions}})$')
GROUP BY
  path
ORDER BY
  count DESC
LIMIT
  {{limit}};