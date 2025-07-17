#DEV

build-dev: 
	docker build -t github.com/rohit21755/go_webrtc -f containers/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn .

clean-dev:
	docker-compose -f containers/composes/dc.dev.yml.down

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up
