SERVICES = am_migrations pg_migrations db_systemids
APP_ENV = dev
systemids:
	docker build -t db_systemids -f Dockerfile.db_systemids .
	
pgm:
	docker build -t pg_migrations -f Dockerfile.pgm .

aggregates:
	docker build -t aggregates -f Dockerfile.aggregates .

amm:
	docker build -t am_migrations -f Dockerfile.amm .

all: pgm amm systemids aggregates

pushamm: 
	docker tag am_migrations:latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/am_migrations:latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/am_migrations:latest

pushecr:
	$(foreach var,$(SERVICES),docker tag $(var):latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest;)

