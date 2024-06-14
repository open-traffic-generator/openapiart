import importlib
import pytest

module = importlib.import_module("pyapi")


def test_add(api):
    config = api.test_config()

    config.native_features.iter_object.add(
        str_len="200", integer64_1=2132433546, integer64_2=5645336
    )
    print(config)
    assert config.native_features.iter_object[0].str_len == "200"
    yaml = config.serialize(encoding=config.YAML)
    config.deserialize(yaml)


def test_remove(api):
    config = api.test_config()
    config.native_features.iter_object.add(
        str_len="200", integer64_1=2132433546, integer64_2=5645336
    )
    config.native_features.iter_object.add(
        str_len="300", integer64_1=3892433546, integer64_2=989645336
    )
    config.native_features.iter_object.add(
        str_len="400", integer64_1=4789678546, integer64_2=4567645336
    )
    assert len(config.native_features.iter_object) == 3
    assert config.native_features.iter_object[0].str_len == "200"
    assert config.native_features.iter_object[1].str_len == "300"
    assert config.native_features.iter_object[2].str_len == "400"

    config.native_features.iter_object.remove(1)
    config.native_features.iter_object.add(
        str_len="500", integer64_1=4789678546, integer64_2=4567645336
    )
    assert len(config.native_features.iter_object) == 3
    assert config.native_features.iter_object[0].str_len == "200"
    assert config.native_features.iter_object[1].str_len == "400"
    assert config.native_features.iter_object[2].str_len == "500"


def test_append(api):
    config = api.test_config()
    config.native_features.iter_object.add(
        str_len="200", integer64_1=2132433546, integer64_2=5645336
    )
    config.native_features.iter_object.add(
        str_len="300", integer64_1=3892433546, integer64_2=989645336
    )
    config.native_features.iter_object.add(
        str_len="400", integer64_1=4789678546, integer64_2=4567645336
    )
    assert len(config.native_features.iter_object) == 3
    assert config.native_features.iter_object[0].str_len == "200"
    assert config.native_features.iter_object[1].str_len == "300"
    assert config.native_features.iter_object[2].str_len == "400"

    mixed_object = module.MixedObject()
    mixed_object.str_len = "500"
    mixed_object.integer64_1 = 4789678546
    mixed_object.integer64_2 = 4567645336
    config.native_features.iter_object.append(mixed_object)

    assert len(config.native_features.iter_object) == 4
    assert config.native_features.iter_object[0].str_len == "200"
    assert config.native_features.iter_object[1].str_len == "300"
    assert config.native_features.iter_object[2].str_len == "400"
    assert config.native_features.iter_object[3].str_len == "500"


def test_clear(api):
    config = api.test_config()
    config.native_features.iter_object.add(
        str_len="200", integer64_1=2132433546, integer64_2=5645336
    )
    config.native_features.iter_object.add(
        str_len="300", integer64_1=3892433546, integer64_2=989645336
    )
    config.native_features.iter_object.add(
        str_len="400", integer64_1=4789678546, integer64_2=4567645336
    )
    assert len(config.native_features.iter_object) == 3
    assert config.native_features.iter_object[0].str_len == "200"
    assert config.native_features.iter_object[1].str_len == "300"
    assert config.native_features.iter_object[2].str_len == "400"

    config.native_features.iter_object.clear()

    assert len(config.native_features.iter_object) == 0


def test_set(api):
    config = api.test_config()
    config.native_features.iter_object.add(
        str_len="200", integer64_1=2132433546, integer64_2=5645336
    )
    config.native_features.iter_object.add(
        str_len="300", integer64_1=3892433546, integer64_2=989645336
    )
    config.native_features.iter_object.add(
        str_len="400", integer64_1=4789678546, integer64_2=4567645336
    )
    assert len(config.native_features.iter_object) == 3
    assert config.native_features.iter_object[0].str_len == "200"
    assert config.native_features.iter_object[1].str_len == "300"
    assert config.native_features.iter_object[2].str_len == "400"

    mixed_object = module.MixedObject()
    mixed_object.str_len = "500"
    mixed_object.integer64_1 = 4789678546
    mixed_object.integer64_2 = 4567645336
    config.native_features.iter_object.set(2, mixed_object)

    assert len(config.native_features.iter_object) == 3
    assert config.native_features.iter_object[0].str_len == "200"
    assert config.native_features.iter_object[1].str_len == "300"
    assert config.native_features.iter_object[2].str_len == "500"


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
