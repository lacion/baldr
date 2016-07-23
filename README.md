# baldr
Baldr is a utility cli for running services inside docker.

### Commands
baldr uses commands to perform diferent actions, most of the time their related to waiting for dependencies to be ready before continuing.

#### MongoDB

the mongodb command will wait for a mongodb instance to ready to accept connections before continuing or it will fails after some time.

		baldr mongodb -m mongodb://user:pass@db1:10013

you can also specify a wait time in ms before retries with -t and the number of times you want it to retry before bailing out.

		 baldr mongodb -m mongodb://user:pass@db1:10013 -t 5000 -r 10

#### etcd3
the etcd3 command will wait for a etcdv3 cluster instance to ready to accept connections before continuing or it will fails after some time.

		baldr etcd3 -e etcd1:2379,etcd2:2379,etcd3:2379

you can also specify a wait time in ms before retries with -t and the number of times you want it to retry before bailing out.

		 baldr etcd3 -e etcd1:2379,etcd2:2379,etcd3:2379 -t 5000 -r 10

#### micro
the micro command will look in the go-micro registry to a micro service if its not registered it will retry a few time untill its there or it bails out.

		baldr micro -s foo.service -g etcd3:2379

you can also specify a wait time in ms before retries with -t and the number of times you want it to retry before bailing out.

		 baldr micro -s foo.service -g etcd3:2379 -t 5000 -r 10