import pytest


def test_config(api):
    config  = api.config()
    config.a = 'asdf'
    config.b = 1.1
    config.c = 1
    config.d = [config.A, config.B, config.C]
    print(config)
    

if __name__ == '__main__':
    pytest.main(['-v', '-s', __file__])
