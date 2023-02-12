SELECT DISTINCT
  path,
  COUNT(path) AS count
FROM
  `bigquery-public-data.github_repos.files`
WHERE id IN (
    SELECT id
    FROM `bigquery-public-data.github_repos.contents`
    TABLESAMPLE SYSTEM ({{sample_rate}} PERCENT)
    WHERE REGEXP_CONTAINS (content, r'{{regex}}')
)
GROUP BY
  path
ORDER BY
  count DESC
LIMIT
  {{limit}};
