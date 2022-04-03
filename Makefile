service_name = go-boilerplate

build:
	@echo " > Building [$(service_name)]..."
	@cd src && go build -o ../bin/$(service_name) && cd ..
	@echo " > Finished building [$(service_name)]"

# Test
test:
	@echo " > Testing starts..."
	@cd src && go test -race -cover && cd ..
	@echo " > Finished testing"

# RUN
run: build
	@echo " > Running [$(service_name)]..."
	@ ./bin/$(service_name)
	@echo " > Finished running [$(service_name)]"

benchmark:
	@wrk -t2 -c100 -d30s http://127.0.0.1:8080/heroes/1