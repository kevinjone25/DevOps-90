# docker network cmd

- docker network list 查看container network 狀況 和 ID
    - bridge 是  based on linux bridge 上
    - scope locally
- docker network inspect bridge 
    - 查看conf detail(包含name, ID, drivers)