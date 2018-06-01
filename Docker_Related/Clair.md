# Scan Images for Vulnerabilities with CoreOS Clair

#### Step 1 - Deploy Postgres
Download Clair's Docker Compose File and Config

Clair requires a Postgres instance for storing the CVE data and it's service that will scan Docker Images for vulnerabilities. This has been defined within a Docker Compose file. Download it with the command below:

`curl -OL https://raw.githubusercontent.com/coreos/clair/master/contrib/compose/docker-compose.yml`

The Clair configuration defines how Images should be scanned. Download it with:

`mkdir clair_config && curl -L https://raw.githubusercontent.com/coreos/clair/master/config.yaml.sample -o clair_config/config.yaml`

##### Update Config

Set the version of Clair to the last stable release and the default database password.

`sed 's/clair-git:latest/clair:v2.0.1/' -i docker-compose.yml && \
  sed 's/host=localhost/host=postgres password=password/' -i clair_config/config.yaml`

##### Start DB

Start the database below.

`docker-compose up -d postgres`

In the next step we'll populate the DB

#### Step 2 - Populate DB

Download and load the CVE details for Clair to use.

```
curl -LO https://gist.githubusercontent.com/BenHall/34ae4e6129d81f871e353c63b6a869a7/raw/5818fba954b0b00352d07771fabab6b9daba5510/clair.sql
docker run -it \
    -v $(pwd):/sql/ \
    --network "${USER}_default" \
    --link clair_postgres:clair_postgres \
    postgres:latest \
        bash -c "PGPASSWORD=password psql -h clair_postgres -U postgres < /sql/clair.sql"
```

**Note:** Clair would do this by default, but can take 10/15 minutes to download.

#### Step 3 - Deploy Clair

With the DB populated, start the Clair service.

`docker-compose up -d clair`

We can now send it Docker Images to scan and return which vulnerabilities it contains.

#### Step 4 - Scan Image

Clair works by accepting Image Layers via a HTTP API. To scan all the layers, we need an way to send each layer and aggregate the respond. Klar is a simple tool to analyze images stored in a private or public Docker registry for security vulnerabilities using Clair.

Download the latest release from Github.

`curl -L https://github.com/optiopay/klar/releases/download/v1.5/klar-1.5-linux-amd64 -o /usr/local/bin/klar && chmod +x $_`

Using klar, we can now point it at images and see what vulnerabilities they contain, for example `quay.io/coreos/clair:v2.0.1.`

`CLAIR_ADDR=http://localhost:6060 CLAIR_OUTPUT=Low CLAIR_THRESHOLD=10 \
  klar quay.io/coreos/clair:v2.0.1`

The output can be tweaked, for example, only reporting High or Critical vulnerabilities. In this case, we return Low and Above.

#### Step 5 - JSON Output

To help with automation, Klar can return the results as JSON. This allows developers plug the scanning into an automated process, such as CI/CD pipeline.

Start by installing jq to provide an easy way to view the JSON output

`pacman -Syy jq`

By setting the `_JSONOUTPUT=true` parameter, the results will be in JSON which can be piped to another process, like `jq`.

`CLAIR_ADDR=http://localhost:6060 CLAIR_OUTPUT=High CLAIR_THRESHOLD=10 JSON_OUTPUT=true klar postgres:latest | jq`

#### Step 6 - Scan Private Image

The previous step scanned images from the public Docker Registry, but the same can be done for private Docker Registries.

Start a simple registry using 

`docker run -d --name registry -p 5000:5000 registry:2`

This will be made available via the Katacoda Proxy. Tag an existing Docker Image with the newly started Registry URL.

d`ocker tag postgres:latest 2886795331-5000-frugo01.environments.katacoda.com/postgres:latest`

Once tagged, the Image can be pushed to the Registry.

`docker push 2886795331-5000-frugo01.environments.katacoda.com/postgres:latest`

In the same way, the private registry image can now be scanned.

```CLAIR_ADDR=http://localhost:6060 \
  CLAIR_OUTPUT=Low CLAIR_THRESHOLD=10 \
  klar 2886795331-5000-frugo01.environments.katacoda.com/postgres:latest
```
