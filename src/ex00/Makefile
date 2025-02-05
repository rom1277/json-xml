build: 
	go build
clean: 
	rm -rf main
	rm -rf test*
rebuild: clean build
test: build test00
test00: 
	./main -f ../testfiles/original_database.xml > test00.json
	./main -f ../testfiles/stolen_database.json > test00.xml 