import pytest


def test_getting_leaf_nodes_without_default(api):
    m = api.metrics_request()

    # setting properties should change choice
    m.port = "p1"
    assert m.choice == "port"
    assert m._properties.get("port", None) == "p1"
    assert m._properties.get("flow", None) is None
    m.flow = "f1"
    assert m.choice == "flow"
    assert m._properties.get("port", None) is None
    assert m._properties.get("flow", None) == "f1"

    # should be able to set choice as well
    m.choice = "port"
    assert m._properties.get("choice") == "port"
    assert m.choice == "port"
    m.choice = "flow"
    assert m._properties.get("choice") == "flow"
    assert m.choice == "flow"


def test_getting_leaf_nodes_with_default(api):
    config = api.prefix_config()

    # checking default values
    assert config.f.choice == "f_a"
    assert config.f._properties.get("choice") == "f_a"

    # fetching properties should not change choice
    config.f.f_a
    assert config.f.choice == "f_a"
    config.f.f_b
    assert config.f.choice == "f_a"

    # setting properties should change choice
    config.f.f_a = "p1"
    assert config.f.choice == "f_a"
    assert config.f._properties.get("f_a", None) == "p1"
    assert config.f._properties.get("f_b", None) is None
    config.f.f_b = 3.45
    assert config.f.choice == "f_b"
    assert config.f._properties.get("f_a", None) is None
    assert config.f._properties.get("f_b", None) == 3.45


def test_get_set_for_choice_heirarchy(api):
    config = api.prefix_config()
    j = config.j.add()

    # check default for parent
    assert j.choice == "j_a"
    assert j._properties.get("choice") == "j_a"

    # fetching properties should not change choice
    j.j_b.f_b
    assert j.j_b.choice == "f_a"
    assert j.j_b._properties.get("f_a") == "some string"
    assert j.j_b._properties.get("f_b") is None
    j.j_b.f_a
    assert j.j_b.choice == "f_a"
    assert j.j_b._properties.get("f_a") == "some string"
    assert j.j_b._properties.get("f_b") is None

    # mix of set and get of properties should handle choice properly
    j.j_b.f_b = 3.4
    assert j.j_b.choice == "f_b"
    assert j.j_b._properties.get("f_b") == 3.4
    assert j.j_b._properties.get("f_a") is None
    j.j_b.f_a
    assert j.j_b.choice == "f_b"
    assert j.j_b._properties.get("f_b") == 3.4
    assert j.j_b._properties.get("f_a") is None

    j.j_b.f_a = "asd"
    assert j.j_b.choice == "f_a"
    assert j.j_b._properties.get("f_a") == "asd"
    assert j.j_b._properties.get("f_b") is None
    j.j_b.f_b
    assert j.j_b.choice == "f_a"
    assert j.j_b._properties.get("f_a") == "asd"
    assert j.j_b._properties.get("f_b") is None


def test_get_set_for_parent_choice_objects(api):

    config = api.prefix_config()
    j = config.j.add()

    # check default for parent
    assert j.choice == "j_a"
    assert j._properties.get("j_a", None) is not None
    assert j._properties.get("j_b", None) is None

    # fetching properties should not change choice
    f = j.j_b
    assert j.choice == "j_a"
    assert j._properties.get("j_a", None) is not None

    # setting properties should change parent choice
    f.f_a = "asd"
    assert j.choice == "j_b"
    assert j._properties.get("j_b", None) is not None
    assert j._properties.get("j_a", None) is None

    j.j_a.e_a = 123
    assert j.choice == "j_a"
    assert j._properties.get("j_a", None) is not None
    assert j._properties.get("j_b", None) is None
