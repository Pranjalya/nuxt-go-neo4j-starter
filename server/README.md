# goneo4jgql

Neo4j backend using Golang, GraphQL and Docker

## Stack

* [Go 1.14](https://golang.org/doc/go1.14)
* [gqlgen](https://github.com/99designs/gqlgen)
* [neo4j-go-driver](https://github.com/neo4j/neo4j-go-driver)
* [Neo4j apoc plugin](https://github.com/neo4j-contrib/neo4j-apoc-procedures/releases/download/3.5.0.9/apoc-3.5.0.9-all.jar)
* [seabolt connector v1.7.4](https://github.com/neo4j-drivers/seabolt)
* [Neo4j Server v3.5.17](https://github.com/neo4j/neo4j/wiki/Neo4j-3.5-changelog#3517)
* [Docker](https://www.docker.com/)


## Introduction

I've implemented a GraphQL API on top of Neo4j using Golang. I have included the [Neo4j movie example dataset](https://neo4j.com/developer/guide-cypher-basics/) that has the following domain model:

![domain](./docs/i/domain.png)

**Note**: You can find the cypher queries I used for the movies datset here, at the end of the [movies.cypher import file](neo4j/import/movies.cypher)


## Setup (Using Docker)

You just need to install [Docker](https://docs.docker.com/get-docker/) and [Docker compose](https://docs.docker.com/compose/) (included with the latest version of Docker).

1 - Clone this repo
```bash
git clone https://github.com/Pranjalya/nuxt-go-neo4j-starter.git
```

2 - Build app docker image
```bash
docker-compose build
```

3 - Startup Neo4j and goneo4jgql containers:
```bash
docker-compose up -d
```

Check the logs
```bash
docker-compose logs -f
```

You should get an output like this one:
```
$ docker-compose logs -f
Attaching to goneo4jgql, neo4j
goneo4jgql | time="2020-04-16T05:54:09Z" level=info msg="Connected to Neo4j Server" neo4j_server_uri="bolt://neo4j:7687" prefix=main
goneo4jgql | time="2020-04-16T05:54:09Z" level=info msg="API Listening" api_url="0.0.0.0:8080" prefix=main
neo4j    | Fetching versions.json for Plugin 'apoc' from https://neo4j-contrib.github.io/neo4j-apoc-procedures/versions.json
neo4j    | Installing Plugin 'apoc' from https://github.com/neo4j-contrib/neo4j-apoc-procedures/releases/download/3.5.0.9/apoc-3.5.0.9-all.jar to /plugins/apoc.jar
neo4j    | Applying default values for plugin apoc to neo4j.conf
neo4j    | Skipping dbms.security.procedures.unrestricted for plugin apoc because it is already set
neo4j    | Active database: graph.db
neo4j    | Directories in use:
neo4j    |   home:         /var/lib/neo4j
neo4j    |   config:       /var/lib/neo4j/conf
neo4j    |   logs:         /logs
neo4j    |   plugins:      /plugins
neo4j    |   import:       /var/lib/neo4j/import
neo4j    |   data:         /var/lib/neo4j/data
neo4j    |   certificates: /var/lib/neo4j/certificates
neo4j    |   run:          /var/lib/neo4j/run
neo4j    | Starting Neo4j.
neo4j    | 2020-04-16 05:54:27.396+0000 INFO  ======== Neo4j 3.5.17 ========
neo4j    | 2020-04-16 05:54:27.470+0000 INFO  Starting...
neo4j    | 2020-04-16 05:54:59.164+0000 INFO  Bolt enabled on 0.0.0.0:7687.
neo4j    | 2020-04-16 05:55:05.645+0000 INFO  Started.
neo4j    | 2020-04-16 05:55:09.186+0000 INFO  Remote interface available at http://localhost:7474/
```

4 - Load movie dataset
```bash
docker-compose exec neo4j /bin/bash -c 'cat /var/lib/neo4j/import/movies.cypher | cypher-shell -u neo4j -p test'
```

That's all. You should be able to login to Neo4j browser at [http://127.0.0.1:7474/browser/](http://127.0.0.1:7474/browser/)

Use the default **credentials** set in [docker-compose.yml file](./docker-compose.yml)
```
user: neo4j
pass: test
```

![browser](./docs/i/neo4j_browser.png)


## GraphQL API Usage

You should be able to access Playground at [http://0.0.0.0:8080/playground](http://0.0.0.0:8080/playground):

![browser](./docs/i/playground.png)


### GraphQL queries examples

**Get the list of movies** 

```graphql
query movies {
  movies {
    title
    released
    tagline
  }
}
```

![browser](./docs/i/movie_list.png)]


**Get the list of movies with title containing "top"** 
```graphql
query movies {
  movies (title: "top") {
    title
    released
    tagline
  }
}
```

**Get cast, directors and writer data**
```graphql
query movies {
  movies (title: "top") {
    title
    released
    tagline
    cast {
      name
    }
    directors {
      name
    }
    writers {
      name
    }
  }
}
```

![browser](./docs/i/cast.png)]


**Get the list of participations for each cast member for Top Gun**
```graphql
query movies {
  movies (title: "top") {
    title
    released
    tagline
    cast {
      name
      participated {
        role
        movie {
          title
        }
      }
    }
  }
}
```

![browser](./docs/i/participations.png)


## Final notes

* I haven't included any dataloader yet, so expect performance issues for complex graphql queries.
* This is a very simple example made as a proof of concept for a neo4j-grapqhl-go stack. I'm not including any interesting query to take advantage of the real power of graph dbs (at least not in this first version).
* I haven't added any graphql depth/complexity limiting mechanism, so take that into consideration when executing complex queries.
* I used Neo4j v3.5 instead of v4 because bolt connector does not support yet the latest v4 protocol.
