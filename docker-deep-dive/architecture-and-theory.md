# Architecture Big Picture

## Container
- Isolated area of an OS with limited resources applied

## Kernal Internals
- Namespaces
  - Multiple isolated virtual operating systems
  - Each linux namespace contains:
    - Process ID (pid)
    - Network (net)
    - Filesystem/mount (mnt)
    - Inter-pro comms (ipc)
    - UTS (uts)
    - User (user)
- Control groups
  - grouping processes
  - Imposing resource limits

## Docker Engine
- a lightweight and powerful open source containerization technology combined with a work flow for building and containerizing your applications.