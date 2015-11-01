#!/usr/bin/python2
''' URL shortener algorithm test.

Each URL will be stored in a database. The ID value for the corresponding
record will be base-62 encoded to create a shorten URL.
'''

import math

ALPHABET = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890'

def base_convert(number, base):
    '''Convert a base-10 number to an arbitrary base.
    
    number  -- Original base-10 value.
    base    -- Integer base value to use for encoding.
    return  -- Array of coefficients in new base.
    ''' 

    if number == 0:
        return [0]

    digits = []
    while number > 0:
        rem = number % base
        digits.append(rem)
        number /= base

    digits.reverse()
    return digits

def encode(number):
    '''Encode a base-10 number in base-62.
    
    number -- Base-10 number to be encoded.
    return -- Encoded base-62 string.
    '''

    return ''.join(ALPHABET[i] for i in base_convert(number, 62))


def decode(key):
    '''Decode a base-64 key and return the base-10 value.
    
    key    -- String value of base-62 value.
    return -- Decoded base-10 integer.
    '''

    indices = [ALPHABET.find(c) for c in key]
    indices.reverse()
    total = 0
    for (i, n) in enumerate(indices):
        total += n * int(math.pow(62, i))
    return total


if __name__ == '__main__':

    from random import randint

    # Test the encoding/decoding functions.

    assert base_convert(0, 62)  == [0]
    assert base_convert(1, 62)  == [1]
    assert base_convert(61, 62) == [61]
    assert base_convert(62, 62) == [1, 0]
    assert base_convert(63, 62) == [1, 1]


    assert encode(0)  == 'a'
    assert encode(1)  == 'b'
    assert encode(61) == '0'
    assert encode(62) == 'ba'
    assert encode(63) == 'bb'

    assert decode('a')  == 0
    assert decode('b')  == 1
    assert decode('0')  == 61
    assert decode('ba') == 62
    assert decode('bb') == 63

    # Test randomly generate values and verify that the original value can be
    # obtained by decoding the ecoded string.
    for i in range(10000):
        val = randint(0, math.pow(62, 6))
        key = encode(val)
        print('{} => {}'.format(val, key))
        assert decode(key) == val
