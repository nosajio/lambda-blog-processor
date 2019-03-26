bpd: build prepare deploy clean

build:
	@echo "\n\nBuilding lambda fn......"
	GOOS=linux go build -o lambda

prepare:
	@echo "\n\nCreating archive......"
	zip lambda.zip ./lambda

deploy:
	@echo "\n\nDeploying to AWS lambda......"
	aws lambda update-function-code \
		--function-name blog-posts-processor \
		--zip-file fileb://./lambda.zip \
		--publish

clean:
	@echo "\n\nCleaning up......"
	rm ./lambda ./lambda.zip