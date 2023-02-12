CREATE TEMPORARY FUNCTION
  getPath(x STRING)
  RETURNS STRING
  LANGUAGE js AS """
  function getPath(s) {
    try {
      return URI(s).pathname();
    } catch (ex) {
      return s;
    }
  }
  return getPath(x);
"""
OPTIONS
  ( library="gs://commonspeak-udf/URI.min.js" );
SELECT DISTINCT
  getPath(url) AS url,
  COUNT(url) AS count
FROM
  `httparchive.response_bodies.{{date}}_desktop`
  TABLESAMPLE SYSTEM ({{sample_rate}} PERCENT)
WHERE 
  REGEXP_CONTAINS (body, r'{{regex}}')
GROUP BY
  url
ORDER BY
  count DESC
LIMIT
  {{limit}};
