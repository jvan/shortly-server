import requests

from nose.tools import assert_equal, assert_in, eq_

from . import construct_url

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

def get_user_format_test():
    '''The /users/:user_id route should return valid json.'''

    url = construct_url('users', '1')

    response = requests.get(url)

    # Verify that the call was successful.
    assert_equal(response.status_code, 200)

    data = response.json()

    # Verify that the respose is ember-data compatible.
    assert_in('user', data)

    # Verify the format of the user object.
    user = data['user']

    assert_in('id', user)
    assert_in('name', user)
    assert_in('links', user)

    # The links object should be a list of integer ids
    links = user['links']
    assert_equal(type(links), list)
    for link in links:
	assert_equal(type(link), int)

    # Verify the testing data.
    assert_equal(user['name'], 'alice@gmail.com')

    assert_equal(len(links), 3)
    assert_equal(links, [1,2,3])
    
