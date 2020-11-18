SELECT
  url
FROM
  `httparchive.technologies.{{date}}_desktop`
WHERE
  app = '{{technology}}'
LIMIT
  {{limit}}