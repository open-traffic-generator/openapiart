import json
import os


CONFIGS_DIR = "tests/json_configs"


def compare_json(d1, d2):
    d1 = sorting(json.loads(d1))
    d2 = sorting(json.loads(d2))
    if d1 == d2:
        return True
    return False


def sorting(item):
    if isinstance(item, dict):
        return sorted((key, sorting(values)) for key, values in item.items())
    if isinstance(item, list):
        return sorted(sorting(x) for x in item)
    else:
        return item


def get_root_dir():
    return os.path.dirname(os.path.dirname(os.path.abspath(__file__)))


def get_test_config_path(config_name):
    return os.path.join(
        os.path.dirname(get_root_dir()), CONFIGS_DIR, config_name)

