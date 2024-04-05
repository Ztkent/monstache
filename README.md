# Monstache
Go daemon that syncs MongoDB to ElasticSearch & OpenSearch in realtime.  
This fork is provides a simple framework to build and run a custom Monstache container.

[Monstache Documentation](https://rwynn.github.io/monstache-site/)   
[Monstache](https://github.com/rwynn/monstache/)   

### Support:
This version of Monstache is designed for:
- MongoDB 3.6+
- Elasticsearch 7.0+
- OpenSearch 2.0+ 

It uses the official [MongoDB Golang driver](https://github.com/mongodb/mongo-go-driver), and the community supported [Elasticsearch driver](https://github.com/olivere/elastic/v7).

## Plugins
Follow the instructions from Monstache as reference when building a new plugin:   
- [Plugin Options](https://rwynn.github.io/monstache-site/advanced/#golang)   
- [Plugin Build](https://github.com/rwynn/monstache/wiki/Go-plugin-guide)

All plugins should be implemented in a `plugin.go` file.   
It is expected in the `/monstache/plugin` directory, this can be changed in the `Dockerfile`.  

Plugins support the following functions:
```go
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error)
func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error)
func Pipeline(ns string, changeStream bool) (stages []interface, err error)
func Process(input*monstachemap.ProcessPluginInput) error
```

`MapperPluginInput` and `MapperPluginOutput` are defined in the `/monstache/monstachemap` package.  
These can be modified to enhance your plugin functionality.   