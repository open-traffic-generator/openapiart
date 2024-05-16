import pytest


def test_warnings_with_api(api, capfd):
    u_conf = api.update_config()
    u_conf.g.gobject(g_a="abcd")

    api.update_configuration(u_conf)
    out, err = capfd.readouterr()
    print(out, err)
    assert (
        "[WARNING]: update_configuration api is deprecated, please use post instead"
        in err
    )
    assert api.__warnings__ != []
    assert len(api.__warnings__) == 1
    assert (
        api.__warnings__[0]
        == "update_configuration api is deprecated, please use post instead"
    )


def test_warning_for_primitive_attr(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.space_1 = 56
    conf.str_len = "1245"
    conf.hex_slice = ["0xff", "0xab"]
    s_obj = conf.serialize(conf.DICT)

    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: space_1 property in schema PrefixConfig is deprecated, Information TBD"
        in err
    )
    assert (
        "[WARNING]: a property in schema PrefixConfig is under_review, Information TBD"
        in err
    )
    assert (
        "[WARNING]: str_len property in schema PrefixConfig is under_review, Information TBD"
        in err
    )
    assert (
        "[WARNING]: hex_slice property in schema PrefixConfig is under_review, Information TBD"
        in err
    )

    conf.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: space_1 property in schema PrefixConfig is deprecated, Information TBD"
        in err
    )
    assert (
        "[WARNING]: a property in schema PrefixConfig is under_review, Information TBD"
        in err
    )
    assert (
        "[WARNING]: str_len property in schema PrefixConfig is under_review, Information TBD"
        in err
    )
    assert (
        "[WARNING]: hex_slice property in schema PrefixConfig is under_review, Information TBD"
        in err
    )


def test_warnings_for_non_primitive_attr(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.e.e_a = 100
    conf.e.e_b = 4.5

    s_obj = conf.serialize(conf.DICT)
    out, err = capsys.readouterr()

    assert out == ""
    assert (
        "e property in schema PrefixConfig is deprecated, Information TBD"
        in err
    )

    conf.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "e property in schema PrefixConfig is deprecated, Information TBD"
        in err
    )


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

    s_obj = conf.serialize(conf.DICT)
    out, err = capsys.readouterr()

    assert out == ""
    assert (
        err.count(
            "[WARNING]: g_c property in schema GObject is deprecated, Information TBD"
        )
        == 3
    )

    conf.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        err.count(
            "[WARNING]: g_c property in schema GObject is deprecated, Information TBD"
        )
        == 3
    )


def test_warnings_for_x_enmu_attr(api, capsys):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.response = "status_404"
    s_obj = conf.serialize(conf.DICT)

    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: STATUS_404 enum in property response is deprecated, new code will be coming soon"
        in err
    )

    conf.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: STATUS_404 enum in property response is deprecated, new code will be coming soon"
        in err
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
    s_obj = conf.serialize(conf.DICT)

    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: J_B enum in property choice is deprecated, use j_a instead"
        in err
    )

    conf.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: J_B enum in property choice is deprecated, use j_a instead"
        in err
    )


def test_warnings_for_schema(api, capsys):
    conf = api.update_config()
    conf.g.add(g_c=5.67)
    conf.g.add(g_a="asd")
    s_obj = conf.serialize(conf.DICT)

    out, err = capsys.readouterr()
    assert out == ""
    assert (
        err.count(
            "[WARNING]: UpdateConfig is under_review, the whole schema is being reviewed"
        )
        == 1
    )
    assert (
        err.count(
            "[WARNING]: GObject is deprecated, new schema Jobject to be used"
        )
        == 2
    )
    assert (
        err.count(
            "[WARNING]: g_c property in schema GObject is deprecated, Information TBD"
        )
        == 2
    )

    conf.deserialize(s_obj)
    out, err = capsys.readouterr()

    assert out == ""
    assert (
        err.count(
            "[WARNING]: UpdateConfig is under_review, the whole schema is being reviewed"
        )
        == 1
    )
    assert (
        err.count(
            "[WARNING]: GObject is deprecated, new schema Jobject to be used"
        )
        == 2
    )
    assert (
        err.count(
            "[WARNING]: g_c property in schema GObject is deprecated, Information TBD"
        )
        == 2
    )
