# Scan Images for Vulnerabilities with CoreOS Clair

#### Step 1 - Deploy Postgres
Download Clair's Docker Compose File and Config

Clair requires a Postgres instance for storing the CVE data and it's service that will scan Docker Images for vulnerabilities. This has been defined within a Docker Compose file. Download it with the command below:

`curl -OL https://raw.githubusercontent.com/coreos/clair/master/contrib/compose/docker-compose.yml`

The Clair configuration defines how Images should be scanned. Download it with:

`mkdir clair_config && curl -L https://raw.githubusercontent.com/coreos/clair/master/config.yaml.sample -o clair_config/config.yaml`

Update Config

Set the version of Clair to the last stable release and the default database password.

`sed 's/clair-git:latest/clair:v2.0.1/' -i docker-compose.yml && \
  sed 's/host=localhost/host=postgres password=password/' -i clair_config/config.yaml`

Start DB

Start the database below.

`docker-compose up -d postgres`

In the next step we'll populate the DB

