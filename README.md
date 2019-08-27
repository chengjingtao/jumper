# Description
Centralized SSH Servers info Management 
# Features
- manage servers infoes in center (git repository)
- update servers infoes from center
- jump to target server
# Using
- jump repo add default https://github.co/example/servers.git
- jump repo remove default
- jump repo update default
- jump repo list
- jump repo inspect default 
  show areas in current repo
- jump to {area}/{name}
  will execute ssh user@ip
# Examples
```
➜ jump repo list           
default  https://github.com/chengjingtao/servers.git
➜ jump repo update default
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Compressing objects: 100% (4/4), done.
Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
➜ jump repo inspect default
AREA: devops
 => int root@1.1.1.1
➜ jump to devops/int
/usr/bin/ssh ssh root@1.1.1.1
```


docker run --rm \
    -v $(pwd):/workspace \
    -e DOCKER_CONFIG=/root/.docker \
    -v ${HOME}/.docker:/root/.docker \
    index.alauda.cn/alaudak8s/kaniko-project-executor:latest \
    --dockerfile /workspace/images/Dockerfile --destination index.alauda.cn/alaudak8s/jumper:v0.0.1 --context /workspace