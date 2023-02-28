import pytest


def test_warnings_with_api(api, capfd):

    u_conf = api.update_config()
    u_conf.g.gobject(g_a="abcd")

    api.update_configuration(u_conf)
    out, err = capfd.readouterr()
    print(out, err)
    assert "update_configuration is deprecated, please use post instead" in out
    # assert api.__warnings__ != []
    # assert len(api.__warnings__) == 1


def test_warning_for_primitive_attr(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.space_1 = 56
    conf.str_len = "1245"
    conf.hex_slice = ["str1", "str2"]
    conf.serialize(conf.DICT)

    out, err = capsys.readouterr()
    assert err == ""
    assert "[WARNING]: space_1 is deprecated, Information TBD" in out
    assert "[WARNING]: a is under-review, Information TBD" in out
    assert "[WARNING]: str_len is under-review, Information TBD" in out
    assert "[WARNING]: hex_slice is under-review, Information TBD" in out


def test_warnings_for_non_primitive_attr(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.e.e_a = 100
    conf.e.e_b = 4.5

    conf.serialize(conf.DICT)
    out, err = capsys.readouterr()

    assert err == ""
    assert "e is deprecated, Information TBD" in out


def test_warnings_for_iter_items(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.g.add(name="a", g_c=4.56)
    conf.g.add(name="b", g_c=5.46)
    conf.g.add(name="c", g_c=6.54)

    conf.serialize(conf.DICT)
    out, err = capsys.readouterr()

    assert err == ""
    assert out.count("[WARNING]: g_c is deprecated, Information TBD") == 3


def test_warnings_for_x_enmu_attr(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.response = "status_404"
    conf.serialize(conf.DICT)

    out, err = capsys.readouterr()
    assert err == ""
    assert (
        "[WARNING]: STATUS_404 is deprecated, new code will be coming soon"
        in out
    )


def test_warnings_for_choice_attr(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    j = conf.j.add()
    j.j_b.f_a = "some string"
    conf.serialize(conf.DICT)

    out, err = capsys.readouterr()
    assert err == ""
    assert "[WARNING]: J_B is deprecated, use j_a instead" in out
