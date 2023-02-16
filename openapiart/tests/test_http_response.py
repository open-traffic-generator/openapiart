import json


def test_valid_version_check(api):
    try:
        api._version_check = True
        config = api.prefix_config()
        config.a = "asdf"
        config.b = 1.1
        config.c = 1
        config.required_object.e_a = 1.1
        config.required_object.e_b = 1.2
        config.d_values = [config.A, config.B, config.C]
        config.level.l1_p1.l2_p1.l3_p1 = "test"
        config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
        api.set_config(config)
    finally:
        api._version_check = False


def test_invalid_version_check(api):
    try:
        api.get_local_version().api_spec_version = "0.2.1"
        api._version_check = True
        config = api.prefix_config()
        config.a = "asdf"
        config.b = 1.1
        config.c = 1
        config.required_object.e_a = 1.1
        config.required_object.e_b = 1.2
        config.d_values = [config.A, config.B, config.C]
        config.level.l1_p1.l2_p1.l3_p1 = "test"
        config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
        api.set_config(config)
        raise Exception("expected version error")
    except Exception:
        pass
    finally:
        api.get_local_version().api_spec_version = "0.1.0"
        api._version_check = False
