This is a pratice session from Liz Rice's famous [Building a container from scratch in Go](https://www.youtube.com/watch?v=Utf-A4rODH8&list=PLsGzjZLGKzelYamS5vfclKVNy9CrzT0t7&index=3) talk at Container Camp

Learned about:
- Containers
- Namespaces
- Control groups (cgroups -> about limiting resources)
  - CPU
  - Memory
  - Disk I/O
  - Network
  - Device Permissions
- Process IDs -> `ps` doesn't directly look at the list of running processes, it looks in `/proc`
- Linux syscall definitions
- File system
- Container Image
