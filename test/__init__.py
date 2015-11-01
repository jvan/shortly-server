import os

API_ROOT = 'http://localhost:1323'

def construct_url(*args):
    ''' Returns a valid url for api call.

    Examples:
        construct_url('users')      => https://54.165.29.115/api/users
        construct_url('users', '1') => https://54.165.29.115/api/users/1
    '''

    return os.path.join(API_ROOT, *args)

