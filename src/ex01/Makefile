build: 
	go build
clean: 
	rm -rf main
	rm -rf test*
rebuild: clean build
test: build test01
test01:
	./main --old ../testfiles/original_database.xml --new ../testfiles/stolen_database.json > test01.txt