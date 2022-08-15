include .env
export DISPLAY=:0

# ==================
# = Makefile rules =
# ==================

run:
	@echo "Running..."
	go run cmd/app/main.go

build:
	@echo "Maintaining..."
# @echo "Building image..."
# sudo docker buildx build -t ecom-api:0.0.1 --load .

compose-up:
	@echo "Maintaining..."
# sudo docker run -p 5001:8080 ecom-api:0.0.1

remove:
	@echo "Removing image..."
	sudo docker image rm -f ecom-api:0.0.1