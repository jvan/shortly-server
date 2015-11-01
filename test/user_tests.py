import requests
import os

from nose.tools import assert_equal, assert_in, eq_

API_ROOT = 'http://localhost:1323'

def construct_url(*args):
    ''' Returns a valid url for api call.

    Examples:
        construct_url('users')      => https://54.165.29.115/api/users
        construct_url('users', '1') => https://54.165.29.115/api/users/1
    '''

    return os.path.join(API_ROOT, *args)


def get_all_users_format_test():
    '''The /users/ route should return valid json.'''

    url = construct_url('users')
    response = requests.get(url)

    # Verify that the call was successful.
    assert_equal(response.status_code, 200)

    data = response.json()

    # Verify that the respose is ember-data compatible.
    assert_in('users', data)

    users = data['users']

    # Verify the testing data.
    assert_equal(len(users), 3)

    user_names = [user['name'] for user in users]

    assert_in('alice@gmail.com', user_names)
    assert_in('bob@gmail.com',   user_names)
    assert_in('bob@gmail.com',   user_names)

    # Verify the format of the individual user objects.
    user = users[0]

    assert_in('id',   user)
    assert_in('name', user)
