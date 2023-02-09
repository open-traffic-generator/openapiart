import pytest


def test_strict_choice_type_leaf_nodes_with_set(api):
    config = api.prefix_config()

    # checking the default choice and user_choice
    assert config.f._user_choice is None
    assert config.f.choice == "f_a"

    # setting the value to a particular choice
    config.f.f_a = "test"
    assert config.f._user_choice == "f_a"
    assert config.f.choice == "f_a"
    choice_error = "Cannot set/retrieve f_b, as f_a was already set earlier."

    # if we want to set the choice again we hsould have an exception
    with pytest.raises(Exception) as execinfo:
        config.f.f_b = 3.4
    assert execinfo.value.args[0] == choice_error


def test_strict_choice_type_first_user_input(api):
    config = api.prefix_config()
    assert config.f._user_choice is None
    assert config.f.choice == "f_a"

    # Although we have a default choice thats done internally by code base.
    # So the first choice should be allowed to set by the user.
    config.f.f_b = 3.45
    assert config.f._user_choice == "f_b"
    assert config.f.choice == "f_b"


def test_strict_choice_type_leaf_nodes_get(api):
    config = api.prefix_config()
    assert config.f._user_choice is None
    assert config.f.choice == "f_a"

    config.f.f_b = 3.45
    assert config.f._user_choice == "f_b"
    assert config.f.choice == "f_b"

    # The same choice should be allowed to fetch even after user choice is set.
    config.f.f_b

    # However if another choice is not allowed to be fetched if suer choice is set.
    choice_error = "Cannot set/retrieve f_a, as f_b was already set earlier."
    with pytest.raises(Exception) as execinfo:
        config.f.f_a = 3.4
    assert execinfo.value.args[0] == choice_error


def test_strict_choice_type_leaf_nodes_no_default(api):
    m = api.metrics_request()

    # ensuring no default value assigned
    assert m._user_choice is None

    # ensuring choice set properly
    m.port = "abcd"
    assert m.choice == "port"
    assert m._user_choice == "port"

    # Trying to override the choice should throw exception
    choice_error = "Cannot set/retrieve flow, as port was already set earlier."
    with pytest.raises(Exception) as execinfo:
        m.flow = "f1"
    assert execinfo.value.args[0] == choice_error


def test_strict_choice_type_with_pattern_nodes(api):
    config = api.prefix_config()
    ip = config.ipv4_pattern.ipv4

    # check default values
    assert ip._user_choice is None
    assert ip.choice == "value"

    ip.increment.count = 5
    ip.increment.step = "1.1.1.1"
    ip.increment.start = "1.1.1.1"
    ip.increment.count = 5

    # check set values
    assert ip._user_choice == "increment"
    assert ip.choice == "increment"

    # Trying to override the choice should throw exception
    choice_error = (
        "Cannot set/retrieve values, as increment was already set earlier."
    )
    with pytest.raises(Exception) as execinfo:
        ip.values = ["1.1.1.1"]
    assert execinfo.value.args[0] == choice_error
