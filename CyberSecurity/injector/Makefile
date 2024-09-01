testfiles:
	@echo "Building executables"
	@echo
	go build -o 01 ./testFiles/01.go
	@echo
	@echo "Contents of 01 executable:"
	./01
	@echo
	go build -o HelloWorld ./testFiles/HelloWorld.go
	@echo 
	@echo "Contents of HelloWorld executable:"
	./HelloWorld
	@echo

test: 
	go run injector.go 01 HelloWorld
	@echo 
	@echo "Contents of a new file:"
	@echo
	./result/result
	@echo