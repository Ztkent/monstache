# Monstache
Go daemon that syncs MongoDB to ElasticSearch & OpenSearch in realtime.  
This fork is provides a simple framework to build and run a custom Monstache container.

[Monstache](https://github.com/rwynn/monstache/)   
[Monstache Documentation](https://rwynn.github.io/monstache-site/)   

### Support:
This version of Monstache is designed for:
- MongoDB 3.6+
- Elasticsearch 7.0+
- OpenSearch 2.0+ 

It uses the official [MongoDB Golang driver](https://github.com/mongodb/mongo-go-driver), and the community supported [Elasticsearch driver](https://github.com/olivere/elastic/v7).

## Plugins
To simplify the process of building plugins, this project provides a new configuration option: 
```
mapper-plugin-path = "local"
```
This option allows you to build plugins alongside the Monstache container.   
Elimating the need for dynamic loading improves plugin compatibilty and performance.

## Plugin Options
Plugins support the following functions:
```go
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error)
func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error)
func Pipeline(ns string, changeStream bool) (stages []interface, err error)
func Process(input*monstachemap.ProcessPluginInput) error
```
All plugin options are documented in `monstache/plugin/plugin.go`.   


## Compiled Plugins
Compiling plugins is still supported, you can follow the instructions from Monstache as reference:   
- [Plugin Build](https://github.com/rwynn/monstache/wiki/Go-plugin-guide)    
- [Plugin Options](https://rwynn.github.io/monstache-site/advanced/#golang)   
