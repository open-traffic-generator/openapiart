import pytest


def test_config(api):
    config  = api.config()
    config.a = 'asdf'
    config.b = 1.1
    config.c = 1
    print(config)
    

if __name__ == '__main__':
    pytest.main(['-v', '-s', __file__])
