# The event loop, the Gil and multiprocessing
### Sync
```sh
$ time python3.6 sync_bub_sort.py -j 5
id: 1 with 1000 items
id: 2 with 1000 items
id: 3 with 1000 items
id: 4 with 1000 items
id: 5 with 1000 items
done

real    0m1,374s
user    0m1,359s
sys     0m0,013s
```

This won't scale, and since each item has can be done independently it will be easy to compute it parallelized right?
Let's go over three different ways to accomplice that with python.

### Threading 
```sh
$ time python3.6 threading_bub_sort.py -j 5
done you have played yourself
id: 1 with 1000 items
id: 3 with 1000 items
id: 4 with 1000 items
id: 5 with 1000 items
id: 2 with 1000 items

real    0m1,420s
user    0m1,409s
sys     0m0,024s
```
Woops, that is not faster at all.
The GIL makes sure that only one thread is processed at a time.
But threading is old, surely this new python asyncio can help.
There are many blog post over how fast it is!

### Async
```sh
$ time python3.6 async_bub_sort.py -j 5
id: 2 with 1000 items
id: 3 with 1000 items
id: 4 with 1000 items
id: 5 with 1000 items
id: 1 with 1000 items

real    0m1,516s
user    0m1,487s
sys     0m0,028s
```
Again not faster...
Python asyncio does not make python faster, but it does make it more efficient doing io tasks. But this doesn't help with computing.

### Multi processing
```sh
$ time python3.6 multi_bub_sort.py -j 5
id: 3 with 1000 items
id: 1 with 1000 items
id: 4 with 1000 items
id: 2 with 1000 items
id: 5 with 1000 items
done

real    0m0,881s
user    0m2,329s
sys     0m0,036s
```
Finally it's faster.
Multiprocessing is the right choice when handling heavy cpu load.

TODO right more documentation.

