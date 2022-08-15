include .env
export

# ==================
# = Makefile rules =
# ==================

run:
	go run cmd/app/main.go
	