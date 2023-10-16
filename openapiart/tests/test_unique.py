import pytest


# @pytest.mark.skip(reason="shall be restored")
def test_unique(config):
    # Update: There is no global and local diff
    # everything is considered as global now
    # *************** global unique ****************
    # Two similar objects with same Name.
    config.w_list.wobject(w_name="global_unique_similar_obj")
    config.w_list.wobject(w_name="global_unique_similar_obj")
    try:
        config.serialize()
        pytest.fail("validation failed")
    except Exception as e:
        if "global_unique_similar_obj already exists" not in str(e):
            pytest.fail("global_unique_similar_obj validation failed")

    # Two similar objects with different name
    config.w_list[1].w_name = "global_unique_similar_obj1"
    config.serialize()

    # Two different objects with same name
    config.name = "global_unique"
    config.w_list.wobject(w_name="global_unique")
    try:
        config.serialize()
        pytest.fail("validation failed")
    except Exception as e:
        if "global_unique already exists" not in str(e):
            pytest.fail("global_unique validation failed")

    #  Two different objects with different name
    config.name = "global_unique1"
    config.serialize()

    # *************** local unique ****************
    # Two similar objects with same Name.
    config.x_list.zobject(name="local_unique_similar")
    config.x_list.zobject(name="local_unique_similar")
    try:
        config.serialize()
        pytest.fail("validation failed")
    except Exception as e:
        if "local_unique_similar already exists" not in str(e):
            pytest.fail("local_unique_similar validation failed")

    # Two similar objects with different name
    config.x_list[1].name = "local_unique_similar1"
    config.serialize()

    # Two different objects with same name
    config.name = "local_global_mix"
    config.x_list.zobject(name="local_global_mix")
    try:
        config.serialize()
        pytest.fail("validation failed")
    except Exception as e:
        if "local_global_mix already exists" not in str(e):
            pytest.fail("local_unique_similar validation failed")

    # check local unique for contents of a list
    config.x_list[2].name = "random"
    config.serialize()

    config.local_unique_obj_list.localuniqueobj(first_name="str1")
    config.local_unique_obj_list.localuniqueobj(first_name="str1")
    with pytest.raises(Exception) as execinfo:
        config.serialize()
    assert execinfo.value.args[0] in "first_name with str1 already exists"

    config.local_unique_obj_list[-1].first_name = "str2"
    s_obj = config.serialize(encoding="dict")

    with pytest.raises(Exception) as execinfo:
        s_obj["local_unique_obj_list"][-1]["first_name"] = "str1"
        config.deserialize(s_obj)
    assert execinfo.value.args[0] in "first_name with str1 already exists"

    u1 = config.local_unique_obj_list[0]
    u2 = config.local_unique_obj_list[1]
    u1.first_name = "f1"
    u2.first_name = "f1"
    y1 = u1.intermediate_list.add(name="n1")
    y1.leaf_list.add(name="n11")
    y1.leaf_list.add(name="n11")
    u2.intermediate_list.add(name="n2")
    u2.intermediate_list.add(name="n2")

    with pytest.raises(Exception) as execinfo:
        config.serialize()
    assert "name with n11 already exists" in execinfo.value.args[0]
    assert "name with n2 already exists" in execinfo.value.args[0]
    assert "first_name with f1 already exists" in execinfo.value.args[0]

    u1.first_name = "f11"
    config.local_unique_obj_list[0].intermediate_list[0].leaf_list[
        0
    ].name = "n12"
    config.local_unique_obj_list[1].intermediate_list[0].name = "n1"
    s_obj = config.serialize(encoding="dict")
    s_obj["local_unique_obj_list"][0]["first_name"] = "f1"
    s_obj["local_unique_obj_list"][0]["intermediate_list"][0]["leaf_list"][0][
        "name"
    ] = "n11"
    s_obj["local_unique_obj_list"][1]["intermediate_list"][0]["name"] = "n2"
    with pytest.raises(Exception) as execinfo:
        config.deserialize(s_obj)
    assert "name with n11 already exists" in execinfo.value.args[0]
    assert "name with n2 already exists" in execinfo.value.args[0]
    assert "first_name with f1 already exists" in execinfo.value.args[0]

    # **********************************************

    # config.z_object.name = "local_unique"
    # config.y_object.y_name = "123"
    # config.x_list.zobject(name="local_unique")
    # try:
    #     config.validate()
    # except Exception as e:
    #     if "local_unique already exists" not in str(
    #         e
    #     ) or "y_name is not a valid type" not in str(e):
    #         pytest.fail("validation failed")
    # try:
    #     config.serialize()
    # except Exception as e:
    #     if "local_unique already exists" not in str(
    #         e
    #     ) or "y_name is not a valid type" not in str(e):
    #         pytest.fail("validation failed")

    # config.x_list[0].name = "local_unique_1"
    # config.y_object.y_name = "local_unique_1"
    # data = config.serialize(config.DICT)
    # data["y_object"]["y_name"] = "local_unique_decode"
    # data["z_object"]["name"] = "local_unique_decode"
    # try:
    #     config.deserialize(data)
    # except Exception as e:
    #     if "local_unique_decode already exists" not in str(
    #         e
    #     ) or "y_name is not a valid type" not in str(e):
    #         pytest.fail("validation failed")

    # config.name = "global_local_same_name_check"
    # config.x_list.zobject(name="global_local_same_name_check")
    # try:
    #     config.validate()
    # except Exception as e:
    #     pytest.fail("validation failed\n {e}".format(e=e))
