CARTS_BINARY=cartsServiceApp
ORDERS_BINARY=ordersServiceApp
PRODUCTS_BINARY=productsServiceApp
REVIEWS_BINARY=reviewsServiceApp
USERS_BINARY=usersServiceApp
BROKER_BINARY=brokerServiceApp
# build services
build_broker:
	@echo "Building Broker"
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!, Built Broker"
build_carts:
	@echo "Building Carts"
	cd ../carts-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CARTS_BINARY} ./cmd/api
	@echo "Done!, Built Carts"
build_orders:
	@echo "Building Orders"
	cd ../orders-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ORDERS_BINARY} ./cmd/api
	@echo "Done!, Built Orders"
build_products:
	@echo "Building Products"
	cd ../products-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PRODUCTS_BINARY} ./cmd/api
	@echo "Done!, Built Products"
build_reviews:
	@echo "Building Reviews"
	cd ../reviews-service && env GOOS=linux CGO_ENABLED=0 go build -o ${REVIEWS_BINARY} ./cmd/api
	@echo "Done!, Built Reviews"
build_users:
	@echo "Building Users"
	cd ../users-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USERS_BINARY} ./cmd/api
	@echo "Done!, Built Users"

build_all: build_broker build_carts build_orders build_products build_reviews build_users
##
