sudo rm -rf ./postgres-data
sudo rm -rf ../postgres-data

docker-compose -f ../docker-compose.yml down --volumes

docker-compose -f ../docker-compose.yml up

docker-compose -f docker-compose.yml down --volumes

docker-compose -f docker-compose.yml up