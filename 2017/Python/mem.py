# -*- coding: utf-8 -*-

# Memoizer intended to be used as a decorator
# for module clients

def memoize_me_please(func):
    mem = {}
    def aux(*args):
        if args not in mem:
            mem[args]=func(*args)
        return mem[args]
    return aux