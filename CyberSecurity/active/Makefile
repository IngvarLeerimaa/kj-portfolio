server:
	@echo "Starting test server"
	go run test_server.go

u:
	@echo "Testing UDP server"
	python3 tinyscanner.py -u 127.0.0.1 -p 8080-8082

t: 
	@echo "Testing TCP server"
	python3 tinyscanner.py -t 127.0.0.1 -p 8080-8082

all: 
	@echo "Testing both TCP and UDP server"
	python3 tinyscanner.py -u 127.0.0.1 -p 8080-8082; python3 tinyscanner.py -t 127.0.0.1 -p 8080-8082
