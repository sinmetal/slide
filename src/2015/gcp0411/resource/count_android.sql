SELECT count(insertId) AS count
FROM
(TABLE_DATE_RANGE(gaelog_from_bqstreaming.appengine_googleapis_com_request_log_, 
                    TIMESTAMP('2015-04-01'), 
                    TIMESTAMP('2015-04-30'))) 
WHERE protoPayload.userAgent CONTAINS "Android"
LIMIT 1000