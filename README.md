### Imagem Docker Hub

- [https://hub.docker.com/r/rschiavi0/go-hpa](https://hub.docker.com/r/rschiavi0/go-hpa)


###### Testar hpa
- Criar pod: kubectl run --rm -it loader --image=busybox /bin/sh
- Executar: while true; do wget -q -O- http://go-hpa.default.svc.cluster.local:8000; done;
