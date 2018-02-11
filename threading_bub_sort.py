import sys
import random
import threading


def bub_sort(l):
    for i in range(len(l)-1):
        for j in range((len(l)-(i+1))):
            if l[j] > l[j+1]:
                l[j], l[j+1] = l[j+1], l[j]


def run_bub_sort(job_id, size):
    l = [random.randrange(10) for _ in range(size)]
    bub_sort(l)
    print(f'id: {job_id} with {len(l)} items')
    return l

if __name__ == "__main__":
    size = int(sys.argv[sys.argv.index('-s')+1]) if '-s' in sys.argv else 1000
    jobs = int(sys.argv[sys.argv.index('-j')+1]) if '-j' in sys.argv else 10

    threads = []
    for job_id in range(1, jobs+1):
        t = threading.Thread(target=run_bub_sort, args=(job_id, size,))
        threads.append(t)
        t.deamon = True
        t.start()

    print("done you have played yourself")
