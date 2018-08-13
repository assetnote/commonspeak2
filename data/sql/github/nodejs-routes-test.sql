CREATE TEMPORARY FUNCTION parseRoutes(routesCode STRING)
RETURNS ARRAY<STRING> LANGUAGE js AS """

  var regex = /("|')[a-zA-Z_0-9\\/\\:!@$%^&*()+=]*('|")/mgi;
  if (routesCode) {
    var result = routesCode.match(regex);
  } else {
    var result = [];
  }
  return result;
  """;

SELECT
  route,
  COUNT(route) AS count
FROM (
  SELECT
    parseRoutes(c.content) AS route
  FROM
    `bigquery-public-data.github_repos.sample_contents` c
  JOIN (
    SELECT
      id
    FROM
      `bigquery-public-data.github_repos.sample_files`
    WHERE
      path = "config/routes.rb" ) f
  ON
    f.id = c.id ) routes,
  UNNEST(route) AS route
WHERE
  route IS NOT NULL
GROUP BY
  route
ORDER BY
  count DESC
LIMIT
  {{limit}}