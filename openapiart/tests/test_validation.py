import os
import pytest
import openapiart


def test_validation():
    rootfolder = os.path.dirname(__file__)
    outputfolder = os.path.join(rootfolder, "unittest-tmp")

    try:
        openapiart.OpenApiArt(
            api_files=[
                os.path.join(rootfolder, "config/negative_nested_config.yaml")
            ],
            artifact_dir=outputfolder,
            protobuf_name="unittest-tmp",
        )
        assert False
    except Exception as error:
        message = str(error)
        assert message.find("nested_component_1") > 0
        assert message.find("nested_component_2") > 0


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
