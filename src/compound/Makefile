build: 
	go build
clean: 
	rm -rf main
	rm -rf test*
rebuild: clean build
test: build test00 test01 test02
	
test00: 
	./main -f ../testfiles/original_database.xml > test00.json
	./main -f ../testfiles/stolen_database.json > test00.xml 
test01:
	./main --old ../testfiles/original_database.xml --new ../testfiles/stolen_database.json > test01.txt
test02:
	./main --old ../testfiles/snapshot1.txt --new ../testfiles/snapshot2.txt > test02.1.txt
	./main --old ../testfiles/snapshot3.txt --new ../testfiles/snapshot4.txt > test02.2.txt