redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
/usr/bin/redis-benchmark -h 9.142.209.107 -p 12378  -a ykz@1234   -c 50 -n 10000   -d 10  -t get,set
redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
/usr/bin/redis-benchmark -h 9.142.209.107 -p 12378  -a ykz@1234   -c 50 -n 10000   -d 20  -t get,set
redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
/usr/bin/redis-benchmark -h 9.142.209.107 -p 12378  -a ykz@1234   -c 50 -n 10000   -d 50  -t get,set
redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
/usr/bin/redis-benchmark -h 9.142.209.107 -p 12378  -a ykz@1234   -c 50 -n 10000   -d 100  -t get,set
redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
/usr/bin/redis-benchmark -h 9.142.209.107 -p 12378  -a ykz@1234   -c 50 -n 10000   -d 200  -t get,set
redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
/usr/bin/redis-benchmark -h 9.142.209.107 -p 12378  -a ykz@1234   -c 50 -n 10000   -d 1000  -t get,set
redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
/usr/bin/redis-benchmark -h 9.142.209.107 -p 12378  -a ykz@1234   -c 50 -n 10000   -d 5000  -t get,set
redis-cli -h 9.142.209.107 -p  12378  -a ykz@1234 info memory  |grep  -w used_memory
