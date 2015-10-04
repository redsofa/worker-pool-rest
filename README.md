# worker-pool-rest


curl \
--verbose \
--silent \
--insecure \
--request POST \
--header "Content-Type: application/json" \
--data '
{
  "table": [
  		{ "Index": 0, "NumA": 1, "NumB": 2 },
        { "Index": 1, "NumA": 3,"NumB": 4}
    ]
}
' \
http://localhost:8080/calcs