Integrating and improving with  https://github.com/rdegges/ipify-api

### Building

```
git clone https://github.com/TelenLiu/ip.git 
```

```
go mod tidy
go build
```



### Docker Build

```
docker buildx build --platform linux/amd64,linux/arm64,linux/386,linux/arm/v7 . -t telenliu/ip:23.4.18 --push
```



###  Docker Run

```
docker run -d -p 3000:3000 --restart=always telenliu/ip
```





### API Get

Integrating  https://www.ipify.org/，replae [api64.ipify.org] for your host

```
https://api64.ipify.org	  98.207.254.136 or 2a00:1450:400f:80d::200e
https://api64.ipify.org?format=json	 {"ip":"98.207.254.136"} or {"ip":"2a00:1450:400f:80d::200e"}
https://api64.ipify.org?format=jsonp	 	callback({"ip":"98.207.254.136"}); or callback({"ip":"2a00:1450:400f:80d::200e"});
https://api64.ipify.org?format=jsonp&callback=getip	 getip({"ip":"98.207.254.136"}); or getip({"ip":"2a00:1450:400f:80d::200e"});
```



