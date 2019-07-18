# Nacelle Example

A small example application to showcase the basic features of the [nacelle](https://nacelle.dev) microservice framework.

## Overview

This example application is organized as an HTTP and gRPC API (with equivalent functionality) that publish a message via Redis Pub/Sub to a worker process. This message contains a URL, and the worker performs an HTTP GET on the given URL. The response and/or error produced by this action is stored into a unique key in Redis, which can be fetched by a subsequent request to either API.

The application entrypoints (the HTTP API, the gRPC API, and the worker) are located in the `cmd` directory.

The **main** function for each API boots nacelle with a initializer that dials Redis and a server initializer for the process provided by this library. The connection created by the former is injected into the later.

The **main** function for the worker boots with an initializer for a Redis Pub/Sub client, which adds additional subscription/unsubscription logic to the Redis connection.

## Building and Running

It is suggested that you build and run with Docker. Simply run `docker-compose up`. This will compile the three commands via multi-stage builds and start containers for the two APIs, a container for the worker, and a container for the Redis dependency.

## Usage

The following shows an example of a successful request to the [GitHub API](https://developer.github.com/v3) user's list. The following output has been trimmed to pretty-print the first two results.

```bash
$ curl -i http://localhost:5000/ -X POST -d 'http://github.com/api/users'
HTTP/1.1 202 Accepted
Date: Fri, 21 Jun 2019 02:58:14 GMT
Content-Length: 45
Content-Type: text/plain; charset=utf-8

{"id":"b70d1b06-4f8d-4d07-bbd3-fa753bee7474"}
```

```bash
$ curl -s http://localhost:5000/2166ac52-e1aa-423f-a573-3e881b881bd1 | jq -r '.body' | jq '.[:2]'
[
  {
    "login": "mojombo",
    "id": 1,
    "node_id": "MDQ6VXNlcjE=",
    "avatar_url": "https://avatars0.githubusercontent.com/u/1?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/mojombo",
    "html_url": "https://github.com/mojombo",
    "followers_url": "https://api.github.com/users/mojombo/followers",
    "following_url": "https://api.github.com/users/mojombo/following{/other_user}",
    "gists_url": "https://api.github.com/users/mojombo/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/mojombo/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/mojombo/subscriptions",
    "organizations_url": "https://api.github.com/users/mojombo/orgs",
    "repos_url": "https://api.github.com/users/mojombo/repos",
    "events_url": "https://api.github.com/users/mojombo/events{/privacy}",
    "received_events_url": "https://api.github.com/users/mojombo/received_events",
    "type": "User",
    "site_admin": false
  },
  {
    "login": "defunkt",
    "id": 2,
    "node_id": "MDQ6VXNlcjI=",
    "avatar_url": "https://avatars0.githubusercontent.com/u/2?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/defunkt",
    "html_url": "https://github.com/defunkt",
    "followers_url": "https://api.github.com/users/defunkt/followers",
    "following_url": "https://api.github.com/users/defunkt/following{/other_user}",
    "gists_url": "https://api.github.com/users/defunkt/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/defunkt/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/defunkt/subscriptions",
    "organizations_url": "https://api.github.com/users/defunkt/orgs",
    "repos_url": "https://api.github.com/users/defunkt/repos",
    "events_url": "https://api.github.com/users/defunkt/events{/privacy}",
    "received_events_url": "https://api.github.com/users/defunkt/received_events",
    "type": "User",
    "site_admin": false
  }
]
```

The following shows an example of a failed request.

```bash
curl -i http://localhost:5000/ -X POST -d 'not even a url'
HTTP/1.1 202 Accepted
Date: Fri, 21 Jun 2019 03:04:22 GMT
Content-Length: 45
Content-Type: text/plain; charset=utf-8

{"id":"cde60fec-1ec5-4945-83fd-bcfd86040f44"}
```

```bash
curl -s http://localhost:5000/47c99f6b-000b-4ac9-8d98-120186f105f6 | jq
{
  "body": "",
  "error": "Get not%20even%20a%20url: unsupported protocol scheme \"\""
}
```
