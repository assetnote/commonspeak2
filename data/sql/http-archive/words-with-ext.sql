CREATE TEMPORARY FUNCTION
  getFilename(x STRING)
  RETURNS STRING
  LANGUAGE js AS """
  function getFilename(s) {
    try {
      return URI(s).filename();
    } catch (ex) {
      return s;
    }
  }
  return getFilename(x);
"""
OPTIONS
  ( library="gs://commonspeak-udf/URI.min.js" );
SELECT
  getFilename(url) AS url,
  COUNT(url) AS count
FROM
  `httparchive.requests.{{date}}_desktop`
WHERE REGEXP_CONTAINS (url, r'({{extensions}})$')
GROUP BY
  url
ORDER BY
  count DESC
LIMIT {{limit}};