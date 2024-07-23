test:
	@docker build -t goexpdt-tests --progress plain \
		--no-cache-filter=run-tests-stage --target run-tests-stage .
