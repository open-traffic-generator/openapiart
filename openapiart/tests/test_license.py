import pytest


def test_license(proto_file_name):
    """Verify the license is set to NO-LICENSE-PRESENT"""
    with open(proto_file_name) as fp:
        contents = fp.read()

    assert "NO-LICENSE-PRESENT" in contents


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
