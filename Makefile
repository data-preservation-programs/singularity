build:
	go build -o singularity .

buildall:
	go build ./...

generate:
	go generate ./...

lint:
	gofmt -s -w .
	golangci-lint run --fix
	staticcheck ./...

test:
	go test -coverprofile=coverage.out -coverpkg=./... ./...

diagram: build
	./singularity admin init
	schemacrawler.sh --server=sqlite --database=./singularity.db --command=schema --output-format=svg --output-file=docs/database-diagram.svg --info-level=maximum

DIR_NAME := $(shell ls -d LanguageTool-* 2>/dev/null | head -n 1)

languagetool:
	if [ -z "$(DIR_NAME)" ]; then \
	    echo "LanguageTool is not installed. Installing..." ; \
		curl -L https://raw.githubusercontent.com/languagetool-org/languagetool/master/install.sh | bash ; \
	else \
		echo "LanguageTool seems to be already installed in $$DIR_NAME. Skipping installation." ; \
	fi
	-cp -f ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt.bak ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt
	cp -f ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt.bak
	echo >>./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt
	cat ./spelling.txt >> ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt
	java -jar ./$(DIR_NAME)/languagetool-commandline.jar -l en-US docs/en/README.md

