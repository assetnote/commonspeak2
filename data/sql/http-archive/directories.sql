CREATE TEMPORARY FUNCTION
  getDir(x STRING)
  RETURNS STRING
  LANGUAGE js AS """
  function getDir(s) {
    try {
      return URI(s).directory();
    } catch (ex) {
      return s;
    }
  }
  return getDir(x);
"""
OPTIONS
  ( library="gs://commonspeak-udf/URI.min.js" );
SELECT
  getDir(url) AS url,
  COUNT(url) AS count
FROM
  `httparchive.requests.{{date}}_desktop`
GROUP BY
  url
ORDER BY
  count DESC
LIMIT {{limit}};