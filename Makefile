mysql-up:
	# echo "service mysql start"
	sudo docker run -p3306:3306 -e MYSQL_ROOT_PASSWORD=QMKAJNNjNK9vBO88 --name mysql mysql
redis-up:
	# echo "service redis start"
	docker run -d --name redis redis 
build-docker:
	# echo "build docker"
	sudo docker build -f ./docker/Docker .
