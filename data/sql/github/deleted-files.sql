CREATE TEMPORARY FUNCTION
  extract_path(subject string)
  RETURNS STRING
  LANGUAGE js AS """
  return subject.split(' ')[1]
""";
SELECT
  path,
  COUNT(path) AS count
FROM (
  SELECT
    extract_path(subject) AS path
  FROM
    `bigquery-public-data.github_repos.commits`
  WHERE
    REGEXP_CONTAINS(subject, '^(remov|delet|eras)(e|ed|es|ing) [^\\s]*$') )
GROUP BY
  path
ORDER BY
  count DESC