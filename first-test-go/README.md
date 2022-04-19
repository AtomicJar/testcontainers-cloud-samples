# First Test with Golang
This is an example repository with a simple test for demonstrating [Testcontainers Cloud](https://testcontainers.cloud).

Close the repository:
```bash
git clone https://github.com/atomicjar/testcontainers-cloud-samples.git

cd testcontainers-cloud-samples/first-test-go
```

Install the Golang dependencies:
```
go get
```

Run the tests:
```bash
go test . -v
```

If the test succeeds, you should see logs from `testcontainers` and the created container:
```bash
=== RUN   TestWithRedis
2022/04/19 15:40:36 Starting container id: 396fd04ae829 image: docker.io/testcontainers/ryuk:0.3.3
2022/04/19 15:40:37 Waiting for container id 396fd04ae829 image: docker.io/testcontainers/ryuk:0.3.3
2022/04/19 15:40:37 Container is ready id: 396fd04ae829 image: docker.io/testcontainers/ryuk:0.3.3
2022/04/19 15:40:37 Starting container id: 37b6361734ff image: redis:latest
2022/04/19 15:40:37 Waiting for container id 37b6361734ff image: redis:latest
2022/04/19 15:40:38 Container is ready id: 37b6361734ff image: redis:latest
    main_test.go:40: M1:C 19 Apr 2022 13:40:38.036 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
        n1:C 19 Apr 2022 13:40:38.036 # Redis version=6.2.6, bits=64, commit=00000000, modified=0, pid=1, just started
        �1:C 19 Apr 2022 13:40:38.036 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
        D1:M 19 Apr 2022 13:40:38.037 * monotonic clock: POSIX clock_gettime
        C1:M 19 Apr 2022 13:40:38.038 * Running mode=standalone, port=6379.
        21:M 19 Apr 2022 13:40:38.038 # Server initialized
        1:M 19 Apr 2022 13:40:38.038 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
        ;1:M 19 Apr 2022 13:40:38.038 * Ready to accept connections
--- PASS: TestWithRedis (1.97s)
PASS
ok      github.com/AtomicJar/testcontainers-cloud-samples       2.289s
```