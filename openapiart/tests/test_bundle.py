import pytest


def test_config(api):
    config  = api.config()
    config.a = 'asdf'
    config.b = 1.1
    config.c = 1
    config.d = [config.A, config.B, config.C]
    config.e.e_a = 1.1
    config.e.e_b = 1.2
    config.f.f_a = 'a'
    print(config)
    

if __name__ == '__main__':
    pytest.main(['-v', '-s', __file__])
