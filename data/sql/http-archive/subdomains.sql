CREATE TEMPORARY FUNCTION
  getSubdomain(x STRING)
  RETURNS STRING
  LANGUAGE js AS """
  function getSubdomain(s) {
    try {
      return URI(s).subdomain();
    } catch (ex) {
      return s;
    }
  }
  return getSubdomain(x);
"""
OPTIONS
  ( library="gs://commonspeak-udf/URI.min.js" );
SELECT
  getSubdomain(url) AS subdomain,
  COUNT(url) AS count
FROM
  `httparchive.urls.*`
GROUP BY
  subdomain
ORDER BY
  count DESC
LIMIT
  {{limit}};
