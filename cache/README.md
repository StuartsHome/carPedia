## Redis


- Start Redis server via launchctl
$ launchctl load ~/Library/LaunchAgents/homebrew.mxcl.redis.plist

- Start Redis server using configuration file
$ redis-server /usr/local/etc/redis.conf

- To test if Redis server is running
$ redis-cli ping

- Location of Redis configuration file
$ /usr/local/etc/redis.conf

- Shutdown Redis
$ redis-cli shutdown

## To use
1. $ redis-cli
2. $ SELECT 10
3. $ SET key1 value
4. $ GET key1
5. $ APPEND key1 1
6. $ KEYS *
7. $ SET KEY2 value3 EX 10  