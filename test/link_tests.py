import requests

from nose.tools import assert_equal, assert_in, eq_

from . import construct_url

def get_link_format_test():
    '''The /link/:link_id route should return valid json'''

    url = construct_url('links', '1')

    response = requests.get(url)

    # Verify that the call was successful.
    assert_equal(response.status_code, 200)

    data = response.json()

    # Verify that the respose is ember-data compatible.
    assert_in('link', data)

    # Verify the format of the link object.
    link = data['link']

    assert_in('id',  link)
    assert_in('url', link)
    assert_in('key', link)

    assert_equal(type(link['url']), unicode)
    assert_equal(type(link['key']), unicode)

    # Verify the testing data.
    assert_equal(link['url'], 'http://duckduckgo.com')
    assert_equal(link['key'], 'b')
