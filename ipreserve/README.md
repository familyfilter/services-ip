Notes:

    GO is installed at /usr/local/go
    Set GOROOT: export GOROOT=/usr/local/go
    Set GOPATH: export GOPATH=/home/ubuntu/workspace/ipreserve
    To Run:
        cd /home/ubuntu/workspace/ipreserve
        go run src/main/ipreserve.go
       
    ETCD is installed at /home/ubuntu/etcd-v2.2.2-linux-amd64
    To Run:
        cd /home/ubuntu/etcd-v2.2.2-linux-amd64
        ./etcd

Usage:

    Right now ETCD does not need to be started to run the GO program but eventually
    it will be.  The GO program can be tested with the following calls:
        curl 0.0.0.0:8080/reserveipv6addr
        curl 0.0.0.0:8080/reserveipv6multiaddr
    