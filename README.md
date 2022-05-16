# PDDikti Proxy API

Proyek ini merupakan eksperimen _fine tuning_ pribadi dari API PDDikti yang asli.
## Optimizations

- Server caching that significantly improve similar request
- Additional filtering
- Personalized Rest API URL Pattern
## Dokumentasi API

#### Daftar Perguruan Tinggi

```http
  GET /api/v1/perguruan-tinggi
```

#### Data Tunggal Perguruan Tinggi

```http
  GET /api/v1/perguruan-tinggi/:mode/:query
```

| Parameter | Type     | Description                                         |
| :-------- | :------- | :---------------------------------------------------|
| `mode`    | `string` | `{k : Berdasarkan kode_pt, n : Berdasarkan nama_pt}`|
| `query`   | `string` |                                                     |

#### Cari Perguruan Tinggi Berdasarkan Nama

```http
  GET /api/v1/perguruan-tinggi/search/:query
```

| Parameter | Type     | Description |
| :-------- | :------- | :-----------|
| `query`   | `string` |             |

#### Daftar Prodi Perguruan Tinggi

```http
  GET /api/v1/perguruan-tinggi/k/:query/daftar-prodi
```

| Parameter | Type     | Description                |
| :-------- | :------- | :--------------------------|
| `mode`    | `string` | `{k : Berdasarkan kode_pt}`|
| `query`   | `string` |                            |


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
