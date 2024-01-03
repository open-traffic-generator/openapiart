import pytest


def test_boundary_val_scham(api):

    #  This test checks the values in boundary val schema.
    #  Objective is to check if values are set properly no problem in serialize and deserialize.

    config = api.test_config()
    b_val = config.native_features.boundary_val

    assert b_val.int_val == 50
    assert b_val.num_val == 50.05

    b_val.int_val = 100
    b_val.num_val = 100.0

    s_obj = config.serialize()
    config.deserialize(s_obj)
    b_val = config.native_features.boundary_val

    assert b_val.int_val == 100
    assert b_val.num_val == 100.0


def test_error_for_min_max(api):

    # This test basically checks for errors if value not in range of min max

    config = api.test_config()
    b_val = config.native_features.boundary_val

    b_val.int_val = 3

    with pytest.raises(Exception) as execinfo:
        config.serialize()

    assert "expected min 5, expected max 100" in execinfo.value.args[0]

    b_val.int_val = 200

    with pytest.raises(Exception) as execinfo:
        config.serialize()

    assert "expected min 5, expected max 100" in execinfo.value.args[0]


def test_boundary_val_array_schema(api):

    #  This test checks the values in boundary val schema.
    #  Objective is to check if values are set properly no problem in serialize and deserialize.

    config = api.test_config()
    b_val = config.native_features.boundary_val_array

    b_val.int_vals = [50, 51]
    b_val.num_vals = [50.05, 78.90]

    s_obj = config.serialize()
    config.deserialize(s_obj)
    b_val = config.native_features.boundary_val_array

    assert b_val.int_vals == [50, 51]
    assert b_val.num_vals == [50.05, 78.90]


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
