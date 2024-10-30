run:
	docker run -d -p 8088:80 --name kamago kamago
build:
	docker build -t kamago .


deploy:
	kamal deploy
