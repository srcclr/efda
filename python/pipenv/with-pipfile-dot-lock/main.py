import hashlib

from jwt import algorithms

def inner_func():
    algorithms.HMACAlgorithm(hashlib.sha256).prepare_key('')

if __name__ == '__main__':
    inner_func()
