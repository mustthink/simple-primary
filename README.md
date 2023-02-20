# Check Primary Handler

### Request example
`curl  -H 'Content-Type: application/json' --data '["1", "2", "3", "4", "5", "6"]' localhost:8080/check`

### Response example
`["false","true","true","false","true","false"]`