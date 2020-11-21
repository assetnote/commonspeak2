CREATE TEMPORARY FUNCTION
getFP(x STRING)
RETURNS STRING
LANGUAGE js AS
"""
var path=URI(x).pathname();
if (path.startsWith("/api/") || path.startsWith("/v1/") || path.startsWith("/v2/") || path.startsWith("/rest/")){
    res=path;
}else{
 res = "";
}

return res;
"""
OPTIONS
  ( library="gs://commonspeak-udf/URI.min.js" );
SELECT
  getFP(url) AS fullpath,
  COUNT(url) AS count
FROM
  `httparchive.requests.{{date}}_desktop`
GROUP BY
  fullpath
ORDER BY
  count DESC
LIMIT {{limit}};