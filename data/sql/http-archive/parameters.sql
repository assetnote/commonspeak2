CREATE TEMPORARY FUNCTION
getQuery(x STRING)
RETURNS ARRAY<STRING>
LANGUAGE js AS
"""
var query=URI(x).search(true);
var keys = [];
for(var k in query) keys.push(k);
return keys;
"""
OPTIONS
  ( library="gs://commonspeak-udf/URI.min.js" );
SELECT
  COUNT(parameter) AS paramCount,
  parameter,
FROM
  `httparchive.requests.{{date}}_desktop`,
  UNNEST(getQuery(url)) as parameter
GROUP BY
  parameter
ORDER BY
  paramCount DESC
LIMIT {{limit}};