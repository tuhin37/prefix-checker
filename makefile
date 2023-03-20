docker-build:
	docker build --no-cache -t fidays/check-prefix:latest -f ./build/Dockerfile .

docker-exec:
	docker run --rm -it -p 5000:5000 fidays/check-prefix:latest sh

docker-rmi:
	docker rmi fidays/check-prefix:latest

docker-run:
	docker run --rm -v /home/drag/programming/interviews/truecaller/prefix-wordlists/:/app/prefix-wordlists/ -it -p 5000:5000 fidays/check-prefix:latest
