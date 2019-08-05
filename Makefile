SERVICES = am_migrations pg_migrations db_systemids
APP_ENV = prod

mailreports:
	docker build -t db_mailreports -f Dockerfile.db_mailreports .

systemids:
	docker build -t db_systemids -f Dockerfile.db_systemids .
	
pgm:
	docker build -t pg_migrations -f Dockerfile.pgm .

aggregates:
	docker build -t db_aggregates -f Dockerfile.db_aggregates .

techupdate:
	docker build -t db_techupdate -f Dockerfile.db_techupdate .

amm:
	docker build -t am_migrations -f Dockerfile.amm .

all: pgm amm systemids aggregates techupdate mailreports

pushmailreports:
	docker tag db_mailreports:latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/db_mailreports:latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/db_mailreports:latest

pushamm: 
	docker tag am_migrations:latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/am_migrations:latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/am_migrations:latest

pushaggregates:
	docker tag db_aggregates:latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/db_aggregates:latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/db_aggregates:latest
	
pushtechupdate:
	docker tag db_techupdate:latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/db_techupdate:latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/db_techupdate:latest	

pushecr:
	$(foreach var,$(SERVICES),docker tag $(var):latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest;)

