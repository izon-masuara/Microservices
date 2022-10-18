# To migration database
cld:
	sudo docker-compose down
	sudo docker rmi microservices_nginx microservices_api_gt microservices_api_files microservices_api_user microservices_api_analysis