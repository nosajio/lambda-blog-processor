bpd: build prepare deploy clean

build:
	@echo "Building lambda fn......"
	GOOS=linux go build -o lambda
	@echo "\n"

prepare:
	@echo "Creating archive......"
	zip lambda.zip ./lambda
	@echo "\n"

deploy:
	@echo "Deploying to AWS lambda......"
	aws lambda update-function-code \
		--function-name blog-posts-processor \
		--zip-file fileb://./lambda.zip \
		--publish \
		--debug
	@echo "\n"

clean:
	@echo "Cleaning up......"
	rm ./lambda ./lambda.zip
	@echo "\n"