build: 
	go build
clean: 
	rm -rf main
	rm -rf test*
rebuild: clean build
test: build test02
test02:
	./main --old ../testfiles/snapshot1.txt --new ../testfiles/snapshot2.txt > test02.1.txt
	./main --old ../testfiles/snapshot3.txt --new ../testfiles/snapshot4.txt > test02.2.txt