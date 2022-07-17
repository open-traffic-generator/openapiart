import os
import fnmatch
import jsonpath_ng
from ruamel import yaml


class AutoFieldUid(object):
    """This utility will use to fill x-field-uid and x-enum-values

        Args
    ----
        parent_folders (list): Parent or top level folder of entire yaml
        output_dir (str): Output directory
    """

    _FIELD_UID = "x-field-uid"

    def __init__(self, parent_folders):
        self._files = []
        for path in parent_folders:
            for r, d, f in os.walk(path):
                for file in f:
                    if fnmatch.fnmatch(file, '*.yaml'):
                        self._files.append(os.path.join(r, file))

    def annotate(self):
        for filename in self._files:
            with open(filename) as fid:
                yobject = yaml.load(
                    fid, Loader=yaml.RoundTripLoader, preserve_quotes=True
                )
            self._annotate_msg_fields(yobject)
            self._annnotate_response_fields(yobject)
            self._dump_file(
                filename, yobject
            )

    def _dump_file(self, output_filename, content):
        with open(output_filename, "w") as fp:
            yaml.dump(content, fp, Dumper=yaml.RoundTripDumper)

    def _annnotate_response_fields(self, yobject):
        path_object = yobject.get("paths")
        if path_object is None:
            return
        print("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
        for response in jsonpath_ng.parse("$..responses").find(
                path_object
        ):
            for code, code_schema in response.value.items():
                print(code)

    def _annotate_enum_fields(self, property_object):
        if "type" in property_object:
            type = property_object["type"]
            if type == "string" and "enum" in property_object:
                x_enum = dict()
                for idx, enum in enumerate(property_object["enum"]):
                    x_enum[enum] = {
                        AutoFieldUid._FIELD_UID: idx + 1
                    }
                property_object.update({"x-enum": x_enum})
                property_object.pop("enum")
            elif type == "array" and "items" in property_object:
                self._annotate_enum_fields(
                    property_object["items"]
                )

    def _annotate_msg_fields(self, yobject):
        components_object = yobject.get("components")
        if components_object is None:
            return
        schema_objects = components_object.get("schemas")
        if schema_objects is None:
            return
        for schema_name, schema_object in schema_objects.items():
            # ignore content field as it always contain single value
            if "properties" not in schema_object:
                continue
            id = 0
            for property_name, property_object in schema_object["properties"].items():
                id += 1
                if not isinstance(property_object, dict):
                    print("schema %s do not have dict of %s" % (
                        schema_name, property_name
                    ))
                    continue
                property_object[AutoFieldUid._FIELD_UID] = id
                self._annotate_enum_fields(
                    property_object
                )


if __name__ == "__main__":
    # api_files = [
    #     "D:/OTG/Codebase/openapiart/openapiart/tests/api/info.yaml",
    #     "D:/OTG/Codebase/openapiart/openapiart/tests/common/common.yaml",
    #     "D:/OTG/Codebase/openapiart/openapiart/tests/api/api.yaml"
    # ]
    parent_folders = [
        # "D:/OTG/Codebase/models"
        "D:/OTG/Codebase/openapiart/openapiart/tests"
    ]
    # AutoFieldUid(parent_folders, output_dir="D:/OTG/Codebase/openapiart/mmmm").annotate()
    AutoFieldUid(parent_folders).annotate()
