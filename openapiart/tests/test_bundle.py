import pytest
import jsonpath_ng as jp
import json


def test_config(api):
    config  = api.config()
    config.a = 'asdf'
    config.b = 1.1
    config.c = 1
    config.d_values = [config.A, config.B, config.C]
    config.e.e_a = 1.1
    config.e.e_b = 1.2
    config.f.f_a = 'a'
    djson = json.loads(config.serialize(config.JSON))
    assert(jp.parse('$.a').find(djson)[0].value == config.a)
    assert(jp.parse('$.f.f_a').find(djson)[0].value == config.f.f_a)
    print(config)
    

if __name__ == '__main__':
    pytest.main(['-v', '-s', __file__])
