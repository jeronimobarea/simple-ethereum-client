
tests:
	docker build . -t go_test_env
	docker run \
    	--rm \
    	go_test_env
