
# PDDikti Proxy API

This is a proxy API from the original PDDikti API with custom Rest Api personalized url pattern and some optimizations.
## Optimizations

- Server caching that significantly improve similar request
- Additional filtering
- Personalized Rest API URL Pattern
## API Documentation
https://postman.com/noobscoder/workspace/pddikti-proxy

#### Get Universities

```http
  GET /api/v1/universities
```

#### Get University

```http
  GET /api/v1/university/:mode/:query
```

| Parameter | Type     | Description                                                          |
| :-------- | :------- | :--------------------------------------------------------------------|
| `mode`    | `string` | `{c : kode_pt (university's code), n : nama_pt (university's name) }`|
| `query`   | `string` |                                                                      |

#### Search by University's Name

```http
  GET /api/v1/university/search/:query
```

| Parameter | Type     | Description |
| :-------- | :------- | :-----------|
| `query`   | `string` |             |

#### University's Majors

```http
  GET /api/v1/university/k/:query/majors
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :----------------------------------|
| `mode`    | `string` | `{c : kode_pt (university's code)}`|
| `query`   | `string` |                                    |


## Acknowledgements

 - [Situs PDDikti asli](https://pddikti.kemdikbud.go.id/)
 
## Tech Stack

Golang with GoFiber API Framework


## Lessons Learned

- Memory safety: Value, Reference, Pointer, Mutex, Channel
- Cron Job/Scheduler
## Authors

Hi, I am [@luthfiyufajjiru](www.linkedin.com/in/yufajjiru) and I am a software engineer at AILIMA.
Have been working on client side and/or server side, even work with many database system like Postgresql, Neo4j, and MongoDb.
I also having a good understanding on IT Infrastructure in many cloud providers. Have been design a private network and the
load balancing in cloud provider, dockerize an application, and manage container in Kubernetes or Podman.
