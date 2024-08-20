aws --endpoint-url=http://localhost:4566 secretsmanager create-secret \
  --name notes-database-secret \
  --secret-string '{"username":"postgres","password":"postgres","engine":"postgres","host":"localhost","port":5432,"dbname":"postgres","dbInstanceIdentifier":"notes"}' \
  --region us-east-2
