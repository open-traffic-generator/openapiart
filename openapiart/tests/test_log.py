import pytest
import sys
import os

sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art"))
import sanity

def test_log():
    sanity_api = sanity.api()
    sanity_api.logger.info("Testing logger")
    sanity_api.logger.debug("Testing logger")